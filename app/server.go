// Package app contains implementation for the server
// and routing patterns to the handlers
package app

import (
	"github.com/justinas/alice"
	"log"
	"net/http"
	"os"
)

// router creates to handle direction of requests to respective
// handlers
func router() *http.ServeMux {
	// create a new multiplexer
	mux := http.NewServeMux()
	uploadHandlers := alice.New(fileTypeHandler, checkAuthTokenHandler, fileSizeHandler, imageContentHandler)
	mux.HandleFunc("/", templateHandler)
	mux.Handle("/upload", uploadHandlers.ThenFunc(uploadImageHandler))
	return mux
}

// Server creates the webserver
func Server() error {
	port := os.Getenv("PORT")
	log.Printf("Starting server on port:%s... \n", port)
	err := http.ListenAndServe(":"+port, router())
	return err
}