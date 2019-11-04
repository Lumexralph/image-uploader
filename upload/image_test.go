package upload

import (
	"fmt"
	"os"
	"testing"
)

func TestImageCreation(t *testing.T) {
	// setup
	const imgDir = "testdir"
	// - create the image storage temporary directory
	err := os.Mkdir(imgDir, 0777)
	if err != nil {
		t.Fatal(err)
	}
	// run assertions
	cases := []struct {
		name, want string
		size            int
	}{
		{"sample.png", "testdir/sample.png", 100},
		{"sample2.jpg", "testdir/sample2.jpg", 1000},
		{"sample3.jpeg", "testdir/sample3.jpeg", 1000},
		{"sample4.gif", "testdir/sample4.gif", 1000},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Create file %s", tc.name), func(t *testing.T) {
			// open the image
			f, err := os.Open("../example.png")
			if err != nil {
				t.Fatal(err)
			}
			img := New(tc.name, tc.size, f)
			got, err := img.Store(imgDir)

			if err != nil {
				t.Errorf("img.Store(%s) could not store image", imgDir)
			}

			if got != tc.want {
				t.Errorf("img.Store(%s) got %s; want %s", imgDir, got, tc.want)
			}
		})
	}

	// teardown destroy the directory with the files
	err = os.RemoveAll(imgDir)
	if err != nil {
		t.Fatal(err)
	}
}
