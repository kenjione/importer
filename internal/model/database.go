package model

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func CreateSchema(db *pg.DB) error {
	return db.CreateTable((*Location)(nil), &orm.CreateTableOptions{
		Temp: false,
	})
}
