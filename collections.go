package minidb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

func parseCollection(folderPath, filename string) *MiniCollections {
	cols := &MiniCollections{
		path:     folderPath,
		filename: filename,
		db:       path.Join(folderPath, filename),
		store:    []interface{}{},
		mutex:    &sync.Mutex{},
		mutexes:  make(map[int]*sync.Mutex),
	}

	var content []byte

	if initialData, err := json.Marshal(&cols.store); err != nil {
		content = initialData
	}

	// create the folder
	if _, err := os.Stat(folderPath); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(folderPath, 0755)
	} else {
		logError(err, "error trying to read / check folder")
	}

	// create the json db file
	if _, err := os.Stat(cols.db); errors.Is(err, os.ErrNotExist) {
		cols.writeToDB()
	} else {
		data, err := ioutil.ReadFile(cols.db)
		logError(err, err)

		content = data
	}

	json.Unmarshal(content, &cols.store)

	return cols
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
