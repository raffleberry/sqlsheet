package main

import (
	"github.com/raffleberry/sqlsheet/pkg/store"
	"github.com/raffleberry/sqlsheet/web/api"
)

type Test struct {
	Title string
}

func main() {
	store.Connect()
	api.Start()

}
