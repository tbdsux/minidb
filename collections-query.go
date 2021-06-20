package minidb

import (
	"errors"
	"strings"
)

// Remove removes the first element that is equal to v. It wraps around RemoveMany().
func (c *MiniCollections) Remove(v interface{}) error {
	return c.RemoveMany(v, 1)
}

// RemoveAll removes all elements that are equal to v. It wraps around RemoveMany().
func (c *MiniCollections) RemoveAll(v interface{}) error {
	return c.RemoveMany(v, -1)
}

// RemoveMany removes the number of elements corresponding to l. Use RemoveAll() for removing
// all elements and Remove() for a single element.
// `l` cannot be less than -1 or equal to 0.
func (c *MiniCollections) RemoveMany(v interface{}, l int) error {
	if l < -1 || l == 0 {
		return errors.New("l cannot be less than -1 or equal to 0")
	}

	d := c.getOrCreateMutex(len(c.content) + l)
	d.Lock()
	defer d.Unlock()

	values := []interface{}{}
	removed := []interface{}{}

	for index, x := range c.content {
		// -1 means remove all similar values
		if l > 0 {
			// check if total is similar
			if len(removed) == l {
				values = append(values, c.content[index:]...)
				break
			}
		}

		// if x is same with v, append it to removed
		// otherwise, to the new values
		if v == x {
			removed = append(removed, x)
		} else {
			values = append(values, x)
		}
	}

	// write new values
	c.content = values
	c.writeToDB()

	return nil
}

// Push adds an item to the store slice.
func (c *MiniCollections) Push(v ...interface{}) {
	d := c.getOrCreateMutex(len(c.content) + 1)
	d.Lock()
	defer d.Unlock()

	c.content = append(c.content, v...)
	c.writeToDB()
}

// List returns all of the contents of the collection
func (c *MiniCollections) List() []interface{} {
	return c.content
}

// First returns the first element of the collections.
func (c *MiniCollections) First() interface{} {
	return c.content[0]
}

// Last returns the last element of the collections.
func (c *MiniCollections) Last() interface{} {
	return c.content[len(c.content)-1]
}

// Find attemps to find the value from the store and returns it's index.
// If it doesn't exist, it will return -1.
func (c *MiniCollections) Find(v interface{}) int {
	for index, value := range c.content {
		if value == v {
			return index
		}
	}

	return -1
}

// FindAll returns all values that matches v.
func (c *MiniCollections) FindAll(v interface{}) []interface{} {
	values := []interface{}{}

	for _, value := range c.content {
		if value == v {
			values = append(values, value)
		}
	}

	return values
}

// MatchString returns the first element that contains v.
func (c *MiniCollections) MatchString(v string) (string, error) {
	for _, value := range c.content {
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

	for _, value := range c.content {
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

	for _, i := range c.content {
		if v, ok := i.(string); ok {
			values = append(values, v)
		}
	}

	return values
}

// FilterInt returns all int elements.
func (c *MiniCollections) FilterInt() []int {
	values := []int{}

	for _, i := range c.content {
		if v, ok := i.(int); ok {
			values = append(values, v)
		}
	}

	return values
}

// FilterBool returns all bool elements.
func (c *MiniCollections) FilterBool() []bool {
	values := []bool{}

	for _, i := range c.content {
		if v, ok := i.(bool); ok {
			values = append(values, v)
		}
	}

	return values
}

// FilterFloat returns all float elements (uses float64). Some integers (int) might be included.
func (c *MiniCollections) FilterFloat() []float64 {
	values := []float64{}

	for _, i := range c.content {
		if v, ok := i.(float64); ok {
			values = append(values, v)
		}
	}

	return values
}

// Filter accepts a function that return the values that returns true with it. // TODO
func (c *MiniCollections) Filter(f func(interface{}) bool) []interface{} {
	values := []interface{}{}

	for _, i := range c.content {
		if f(i) {
			values = append(values, i)
		}
	}

	return values
}

// FilterStringFunc returns all string elements.
func (c *MiniCollections) FilterStringFunc(f func(x string) bool) []string {
	values := []string{}

	for _, i := range c.content {
		if v, ok := i.(string); ok {
			if f(v) {
				values = append(values, v)
			}
		}
	}

	return values
}

// FilterIntFunc returns all int elements.
func (c *MiniCollections) FilterIntFunc(f func(x int) bool) []int {
	values := []int{}

	for _, i := range c.content {
		if v, ok := i.(int); ok {
			if f(v) {
				values = append(values, v)
			}
		}
	}

	return values
}

// FilterBoolFunc returns all bool elements.
func (c *MiniCollections) FilterBoolFunc(f func(x bool) bool) []bool {
	values := []bool{}

	for _, i := range c.content {
		if v, ok := i.(bool); ok {
			if f(v) {
				values = append(values, v)
			}
		}
	}

	return values
}

// FilterFloatFunc returns all float elements (uses float64). Some integers (int) might be included.
func (c *MiniCollections) FilterFloatFunc(f func(x float64) bool) []float64 {
	values := []float64{}

	for _, i := range c.content {
		if v, ok := i.(float64); ok {
			if f(v) {
				values = append(values, v)
			}
		}
	}

	return values
}
