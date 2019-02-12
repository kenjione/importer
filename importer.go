package importer

import (
	"github.com/kenjione/importer/internal"
	"github.com/kenjione/importer/internal/model"

	"github.com/go-pg/pg"
)

type Importer struct {
	Conn *pg.DB
}

func NewImporter(cnf *Config) *Importer {
	conn := pg.Connect(&pg.Options{
		Database: cnf.DatabaseHost,
		User:     cnf.DatabaseUser,
		Password: cnf.DatabasePassword,
	})

	model.CreateSchema(conn)

	return &Importer{
		Conn: conn,
	}
}

func (i *Importer) Close() {
	i.Conn.Close()
}

func (i *Importer) Parse(filename string) *internal.Stats {
	parser := internal.NewLocationParser(filename, i.Conn)
	defer parser.File.Close()

	return parser.Parse()
}

func (i *Importer) FindByIP(ip string) ([]byte, error) {
	repo := model.NewRepository(i.Conn)

	return repo.FindByIP(ip)
}
