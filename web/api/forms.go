package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/raffleberry/sqlsheet/pkg/db"
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

func validateFormCreate(f url.Values) error {
	errStr := ""

	tables, err := db.ListTables()

	if err != nil {
		return err
	}

	for f, t := range f {
		ta := strings.Split(f, ".")
		if len(ta) != 2 {
			errStr += fmt.Sprintf("`<table>.<column>` expected found: `%s`, ", f)
			continue
		}
		table, col := ta[0], ta[1]

		if !slices.Contains(tables, table) {
			errStr += fmt.Sprintf("`%s` table from `%s` doesn't exist, ", table, f)
			continue
		}

		colNs, _, err := db.ListColumns(table)
		if err != nil {
			return err
		}

		if !slices.Contains(colNs, col) {
			errStr += fmt.Sprintf("`%s` column in `%s` table doesn't exist, ", col, table)
		}

		if !db.IsDataTypeSupported(t[0]) {
			errStr += fmt.Sprintf("`%s` with datatype:`%s` isn't supported, ", f, t[0])
		}
	}

	if errStr != "" {
		return errors.New(errStr)
	}

	return nil
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
				err := validateFormCreate(r.Form)
				if err != nil {
					log.Printf("ERR : %s\n", err)
					tmpl.Use(w, "form-create:save", err)
				} else {
					tmpl.Use(w, "form-create:save", r.Form)
				}
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
