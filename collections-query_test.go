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

func TestFind_Collections(t *testing.T) {
	filename := "findcols.json"

	defer cleanFileAfter(filename, t)

	db := NewMiniCollections(filename)

	db.Push("hello")
	db.Push(1000)
	db.Push(false)
	db.Push("THIN")
	db.Push("sample")

	if db.Find("THIN") != 3 || db.Find("hello") != 0 || db.Find(1000) != 1 {
		t.Fatal("wrong index value from find")
	}
}

func TestFindAll_Collections(t *testing.T) {
	filename := "findall.json"
	defer cleanFileAfter(filename, t)

	db := NewMiniCollections(filename)

	db.Push("hello")
	db.Push(1)
	db.Push(true)
	db.Push("hello")

	if len(db.FindAll("hello")) != 2 {
		t.Fatal("wrong return length from db.FindAll")
	}
}

func TestMatchString_Collections(t *testing.T) {
	filename := "matchstring.json"
	defer cleanFileAfter(filename, t)

	db := NewMiniCollections(filename)

	db.Push("hello")
	db.Push(1)
	db.Push(true)
	db.Push("hello")

	if value, _ := db.MatchString("he"); value != "hello" {
		t.Fatal("wrong match")
	}
}

func TestMatchStringAll_Collections(t *testing.T) {
	filename := "matchstring.json"
	defer cleanFileAfter(filename, t)

	db := NewMiniCollections(filename)

	db.Push("hellox")
	db.Push(1)
	db.Push(true)
	db.Push("hello_world")

	sampleReturn := []string{"hellox", "hello_world"}
	if values, _ := db.MatchStringAll("hello"); values[0] != sampleReturn[0] || values[1] != sampleReturn[1] {
		t.Fatal("wrong return value from match string all")
	}
}
