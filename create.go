package minidb

import (
	"encoding/json"
)

// return the marshalled db.store
func (db *MiniDB) marshalStore() []byte {
	d, err := json.Marshal(db.store)
	logError(err, "error tring to marshall struct")

	return d
}
