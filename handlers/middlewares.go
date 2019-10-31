package handlers

import (
	"os"
	"image"
	"log"
	"net/http"
)

// check the content type
func contentTypeHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		_, ok := r.Header["Content-Type"]
		if !ok {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// check if it is an image
func imageContentHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		r.ParseForm()
		file, _, err := r.FormFile("data")
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(400), 400)
			return
		}

		// verify the image
		_, _, err = image.Decode(file)
		if err != nil {
			log.Printf("could not decode the file into an image")
			http.Error(w, "could not decode body image", 400)
			return
		}
	}

	return http.HandlerFunc(fn)
}

// parse body of the request interface
// parse the html with auth field, check that token are the same
func formParser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		r.ParseForm()
		//Call to ParseForm makes form fields available.
		err := r.ParseForm()
		if err != nil {
			log.Println("Error parsing the form", err)
			http.Error(w, http.StatusText(500), 500)
		}

		token := r.PostFormValue("auth")
		envToken := os.Getenv("TOKEN")
		if token != envToken {
			http.Error(w, http.StatusText(403), 403)
			return
		}
	}

	return http.HandlerFunc(fn)
}

// check the file size
func fileSizeHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		file, _, err := r.FormFile("data")
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(400), 400)
			return
		}

		// verify the size of the file is not above 8MB
		size, err := file.Seek(0, 0)
		sizeLimit := int64(8 << 20)
		if err != nil {
			log.Fatalln("Error getting the image size", err)
		}
		if size > sizeLimit {
			http.Error(w, "Image size should not be above 8MB", 403)
			return
		}
	}

	return http.HandlerFunc(fn)
}
