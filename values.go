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

//
func (db *MiniDB) recoverAssertion() {
	if r := recover(); r != nil {
		log.Fatalln(r)
	}
}

// getValue tries to get the key from the map if exists. If value is nil,
//  It will log error that the key is unknown.
func (db *MiniDB) getValue(key string) interface{} {
	value, ok := db.store.Values[key]

	if !ok {
		log.Fatalf("Unknown key: %s", key)
	}

	return value
}

// GetBool finds the key with bool value and returns if exits.
func (db *MiniDB) GetBool(key string) bool {
	defer db.recoverAssertion()

	return db.getValue(key).(bool)
}

// GetString finds the key with the string value and returns if exists.
func (db *MiniDB) GetString(key string) string {
	defer db.recoverAssertion()

	return db.getValue(key).(string)
}
