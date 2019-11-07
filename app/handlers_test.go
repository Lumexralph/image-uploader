package app

import (
	"log"
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestTemplateHandlerWorks(t *testing.T) {
	// create a request to pass to the handler
	req := httptest.NewRequest("GET", "/", nil)

	// create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(templateHandler)

	// make a request to the handler calling the ServeHTTP method
	handler.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("templateHandler(%v, %+v) for url /; returned wrong status code: got %v want %v", rr, req, status, http.StatusOK)
	}
}

// newImageUploadRequest creates a request with the provided uri and parameters
func newImageUploadRequest(uri, formField, fileName, fileExtension string) (r *http.Request, retErr error) {
	// create a buffer to write the create image to
	buf := new(bytes.Buffer)
	// create a multipart form data to be attached to the request
	w := multipart.NewWriter(buf)
	defer func () {
		if err := w.Close(); err != nil {
			log.Printf("Error occurred closing the multipart file: %v", err)
		}
	}()

	// create the data form field.
	formData, err := w.CreateFormFile(formField, fileName+fileExtension)
	if err != nil {
		retErr = err
	}
	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, 256, 256))
	switch fileExtension {
	case "png":
		if err := png.Encode(formData, img); err != nil {
			retErr = err
		}
	case "jpg":
		if err := jpeg.Encode(formData, img, nil); err != nil {
			retErr = err
		}
	case "gif":
		if err := gif.Encode(formData, img, nil); err != nil {
			retErr = err
		}
	}

	r = httptest.NewRequest("POST", uri, buf)
	r.Header.Add("Content-Type", w.FormDataContentType())
	return r, retErr
}

func TestUploadImageHandlerWorks(t *testing.T) {
	cases := []string{".png", ".jpg", ".gif"}
	fh := &fileHandler{db: &mockDB{}}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Upload %s image", tc), func(t *testing.T) {
			r, err := newImageUploadRequest("/upload", "data", "test-sample", tc)
			if err != nil {
				t.Fatalf("newImageUploadRequest(%s, %s, %s, %s) failed to create form request, ", "/upload", "data", "test-sample", tc)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(fh.uploadImageHandler)
			handler.ServeHTTP(rr, r)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("uploadImageHandler(%v, %+v) for POST /upload; returned wrong status code: got %v want %v", rr, r, status, http.StatusOK)
			}
		})
	}
}
