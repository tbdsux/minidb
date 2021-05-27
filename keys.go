package minidb

import (
	"encoding/json"
	"path"
	"sync"
)

// this is helper for creating a new key db
func newMiniDB(dir, filename string) *MiniDB {
	db := &MiniDB{
		path:     dir,
		filename: filename,
		store: MiniDBStore{
			Keys:        map[string]string{},
			Collections: map[string]string{},
			Store:       map[string]string{},
		},
		mutexes: make(map[string]*sync.Mutex),
		BaseMiniDB: BaseMiniDB{
			db: path.Join(dir, filename),

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

// Key creates a new key in the json.
// It is better to use this for nesting and only if needed.
func (db *MiniDB) Key(key string) *MiniDB {
	d := db.getOrCreateMutex("keys_" + key)
	d.Lock()
	defer d.Unlock()

	// if the key exists, get the file's name,
	// otherwise, create a new one
	filename, ok := db.store.Keys[key]
	if !ok {
		filename = generateFileName("key")
	}

	db.store.Keys[key] = filename
	db.writeToDB()

	return newMiniDB(db.path, filename)
}
