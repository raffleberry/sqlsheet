package main

import (
	"log"

	"github.com/raffleberry/sqlsheet/web/api"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	api.Start()

}
