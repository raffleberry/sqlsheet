package api

import (
	"fmt"
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

	log.Printf("%v: %v\n", r.Method, r.URL.Path)

	switch el {
	case "save":
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				log.Printf("ERR : %s\n", err)
				tmpl.Use(w, "form-create:save", err)
			} else {
				for dt, fn := range r.Form {
					log.Println(dt, fn)
				}
				tmpl.Use(w, "form-create:save", r.Form)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "400: Bad Request")
		}
	case "":
		log.Println("create")
		err := tmpl.Use(w, "form-create", "")

		if err != nil {
			log.Println("ERROR", err)
		}
	default:
		log.Printf("%v Bad Request\n", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "400: Bad Request")
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
