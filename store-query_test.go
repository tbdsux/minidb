package minidb

import "testing"

func TestGetQuery(t *testing.T) {
	defer cleanFileAfter("getstore.json", t)

	db := NewMiniStore("getstore.json")
	db.Set("hello", "world")

	if db.GetString("hello") != "world" {
		t.Fatal("`hello` key is not equal to world")
	}
}

func TestRemoveQuery(t *testing.T) {
	filename := "removestore.json"

	defer cleanFileAfter(filename, t)

	db := NewMiniStore(filename)
	db.Set("value", false)
	db.Set("string", "123")

	err := db.Remove("value")
	if err != nil {
		t.Fatal("key is not removed")
	}
}

func TestUpdateQuery(t *testing.T) {
	filename := "updatestore.json"

	defer cleanFileAfter(filename, t)

	db := NewMiniStore(filename)
	db.Set("value", false)
	db.Set("string", "123")

	db.Update("value", true)
	if db.GetBool("value") != true {
		t.Fatal("update is not working ")
	}
}
