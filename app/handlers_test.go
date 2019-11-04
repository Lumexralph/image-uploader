package app

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
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

// createUploadRequest creates a request with the provided uri and parameters
func newImageUploadRequest(uri, formField, fileName, fileExtension string) (r *http.Request, retErr error) {
	// Set up a pipe to avoid buffering.
	pr, pw := io.Pipe()
	// the writer will transform whatever to
	// multipart form data and pass it to the pipe.
	w := multipart.NewWriter(pw)
	go func() {
		defer w.Close()
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
	}()

	r = httptest.NewRequest("POST", uri, pr)
	r.Header.Add("Content-Type", w.FormDataContentType())
	return r, retErr
}

func TestUploadImageHandlerWorks(t *testing.T) {
	cases := []string{".png", ".jpg", ".gif"}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Upload %s image", tc), func(t *testing.T) {
			r, err := newImageUploadRequest("/upload", "data", "test-sample", tc)
			if err != nil {
				t.Error(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(uploadImageHandler)
			handler.ServeHTTP(rr, r)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("uploadImageHandler(%v, %+v) for POST /upload; returned wrong status code: got %v want %v", rr, r, status, http.StatusOK)
			}
		})
	}
}
