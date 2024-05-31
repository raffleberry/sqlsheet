package server_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/raffleberry/sqlsheet/pkg/utils"
	"github.com/raffleberry/sqlsheet/web/server"
)

func TestNew(t *testing.T) {
	ready, done, s := server.New(0, nil)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	defer cancel()
	select {
	case <-ctx.Done():
		t.Fatalf("Timedout waiting for server to START - %s", ctx.Err().Error())
	case <-ready:
		cancel()
	}

	// TEST
	addr, err := net.ResolveTCPAddr("tcp4", s.Addr)
	utils.TPanic(t, err)
	t.Logf("IP: %v", addr.IP)
	t.Logf("Port: %v", addr.Port)

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
