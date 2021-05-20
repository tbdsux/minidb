package minidb

import (
	"sync"
)

type (
	// MiniDB is the base store file.
	MiniDB struct {
		path     string
		filename string
		db       string // combined path and filename
		store    MiniDBStore
		mutex    *sync.Mutex
		mutexes  map[string]*sync.Mutex
	}

	// MiniDBStore is the types of MiniDB.store
	MiniDBStore struct {
		Keys        map[string]string      `json:"keys"`
		Collections map[string]string      `json:"collections"`
		Values      map[string]interface{} `json:"values"`
	}

	MiniDBCollectionsStore = []interface{}
	// MiniCollections is a new collections store.
	MiniCollections struct {
		path     string
		filename string
		db       string // combined path and filename
		store    MiniDBCollectionsStore
		mutex    *sync.Mutex
		mutexes  map[int]*sync.Mutex
	}
)

// New creates a new MiniDB struct.
func New(folderPath string) *MiniDB {
	return parseNew(folderPath, "__default.json")
}
