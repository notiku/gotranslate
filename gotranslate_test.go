package gotranslate_test

import (
	"testing"

	"github.com/notiku/gotranslate"
)

func TestTranslate(t *testing.T) {
	// Call the Translate function
	source := "Hello, World!"
	from := "en"
	to := "es"
	resp, err := gotranslate.Translate(source, from, to)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check the results
	expectedTranslated := "¡Hola Mundo!"
	expectedFrom := "en"
	if resp.Translated != expectedTranslated {
		t.Errorf("expected translated text %q, got %q", expectedTranslated, resp.Translated)
	}
	if resp.From != expectedFrom {
		t.Errorf("expected from language %q, got %q", expectedFrom, resp.From)
	}
	if resp.To != to {
		t.Errorf("expected to language %q, got %q", to, resp.To)
	}
}
