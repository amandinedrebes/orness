package api

import (
	"log"
	"net/http"

	"github.com/amandinedrebes/orness/internal/core"
	"github.com/gorilla/mux"
)

//InitAPI start API connection
func InitAPI(ip, port string) {
	apiF := API{
		APIIP:    ip,
		APIPort:  port,
		Database: core.NewDatabase(),
	}
	swagger(&apiF)
}

func swagger(r *API) {
	router := mux.NewRouter()

	//display swaggerUI path specification
	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./internal/website/swaggerui/")))
	router.PathPrefix("/swaggerui/").Handler(sh)

	router.HandleFunc(URLNotes, func(w http.ResponseWriter, req *http.Request) {
		readNotes(r, w, req)
	}).Methods(http.MethodGet)

	router.HandleFunc(URLNotes, func(w http.ResponseWriter, req *http.Request) {
		addNotes(r, w, req)
	}).Methods(http.MethodPost)

	functions := []string{URLNotes}
	for _, f := range functions {
		router.HandleFunc(f, func(w http.ResponseWriter, req *http.Request) {
			r.SetDefaultHeader(w, req)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Connection", "close")
		}).Methods(http.MethodOptions)
	}

	router.Use(mux.CORSMethodMiddleware(router))

	log.Fatal(http.ListenAndServe(r.APIIP+":"+r.APIPort, router))
}
