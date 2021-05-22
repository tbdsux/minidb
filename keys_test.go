package minidb

import (
	"path"
	"testing"
)

func TestKeys_FileContent(t *testing.T) {
	f := "keycontent"
	New(f)

	checkFileContent(path.Join(f, "__default.json"), `{"keys":{},"collections":{},"values":{}}`, t)
	cleanFileAfter(path.Join(f, "__default.json"), t)
}
