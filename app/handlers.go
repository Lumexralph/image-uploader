package app

import (
	handlerchain "github.com/justinas/alice"
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
	data := &template.Data{
		Title: "brankas",
		Token: token,
	}
	err := template.Process(template.FormPage, data, w)
	if err != nil {
		log.Printf("Error occurred process the template: %v", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}

// uploadImageHandler will handle storing the image to a temp directory
func uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	file, h, err := r.FormFile("data")
	if err != nil {
		log.Printf("Error occurred getting file from form: %v", err)
	}
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

func uploadHandlers(handler func(http.ResponseWriter, *http.Request)) http. Handler {
	// chain all handlers together from left to right
	uploadHandlers := handlerchain.New(parsePOSTHandler, fileTypeHandler, checkAuthTokenHandler, fileSizeHandler, imageContentHandler)

	return uploadHandlers.ThenFunc(handler)
}