package app

import (
	"image"
	"os"
	"strings"
	"log"
	"net/http"
	// needed to register the image format
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// parsePOSTHandler parses the form attached to the request
func parsePOSTHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		if err := r.ParseForm(); err != nil {
			log.Println("Error parsing the form", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// check the content type
func fileTypeHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		_, header, err := r.FormFile("data")
		if err != nil {
			log.Printf("Error occurred getting file from form: %v", err)
			http.Error(w, http.StatusText(400), 400)
			return
		}
		format := []string{".jpeg", ".jpg", ".png", ".gif"}
		var valid bool
		for _, imgFormat := range format {
			// check if the file is in jpeg, png, jpg or gif format
			if strings.HasSuffix(header.Filename, imgFormat) {
				valid = true
				break
			}
		}
		if !valid {
			http.Error(w, "Please supply image of jpg, jpeg, png or gif format", 400)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// parse body of the request interface
// parse the html with auth field, check that token are the same
func checkAuthTokenHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if token := r.PostFormValue("auth"); token != os.Getenv("TOKEN") {
			http.Error(w, http.StatusText(403), 403)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// check if it is an image
func imageContentHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("data")
		if err != nil {
			log.Printf("Error occurred getting file from form: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		// verify the image
		if _, _, err = image.Decode(file); err != nil {
			log.Println("Decoding the image: ", err.Error())
			// close the opened file to free resources
			file.Close()
			http.Error(w, err.Error(), 400)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// check the file size.
func fileSizeHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		_, header, err := r.FormFile("data")
		if err != nil {
			log.Printf("Error occurred getting file from form: %v", err)
			http.Error(w, http.StatusText(400), 400)
			return
		}
		// verify the size of the file is not above 8MB
		sizeLimit := int64(8 << 20)
		if header.Size > sizeLimit {
			http.Error(w, "Image size should not be above 8MB", 403)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
