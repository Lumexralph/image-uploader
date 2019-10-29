// Test that the template is generated and token is added to the
// template on generation
package main

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
	"testing"
)

const wantTemplate = `
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

func TestGenerateTemplate(t *testing.T) {
	cases := []struct {
		substr string
		want   bool
	}{
		{`<input type="submit" value="Upload">`, true},
		{`<form action="" method="post">`, true},
		{`<!DOCTYPE html>`, true},
		{`<label for="data">Name:</label>`, true},
		{`<input id="data" type="file" name="data">`, true},
		{`<input id="auth" type="text" name="auth" value="{{ .Token }}" hidden>`, true},
	}
	gotTemplate := generateTemplate()

	for _, tc := range cases {
		t.Run(tc.substr, func(t *testing.T) {
			if got := strings.Contains(gotTemplate, tc.substr); got != tc.want {
				t.Fatalf("generateTemplate() got %+v; want %+v", gotTemplate, wantTemplate)
			}
		})
	}
}

func TestTemplateProcessor(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal("Could not load env. file \n")
	}

	token := os.Getenv("TEST_TOKEN")
	cases := []TemplateData{TemplateData{"image upload", ""}, TemplateData{"brankas", token}}

	for _, tc := range cases {
		t.Run(tc.Title, func(t *testing.T){
			if err := templateProcessor(wantTemplate, tc); err != nil {
				t.Fatalf("templateProcessor() got %+v; could not process template", err)
			}
		})
	}
}
