package store

import (
	"github.com/raffleberry/sqlsheet/pkg/db"
	"github.com/raffleberry/sqlsheet/pkg/utils"
)

func Connect() {
	db.Init()
	utils.Panic(db.Err)
}

func ViewAll() ([]View, error) {
	var views []View
	rows, err := db.Conn.Query("SELECT * from views;")
	if err != nil {
		return []View{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var v View
		if err := v.Scan(rows); err != nil {
			return []View{}, err
		}
		views = append(views, v)
	}
	if err = rows.Err(); err != nil {
		return []View{}, err
	}
	return views, nil
}

func FormAll() ([]Form, error) {
	var forms []Form
	rows, err := db.Conn.Query("SELECT * from forms;")
	if err != nil {
		return []Form{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var v Form
		if err := v.Scan(rows); err != nil {
			return []Form{}, err
		}
		forms = append(forms, v)
	}
	if err = rows.Err(); err != nil {
		return []Form{}, err
	}
	return forms, nil
}
