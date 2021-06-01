package minidb

import (
	"io/ioutil"
	"testing"
)

/* start test structs */
type TestStoreReadFunc struct {
	Value string `json:"value"`
	Bool  bool   `json:"bool"`
}

type TestSampleMapReadKey struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

/* end test structs */

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

func TestRead_Store(t *testing.T) {
	filename := ("readstore.json")

	defer cleanFileAfter(filename, t)

	db := NewMiniStore(filename)
	db.Set("value", "hello world")
	db.Set("bool", true)
	db.Set("user", map[string]interface{}{
		"name": "John",
		"age":  20,
	})

	reader := TestStoreReadFunc{}
	db.Read(&reader)

	user := TestSampleMapReadKey{}
	db.ReadKey("user", &user)

	if reader.Value != "hello world" {
		t.Fatal("wrong value from Read function")
	}

	if user.Name != "John" {
		t.Fatal("wrong value from ReadKey function")
	}
}
