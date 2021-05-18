package minidb

import "log"

// getValue tries to get the key from the map if exists. If value is nil,
//  It will log error that the key is unknown.
func (db *MiniDB) getValue(key string) interface{} {
	value := db.store[key]
	if value == nil {
		log.Fatalf("Unknown key: %s", key)
	}

	return value
}

// GetBool finds the key with bool value and returns if exits.
func (db *MiniDB) GetBool(key string) bool {
	return db.getValue(key).(bool)
}
