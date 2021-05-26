package minidb

import (
	"encoding/json"
	"path"
	"sync"
)

func newMiniCollection(filename string) *MiniCollections {
	db := &MiniCollections{
		store:   []interface{}{},
		mutexes: make(map[int]*sync.Mutex),
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

// Collections creates a new key with an array / slice value.
func (db *MiniDB) Collections(key string) *MiniCollections {
	d := db.getOrCreateMutex("cols_" + key)
	d.Lock()
	defer d.Unlock()

	// if the key exists, get the file's name,
	// otherwise, create a new one
	filename, ok := db.store.Collections[key]
	if !ok {
		filename = generateFileName("cols")
	}

	db.store.Collections[key] = filename
	db.writeToDB()

	return newMiniCollection(path.Join(db.path, filename))
}
