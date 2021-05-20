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
func (c *MiniCollections) getOrCreateMutex(key int) *sync.Mutex {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	m, ok := c.mutexes[key]

	// if the mutex doesn't exist make it
	if !ok {
		m = &sync.Mutex{}
		c.mutexes[key] = m
	}

	return m
}
