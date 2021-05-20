package minidb

// TODO::
// Key creates a new key in the json.
func (db *MiniDB) Key(key string) *MiniDB {
	d := db.getOrCreateMutex(key)
	d.Lock()
	defer d.Unlock()

	// if the key exists, get the file's name,
	// otherwise, create a new one
	filename, ok := db.store.Keys[key]
	if !ok {
		filename = generateId() + ".json"
	}

	db.store.Keys[key] = filename
	db.writeToDB()

	return parseNew(db.path, filename)
}
