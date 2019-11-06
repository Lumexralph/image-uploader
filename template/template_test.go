// Test that the template is generated and token is added to the
// template on generation
package template

import (
	"bytes"
	"strings"
	"testing"
)

func TestTemplateProcess(t *testing.T) {
	gotData := new(bytes.Buffer)

	wantData := &Data{"brankas", "12345"}
	err := Process(FormPage, wantData, gotData)

	if err != nil {
		t.Fatalf("TemplateProcessor(%s, %+v) got %+v; could not process template \n", FormPage, wantData, err)
	}

	wantTemplate := `<title>brankas</title>`
	if !strings.Contains(string(gotData.String()), wantTemplate) {
		t.Fatalf("TemplateProcessor(%s, %+v) got %+v; want %+v, could not find title  \n", FormPage, wantData, gotData.String(), wantTemplate)
	}

	wantTemplate = `<input id="auth" type="hidden" name="auth" value="12345">`
	if !strings.Contains(string(gotData.String()), wantTemplate) {
		t.Fatalf("TemplateProcessor(%s, %+v) got %+v; want %+v, could not find title  \n", FormPage, wantData, gotData.String(), wantTemplate)
	}

	wantTemplate = `<form enctype="multipart/form-data" action="/upload" method="post">
						<label for="data">Upload an Image:</label>
						<br />
						<input id="data" type="file" name="data" accept=".jpg, .jpeg, .png, .gif">
						<input id="auth" type="hidden" name="auth" value="12345">
						<br/>
						<input type="submit" value="Upload">
					</form>`
	if !strings.Contains(string(gotData.String()), wantTemplate) {
		t.Fatalf("TemplateProcessor(%s, %+v) got %+v; want %+v, could not find input template with token \n", FormPage, wantData, gotData.String(), wantTemplate)
	}
}
