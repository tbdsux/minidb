package minidb

import "testing"

func TestIsPathExists(t *testing.T) {
	if !isPathExists("minidb.go") {
		t.Fatal("path: does not exist but is exists")
	}
}
