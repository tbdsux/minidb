package minidb

import "log"

// KeyValue creates a new key with a given value in the json.
func (db *MiniDB) KeyValue(key string, value interface{}) {
	d := db.getOrCreateMutex(key)
	d.Lock()
	defer d.Unlock()

	db.store.Values[key] = value
	db.writeToDB()
}

// getValue tries to get the key from the map if exists. If value is nil,
// It will log error that the key is unknown.
func (db *MiniDB) getValue(key string) interface{} {
	value, ok := db.store.Values[key]

	if !ok {
		log.Fatalf("Unknown key: %s", key)
	}

	return value
}
