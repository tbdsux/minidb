# minidb

[![Test Status](https://github.com/TheBoringDude/minidb/workflows/Test/badge.svg)](https://github.com/TheBoringDude/minidb/actions)

a simple multi-flat-files json database

This has a really weird structure and api and **I know it :happy:**

## Install

```bash
go get -u github.com/TheBoringDude/minidb
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/TheBoringDude/minidb"
)

func main() {
	db := mindb.New('db')

    fmt.Println(db)
}

```

All operations are just appending and setting in a `map[string]interface{}` or `append([]interface{}, interface{})`.

In all operations, it writes to the json file. I think it is not a good idea?

#### Full Doc: https://pkg.go.dev/github.com/TheBoringDude/minidb

### MiniDB

It takes a directory and manages all files within it. It is better only to use this when trying to manage multiple json files.

New files are created with a **`random uuid`** using [**`ksuid`**](https://github.com/segmentio/ksuid) so each file created by the library is unique.

#### NOTE: this creates many json files in a specified directory

```go
db := minidb.New("dirfolder")

// db.Keys("key"), nested minidbs
// db.Collections("key"), a json collections, []
// db.Store("key"), a simple json key-value store (not meant with nested maps)

cols := db.Collections("posts")
cols.Push(map[string]string{
    "title": "Hello World",
    "content": "This is just something, maybe a content or not. I don't know how it works though.",
})

// multiple elements is possible
cols.Push(100, 20, "sample", false, []int{1,2,3,4,5})

fmt.Println(cols)
```

### MiniCollections

A simple collections json db file.

```go
db := minidb.NewCollections("cols.json")
db.Push(1)

fmt.Println(1)
```

### MiniStore

A simple key-value store json db file.

```go
db := minidb.NewStore("store.json")
db.Set("key", "value")

fmt.Println(db.GetString("key"))
```

## TODO

- Improve concurrency support.
- fixes, improvements
- more changes..
- ...future development

##

#### &copy; 2021 | [License](./LICENSE)
