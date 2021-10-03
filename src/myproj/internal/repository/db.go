package repository

import (
	"github.com/jmoiron/sqlx"
)
type PostgresDB struct {
	DB *sqlx.DB
}

var pgDB PostgresDB

func InitDB(db *sqlx.DB){
	pgDB.DB = db
}

func DB()PostgresDB{
	return pgDB
}