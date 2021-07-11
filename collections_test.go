package minidb

import (
	"testing"
)

func TestCollections_FileContent(t *testing.T) {
	filename := "content.json"
	NewCollections(filename)

	checkFileContent(filename, "[]", t)
	cleanFileAfter(filename, t)
}

func TestCollectionsWithDefault(t *testing.T) {
	dbname := "colsdef"
	db := New(dbname)

	values := []interface{}{1, 2, "hello", "3"}
	cols := db.CollectionsWithDefault("def", values)

	if cols.Find("hello") == -1 {
		t.Fatal("collections should have a default value but has none")
	}

	cleanFileAfter(dbname, t)
}
