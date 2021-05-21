package minidb

import (
	"os"
	"path"
	"testing"
)

// removes the f which could be the created file or folder
func cleanFileAfter(f string, t *testing.T) {
	if !isPathExists(path.Join(f, "__default.json")) {
		t.Fatal("Default json db file: `__default.json` does not exist!")
	} else {
		// clean dir
		os.RemoveAll(f)
	}
}

func TestNew(t *testing.T) {
	dirname := "sampledir"
	New(dirname)

	cleanFileAfter(dirname, t)
}

func TestNewMiniCollections(t *testing.T) {
	file := "cols.json"
	NewMiniCollections(file)

	cleanFileAfter(file, t)
}

func TestNewMiniStore(t *testing.T) {
	file := "store.json"
	newMiniStore(file)

	cleanFileAfter(file, t)
}
