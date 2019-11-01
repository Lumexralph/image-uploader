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

// Processor adds a generated token from the environment to the template
func Processor(tpl string, d Data) (string, error) {
	var b bytes.Buffer

	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		return "", err
	}

	err = t.Execute(&b, d)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
