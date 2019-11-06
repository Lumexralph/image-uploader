// Package upload contains the implementation needed
// to create an image data depending on the type png, jpg or gif
// it creates the image and appends it to a tmp directory
package upload

import (
	"mime/multipart"
	"path/filepath"
	"io"
	"log"
	"os"
)

// Data to model the image data to be uploaded
type Data struct {
	name string
	size int
	file multipart.File
}

// New method will create a new image data
func New(name string, size int, f multipart.File) *Data {
	return &Data{
		name,
		size,
		f,
	}
}

// Store method will create the image file and store it in a tmp directory
func (img *Data) Store(imageDir string) (fpath string, retErr error) {
	// create a destination file
	fpath = filepath.Join(imageDir, img.name)
	dst, err := os.Create(fpath)
	if err != nil {
		return fpath, err
	}
	defer img.file.Close()
	defer func(){
		err := dst.Close()
		if retErr == nil {
		  retErr = err
		}
	  }()

	_, err = io.Copy(dst, img.file)
	if err != nil {
		return fpath, err
	}

	log.Printf("Image %s successfully created in the directory %s", img.name, imageDir)
	return fpath, nil
}
