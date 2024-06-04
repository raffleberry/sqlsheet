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

const TableLs = "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_TYPE = 'BASE TABLE' AND TABLE_SCHEMA='sqlsheet';"

func ListTables() ([]string, error) {
	tables := make([]string, 0)
	rows, err := Conn.Query(TableLs)
	if err != nil {
		return tables, err
	}

	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			return tables, err
		}
		tables = append(tables, table)
	}

	if err := rows.Err(); err != nil {
		return tables, err
	}

	return tables, nil
}
