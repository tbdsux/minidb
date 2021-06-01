package minidb

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/segmentio/ksuid"
)

// checks that the path or file exists
func isPathExists(f string) bool {
	if _, err := os.Stat(f); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

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
