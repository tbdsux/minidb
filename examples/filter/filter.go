package main

import (
	"fmt"
	"strings"

	"github.com/TheBoringDude/minidb"
)

func main() {
	db := minidb.NewMiniCollections("cols.json")

	db.Push("sample", "sam", "hello", 100, true)

	search := db.Filter(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			if strings.Contains(s, "sam") {
				return true
			}

		}

		return false
	})

	fmt.Println(search)
}
