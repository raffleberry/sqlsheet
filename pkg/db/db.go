package db

import (
	"database/sql"

	"github.com/raffleberry/sqlsheet/pkg/utils"
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
		user:   "sqlsheet",
		passwd: "sqlsheet",
		addr:   "127.0.0.1:3306",
		dbname: "sqlsheet",
	}
	err := Db.Init()
	utils.Panic(err)
}
