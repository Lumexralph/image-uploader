// Test that the template is generated and token is added to the
// template on generation
package htmlform

import (
	"testing"
)

func TestTemplateProcessor(t *testing.T) {
	cases := []TemplateData{TemplateData{"image upload", "abcde"}, TemplateData{"brankas", "12345"}}

	for _, tc := range cases {
		t.Run(tc.Title, func(t *testing.T){
			if _, err := TemplateProcessor(Tpl, tc); err != nil {
				t.Fatalf("templateProcessor() got %+v; could not process template", err)
			}
		})
	}
}
