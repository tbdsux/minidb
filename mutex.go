package minidb

import "sync"

// referred from ::> https://github.com/sdomino/scribble/blob/master/scribble.go#L254
func (db *MiniDB) getOrCreateMutex(key string) *sync.Mutex {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	m, ok := db.mutexes[key]

	// if the mutex doesn't exist make it
	if !ok {
		m = &sync.Mutex{}
		db.mutexes[key] = m
	}

	return m
}

// referred from ::> https://github.com/sdomino/scribble/blob/master/scribble.go#L254
func (db *MiniCollections) getOrCreateMutex(key int) *sync.Mutex {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	m, ok := db.mutexes[key]

	// if the mutex doesn't exist make it
	if !ok {
		m = &sync.Mutex{}
		db.mutexes[key] = m
	}

	return m
}

// referred from ::> https://github.com/sdomino/scribble/blob/master/scribble.go#L254
func (db *MiniStore) getOrCreateMutex(key string) *sync.Mutex {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	m, ok := db.mutexes[key]

	// if the mutex doesn't exist make it
	if !ok {
		m = &sync.Mutex{}
		db.mutexes[key] = m
	}

	return m
}
