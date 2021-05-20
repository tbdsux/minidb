package minidb

import (
	"io/ioutil"
)

// writeToDB write the db.store to the defined json db file.
func (db *MiniDB) writeToDB() {
	data := db.marshalStore()

	ioutil.WriteFile(db.db, []byte(string(data)), 0755)
}
