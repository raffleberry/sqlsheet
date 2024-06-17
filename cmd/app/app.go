package main

import (
	"github.com/raffleberry/sqlsheet/pkg/store"
	"github.com/raffleberry/sqlsheet/web/api"
)

func main() {
	store.Connect()
	api.Start()

}
