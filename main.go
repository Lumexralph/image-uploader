// Package main to start the whole project to get token from the
// environment, generate an html template with the data, parse
// the html form on sudmission, upload the file to a directory
// and then save the metadata to the database
package main

import (
	"github.com/Lumexralph/image-uploader/app"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load env. file...")
	}
	// start the server
	err = app.Server()
	if err != nil {
		log.Println("Error occurred starting the server \n", err)
	}
}
