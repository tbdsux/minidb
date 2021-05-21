package minidb

import (
	"encoding/json"
	"log"
	"path"
	"sync"
)

func newMiniStore(filename string) *MiniStore {
	db := &MiniStore{
		store:   map[string]interface{}{},
		mutexes: make(map[string]*sync.Mutex),
		BaseMiniDB: BaseMiniDB{
			db:    filename,
			mutex: &sync.Mutex{},
		},
	}
	if content, f := ensureInitialDB(db.db); f {
		db.writeToDB()
	} else {
		json.Unmarshal(content, &db.store)
	}

	return db
}

// Store creates a new key with a given value in the json.
func (db *MiniDB) Store(key string) *MiniStore {
	d := db.getOrCreateMutex(key)
	d.Lock()
	defer d.Unlock()

	// if the key exists, get the file's name,
	// otherwise, create a new one
	filename, ok := db.store.Keys[key]
	if !ok {
		filename = generateFileName("store")
	}

	db.store.Keys[key] = filename
	db.writeToDB()

	return newMiniStore(path.Join(db.path, filename))
}

// getValue tries to get the key from the map if exists. If value is nil,
// It will log error that the key is unknown.
func (db *MiniStore) getValue(key string) interface{} {
	d := db.getOrCreateMutex(key)
	d.Lock()
	defer d.Unlock()

	value, ok := db.store[key]

	if !ok {
		log.Fatalf("Unknown key: %s", key)
	}

	return value
}

// Set sets the store[key] to v.
func (db *MiniStore) Set(key string, v interface{}) {
	d := db.getOrCreateMutex(key)
	d.Lock()
	defer d.Unlock()

	db.store[key] = v
	db.writeToDB()
}
