package main

import (
	"log"

	"github.com/raffleberry/sqlsheet/pkg/store"
	"github.com/raffleberry/sqlsheet/pkg/utils"
	"github.com/raffleberry/sqlsheet/web/api"
)

type Test struct {
	Title string
}

func main() {
	store.Connect()

	forms, err := store.FormAll()
	utils.Panic(err)
	log.Println(forms)

	views, err := store.ViewAll()
	utils.Panic(err)
	log.Println(views)

	api.Start()

}
