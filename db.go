package minidb

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// ensures the db, creates the path if doesn't exist and reads the file db if exists
// returns true if file does not exist, otherwise, false
func ensureInitialDB(path string) ([]byte, bool) {
	// read the json db file if exists
	if isPathExists(path) {
		data, err := ioutil.ReadFile(path)
		logError(err, err)

		return data, false
	}

	err := os.MkdirAll(filepath.Dir(path), 0755)
	logError(err, "error trying to write to db file path")

	return make([]byte, 0), true
}
