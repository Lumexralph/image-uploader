package main

import (
	"log"
	"html/template"
	"os"
)

// TemplateData will model data to be infused into the template
type TemplateData struct {
	Title, Token string
}

// generateTemplate function will generate the template needed
// for user to upload image
func generateTemplate() string {
	const tpl = `
				<!DOCTYPE html>
				<html>
					<head>
						<meta charset="UTF-8">
						<title>{{.Title}}</title>
					</head>
					<body>
						<form action="" method="post">
							<label for="data">Name:</label>
							<input id="data" type="file" name="data">
							<input id="auth" type="text" name="auth" value="{{ .Token }}" hidden>
							<br/>
							<input type="submit" value="Upload">
						</form>
					</body>
				</html>`

	return tpl
}

// templateProcessor adds a generated token from the environment to the template
func templateProcessor(tpl string, data TemplateData) error {
	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		return err
	}
	err = t.Execute(os.Stdout, data)

	return nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
