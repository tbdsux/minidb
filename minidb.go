package minidb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
)

type MiniDB struct {
	path  string
	store BaseMiniDB
}

type BaseMiniDB struct {
	Keys        map[string]string      `json:"keys"`
	Collections map[string]string      `json:"collections"`
	Values      map[string]interface{} `json:"values"`
}

// New creates a new MiniDB struct.
func New(folderPath string) *MiniDB {
	db := &MiniDB{
		store: BaseMiniDB{
			Keys:        map[string]string{},
			Collections: map[string]string{},
			Values:      map[string]interface{}{},
		},
	}

	var content []byte = []byte("")

	if initialData, err := json.Marshal(&db.store); err != nil {
		content = initialData
	}

	defaultDb := path.Join(folderPath, "__default.json")

	if _, err := os.Stat(folderPath); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(folderPath, 0755)
		ioutil.WriteFile(defaultDb, content, 0755)
	} else if errors.Is(err, os.ErrPermission) {
		logError(err, "Permission DENIED!")
	} else {
		data, err := ioutil.ReadFile(defaultDb)
		logError(err, err)

		content = data
	}

	db.path = defaultDb

	json.Unmarshal(content, &db.store)

	return db
}

// writeToDB write the db.store to the defined json db file.
func (db *MiniDB) writeToDB() {
	m, err := json.Marshal((db.store))
	logError(err, "error while trying to write to db")

	ioutil.WriteFile(db.path, []byte(string(m)), 0755)
}

// TODO::
// Key creates a new key in the json.
// func (db *MiniDB) Key(key string) {
// 	db.store[key] = map[string]interface{}{}

// 	db.writeToDB()
// }

// KeyValue creates a new key with a given value in the json.
func (db *MiniDB) KeyValue(key string, value interface{}) {
	db.store.Values[key] = value

	db.writeToDB()
}

// Collections creates a new key with an array / slice value.
// func (db *MiniDB) Collections(key string) {
// 	db.store.collections[key] = []interface{}{}

// 	db.writeToDB()
// }
