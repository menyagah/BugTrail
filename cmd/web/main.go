package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)


func main(){

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /bugtrail/view/{id}", bugtrailView)
	mux.HandleFunc("GET /bugtrail/create", bugtrailCreate)
	mux.HandleFunc("POST /bugtrail/create", bugtrailCreatePost)
	
	logger.Info("starting server", slog.String("addr", *addr))

	
	err := http.ListenAndServe(*addr, mux)

	logger.Error(err.Error())
	os.Exit(1)
}

