package tmpl

import (
	"embed"
	"html/template"
	"io"
	"os"
)

type Pet struct {
	Name   string
	Sex    string
	Intact bool
	Age    string
	Breed  string
}

//go:embed templates.tmpl
var Tmpl embed.FS
var T *template.Template
var Err error

func init() {
	T, Err = template.ParseFS(Tmpl, "templates.tmpl")
	if Err != nil {
		panic(Err)
	}
}

func Use(wr io.Writer, name string, data any) error {
	err := T.ExecuteTemplate(os.Stdout, name, data)
	if err != nil {
		return err
	}
	return nil
}
