package minidb

import (
	"errors"
	"os"
	"path"
)

// FindKey gets the key in the keys map and returns its corresponding filename.
// It returns nil if it exists.
func (db *MiniDB) FindKey(key string) (string, error) {
	filename, ok := db.store.Keys[key]

	if !ok {
		return "", errors.New("the key does not exist")
	}

	return filename, nil
}

// FindCollection gets the key in the keys map and returns its corresponding filename.
// It returns nil if it exists.
func (db *MiniDB) FindCollection(key string) (string, error) {
	filename, ok := db.store.Keys[key]

	if !ok {
		return "", errors.New("the key does not exist")
	}

	return filename, nil
}

// FindStore gets the key in the keys map and returns its corresponding filename.
// It returns nil if it exists.
func (db *MiniDB) FindStore(key string) (string, error) {
	filename, ok := db.store.Keys[key]

	if !ok {
		return "", errors.New("the key does not exist")
	}

	return filename, nil
}

// RemoveCollection removes the collection key and the files corresponding to it.
// It returns nil if it is successful.
func (db *MiniDB) RemoveCollection(key string) error {
	d := db.getOrCreateMutex("delete_cols" + key)
	d.Lock()
	defer d.Unlock()

	// get the filename if it exists
	filename, ok := db.store.Collections[key]
	if !ok {
		return errors.New("collections key does not exist")
	}

	// remove the key and the filename
	delete(db.store.Collections, key)

	return os.RemoveAll(path.Join(db.path, filename))
}

// RemoveStore removes the store key and the files corresponding to it.
// It returns nil if it is successful.
func (db *MiniDB) RemoveStore(key string) error {
	d := db.getOrCreateMutex("delete_store" + key)
	d.Lock()
	defer d.Unlock()

	// get the filename if it exists
	filename, ok := db.store.Store[key]
	if !ok {
		return errors.New("collections key does not exist")
	}

	// remove the key and the filename
	delete(db.store.Store, key)

	return os.RemoveAll(path.Join(db.path, filename))
}

// RemoveKey removes the key and the files corresponding to it.
// It returns nil if it is successful.
func (db *MiniDB) RemoveKey(key string) error {
	d := db.getOrCreateMutex("delete_key" + key)
	d.Lock()
	defer d.Unlock()

	// get the filename if it exists
	filename, ok := db.store.Keys[key]
	if !ok {
		return errors.New("collections key does not exist")
	}

	// remove the key and the filename
	delete(db.store.Keys, key)

	return os.RemoveAll(path.Join(db.path, filename))
}
