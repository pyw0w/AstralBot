package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	Connect() error
	GetDB() interface{}
}

type SQLiteDB struct {
	ConnectionString string
	DB               *sql.DB
}

func (db *SQLiteDB) Connect() error {
	database, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		return err
	}
	db.DB = database
	return nil
}

func (db *SQLiteDB) GetDB() interface{} {
	return db.DB
}

func NewDatabase(dbType, connectionString string) (Database, error) {
	switch dbType {
	case "sqlite":
		return &SQLiteDB{ConnectionString: connectionString}, nil
	default:
		return nil, fmt.Errorf("неизвестный тип базы данных: %s", dbType)
	}
}
