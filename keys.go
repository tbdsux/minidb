package minidb

import (
	"encoding/json"
	"path"
	"sync"
)

// Key creates a new key in the json.
func (db *MiniDB) Key(key string) *MiniDB {
	d := db.getOrCreateMutex(key)
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

// this is helper for creating a new key db
func newMiniDB(folderPath, filename string) *MiniDB {
	db := &MiniDB{
		store: MiniDBStore{
			Keys:        map[string]string{},
			Collections: map[string]string{},
			Values:      map[string]string{},
		},
		mutexes: make(map[string]*sync.Mutex),
		BaseMiniDB: BaseMiniDB{
			db:       path.Join(folderPath, filename),
			path:     folderPath,
			filename: filename,
			mutex:    &sync.Mutex{},
		},
	}
	if content, f := ensureInitialDB(folderPath, db.db); f {
		db.writeToDB()
	} else {
		json.Unmarshal(content, &db.store)
	}

	return db
}
