package store

import "database/sql"

type View struct {
	ID    int64
	Query string
}

func (v *View) Scan(row *sql.Rows) error {
	err := row.Scan(&v.ID, &v.Query)
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

func (v *Form) Scan(row *sql.Rows) error {
	err := row.Scan(&v.ID, &v.ColumnNames, &v.ColumnTypes)
	if err != nil {
		return err
	}
	return nil
}
