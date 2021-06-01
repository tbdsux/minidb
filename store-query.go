package minidb

import "errors"

// Update updates the key's value. It returns nil if updated.
func (db *MiniStore) Update(key string, v interface{}) error {
	d := db.getOrCreateMutex("store_update_" + key)
	d.Lock()
	defer d.Unlock()

	if _, ok := db.content[key]; !ok {
		return errors.New("unknown key")
	}

	db.content[key] = v

	db.writeToDB()

	return nil
}

// Remove attemps to remove the key from the db if it exists.
// It returns nil if it is removed
func (db *MiniStore) Remove(key string) error {
	d := db.getOrCreateMutex("store_remove_" + key)
	d.Lock()
	defer d.Unlock()

	if _, ok := db.content[key]; !ok {
		return errors.New("key does not exists")
	}

	// remove
	delete(db.content, key)

	db.writeToDB()

	return nil
}

// GetBool finds the key with bool value and returns if exits.
// It panics if there is an error in type assertion.
func (db *MiniStore) GetBool(key string) bool {
	return db.getValue(key).(bool)
}

// GetString finds the key with the string value and returns if exists.
// It panics if there is an error in type assertion.
func (db *MiniStore) GetString(key string) string {
	return db.getValue(key).(string)
}

// GetInt finds the key with the int value and returns if exists.
// It panics if there is an error in type assertion.
func (db *MiniStore) GetInt(key string) int {
	return db.getValue(key).(int)
}

// GetFloat32 finds the key with the float32 value and returns if exists.
// It panics if there is an error in type assertion.
func (db *MiniStore) GetFloat32(key string) float32 {
	return db.getValue(key).(float32)
}

// GetFloat64 finds the key with the float64 value and returns if exists.
// It panics if there is an error in type assertion.
func (db *MiniStore) GetFloat64(key string) float64 {
	return db.getValue(key).(float64)
}

// GetBoolSlice finds the key with the []bool value and returns if exits.
// It panics if there is an error in type assertion.
func (db *MiniStore) GetBoolSlice(key string) []bool {
	return db.getValue(key).([]bool)
}

// GetStringSlice finds the key with the []string value and returns if exists.
// It panics if there is an error in type assertion.
func (db *MiniStore) GetStringSlice(key string) []string {
	return db.getValue(key).([]string)
}

// GetIntSlice finds the key with the []int value and returns if exists.
// It panics if there is an error in type assertion.
func (db *MiniStore) GetIntSlice(key string) []int {
	return db.getValue(key).([]int)
}

// GetFloat32Slice finds the key with the []float32 value and returns if exists.
// It panics if there is an error in type assertion.
func (db *MiniStore) GetFloat32Slice(key string) []float32 {
	return db.getValue(key).([]float32)
}

// GetFloat64Slice finds the key with the []float64 value and returns if exists.
// It panics if there is an error in type assertion.
func (db *MiniStore) GetFloat64Slice(key string) []float64 {
	return db.getValue(key).([]float64)
}

// Get finds the key and returns an interface value.
// It also returns false if the key does not exist and true, otherwise.
func (db *MiniStore) Get(key string) (interface{}, bool) {
	return db.getValueOK(key)
}

// IsExists asserts if the key exists. You should use Get() if you want to get
// the Raw value and if it exists.
func (db *MiniStore) IsExists(key string) bool {
	_, ok := db.getValueOK(key)
	return ok
}

// List returns the content of db.store
func (db *MiniStore) List() map[string]interface{} {
	return db.content
}
