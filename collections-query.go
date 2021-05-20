package minidb

func (c *MiniCollections) Push(v interface{}) {
	d := c.getOrCreateMutex(len(c.store) + 1)
	d.Lock()
	defer d.Unlock()

	c.store = append(c.store, v)
	c.writeToDB()
}
