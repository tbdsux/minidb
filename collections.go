package minidb

import (
	"encoding/json"
	"path"
	"sync"
)

// base function for creating a new collection
func minicollection(filename string) *MiniCollections {
	db := &MiniCollections{
		content:   []interface{}{},
		mutexes: make(map[int]*sync.Mutex),
		BaseMiniDB: BaseMiniDB{
			db:    filename,
			mutex: &sync.Mutex{},
		},
	}

	if content, f := ensureInitialDB(db.db); f {
		db.writeToDB()
	} else {
		json.Unmarshal(content, &db.content)
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
	filename, ok := db.content.Collections[key]
	if !ok {
		filename = generateFileName("cols")
	}

	db.content.Collections[key] = filename
	db.writeToDB()

	return minicollection(path.Join(db.path, filename))
}
