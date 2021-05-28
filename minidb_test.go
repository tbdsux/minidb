package minidb

import (
	"reflect"
	"sort"
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

func TestListCollections(t *testing.T) {
	dir := "many-cols"

	defer cleanFileAfter(dir, t)

	db := New(dir)
	db.Collections("new")
	db.Collections("sample")
	db.Collections("zoo")

	listCols := db.ListCollections()
	sort.Strings(listCols) // need to sort like the values below, since it can be different at any time
	sampleReturn := []string{"new", "sample", "zoo"}

	if !reflect.DeepEqual(listCols, sampleReturn) {
		t.Fatal("list collections wrong return values")
	}
}

func TestListStores(t *testing.T) {
	dir := "many-stores"

	defer cleanFileAfter(dir, t)

	db := New(dir)
	db.Store("new")
	db.Store("sample")
	db.Store("zoo")

	listStores := db.ListStores()
	sort.Strings(listStores) // need to sort like the values below, since it can be different at any time
	sampleReturn := []string{"new", "sample", "zoo"}

	if !reflect.DeepEqual(listStores, sampleReturn) {
		t.Fatal("list stores wrong return values")
	}
}
