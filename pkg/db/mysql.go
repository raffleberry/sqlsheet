package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var Conn *sql.DB
var Err error

func Init() {
	cfg := mysql.Config{
		User:   "sqlsheet",
		Passwd: "sqlsheet",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "sqlsheet",
	}

	Conn, Err = sql.Open("mysql", cfg.FormatDSN())

	if Err != nil {
		log.Println(Err)
	}

	pingErr := Conn.Ping()
	if pingErr != nil {
		log.Println(pingErr)
	}
	log.Println("Connected to MySql!")
}
