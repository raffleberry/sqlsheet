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

	switch idstr {
	case "":

		forms, err := store.FormAll()
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = tmpl.Use(w, "form-all", forms)

		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	default:

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

	el := r.URL.Path[len(formCreatePath):]

	log.Println(el)

	switch el {
	case "":
		log.Println("create")
		err := tmpl.Use(w, "form-create", "")

		if err != nil {
			log.Println("ERROR", err)
		}

	case "new-field":
		log.Println("new-field")
		err := tmpl.Use(w, "form-create:new-field", "")

		if err != nil {
			log.Println("ERROR", err)
		}

	default:

	}

}

const formEditPath = "/form/edit/"

func formEdit(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Path[len(formEditPath):]

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

	log.Println("EDITING: ", form)

	form.Name = "EDITING - " + form.Name

	err = tmpl.Use(w, "form", form)

	if err != nil {
		log.Println("ERROR", err)
	}

}
