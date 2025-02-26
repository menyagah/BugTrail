package main

import (
	"log"
	"net/http"
)


func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /bugtrail/view/{id}", bugtrailView)
	mux.HandleFunc("GET /bugtrail/create", bugtrailCreate)
	mux.HandleFunc("POST /bugtrail/create", bugtrailCreatePost)
	

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

