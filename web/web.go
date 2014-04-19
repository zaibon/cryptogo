package web

import (
	"net/http"
	"text/template"

	"github.com/Zaibon/cryptogo/core"
)

type webHandler func(http.ResponseWriter, *http.Request, *core.Config)

var (
	Templates map[string]*template.Template
)

func init() {
	if err := BuildTemplate("templates"); err != nil {
		panic(err)
	}
}

func MakeHandler(fn webHandler, conf *core.Config) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		fn(rw, req, conf)
	}
}
