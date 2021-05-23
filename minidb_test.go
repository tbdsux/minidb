package minidb

import (
	"testing"
)

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
