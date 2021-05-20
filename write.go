package minidb

import (
	"io/ioutil"
)

// writeToDB write the db.store to the defined json db file.
func (db *MiniDB) writeToDB() {
	data := db.marshalStore()

	err := ioutil.WriteFile(db.db, data, 0755)
	logError(err, "Error writing to DB.")
}
