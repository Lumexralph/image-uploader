// Package app contains implementation for the server
// and routing patterns to the handlers
// Package app provides the router, middlewares, server and
// handler for the image upload
package app

import (
	"log"
	"net/http"
	"os"
)

// Router creates to handle direction of requests to respective
// handlers
func Router(db *DB) *http.ServeMux {
	// initialize the handler with the db
	fh := &fileHandler{db}

	// create a new multiplexer
	mux := http.NewServeMux()

	mux.HandleFunc("/", templateHandler)
	mux.Handle("/upload", uploadHandlers(fh.uploadImageHandler))

	return mux
}

// Server creates the webserver
func Server(router *http.ServeMux) error {
	port := os.Getenv("PORT")
	log.Printf("Starting server on port:%s... \n", port)
	return http.ListenAndServe(":"+port, router)
}
