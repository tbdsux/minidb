package minidb

import (
	"path"
	"testing"
)

func TestNew(t *testing.T) {
	dirname := "sampledir"
	New(dirname)

	cleanFileAfter(path.Join(dirname, "__default.json"), t)
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
