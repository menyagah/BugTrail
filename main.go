package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from bugtrail"))
}

func bugtrailView(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("This is a view"))
}

func bugtrailCreate(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("This is the creation bit"))
}



func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/bugtrail/view", bugtrailView)
	mux.HandleFunc("/bugtrail/create", bugtrailCreate)

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

