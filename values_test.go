package minidb

import (
	"io/ioutil"
	"testing"
)

func TestSet_Store(t *testing.T) {
	defer cleanFileAfter("setstore.json", t)

	db := NewMiniStore("setstore.json")
	db.Set("hello", "world")

	if check, err := ioutil.ReadFile("setstore.json"); err == nil {
		if string(check) != `{"hello":"world"}` {
			t.Fatal("write is not similar to the output file")
		}
	} else {
		t.Fatal(err)
	}
}
