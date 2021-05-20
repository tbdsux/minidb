package minidb

// TODO::
// Key creates a new key in the json.
func (db *MiniDB) Key(key string) *MiniDB {
	defer db.writeToDB()

	filename := generateId() + ".json"

	db.store.Keys[key] = filename

	return parseNew(db.path, filename)
}
