// Package main to start the whole project to get token from the
// environment, generate an html template with the data, parse
// the html form on sudmission, upload the file to a directory
// and then save the metadata to the database
package main

import (
	"os"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	checkError(err)

	// generate the data and token
	token := os.Getenv("TOKEN")
	data := TemplateData{"brankas", token}
	tpl := generateTemplate()
	err = templateProcessor(tpl, data)
	checkError(err)
}