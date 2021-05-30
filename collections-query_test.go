package minidb

import (
	"reflect"
	"testing"
)

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

func TestFilter(t *testing.T) {
	filename := "filter.json"
	defer cleanFileAfter(filename, t)

	db := NewMiniCollections(filename)

	db.Push("hellox")
	db.Push("sample")
	db.Push(1)
	db.Push(100)
	db.Push(false)

	frString := []string{"hellox", "sample"}
	if !reflect.DeepEqual(db.FilterString(), frString) {
		t.Fatal("wrong filter string output")
	}

	frInt := []int{1, 100}
	if !reflect.DeepEqual(db.FilterInt(), frInt) {
		t.Fatal("wrong filter int output")
	}

	frBool := []bool{false}
	if !reflect.DeepEqual(db.FilterBool(), frBool) {
		t.Fatal("wrong filter bool output")
	}
}
