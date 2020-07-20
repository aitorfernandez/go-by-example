package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aitorfernandez/go-by-example/rest/mgt"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	p := mgt.NewPost()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})
	api.HandleFunc("/posts/{id}", p.GetPost).Methods(http.MethodGet)
	// and so on...

	log.Fatalln(http.ListenAndServe(":8080", r))
}
