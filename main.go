// Package main to start the whole project to get token from the
// environment, generate an html template with the data, parse
// the html form on sudmission, upload the file to a directory
// and then save the metadata to the database
package main

import (
	"github.com/Lumexralph/image-uploader/template"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load env. file...")
	}

	// generate the data and token
	token := os.Getenv("TOKEN")
	data := template.Data{
		Title: "brankas",
		Token: token,
	}
	tmpl, err := template.Processor(template.Tpl, data)
	fmt.Println(tmpl)
package main

import (
	"log"
	"os"
	"github.com/Lumexralph/image-uploader/image"
)


func main() {
	// open the image
	f, err := os.Open("example.png")
	if err != nil {
		log.Println(err)
	}

	// create a new image
	img := image.New("sample2", "png", 100, f)
	// store in a temp directory
	fpath, err := img.Store()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(fpath)

}
