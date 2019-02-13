package internal

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg"
	"github.com/kenjione/importer/internal/model"
)

type Stats struct {
	Duration string
	Accepted int
	Invalid  int
	NotSaved int
	List     []*model.Location
}

type Parser struct {
	File *os.File
	Repo *model.Repository
}

func NewLocationParser(filename string, conn *pg.DB) *Parser {
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	repo := model.NewRepository(conn)

	return &Parser{
		File: csvFile,
		Repo: repo,
	}
}

func (p *Parser) Parse() *Stats {
	reader := csv.NewReader(p.File)
	stats := &Stats{}

	start := time.Now()
	total := 0

	fmt.Println("Import started")

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		total++

		if total%10000 == 0 {
			fmt.Printf("%+v records handled!", total)
		}

		if err != nil {
			log.Fatal(err)
		}

		l := &model.Location{
			IPaddress:    line[0],
			CountryCode:  line[1],
			Country:      line[2],
			City:         line[3],
			Latitude:     line[4],
			Longitude:    line[5],
			MysteryValue: line[6],
		}

		if err := l.Validate(); err != nil {
			stats.Invalid++
			continue
		}

		if err := p.Repo.Store(l); err != nil {
			stats.NotSaved++
			continue
		}

		stats.Accepted++
	}

	stats.Duration = time.Since(start).String()
	return stats
}
