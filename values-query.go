package minidb

// GetBool finds the key with bool value and returns if exits.
func (db *MiniDB) GetBool(key string) bool {
	defer recoverAssertion()

	return db.getValue(key).(bool)
}

// GetString finds the key with the string value and returns if exists.
func (db *MiniDB) GetString(key string) string {
	defer recoverAssertion()

	return db.getValue(key).(string)
}

// GetInt finds the key with the int value and returns if exists.
func (db *MiniDB) GetInt(key string) int {
	defer recoverAssertion()

	return db.getValue(key).(int)
}

// GetFloat32 finds the key with the float32 value and returns if exists.
func (db *MiniDB) GetFloat32(key string) float32 {
	defer recoverAssertion()

	return db.getValue(key).(float32)
}

// GetFloat64 finds the key with the float64 value and returns if exists.
func (db *MiniDB) GetFloat64(key string) float64 {
	defer recoverAssertion()

	return db.getValue(key).(float64)
}

// GetBoolSlice finds the key with the []bool value and returns if exits.
func (db *MiniDB) GetBoolSlice(key string) []bool {
	defer recoverAssertion()

	return db.getValue(key).([]bool)
}

// GetStringSlice finds the key with the []string value and returns if exists.
func (db *MiniDB) GetStringSlice(key string) []string {
	defer recoverAssertion()

	return db.getValue(key).([]string)
}

// GetIntSlice finds the key with the []int value and returns if exists.
func (db *MiniDB) GetIntSlice(key string) []int {
	defer recoverAssertion()

	return db.getValue(key).([]int)
}

// GetFloat32Slice finds the key with the []float32 value and returns if exists.
func (db *MiniDB) GetFloat32Slice(key string) []float32 {
	defer recoverAssertion()

	return db.getValue(key).([]float32)
}

// GetFloat64Slice finds the key with the []float64 value and returns if exists.
func (db *MiniDB) GetFloat64Slice(key string) []float64 {
	defer recoverAssertion()

	return db.getValue(key).([]float64)
}

// Get finds the key and returns an interface value.
func (db *MiniDB) Get(key string) interface{} {
	return db.getValue(key)
}
