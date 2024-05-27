package server

import (
	"fmt"
	"net"
	"net/http"
)

// start a http server and returns <-ready, <-done, server
func Start(port int) (<-chan bool, <-chan bool, *http.Server) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	})

	ready := make(chan bool, 1)
	done := make(chan bool, 1)

	server := http.Server{}

	go func() {
		defer close(ready)
		defer close(done)

		l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

		if err != nil {
			panic(err)
		}

		server.Addr = l.Addr().String()

		ready <- true

		err = server.Serve(l)

		if err != nil && err != http.ErrServerClosed {
			fmt.Println(err.Error())
		}
	}()

	return ready, done, &server
}
