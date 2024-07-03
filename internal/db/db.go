package db

import (
	"database/sql"
	"os"

	"github.com/raffleberry/sqlsheet/internal/utils"
)

type database interface {
	Init() error
	Conn() *sql.DB
	GetTables() ([]string, error)
	GetColumns(string) ([]Column, error)
	IsDataTypeSupported(string) bool
	GetViews() ([]View, error)
	GetViewById(int) (View, error)
}

var Db database

func init() {
	Db = &mySql{
		user:   os.Getenv("MYSQL_USER"),
		passwd: os.Getenv("MYSQL_PASS"),
		addr:   os.Getenv("MYSQL_ADDR"),
		dbname: os.Getenv("MYSQL_DB"),
	}
	err := Db.Init()
	utils.Panic(err)
}
