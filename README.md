# minidb

[![Test Status](https://github.com/TheBoringDude/minidb/workflows/Test/badge.svg)](https://github.com/TheBoringDude/minidb/actions)

a simple multi-flat-files json database

This has a really weird structure and api and **I know it :happy:**

## Install

```bash
go get -u github.com/TheBoringDude/minidb
```

## Usage

All operations are just appending and setting in a `map[string]interface{}` or `append([]interface{}, interface{})`.

In all operations, it writes to the json file. I think it is not a good idea?

### MiniDB

It takes a directory and manages all files within it. It is better only to use this when trying to manage multiple json files.

New files are created with a **`random uuid`** using [**`ksuid`**](https://github.com/segmentio/ksuid)

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

fmt.Println(cols)
```

### MiniCollections

A simple collections json db file.

```go
db := minidb.NewMiniCollections("cols.json")
db.Push(1)

fmt.Println(1)
```

### MiniStore

A simple key-value store json db file.

```go
db := minidb.NewMiniStore("store.json")
db.Set("key", "value")

fmt.Println(db.GetString("key"))
```

## TODO

-   Improve concurrency support.
-   fixes, improvements
-   more changes..
-   ...future development

##

#### [License](./LICENSE)
