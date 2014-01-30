package edit

import (
	"testing"
)

func TestGetEditorDefault(t *testing.T) {
	s := GetEditorDefault("something")
	if s != "/usr/bin/vim" {
		t.Errorf("Expected \"/usr/bin/vim\", got \"%s\".", s)
	}
}

func TestGetEditor(t *testing.T) {
	// TODO
}

/*
func TestEditStringDefault(t *testing.T) {
	s, err := EditStringDefault("String", "/usr/bin/nano")
	if err != nil {
		// TODO
	}

	if s != "New String\n" {
		t.Errorf("Expected \"New String\\n\", go \t \"%s\" (did you mistype?).", s)
	}
}*/
