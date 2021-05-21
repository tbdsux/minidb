package minidb

import (
	"io/ioutil"
	"testing"
)

func TestSet_Store(t *testing.T) {
	defer cleanFileAfter("store.json", t)

	db := NewMiniStore("store.json")
	db.Set("hello", "world")

	if check, err := ioutil.ReadFile("store.json"); err != nil {
		if string(check) != `{"hello":"world"}` {
			t.Fatal("write is not similar to the output file")
		}
	} else {
		t.Fatal(err)
	}
}
