package minidb

import (
	"sync"
)

type (
	// BaseMiniDB is the base db structure.
	BaseMiniDB struct {
		path     string
		filename string
		db       string // combined path and filename
		mutex    *sync.Mutex
	}

	// MiniDB is the base store file.
	MiniDB struct {
		store   MiniDBStore
		mutexes map[string]*sync.Mutex
		BaseMiniDB
	}

	// MiniDBStore is the types of MiniDB.store
	MiniDBStore struct {
		Keys        map[string]string      `json:"keys"`
		Collections map[string]string      `json:"collections"`
		Values      map[string]interface{} `json:"values"`
	}

	// MiniCollections is a new collections store.
	MiniCollections struct {
		store   []interface{}
		mutexes map[int]*sync.Mutex
		BaseMiniDB
	}
)

// New creates a new MiniDB struct.
func New(folderPath string) *MiniDB {
	return parseNew(folderPath, "__default.json")
}
