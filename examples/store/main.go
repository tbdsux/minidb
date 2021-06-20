package main

import (
	"fmt"
	"log"

	"github.com/TheBoringDude/minidb"
)

func main() {
	app := minidb.New("db")

	store := app.Store("my-store")

	store.Set("hello", "world")
	store.Set("my number", 1)
	store.Set("production", false)
	store.Set("name", "Me")

	t := store.GetString("hello")
	fmt.Println(t)

	u, ok := store.Get("name")
	if !ok {
		log.Fatal("invalid!")
	}

	fmt.Println(u.(string) == "Me")
}
