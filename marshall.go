package minidb

import (
	"encoding/json"
)

// return the marshalled db.store
func marshalStore(v interface{}) []byte {
	d, err := json.Marshal(v)
	logError(err, "error tring to marshall struct")

	return d
}
