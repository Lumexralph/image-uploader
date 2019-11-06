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

// router creates to handle direction of requests to respective
// handlers
func router() *http.ServeMux {
	// create a new multiplexer
	mux := http.NewServeMux()

	mux.HandleFunc("/", templateHandler)
	mux.Handle("/upload", uploadHandlers(uploadImageHandler))

	return mux
}

// Server creates the webserver
func Server() error {
	port := os.Getenv("PORT")
	log.Printf("Starting server on port:%s... \n", port)
	return http.ListenAndServe(":"+port, router())
}
