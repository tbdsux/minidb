package minidb

import (
	"reflect"
	"strings"
	"testing"
)

func TestPush_Collections(t *testing.T) {
	filename := "pushcols.json"

	defer cleanFileAfter(filename, t)

	db := NewCollections(filename)
	db.Push([]int{1, 2, 3, 4, 5}, "sample")

	checkFileContent(filename, `[[1,2,3,4,5],"sample"]`, t)
}

func TestFirstLast_Collections(t *testing.T) {
	filename := "firstcols.json"

	defer cleanFileAfter(filename, t)

	db := NewCollections(filename)

	db.Push("hello", 1000, false)

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

	db := NewCollections(filename)

	db.Push("hello", 1000, false, "THIN", "sample")

	if db.Find("THIN") != 3 || db.Find("hello") != 0 || db.Find(1000) != 1 {
		t.Fatal("wrong index value from find")
	}
}

func TestFindAll_Collections(t *testing.T) {
	filename := "findall.json"
	defer cleanFileAfter(filename, t)

	db := NewCollections(filename)

	db.Push("hello", 1, true, "hello")

	if len(db.FindAll("hello")) != 2 {
		t.Fatal("wrong return length from db.FindAll")
	}
}

func TestMatchString_Collections(t *testing.T) {
	filename := "matchstring.json"
	defer cleanFileAfter(filename, t)

	db := NewCollections(filename)

	db.Push("hello", 1, true, "hello")

	if value, _ := db.MatchString("he"); value != "hello" {
		t.Fatal("wrong match")
	}
}

func TestMatchStringAll_Collections(t *testing.T) {
	filename := "matchstring.json"
	defer cleanFileAfter(filename, t)

	db := NewCollections(filename)

	db.Push("hellox", 1, true, "hello_world")

	sampleReturn := []string{"hellox", "hello_world"}
	if values, _ := db.MatchStringAll("hello"); values[0] != sampleReturn[0] || values[1] != sampleReturn[1] {
		t.Fatal("wrong return value from match string all")
	}
}

func TestFilter(t *testing.T) {
	filename := "filter.json"
	defer cleanFileAfter(filename, t)

	db := NewCollections(filename)

	db.Push("hellox", "sample", 1, 100, false)

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

func TestFilterFunc(t *testing.T) {
	filename := "filterfunc.json"
	defer cleanFileAfter(filename, t)

	db := NewCollections(filename)
	db.Push("hellox", "sample", 1, 100, false)
	db.Push("sample", "another", 100, 9023423489, "he she", false, "hello world")

	strValues := db.FilterStringFunc(func(x string) bool {
		return strings.Contains(x, "he")
	})

	if !reflect.DeepEqual(strValues, []string{"hellox", "another", "he she", "hello world"}) {
		t.Fatal("wrong returned values from filter func")
	}
}

func TestRemove(t *testing.T) {
	filename := "removecols.json"
	defer cleanFileAfter(filename, t)

	db := NewCollections(filename)
	db.Push("sample", "another", 100, 9023423489, "sample", false, "sample")

	db.RemoveAll("sample")
	if len(db.FindAll("sample")) > 0 {
		t.Fatal("some items are not removed")
	}

	db.Remove(100)
	if db.Find(100) != -1 {
		t.Fatal("item is not removed")
	}
}
