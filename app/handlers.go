package app

import (
	"github.com/Lumexralph/image-uploader/upload"
	"github.com/Lumexralph/image-uploader/template"
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
	html, err := template.Process(template.Tpl, data)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	w.Write(html)
}

// uploadImageHandler will handle storing the image to a temp directory
func uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	file, h, err := r.FormFile("data")
	// close the opened file to free resources
	defer file.Close()

	// create a new image
	img := upload.New(h.Filename, int(h.Size), file)
	// store in a temp directory
	// todo: get the destination directory from environment variable
	fpath, err := img.Store("/tmp")
	if err != nil {
		log.Println(err)
		http.Error(w, "Error occurred while uploading image", 500)
		return
	}
	log.Println(fpath)
	w.Write([]byte("Image upload successful"))
}
