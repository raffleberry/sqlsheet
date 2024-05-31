package tmpl

import (
	"embed"
	"html/template"
	"io"

	"github.com/raffleberry/sqlsheet/pkg/utils"
)

//go:embed *.html
var Tmpl embed.FS
var T *template.Template
var Err error

func init() {
	T, Err = template.ParseFS(Tmpl, "*.html")
	utils.Panic(Err)
}

func Use(wr io.Writer, name string, data any) error {
	err := T.ExecuteTemplate(wr, name, data)
	if err != nil {
		return err
	}
	return nil
}
