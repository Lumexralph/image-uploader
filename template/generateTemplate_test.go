// Test that the template is generated and token is added to the
// template on generation
package template

import (
	"strings"
	"testing"
)

func TestTemplateProcessor(t *testing.T) {
	wantData := Data{"brankas", "12345"}
	gotTemplate, err := Processor(Tpl, wantData)

	if err != nil || gotTemplate == "" {
		t.Fatalf("TemplateProcessor(%s, %+v) got %+v; could not process template \n", Tpl, wantData, err)
	}

	wantTemplate := `<title>brankas</title>`
	if !strings.Contains(gotTemplate, wantTemplate) {
		t.Fatalf("TemplateProcessor(%s, %+v) got %+v; want %+v, could not find title  \n", Tpl, wantData, gotTemplate, wantTemplate)
	}

	wantTemplate = `<form action="" method="post">
						<label for="data">Upload an Image:</label>
						<input id="data" type="file" name="data" accept=".jpg, .jpeg, .png, .gif">
						<input id="auth" type="text" name="auth" value="12345" hidden>
						<br/>
						<input type="submit" value="Upload">
					</form>`
	if !strings.Contains(gotTemplate, wantTemplate) {
		t.Fatalf("TemplateProcessor(%s, %+v) got %+v; want %+v, could not find input template with token \n", Tpl, wantData, gotTemplate, wantTemplate)
	}

}