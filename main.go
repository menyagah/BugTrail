package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from bugtrail"))
}

func bugtrailView(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d ", id)
	// msg := fmt.Sprintf("Display a specific snippet with ID %d ", id)
	// w.Write([]byte(msg))
}

func bugtrailCreate(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("This is the creation bit"))
}

func bugtrailCreatePost(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new bugtrail"))
}





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

