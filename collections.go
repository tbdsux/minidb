package minidb

import (
	"encoding/json"
	"path"
	"sync"
)

func parseCollection(folderPath, filename string) *MiniCollections {
	db := &MiniCollections{
		store:   []interface{}{},
		mutexes: make(map[int]*sync.Mutex),
		BaseMiniDB: BaseMiniDB{
			path:     folderPath,
			filename: filename,
			db:       path.Join(folderPath, filename),
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

// Collections creates a new key with an array / slice value.
func (db *MiniDB) Collections(key string) *MiniCollections {
	d := db.getOrCreateMutex(key)
	d.Lock()
	defer d.Unlock()

	// if the key exists, get the file's name,
	// otherwise, create a new one
	filename, ok := db.store.Collections[key]
	if !ok {
		filename = "cols." + generateId() + ".json"
	}

	db.store.Collections[key] = filename
	db.writeToDB()

	return parseCollection(db.path, filename)
}
