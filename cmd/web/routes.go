package main

import "net/http"


func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /bugtrail/view/{id}", app.bugtrailView)
	mux.HandleFunc("GET /bugtrail/create", app.bugtrailCreate)
	mux.HandleFunc("POST /bugtrail/create", app.bugtrailCreatePost)

	return mux
}