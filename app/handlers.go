package app

import (
	imageUploader "github.com/Lumexralph/image-uploader/image"
	"github.com/Lumexralph/image-uploader/template"
	"image"
	"strings"
	// needed to register the image format
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
)

// templateHandler to supply template to client
func templateHandler(w http.ResponseWriter, r *http.Request) {
	// generate the data and token
	token := os.Getenv("TOKEN")
	data := template.Data{
		Title: "brankas",
		Token: token,
	}
	tmpl, err := template.Processor(template.Tpl, data)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(tmpl))
}

// check the content type
func fileTypeHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var valid bool

		if r.Method != "POST" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		r.ParseForm()
		_, header, _ := r.FormFile("data")
		format := []string{".jpeg", ".jpg", ".png", ".gif"}
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
		if r.Method != "POST" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		//Call to ParseForm makes form fields available.
		err := r.ParseForm()
		if err != nil {
			log.Println("Error parsing the form", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		token := r.PostFormValue("auth")
		log.Println(token)
		envToken := os.Getenv("TOKEN")
		if token != envToken {
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
		if r.Method != "POST" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		r.ParseForm()
		file, _, err := r.FormFile("data")
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		// verify the image
		log.Println(file)
		_, _, err = image.Decode(file)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 400)
			return
		}
		next.ServeHTTP(w, r)
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
		_, header, err := r.FormFile("data")
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(400), 400)
			return
		}
		// verify the size of the file is not above 8MB
		sizeLimit := int64(8 << 20)
		size := header.Size
		if err != nil {
			log.Println("Error getting the image size", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		if size > sizeLimit {
			http.Error(w, "Image size should not be above 8MB", 403)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// uploadImageHandler will handle storing the image to a temp directory
func uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	f, h, err := r.FormFile("data")
	// close the opened file to free resources
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	// create a new image
	img := imageUploader.New(h.Filename, int(h.Size), f)

	// store in a temp directory
	// todo: get the destination directory from environment variable
	fpath, err := img.Store("/tmp")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(fpath)
	w.Write([]byte("Image upload successful"))
}
