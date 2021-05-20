package minidb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"sync"

	"github.com/segmentio/ksuid"
)

type MiniDB struct {
	path     string
	filename string
	db       string // combined path and filename
	store    BaseMiniDB
	mutex    *sync.Mutex
	mutexes  map[string]*sync.Mutex
}

type BaseMiniDB struct {
	Keys        map[string]string      `json:"keys"`
	Collections map[string]string      `json:"collections"`
	Values      map[string]interface{} `json:"values"`
}

// New creates a new MiniDB struct.
func New(folderPath string) *MiniDB {
	return parseNew(folderPath, "__default.json")
}

// this is helper for creating a new key db
func parseNew(folderPath, filename string) *MiniDB {
	db := &MiniDB{
		store: BaseMiniDB{
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

	var content []byte = []byte("")

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

// referred from ::> https://github.com/sdomino/scribble/blob/master/scribble.go#L254
func (db *MiniDB) getOrCreateMutex(key string) *sync.Mutex {

	db.mutex.Lock()
	defer db.mutex.Unlock()

	m, ok := db.mutexes[key]

	// if the mutex doesn't exist make it
	if !ok {
		m = &sync.Mutex{}
		db.mutexes[key] = m
	}

	return m
}

// generates a new id with ksuid
func generateId() string {
	return ksuid.New().String()
}
