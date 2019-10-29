// Package main to start the whole project to get token from the
// environment, generate an html template with the data, parse
// the html form on sudmission, upload the file to a directory
// and then save the metadata to the database
package main

import (
	"fmt"
	"github.com/Lumexralph/image-uploader/htmlform"
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
	data := htmlform.TemplateData{
		Title: "brankas",
		Token: token,
	}
	tmpl, err := htmlform.TemplateProcessor(htmlform.Tpl, data)
	fmt.Println(tmpl)

}
