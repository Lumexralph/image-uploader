// Package template handles the generation of a processed html form
package template

import (
	"bytes"
	"html/template"
)

// Tpl is the html template that will be served to client
const Tpl = `
			<!DOCTYPE html>
			<html>
				<head>
					<meta charset="UTF-8">
					<title>{{.Title}}</title>
				</head>
				<body>
					<form enctype="multipart/form-data" action="/upload" method="post">
						<label for="data">Upload an Image:</label>
						<br />
						<input id="data" type="file" name="data" accept=".jpg, .jpeg, .png, .gif">
						<input id="auth" type="hidden" name="auth" value="{{ .Token }}">
						<br/>
						<input type="submit" value="Upload">
					</form>
				</body>
				</html>`

// Data will model data to be infused into the template
type Data struct {
	Title, Token string
}

// Process adds a generated token from the environment to the template
func Process(tpl string, d Data) ([]byte, error) {
	var b bytes.Buffer

	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		return nil, err
	}

	err = t.Execute(&b, d)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
