package api

import (
	"log"
	"net/http"

	"github.com/raffleberry/sqlsheet/internal/db"
	"github.com/raffleberry/sqlsheet/web/tmpl"
)

const formsPath = "/form/"

func forms(w http.ResponseWriter, r *http.Request) {
	tableName := r.URL.Path[len(formsPath):]

	switch tableName {
	case "":

		tables, err := db.Db.GetTables()
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = tmpl.Use(w, "form-all", tables)

		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:

		cols, err := db.Db.GetColumns(tableName)
		if err != nil {
			log.Println("ERROR", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println(cols)

		err = tmpl.Use(w, "form-table",
			map[string]interface{}{
				"Name":    tableName,
				"Columns": cols,
			},
		)

		if err != nil {
			log.Println("ERROR", err)
		}

	}
}
