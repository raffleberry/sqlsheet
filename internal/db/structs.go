package db

import "database/sql"

type Column struct {
	ColumnName      string
	ColumnType      string
	OrdinalPosition int
	ColumnDefault   sql.NullString
	IsNullable      string
}

type Table struct {
	Name    string
	Columns []Column
}

type View struct {
	ID    int64
	Name  string
	Query string
}
