// Package image contains the implementation needed
// to create an image data depending on the type png, jpg or gif
// it creates the image and appends it to a tmp directory
package image

import (
	"mime/multipart"
	"path/filepath"
	"io"
	"log"
	"os"
)

var bufferSize = 1024

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
func (img *Data) Store(imageDir string) (string, error) {
	// create a destination file
	dstFile := filepath.Join(imageDir, img.name)
	dst, err := os.Create(dstFile)
	if err != nil {
		return dstFile, err
	}
	defer img.file.Close()
	defer dst.Close()
	// if we get an image lesser than the buffer size to avoid
	// wasteful memory space
	if img.size < bufferSize {
		bufferSize = img.size
	}
	buf := make([]byte, bufferSize)

	for {
		// read the content of the file sequentially by bufferring
		n, err := img.file.Read(buf)
		if err != nil && err != io.EOF {
			return dstFile, err
		}
		// if it gets to the end of the file
		if err == io.EOF {
			break
		}

		if _, err := dst.Write(buf[:n]); err != nil {
			return dstFile, err
		}
	}
	log.Printf("Image %s successfully created in the directory %s", img.name, imageDir)
	return dstFile, nil
}
