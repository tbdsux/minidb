package minidb

import "strings"

// Push adds an item to the store slice.
func (c *MiniCollections) Push(v interface{}) {
	d := c.getOrCreateMutex(len(c.store) + 1)
	d.Lock()
	defer d.Unlock()

	c.store = append(c.store, v)
	c.writeToDB()
}

// List returns all of the contents of the collection
func (c *MiniCollections) List() []interface{} {
	return c.store
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

// FindAll returns all values that matches v.
func (c *MiniCollections) FindAll(v interface{}) []interface{} {
	values := []interface{}{}

	for _, value := range c.store {
		if value == v {
			values = append(values, value)
		}
	}

	return values
}

// MatchString returns the first element that contains v.
func (c *MiniCollections) MatchString(v string) (string, error) {
	for _, value := range c.store {
		s, ok := value.(string)
		if ok {
			if strings.Contains(s, v) {
				return s, nil
			}
		}
	}

	return "", nil
}

// MatchStringAll returns the first element that contains v.
func (c *MiniCollections) MatchStringAll(v string) ([]string, error) {
	values := []string{}

	for _, value := range c.store {
		s, ok := value.(string)
		if ok {
			if strings.Contains(s, v) {
				values = append(values, s)
			}
		}
	}

	return values, nil
}

// FilterString returns all string elements.
func (c *MiniCollections) FilterString() []string {
	values := []string{}

	for _, i := range c.store {
		if v, ok := i.(string); ok {
			values = append(values, v)
		}
	}

	return values
}

// FilterInt returns all int elements.
func (c *MiniCollections) FilterInt() []int {
	values := []int{}

	for _, i := range c.store {
		if v, ok := i.(int); ok {
			values = append(values, v)
		}
	}

	return values
}

// FilterBool returns all bool elements.
func (c *MiniCollections) FilterBool() []bool {
	values := []bool{}

	for _, i := range c.store {
		if v, ok := i.(bool); ok {
			values = append(values, v)
		}
	}

	return values
}

// FilterFloat returns all float elements (uses float64). Some integers (int) might be included.
func (c *MiniCollections) FilterFloat() []float64 {
	values := []float64{}

	for _, i := range c.store {
		if v, ok := i.(float64); ok {
			values = append(values, v)
		}
	}

	return values
}
