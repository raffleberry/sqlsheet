package web

import (
	"embed"
	"html/template"
	"os"
)

type Pet struct {
	Name   string
	Sex    string
	Intact bool
	Age    string
	Breed  string
}

//go:embed templates
var Tmpl embed.FS

func PrintDogs(dogs []Pet) {
	t, err := template.ParseFS(Tmpl, "templates/hello.tmpl")
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "foo", "World")
	if err != nil {
		panic(err)
	}
}
