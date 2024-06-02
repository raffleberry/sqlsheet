package store

import "database/sql"

type View struct {
	ID    int64
	Name  string
	Query string
}

func (v *View) Scan(rows *sql.Rows) error {
	err := rows.Scan(&v.ID, &v.Name, &v.Query)
	if err != nil {
		return err
	}
	return nil
}

type Form struct {
	ID          int64
	ColumnNames string
	ColumnTypes string
}

func (v *Form) Scan(rows *sql.Rows) error {
	err := rows.Scan(&v.ID, &v.ColumnNames, &v.ColumnTypes)
	if err != nil {
		return err
	}
	return nil
}
