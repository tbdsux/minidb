package minidb

import (
	"encoding/json"
	"log"

	"github.com/segmentio/ksuid"
)

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
