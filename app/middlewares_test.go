package app

import (
	handlerchain "github.com/justinas/alice"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockDB struct{}

func (mdb *mockDB) CreateFileMetaData(table string, fd *fileData) { }

func TestParsePOSTHandlerWithWrongMethod(t *testing.T) {
	fh := &fileHandler{db: &mockDB{}}
	uploadHandlers := handlerchain.New(parsePOSTHandler)
	// make request with wrong method
	r := httptest.NewRequest("GET", "/upload", nil)
	rr := httptest.NewRecorder()
	handler := uploadHandlers.ThenFunc(fh.uploadImageHandler)
	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("parsePOSTHandler(%v, %+v) for GET /upload; returned wrong status code: got %v want %v", rr, r, status, http.StatusBadRequest)
	}
}

func TestParsePOSTHandlerWithWorks(t *testing.T) {
	fh := &fileHandler{db: &mockDB{}}

	uploadHandlers := handlerchain.New(parsePOSTHandler)

	r, err := newImageUploadRequest("/upload", "data", "test-sample", ".png")
	if err != nil {
		t.Fatalf("newImageUploadRequest(%s, %s, %s, %s) failed to create form request", "/upload", "data", "test-sample", ".png")
	}
	rr := httptest.NewRecorder()
	handler := uploadHandlers.ThenFunc(fh.uploadImageHandler)
	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("parsePOSTHandler(%v, %+v) for POST /upload; returned wrong status code: got %v want %v", rr, r, status, http.StatusOK)
	}
}

func TestFileTypeHandlerWorks(t *testing.T) {
	fh := &fileHandler{db: &mockDB{}}

	cases := []string{".png", ".jpg", ".gif"}
	uploadHandlers := handlerchain.New(fileTypeHandler)

	for _, tc := range cases {
		t.Run("Upload %s file", func(t *testing.T) {
			r, err := newImageUploadRequest("/upload", "data", "test-sample", tc)
			if err != nil {
				t.Fatalf("newImageUploadRequest(%s, %s, %s, %s) failed to create form request", "/upload", "data", "test-sample", tc)
			}
			rr := httptest.NewRecorder()
			handler := uploadHandlers.ThenFunc(fh.uploadImageHandler)
			handler.ServeHTTP(rr, r)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("fileTypeHandler(%v, %+v) for POST /upload; returned wrong status code: got %v want %v", rr, r, status, http.StatusOK)
			}
		})
	}
}

func TestFileTypeHandlerWithWrongImageFormat(t *testing.T) {
	fh := &fileHandler{db: &mockDB{}}

	cases := []string{".pdf", ".csv", ".mp4"}
	uploadHandlers := handlerchain.New(fileTypeHandler)

	for _, tc := range cases {
		t.Run("Upload %s file", func(t *testing.T) {
			r, err := newImageUploadRequest("/upload", "data", "test-sample", tc)
			if err != nil {
				t.Fatalf("newImageUploadRequest(%s, %s, %s, %s) failed to create form request", "/upload", "data", "test-sample", tc)
			}
			rr := httptest.NewRecorder()
			handler := uploadHandlers.ThenFunc(fh.uploadImageHandler)
			handler.ServeHTTP(rr, r)

			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("fileTypeHandler(%v, %+v) for POST /upload; returned wrong status code: got %v want %v", rr, r, status, http.StatusBadRequest)
			}
		})
	}
}

func TestCheckAuthTokenHandlerWorks(t *testing.T) {
	fh := &fileHandler{db: &mockDB{}}
	uploadHandlers := handlerchain.New(checkAuthTokenHandler)

	r, err := newImageUploadRequest("/upload", "data", "test-sample", ".png")
	if err != nil {
		t.Fatalf("newImageUploadRequest(%s, %s, %s, %s) failed to create form request", "/upload", "data", "test-sample", ".png")
	}
	rr := httptest.NewRecorder()

	handler := uploadHandlers.ThenFunc(fh.uploadImageHandler)
	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("checkAuthTokenHandler(%v, %+v) for POST /upload; returned wrong status code: got %v want %v", rr, r, status, http.StatusOK)
	}
}

func TestCheckAuthTokenHandlerWithWrongToken(t *testing.T) {
	fh := &fileHandler{db: &mockDB{}}
	uploadHandlers := handlerchain.New(checkAuthTokenHandler)

	r, err := newImageUploadRequest("/upload", "data", "test-sample", ".png")
	if err != nil {
		t.Fatalf("newImageUploadRequest(%s, %s, %s, %s) failed to create form request", "/upload", "data", "test-sample", ".png")
	}
	rr := httptest.NewRecorder()
	r.ParseForm()
	r.PostForm["auth"] = []string{"12345"}
	handler := uploadHandlers.ThenFunc(fh.uploadImageHandler)
	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("checkAuthTokenHandler(%v, %+v) for POST /upload; returned wrong status code: got %v want %v", rr, r, status, http.StatusForbidden)
	}
}

func TestFileSizeHandlerWithLessThan8MBFile(t *testing.T) {
	fh := &fileHandler{db: &mockDB{}}
	uploadHandlers := handlerchain.New(fileSizeHandler)

	r, err := newImageUploadRequest("/upload", "data", "test-sample", ".gif")
	if err != nil {
		t.Fatalf("newImageUploadRequest(%s, %s, %s, %s) failed to create form request", "/upload", "data", "test-sample", ".gif")
	}

	rr := httptest.NewRecorder()
	handler := uploadHandlers.ThenFunc(fh.uploadImageHandler)
	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("fileSizeHandler(%v, %+v) for POST /upload; returned wrong status code: got %v want %v", rr, r, status, http.StatusOK)
	}
}

func TestFileSizeHandlerWithOver8MBFile(t *testing.T) {
	fh := &fileHandler{db: &mockDB{}}
	uploadHandlers := handlerchain.New(fileSizeHandler)

	r, err := newImageUploadRequest("/upload", "data", "test-sample", ".gif")
	if err != nil {
		t.Fatalf("newImageUploadRequest(%s, %s, %s, %s) failed to create form request", "/upload", "data", "test-sample", ".gif")
	}
	_, header, _ := r.FormFile("data")
	header.Size = int64(9 << 20)

	rr := httptest.NewRecorder()
	handler := uploadHandlers.ThenFunc(fh.uploadImageHandler)
	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("fileSizeHandler(%v, %+v) for POST /upload; returned wrong status code: got %v want %v", rr, r, status, http.StatusForbidden)
	}
}

func TestFileImageContentHandlerWithNonImage(t *testing.T) {
	fh := &fileHandler{db: &mockDB{}}
	uploadHandlers := handlerchain.New(imageContentHandler)

	r, err := newImageUploadRequest("/upload", "data", "test-sample", ".gi")
	if err != nil {
		t.Fatalf("newImageUploadRequest(%s, %s, %s, %s) failed to create form request", "/upload", "data", "test-sample", ".gi")
	}

	rr := httptest.NewRecorder()
	handler := uploadHandlers.ThenFunc(fh.uploadImageHandler)
	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("fileSizeHandler(%v, %+v) for POST /upload; returned wrong status code: got %v want %v", rr, r, status, http.StatusBadRequest)
	}
}
