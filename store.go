package minidb

import (
	"encoding/json"
	"errors"
	"log"
	"path"
	"sync"
)

// base functions for creating a new store
func ministore(filename string, defaultValue interface{}) *MiniStore {
	db := &MiniStore{
		content: map[string]interface{}{},
		mutexes: make(map[string]*sync.Mutex),
		BaseMiniDB: BaseMiniDB{
			db:    filename,
			mutex: &sync.Mutex{},
		},
	}

	content, f := ensureInitialDB(db.db, defaultValue, "{}")
	err := json.Unmarshal(content, &db.content)
	logError(err, "(collections) failed to unmarshall content to struct")

	if f {
		db.writeToDB()
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
	filename, ok := db.content.Store[key]
	if !ok {
		filename = generateFileName("store")
	}

	db.content.Store[key] = filename
	db.writeToDB()

	return ministore(path.Join(db.path, filename), nil)
}

// Store creates a new key with a given value in the json. If the key does not exist,
// it will write the defaultValue as the first data to the db.
func (db *MiniDB) StoreWithDefault(key string, defaultValue interface{}) *MiniStore {
	d := db.getOrCreateMutex("store_" + key)
	d.Lock()
	defer d.Unlock()

	// if the key exists, get the file's name,
	// otherwise, create a new one
	filename, ok := db.content.Store[key]
	if !ok {
		filename = generateFileName("store")
	}

	db.content.Store[key] = filename
	db.writeToDB()

	return ministore(path.Join(db.path, filename), defaultValue)
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

	value, ok := db.content[key]
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

	db.content[key] = v
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

	json.Unmarshal(d, &db.content)

	return write(db.db, d)
}

// Read parses the contents of db.store to v which is a struct object.
// It just wraps around `json.Marshal` and `json.Unmarshal`.
func (db *MiniStore) Read(v interface{}) error {
	d, err := json.Marshal(db.content)
	if err != nil {
		return err
	}

	return json.Unmarshal(d, v)
}

// ReadKey parses the contents of db.store[key] to v which is a struct object.
// It is better to use this if the value of key is a map.
// It just wraps around `json.Marshal` and `json.Unmarshal`.
func (db *MiniStore) ReadKey(key string, v interface{}) error {
	value, ok := db.getValueOK(key)
	if !ok {
		return errors.New("key does not exist")
	}

	d, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return json.Unmarshal(d, v)
}
