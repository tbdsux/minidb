package main

import (
	"fmt"

	"github.com/TheBoringDude/minidb"
)

func main() {
	cols := minidb.NewCollectionsWithDefault("test.json", []interface{}{1, "3", 4, "hello"})

	fmt.Println(cols.List())
}
