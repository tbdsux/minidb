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
