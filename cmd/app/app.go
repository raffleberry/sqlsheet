package main

import (
	"os"

	"github.com/raffleberry/sqlsheet/web/tmpl"
)

func main() {
	tmpl.Use(os.Stdout, "foo", "World")
}
