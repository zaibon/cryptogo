package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Zaibon/cryptogo/web"

	"github.com/Zaibon/cryptogo/core"
)

var (
	conf *core.Config
)

func main() {
	var err error
	conf, err = core.ParseConfig()
	if err != nil {
		log.Fatalln("error parsing config : ", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", web.MakeHandler(web.HomeHandler, conf))
	r.HandleFunc("/transactions/{wallet}", web.MakeHandler(web.TransactionsHandler, conf))

	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("static/js/"))))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("static/css/"))))

	log.Fatalln(http.ListenAndServe(":8080", r))
}
