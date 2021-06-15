package minidb

import (
	"encoding/json"

	"github.com/segmentio/ksuid"
)

// generates a new id with ksuid
func generateId() string {
	return ksuid.New().String()
}

// generate a string filename
func generateFileName(k string) string {
	return k + "." + generateId() + ".json"
}

// json marshalling wrapper
func marshalStore(v interface{}) []byte {
	d, err := json.Marshal(v)
	logError(err, "error tring to marshall struct")

	return d
}
