package minidb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/segmentio/ksuid"
)

type MiniDB struct {
	path     string
	filename string
	db       string // combined path and filename
	store    BaseMiniDB
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
		db.createDbFile()
	} else {
		data, err := ioutil.ReadFile(db.db)
		logError(err, err)

		content = data
	}

	json.Unmarshal(content, &db.store)

	return db
}

// generates a new id with ksuid
func generateId() string {
	return ksuid.New().String()
}
