package edit

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func EditString(s string) (string, error) {
	return EditStringDefault(s, "")
}

// second string is a path to the default editor
func EditStringDefault(s string, ft string) (string, error) {
	// create temporary file
	f, err := ioutil.TempFile("", "todo-")
	if err != nil {
		return "", fmt.Errorf("edit: EditString: failed to create temporary file: %v", err)
	}

	_, err = f.WriteString(s)
	if err != nil {
		return "", fmt.Errorf("edit: EditString: failed to write to temporary file: %v", err)
	}

	f.Close()
	fname := f.Name()

	// call editor
	cmd := exec.Command(GetEditorDefault(ft), fname)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		return "", fmt.Errorf("edit: EditString: failed to launch editor: %v", err)
	}

	// read file
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return "", fmt.Errorf("edit: EditString: failed to grab contents of tmp file: %v", err)
	}

	// delete file
	if err = os.Remove(fname); err != nil {
		return string(b), fmt.Errorf("edit: EditString: failed to remove tmp file: %v", err)
		// TODO change this later, and is very hacky!
	}

	return string(b), nil
}

func GetEditor() (string, error) {
	s := GetEditorDefault("")
	if s == "" {
		return "", fmt.Errorf("edit: GetEditor: no editor found.")
	}
	return s, nil
}

func GetEditorDefault(ft string) string {
	ed := os.Getenv("VISUAL")
	if ed == "" {
		ed = os.Getenv("EDITOR")

		// Fail through to vim if there is no editor specified
		if ed == "" {
			ed = ft
		}
	}

	return ed
}
