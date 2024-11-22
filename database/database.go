package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Plugin struct{}

func (p *Plugin) Init() error {
	var err error
	db, err = sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		return err
	}

	log.Println("Database connected")
	return nil
}

func (p *Plugin) Run() error {
	return nil
}

func GetDB() *sqlx.DB {
	return db
}
