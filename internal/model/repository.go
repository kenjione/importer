package model

import (
	"encoding/json"

	"github.com/go-pg/pg"
)

type Repository struct {
	conn *pg.DB
}

func NewRepository(conn *pg.DB) *Repository {
	return &Repository{
		conn: conn,
	}
}

func (r *Repository) FindByIP(ip string) ([]byte, error) {
	l := &Location{}
	err := r.conn.Model(l).Where("ip_address = ?", ip).Limit(1).Select()

	if err != nil {
		return []byte{}, err
	}

	return json.Marshal(l)
}

func (r *Repository) Store(l *Location) error {
	return r.conn.Insert(l)
}
