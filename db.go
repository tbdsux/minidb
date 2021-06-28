package minidb

import (
	"encoding/json"
	"os"
	"path/filepath"

	simplefiletest "github.com/TheBoringDude/simple-filetest"
)

// ensures the db, creates the path if doesn't exist and reads the file db if exists
// returns true if file does not exist, otherwise, false
func ensureInitialDB(path string, defaultValue interface{}, defaultData string) ([]byte, bool) {
	// read the json db file if exists
	if simplefiletest.FileExists(path) {
		data, err := os.ReadFile(path)
		logError(err, err)

		return data, false
	}

	err := os.MkdirAll(filepath.Dir(path), 0755)
	logError(err, "error creating db path")

	// return the default value
	if defaultValue != nil {
		d, err := json.Marshal(defaultValue)
		logError(err, "failed to marshall default value")

		return d, true
	}

	return []byte(defaultData), true
}
