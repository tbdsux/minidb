package minidb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type MiniDB struct {
	filename string
	store    map[string]interface{}
}

// New creates a new MiniDB struct.
func New(filename string) *MiniDB {
	var content []byte = []byte("{}")

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		ioutil.WriteFile(filename, content, 0755)
	} else {
		data, err := ioutil.ReadFile(filename)
		logError(err, err)

		content = data
	}

	db := &MiniDB{
		filename: filename,
	}

	json.Unmarshal(content, &db.store)

	return db
}

// writeToDB write the db.store to the defined json db file.
func (db *MiniDB) writeToDB() {
	m, err := json.Marshal((db.store))
	logError(err, "error while trying to write to db")

	ioutil.WriteFile(db.filename, []byte(string(m)), 0755)
}

// Key creates a new key in the json.
func (db *MiniDB) Key(key string) {
	db.store[key] = map[string]interface{}{}

	db.writeToDB()
}

// KeyValue creates a new key with a given value in the json.
func (db *MiniDB) KeyValue(key string, value interface{}) {
	db.store[key] = value

	db.writeToDB()
}
