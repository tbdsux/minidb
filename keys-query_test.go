package minidb

import (
	"path"
	"testing"
)

func TestRemoveKeys(t *testing.T) {
	f := "removekeys"

	defer cleanFileAfter(f, t)

	db := New(f)
	db.Key("sample")
	db.Collections("cols")
	db.Store("stores")

	sk, _ := db.FindKey("sample")
	if !isPathExists(path.Join(f, sk)) {
		t.Fatal("the key does not exist")
	}

	// attempt to remove
	db.RemoveKey("sample")
	if isPathExists(path.Join(f, sk)) {
		t.Fatal("the key still exists!")
	}
}
