package api

import (
	"embed"
	"log"
	"net/http"

	"github.com/raffleberry/sqlsheet/web/server"
	"github.com/raffleberry/sqlsheet/web/tmpl"
)

//go:embed static/*
var static embed.FS

func Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		err := tmpl.Use(w, "home", url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
	mux.HandleFunc(staticFilesPath, staticFilesHandler)
	mux.HandleFunc(formCreatePath, formCreate)
	mux.HandleFunc(formEditPath, formEdit)
	mux.HandleFunc(viewsPath, views)

	ready, done, s := server.New(5500, mux)
	<-ready

	log.Printf("server started on : http://%v\n", s.Addr)

	<-done
}

var staticFilesPath = "/static/"

func staticFilesHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	http.ServeFileFS(w, r, static, path)
}

var formCreatePath = "/form/create/"

func formCreate(w http.ResponseWriter, r *http.Request) {

}

var formEditPath = "/form/edit/"

func formEdit(w http.ResponseWriter, r *http.Request) {

}

var viewsPath = "/view/"

func views(w http.ResponseWriter, r *http.Request) {
	tmpl.Use(w, "view", struct {
		f string
		b string
	}{"foo", "bar"})
}
