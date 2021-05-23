package minidb

import "errors"

// Update updates the key's value. It returns nil if updated.
func (db *MiniStore) Update(key string, v interface{}) error {
	defer db.writeToDB()

	if _, ok := db.store[key]; !ok {
		return errors.New("unknown key")
	}

	db.store[key] = v

	return nil
}

// Remove attemps to remove the key from the db if it exists.
// It returns nil if it is removed
func (db *MiniStore) Remove(key string) error {
	defer db.writeToDB()

	if _, ok := db.store[key]; !ok {
		return errors.New("key does not exists")
	}

	// remove
	delete(db.store, key)

	return nil
}

// GetBool finds the key with bool value and returns if exits.
func (db *MiniStore) GetBool(key string) bool {
	defer recoverAssertion()

	return db.getValue(key).(bool)
}

// GetString finds the key with the string value and returns if exists.
func (db *MiniStore) GetString(key string) string {
	defer recoverAssertion()

	return db.getValue(key).(string)
}

// GetInt finds the key with the int value and returns if exists.
func (db *MiniStore) GetInt(key string) int {
	defer recoverAssertion()

	return db.getValue(key).(int)
}

// GetFloat32 finds the key with the float32 value and returns if exists.
func (db *MiniStore) GetFloat32(key string) float32 {
	defer recoverAssertion()

	return db.getValue(key).(float32)
}

// GetFloat64 finds the key with the float64 value and returns if exists.
func (db *MiniStore) GetFloat64(key string) float64 {
	defer recoverAssertion()

	return db.getValue(key).(float64)
}

// GetBoolSlice finds the key with the []bool value and returns if exits.
func (db *MiniStore) GetBoolSlice(key string) []bool {
	defer recoverAssertion()

	return db.getValue(key).([]bool)
}

// GetStringSlice finds the key with the []string value and returns if exists.
func (db *MiniStore) GetStringSlice(key string) []string {
	defer recoverAssertion()

	return db.getValue(key).([]string)
}

// GetIntSlice finds the key with the []int value and returns if exists.
func (db *MiniStore) GetIntSlice(key string) []int {
	defer recoverAssertion()

	return db.getValue(key).([]int)
}

// GetFloat32Slice finds the key with the []float32 value and returns if exists.
func (db *MiniStore) GetFloat32Slice(key string) []float32 {
	defer recoverAssertion()

	return db.getValue(key).([]float32)
}

// GetFloat64Slice finds the key with the []float64 value and returns if exists.
func (db *MiniStore) GetFloat64Slice(key string) []float64 {
	defer recoverAssertion()

	return db.getValue(key).([]float64)
}

// Get finds the key and returns an interface value.
func (db *MiniStore) Get(key string) interface{} {
	return db.getValue(key)
}
