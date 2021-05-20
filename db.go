package minidb

import (
	"io/ioutil"
	"os"
)

// ensures the db, creates the path if doesn't exist and reads the file db if exists
// returns true if file does not exist, otherwise, false
func ensureInitialDB(folder, db string) ([]byte, bool) {
	// create the folder
	if !isPathExists(folder) {
		os.MkdirAll(folder, 0755)
	}

	// create the json db file
	if isPathExists(db) {
		data, err := ioutil.ReadFile(db)
		logError(err, err)

		return data, false
	}

	return make([]byte, 0), true
}
