package api

import (
	"embed"
	"log"
	"net/http"
	"path/filepath"

	"github.com/raffleberry/sqlsheet/web/server"
	"github.com/raffleberry/sqlsheet/web/tmpl"
)

//go:embed static/*
var static embed.FS

func Start() {
	mux := http.NewServeMux()

	mux.HandleFunc(StaticFilesPath, StaticFilesHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		err := tmpl.Use(w, "home", url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})

	ready, done, s := server.New(5500, mux)
	<-ready

	log.Printf("server started on : http://%v\n", s.Addr)

	<-done
}

var StaticFilesPath = "/static/"

func StaticFilesHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Clean(r.URL.Path)
	http.ServeFileFS(w, r, static, path)
}
