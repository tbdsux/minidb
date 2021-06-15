package minidb

import (
	"path"
	"testing"

	simplefiletest "github.com/TheBoringDude/simple-filetest"
)

func TestRemoveKeys(t *testing.T) {
	f := "removekeys"

	defer cleanFileAfter(f, t)

	db := New(f)
	db.Key("sample")
	db.Collections("cols")
	db.Store("stores")

	sk, _ := db.FindKey("sample")
	if !simplefiletest.Exists(path.Join(f, sk)) {
		t.Fatal("the key does not exist")
	}

	// attempt to remove
	db.RemoveKey("sample")
	if _, err := db.FindKey("sample"); err == nil {
		t.Fatal("key still exists!")
	}
	if simplefiletest.Exists(path.Join(f, sk)) {
		t.Fatal("the key still exists!")
	}
}
