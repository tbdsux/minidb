package minidb

import (
	"encoding/json"
	"errors"
	"log"
	"path"
	"sync"
)

// base functions for creating a new store
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
	d := db.getOrCreateMutex("store_" + key)
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
// It just wraps around getValueOK.
func (db *MiniStore) getValue(key string) interface{} {
	value, ok := db.getValueOK(key)

	if !ok {
		log.Fatalf("Unknown key: %s", key)
	}

	return value
}

// it just returns the value of the key
func (db *MiniStore) getValueOK(key string) (interface{}, bool) {
	d := db.getOrCreateMutex("get_" + key)
	d.Lock()
	defer d.Unlock()

	value, ok := db.store[key]
	return value, ok
}

// Set sets the store[key] to v.
func (db *MiniStore) Set(key string, v interface{}) error {
	d := db.getOrCreateMutex("write_" + key)
	d.Lock()
	defer d.Unlock()

	_, ok := db.getValueOK(key)
	if ok {
		return errors.New("key already exists")
	}

	db.store[key] = v
	db.writeToDB()

	return nil
}

// Write takes a struct and writes it to the json file.
// It accepts a new struct object and encodes and write it.
func (db *MiniStore) Write(v interface{}) error {
	d, err := json.Marshal(v)
	if err != nil {
		return err
	}

	json.Unmarshal(d, &db.store)

	return write(db.db, d)
}
