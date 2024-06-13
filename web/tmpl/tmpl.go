package tmpl

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/raffleberry/sqlsheet/pkg/utils"
)

//go:embed *.html
var fs embed.FS
var T *template.Template
var Err error

func init() {
	T, Err = template.ParseFS(fs, "*.html")
	utils.Panic(Err)
}

func Use(wr io.Writer, name string, data any) error {
	var err error
	if utils.DEV {
		wd, er := os.Getwd()
		utils.Panic(er)

		t, er := template.ParseGlob(filepath.Join(wd, "web", "tmpl", "*.html"))
		utils.Panic(er)
		err = t.ExecuteTemplate(wr, name, data)

	} else {
		err = T.ExecuteTemplate(wr, name, data)
	}
	if err != nil {
		fmt.Fprint(wr, err.Error())
		return err
	}
	return nil
}
