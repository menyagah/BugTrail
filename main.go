package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from bugtrail"))
}

func bugtrailView(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Display a specific snippet with ID %d ", id)
	w.Write([]byte(msg))
}

func bugtrailCreate(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("This is the creation bit"))
}





func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/bugtrail/view/{id}", bugtrailView)
	mux.HandleFunc("/bugtrail/create", bugtrailCreate)
	

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

