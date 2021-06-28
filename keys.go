package minidb

import (
	"encoding/json"
	"path"
	"sync"
)

// this is helper for creating a new key db
func minidb(dir, filename string) *MiniDB {
	db := &MiniDB{
		path:     dir,
		filename: filename,
		content: MiniDBContent{
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

	if content, f := ensureInitialDB(db.db, nil, "{}"); f {
		db.writeToDB()
	} else {
		json.Unmarshal(content, &db.content)
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
	filename, ok := db.content.Keys[key]
	if !ok {
		filename = generateFileName("key")
	}

	db.content.Keys[key] = filename
	db.writeToDB()

	return minidb(db.path, filename)
}
