package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/raffleberry/sqlsheet/pkg/utils"
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
	utils.Panic(Err)

	pingErr := Conn.Ping()
	utils.Panic(pingErr)

	log.Println("Connected to MySql!")
}
