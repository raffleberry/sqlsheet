package api

import (
	"context"
	"embed"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/raffleberry/sqlsheet/pkg/db"
	"github.com/raffleberry/sqlsheet/pkg/store"
	"github.com/raffleberry/sqlsheet/web/server"
	"github.com/raffleberry/sqlsheet/web/tmpl"
)

//go:embed static
var staticFs embed.FS

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
	mux.Handle(staticFilesPath, http.FileServerFS(staticFs))

	mux.HandleFunc(formsPath, forms)
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

const staticFilesPath = "/static/"

const viewsPath = "/view/"

func views(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Path[len(viewsPath):]
	if len(idstr) == 0 {

		views, err := store.ViewAll()
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

		view, err := store.ViewId(id)
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		rows, err := db.Conn.Query(view.Query)

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
			V  store.View
			Th []string
			Td [][]string
		}{view, cols, td})

		if err != nil {
			log.Println("ERROR", err)
		}
	}
}
