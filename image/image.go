// Package image contains the implementation needed
// to create an image data depending on the type png, jpg or gif
// it creates the image and appends it to a tmp directory
package image

import (
	"path/filepath"
	"io"
	"log"
	"os"
)

var BUFFERSIZE = 1024

// Data to model the image data to be uploaded
type Data struct {
	name, contentType string
	size int
	file *os.File
}

// New method will create a new image data
func New(name, contentType string, size int, f *os.File) *Data {
	return &Data{
		name,
		contentType,
		size,
		f,
	}
}

// Store method will create the image file and store it in a tmp directory
func (img *Data) Store(imageDir string) (string, error) {
	// create a destination file
	dstFile := filepath.Join(imageDir, img.name + "." + img.contentType)
	dst, err := os.Create(dstFile)
	if err != nil {
		return dstFile, err
	}
	defer dst.Close()

	// if we get an image lesser than the buffer size to avoid
	// wasteful memory space
	if img.size < BUFFERSIZE {
		BUFFERSIZE = img.size
	}
	buf := make([]byte, BUFFERSIZE)

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
