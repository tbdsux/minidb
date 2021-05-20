package minidb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// return the marshalled db.store
func (db *MiniDB) marshalStore() []byte {
	d, err := json.Marshal(db.store)
	logError(err, "error tring to marshall struct")

	return d
}

// create the json file, path should be included to be sure
func (db *MiniDB) createDbFile() {
	data := db.marshalStore()

	err := ioutil.WriteFile(db.db, data, 0755)
	logError(err, "Error writing to DB.")

	fmt.Println("created")
}
