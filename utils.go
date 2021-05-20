package minidb

import (
	"encoding/json"
	"errors"
	"log"
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

// return the marshalled db.store
func marshalStore(v interface{}) []byte {
	d, err := json.Marshal(v)
	logError(err, "error tring to marshall struct")

	return d
}

//
func recoverAssertion() {
	if r := recover(); r != nil {
		log.Fatalln(r)
	}
}
