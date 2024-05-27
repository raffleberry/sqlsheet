package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/raffleberry/sqlsheet/web/server"
)

func TestStart(t *testing.T) {
	ready, done, s := server.Start(0)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	defer cancel()
	select {
	case <-ctx.Done():
		t.Fatalf("Timedout waiting for server to START - %s", ctx.Err().Error())
	case <-ready:
		cancel()
	}

	// WRITE TESTS HERE
	t.Log("TODO: Parse Port")

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	s.Shutdown(context.Background())
	select {
	case <-ctx.Done():
		t.Fatalf("Timedout waiting for server to SHUTDOWN - %s", ctx.Err().Error())
	case <-done:
		cancel()
	}

}
