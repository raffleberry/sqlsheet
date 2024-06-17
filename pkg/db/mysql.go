package db

import (
	"database/sql"
	"fmt"
	"log"
	"slices"

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

func ListTables() ([]string, error) {
	const q = "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_TYPE = 'BASE TABLE' AND TABLE_SCHEMA='sqlsheet';"
	tables := make([]string, 0)
	rows, err := Conn.Query(q)
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

func ListColumns(table string) ([]string, []string, error) {
	var q = fmt.Sprintf("SELECT `COLUMN_NAME`, `COLUMN_TYPE` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA`='sqlsheet' AND `TABLE_NAME`='%s';", table)
	col_names := make([]string, 0)
	col_types := make([]string, 0)
	rows, err := Conn.Query(q)
	if err != nil {
		return col_names, col_types, err
	}

	for rows.Next() {
		var col_name string
		var col_type string
		err := rows.Scan(&col_name, &col_type)
		if err != nil {
			return col_names, col_types, err
		}
		col_names = append(col_names, col_name)
		col_types = append(col_types, col_type)
	}

	if err := rows.Err(); err != nil {
		return col_names, col_types, err
	}

	return col_names, col_types, nil

}

func IsDataTypeSupported(dt string) bool {
	var sdt = []string{
		"VARCHAR(128)",
		"VARCHAR(2048)",
		"TEXT",
		"MEDIUMTEXT",
		"LONGTEXT",
		"BOOL",
		"INT",
		"FLOAT",
		"DATE",
		"TIME",
		"YEAR",
		"DATETIME",
		"TIMESTAMP",
	}
	return slices.Contains(sdt, dt)
}
