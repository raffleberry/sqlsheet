package server

import (
	"fmt"
	"net"
	"net/http"

	utils "github.com/raffleberry/sqlsheet/pkg"
)

// port = 0, for a random port
// mux can be nil
// returns [<-ready], [<-done], [*http.Server]
func New(port int, mux http.Handler) (<-chan bool, <-chan bool, *http.Server) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	})

	ready := make(chan bool, 1)
	done := make(chan bool, 1)

	server := http.Server{Handler: mux}

	go func() {
		defer close(ready)
		defer close(done)

		listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", "127.0.0.1", port))
		utils.Panic(err)

		server.Addr = listener.Addr().String()

		ready <- true

		err = server.Serve(listener)

		if err != nil && err != http.ErrServerClosed {
			fmt.Println(err.Error())
		}
	}()

	return ready, done, &server
}
