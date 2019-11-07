// Package main to start the whole project to get token from the
// environment, generate an html template with the data, parse
// the html form on sudmission, upload the file to a directory
// and then save the metadata to the database
package main

import (
	"fmt"
	"github.com/Lumexralph/image-uploader/app"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load env. file...")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	// create database url
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	// create db connection
	db, err := app.NewDB("postgres", connStr)
	if err != nil {
		log.Panic("Error creating database connection", err)
	}
	defer db.Close()

	// create router with the db
	mux := app.Router(db)

	// start the server passing the router
	err = app.Server(mux)
	if err != nil {
		log.Println("Error occurred starting the server \n", err)
	}
}
