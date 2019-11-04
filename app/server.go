// Package app contains implementation for the server
// and routing patterns to the handlers
// Package app provides the router, middlewares, server and
// handler for the image upload
package app

import (
	handlerchain "github.com/justinas/alice"
	"log"
	"net/http"
	"os"
)

// router creates to handle direction of requests to respective
// handlers
func router() *http.ServeMux {
	// create a new multiplexer
	mux := http.NewServeMux()
	// chain all handlers together from left to right
	uploadHandlers := handlerchain.New(parsePOSTHandler, fileTypeHandler, checkAuthTokenHandler, fileSizeHandler, imageContentHandler)

	mux.HandleFunc("/", templateHandler)
	mux.Handle("/upload", uploadHandlers.ThenFunc(uploadImageHandler))

	return mux
}

// Server creates the webserver
func Server() error {
	port := os.Getenv("PORT")
	log.Printf("Starting server on port:%s... \n", port)
	return http.ListenAndServe(":"+port, router())
}
