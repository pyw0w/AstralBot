package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database interface {
	Connect() error
	GetDB() interface{}
}

type MongoDB struct {
	ConnectionString string
	Client           *mongo.Client
}

func (db *MongoDB) Connect() error {
	clientOptions := options.Client().ApplyURI(db.ConnectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	db.Client = client
	return nil
}

func (db *MongoDB) GetDB() interface{} {
	return db.Client
}

type SQLiteDB struct {
	ConnectionString string
	DB               *gorm.DB
}

func (db *SQLiteDB) Connect() error {
	database, err := gorm.Open(sqlite.Open(db.ConnectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	db.DB = database
	return nil
}

func (db *SQLiteDB) GetDB() interface{} {
	return db.DB
}

type MySQLDB struct {
	ConnectionString string
	DB               *gorm.DB
}

func (db *MySQLDB) Connect() error {
	database, err := gorm.Open(mysql.Open(db.ConnectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	db.DB = database
	return nil
}

func (db *MySQLDB) GetDB() interface{} {
	return db.DB
}

func NewDatabase(dbType, connectionString string) (Database, error) {
	switch dbType {
	case "mongodb":
		return &MongoDB{ConnectionString: connectionString}, nil
	case "sqlite":
		return &SQLiteDB{ConnectionString: connectionString}, nil
	case "mysql":
		return &MySQLDB{ConnectionString: connectionString}, nil
	default:
		return nil, fmt.Errorf("неизвестный тип базы данных: %s", dbType)
	}
}
