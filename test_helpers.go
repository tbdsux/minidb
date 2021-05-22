/* these are just testing helpers */

package minidb

import (
	"io/ioutil"
	"os"
	"testing"
)

// removes the f which could be the created file or folder
func cleanFileAfter(f string, t *testing.T) {
	if !isPathExists(f) {
		t.Fatal("Default json db file: `__default.json` does not exist!")
	} else {
		// clean dir
		os.RemoveAll(f)
	}
}

// it asserts the file's content
func checkFileContent(filename, expected string, t *testing.T) {
	if content, err := ioutil.ReadFile(filename); err != nil {
		t.Fatalf("error trying to read ->  %s", filename)
	} else {
		if string(content) != expected {
			t.Fatalf("filename: %s => (content)`%s` is not similar to (expected)`%s`", filename, string(content), expected)
		}
	}
}
