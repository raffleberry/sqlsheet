package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/raffleberry/sqlsheet/internal/db"
	"github.com/raffleberry/sqlsheet/web/tmpl"
)

const viewsPath = "/view/"

func views(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Path[len(viewsPath):]
	if len(idstr) == 0 {

		views, err := db.Db.GetViews()
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = tmpl.Use(w, "views", views)

		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	} else {
		id, err := strconv.Atoi(idstr)
		if err != nil {
			log.Println("ERROR", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		view, err := db.Db.GetViewById(id)
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		rows, err := db.Db.Conn().Query(view.Query)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		defer rows.Close()

		cols, err := rows.Columns()
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		td := make([][]string, 0)

		for rows.Next() {
			vals := make([]string, len(cols))
			pointers := make([]interface{}, len(cols))
			for i, _ := range vals {
				pointers[i] = &vals[i]
			}
			rows.Scan(pointers...)
			td = append(td, vals)
		}

		err = tmpl.Use(w, "view", struct {
			V  db.View
			Th []string
			Td [][]string
		}{view, cols, td})

		if err != nil {
			log.Println("ERROR", err)
		}
	}
}
