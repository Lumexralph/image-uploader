package app

import (
	"github.com/Lumexralph/image-uploader/template"
	"github.com/Lumexralph/image-uploader/upload"
	handlerchain "github.com/justinas/alice"
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
		log.Printf("Error occurred processing the template: %v", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}

// fileHandler will handle operations pertaining to a file
type fileHandler struct {
	db databaseHandler
}

// fileData to hold the metadata information to be saved to the database
type fileData struct {
	name, slug, format, path string
	size int64
}

// uploadImageHandler will handle storing the image to a temp directory
func (fh *fileHandler) uploadImageHandler(w http.ResponseWriter, r *http.Request) {
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
	// create the metadata
	format := h.Filename[(len(h.Filename) - 4):]
	name := h.Filename[:len(h.Filename) - 4]
	slug := name + "-" + randomStringGenerator()
	fd := &fileData{
		name: name,
		slug: slug,
		format: format,
		path: fpath,
		size: h.Size,
	}
	// persist the metadata to the database
	fh.db.CreateFileMetaData("file_metadata", fd)
	w.Write([]byte("Image upload successful"))
}

func uploadHandlers(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	// chain all handlers together from left to right
	uploadHandlers := handlerchain.New(parsePOSTHandler, fileTypeHandler, checkAuthTokenHandler, fileSizeHandler, imageContentHandler)

	return uploadHandlers.ThenFunc(handler)
}
