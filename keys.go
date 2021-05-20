package minidb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
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
		filename = generateId() + ".json"
	}

	db.store.Keys[key] = filename
	db.writeToDB()

	return parseNew(db.path, filename)
}

// this is helper for creating a new key db
func parseNew(folderPath, filename string) *MiniDB {
	db := &MiniDB{
		store: MiniDBStore{
			Keys:        map[string]string{},
			Collections: map[string]string{},
			Values:      map[string]interface{}{},
		},
		db:       path.Join(folderPath, filename),
		path:     folderPath,
		filename: filename,
		mutex:    &sync.Mutex{},
		mutexes:  make(map[string]*sync.Mutex),
	}

	var content []byte

	if initialData, err := json.Marshal(&db.store); err != nil {
		content = initialData
	}

	// create the folder
	if _, err := os.Stat(folderPath); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(folderPath, 0755)
	} else {
		logError(err, "error trying to read / check folder")
	}

	// create the json db file
	if _, err := os.Stat(db.db); errors.Is(err, os.ErrNotExist) {
		db.writeToDB()
	} else {
		data, err := ioutil.ReadFile(db.db)
		logError(err, err)

		content = data
	}

	json.Unmarshal(content, &db.store)

	return db
}
