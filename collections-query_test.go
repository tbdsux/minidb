package minidb

import "testing"

func TestPush_Collections(t *testing.T) {
	filename := "pushcols.json"
	db := NewMiniCollections(filename)

	db.Push([]int{1, 2, 3, 4, 5})

	checkFileContent(filename, "[[1,2,3,4,5]]", t)
	cleanFileAfter(filename, t)
}
