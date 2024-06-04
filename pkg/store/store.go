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
	rows, err := db.Conn.Query("SELECT * FROM views;")
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

func ViewId(id int) (View, error) {
	var view View
	rows, err := db.Conn.Query("SELECT * FROM views WHERE id=?;", id)
	if err != nil {
		return view, err
	}
	defer rows.Close()
	rows.Next()
	err = view.Scan(rows)
	if err != nil {
		return view, err
	}

	return view, nil

}

func FormAll() ([]Form, error) {
	var forms []Form
	rows, err := db.Conn.Query("SELECT * FROM forms;")
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

func FormId(id int) (Form, error) {
	var form Form
	rows, err := db.Conn.Query("SELECT * FROM forms WHERE id=?;", id)
	if err != nil {
		return form, err
	}
	defer rows.Close()
	rows.Next()
	err = form.Scan(rows)
	if err != nil {
		return form, err
	}

	return form, nil

}
