package minidb

import (
	"testing"
)

func TestCollections_FileContent(t *testing.T) {
	filename := "content.json"
	NewMiniCollections(filename)

	checkFileContent(filename, "[]", t)
	cleanFileAfter(filename, t)
}
