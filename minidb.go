package minidb

import (
	"sync"
)

const Version = "0.1.0"

type (
	// BaseMiniDB is the base db structure.
	BaseMiniDB struct {
		db    string // combined path and filename
		mutex *sync.Mutex
	}

	// MiniDB is the base store file.
	MiniDB struct {
		path     string
		filename string
		store    MiniDBStore
		mutexes  map[string]*sync.Mutex
		BaseMiniDB
	}

	// MiniDBStore is the types of MiniDB.store
	MiniDBStore struct {
		Keys        map[string]string `json:"keys"`
		Collections map[string]string `json:"collections"`
		Store       map[string]string `json:"store"`
	}

	// MiniCollections is a collections store.
	MiniCollections struct {
		store   []interface{}
		mutexes map[int]*sync.Mutex
		BaseMiniDB
	}

	// MiniStore is a key-value store.
	MiniStore struct {
		store   map[string]interface{}
		mutexes map[string]*sync.Mutex
		BaseMiniDB
	}
)

// New creates a new MiniDB struct.
// The dir will be created if it doesn't exist and a file named `__default.json` will also be generated.
// It is better to use this in managing multiple json files.
func New(dir string) *MiniDB {
	return newMiniDB(dir, "__default.json")
}

// NewMiniStore creates and returns a new key-store collection json db.
func NewMiniStore(f string) *MiniStore {
	return newMiniStore(f)
}

// NewMiniCollections creates and returns a new collections json db.
func NewMiniCollections(f string) *MiniCollections {
	return newMiniCollection(f)
}

// ListCollections returns the list of collections created.
func (db *MiniDB) ListCollections() []string {
	cols := []string{}

	for i := range db.store.Collections {
		cols = append(cols, i)
	}

	return cols
}

// ListCollections returns the list of collections created.
func (db *MiniDB) ListStores() []string {
	stores := []string{}

	for i := range db.store.Store {
		stores = append(stores, i)
	}

	return stores
}
