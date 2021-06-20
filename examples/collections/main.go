package main

import (
	"fmt"

	"github.com/TheBoringDude/minidb"
)

func main() {
	app := minidb.New("db")

	cols := app.Collections("something to watch")
	cols.Push("How To Train Your Dragon", "Avengers", "Iron Main", "Spiderman", "Kimetsu no Yaiba")

	search, _ := cols.MatchStringAll("How")
	for _, v := range search {
		fmt.Println(v)
	}
}
