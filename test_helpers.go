/* these are just testing helpers */

package minidb

import (
	"os"
	"testing"

	simplefiletest "github.com/TheBoringDude/simple-filetest"
)

// removes the f which could be the created file or folder
func cleanFileAfter(f string, t *testing.T) {
	if !simplefiletest.Exists(f) {
		t.Fatalf("Path / file : `%s` does not exist!\n", f)
	} else {
		// clean dir
		os.RemoveAll(f)
	}
}

// it asserts the file's content
func checkFileContent(filename, expected string, t *testing.T) {
	ok, err := simplefiletest.SimilarContents(filename, expected)
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("filename's content is not similar to expected")
	}
}
