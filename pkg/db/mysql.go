package db

import (
	"database/sql"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/go-sql-driver/mysql"
)

type mySql struct {
	user            string
	passwd          string
	addr            string
	dbname          string
	c               *sql.DB
	Err             error
	internalTables  []string
	internalColumns []string
}

func (m *mySql) Init() error {
	m.internalTables = []string{"views"}
	m.internalColumns = []string{"id"}

	cfg := mysql.Config{
		User:   m.user,
		Passwd: m.passwd,
		Net:    "tcp",
		Addr:   m.addr,
		DBName: m.dbname,
	}

	m.c, m.Err = sql.Open("mysql", cfg.FormatDSN())

	if m.Err != nil {
		return m.Err
	}

	m.Err = m.c.Ping()
	if m.Err != nil {
		return m.Err
	}

	log.Println("Connected to MySql!")
	return m.Err
}

func (m *mySql) Conn() *sql.DB {
	return m.c
}

func (m *mySql) GetTables() ([]string, error) {
	q := fmt.Sprintf("SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_TYPE = 'BASE TABLE' AND TABLE_SCHEMA='%s';", m.dbname)
	tables := make([]string, 0)
	rows, err := m.c.Query(q)
	if err != nil {
		return tables, err
	}

	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			return tables, err
		}
		if slices.Contains(m.internalTables, table) {
			continue
		}
		tables = append(tables, table)
	}

	if err := rows.Err(); err != nil {
		return tables, err
	}

	return tables, nil
}

func (m *mySql) GetColumns(tableName string) ([]Column, error) {
	var q = fmt.Sprintf("SELECT `COLUMN_NAME`, `COLUMN_TYPE`, `ORDINAL_POSITION`, `COLUMN_DEFAULT`, `IS_NULLABLE` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA`='%s' AND `TABLE_NAME`='%s';", m.dbname, tableName)
	rows, err := m.c.Query(q)
	columns := []Column{}
	if err != nil {
		return columns, err
	}

	ok := false

	for rows.Next() {
		ok = true
		c := Column{}
		err := rows.Scan(&c.ColumnName, &c.ColumnType, &c.OrdinalPosition, &c.ColumnDefault, &c.IsNullable)
		if err != nil {
			return columns, err
		}
		if slices.Contains(m.internalColumns, c.ColumnName) {
			continue
		}
		columns = append(columns, c)
	}

	if !ok {
		return columns, sql.ErrNoRows
	}

	if err := rows.Err(); err != nil {
		return columns, err
	}

	slices.SortFunc(columns, func(a Column, b Column) int {
		return a.OrdinalPosition - b.OrdinalPosition
	})

	return columns, nil

}

func (m *mySql) IsDataTypeSupported(dt string) bool {
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
	return slices.Contains(sdt, strings.ToUpper(dt))
}

func (m *mySql) GetViews() ([]View, error) {
	var views []View
	rows, err := m.c.Query("SELECT id, name, query FROM views;")
	if err != nil {
		return []View{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var v View
		if err := rows.Scan(&v.ID, &v.Name, &v.Query); err != nil {
			return []View{}, err
		}
		views = append(views, v)
	}
	if err = rows.Err(); err != nil {
		return []View{}, err
	}
	return views, nil
}

func (m *mySql) GetViewById(id int) (View, error) {
	var v View
	rows, err := m.c.Query("SELECT id, name, query FROM views WHERE id=?;", id)
	if err != nil {
		return v, err
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&v.ID, &v.Name, &v.Query)
	if err != nil {
		return v, err
	}

	return v, nil

}
