package minidb

import "testing"

func TestPush_Collections(t *testing.T) {
	filename := "pushcols.json"

	defer cleanFileAfter(filename, t)

	db := NewMiniCollections(filename)
	db.Push([]int{1, 2, 3, 4, 5})

	checkFileContent(filename, "[[1,2,3,4,5]]", t)
}

func TestFirstLast_Collections(t *testing.T) {
	filename := "firstcols.json"

	defer cleanFileAfter(filename, t)

	db := NewMiniCollections(filename)

	db.Push("hello")
	db.Push(1000)
	db.Push(false)

	if db.First().(string) != "hello" {
		t.Fatal("the first element is not equal to `hello`")
	}

	if db.Last().(bool) != false {
		t.Fatal("the last element is not equal to `false`")
	}
}
