package api

import (
	"context"
	"embed"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/raffleberry/sqlsheet/pkg/store"
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

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	ready, done, s := server.New(5500, mux)
	<-ready

	log.Printf("server started on : http://%v\n", s.Addr)
	log.Println("Ctrl + C to shutdown..")
	<-sigint

	ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(time.Second*3))
	defer cancel()

	s.Shutdown(ctx)

	log.Println("Shutting down..")

	select {
	case <-done:
		log.Println("bye 👋")
	case <-ctx.Done():
		log.Println("shutdown request timedout..")
		log.Println("ok, bye 👋")
	}
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

	views, err := store.ViewAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmpl.Use(w, "view", views)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
