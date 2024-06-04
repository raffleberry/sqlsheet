package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/raffleberry/sqlsheet/pkg/store"
	"github.com/raffleberry/sqlsheet/web/tmpl"
)

const formsPath = "/form/"

func forms(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Path[len(formsPath):]
	if len(idstr) == 0 {

		forms, err := store.FormAll()
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = tmpl.Use(w, "forms", forms)

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

		form, err := store.FormId(id)
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Println(form)

		err = tmpl.Use(w, "form", form)

		if err != nil {
			log.Println("ERROR", err)
		}
	}
}

const formCreatePath = "/form/create/"

func formCreate(w http.ResponseWriter, r *http.Request) {

}

const formEditPath = "/form/edit/"

func formEdit(w http.ResponseWriter, r *http.Request) {

}
