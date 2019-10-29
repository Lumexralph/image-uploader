package htmlform

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
					<form action="" method="post">
						<label for="data">Name:</label>
						<input id="data" type="file" name="data">
						<input id="auth" type="text" name="auth" value="{{ .Token }}" hidden>
						<br/>
						<input type="submit" value="Upload">
					</form>
				</body>
				</html>`

// TemplateData will model data to be infused into the template
type TemplateData struct {
	Title, Token string
}

// TemplateProcessor adds a generated token from the environment to the template
func TemplateProcessor(tpl string, data TemplateData) (string, error) {
	var b bytes.Buffer

	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		return "", err
	}

	err = t.Execute(&b, data)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
