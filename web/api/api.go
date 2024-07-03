package api

import (
	"context"
	"embed"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/raffleberry/sqlsheet/internal/utils"
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

	if utils.DEV {
		wd, er := os.Getwd()
		utils.Panic(er)
		mux.Handle(staticFilesPath, http.FileServer(http.Dir(filepath.Join(wd, "web", "api"))))
	} else {
		mux.Handle(staticFilesPath, http.FileServerFS(staticFs))
	}

	mux.HandleFunc(formsPath, forms)

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
