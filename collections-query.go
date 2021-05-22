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
