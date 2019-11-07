# Image Uploader

Image Uploader is an application that handles upload of an image from user/client, store the image in a file on disk and persist the metadata of the image in the database

## Installing

```
    $ git clone https://github.com/Lumexralph/image-uploader.git
    $ cd image-uploader
```


* Create a `.env` file and copy/paste the environment variables from the `.env.sample` file that's already existent in the root project directory.

* Create a postgreSQL database using the variables set in your `.env` file with the SQL script in the `schema.txt` file

## Running the application

Run the command below to run the application locally.
```sh
  $ go build
  or
  $ go run .
```

-Please open your browser and put the URI `localhost:PORT` PORT being the one you set yourself in the .env file

## Running the tests

Run the command below to run the tests for the application in the packages.

```sh
  $ go test
```

## Built With

The project has been built with the following technologies:

* Go
* Some Go packages which can be refrenced in the `go.sum` file
* PostgresQL

## Copyright

Copyright (c) 2019 LumexRalph. Released under the [Apache License](https://github.com/Lumexralph/image-uploader/blob/master/LICENSE).
