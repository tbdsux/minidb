package minidb

// Push adds an item to the store slice.
func (c *MiniCollections) Push(v interface{}) {
	d := c.getOrCreateMutex(len(c.store) + 1)
	d.Lock()
	defer d.Unlock()

	c.store = append(c.store, v)
	c.writeToDB()
}

// First returns the first element of the collections.
func (c *MiniCollections) First() interface{} {
	return c.store[0]
}

// Last returns the last element of the collections.
func (c *MiniCollections) Last() interface{} {
	return c.store[len(c.store)-1]
}

// Find attemps to find the value from the store and returns it's index.
// If it doesn't exist, it will return -1.
func (c *MiniCollections) Find(v interface{}) int {
	for index, value := range c.store {
		if value == v {
			return index
		}
	}

	return -1
}
