package minidb

import (
	"io/ioutil"
	"testing"
)

func TestCollections_FileContent(t *testing.T) {
	filename := "content.json"
	NewMiniCollections(filename)

	if content, err := ioutil.ReadFile(filename); err != nil {
		t.Fatal("Error reading the collections file!")
	} else {
		if string(content) != "[]" {
			t.Fatal("collections content is not similar to `[]`")
		}
	}
}
