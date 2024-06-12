package lib

import "sync"

type Map[T any] struct {
	container map[string]*T
	lock      sync.RWMutex
}

func (c *Map[T]) Load(name string) *T {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if c.container == nil {
		return nil
	}
	return c.container[name]
}

func (c *Map[T]) Store(name string, value *T) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.container == nil {
		c.container = make(map[string]*T)
	}
	c.container[name] = value
}

func (c *Map[T]) Range(iterator func(name string, item *T) bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if c.container == nil {
		return
	}
	for k, v := range c.container {
		if !iterator(k, v) {
			break
		}
	}
}

func (c *Map[T]) Delete(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.container, name)
}

func (c *Map[T]) DeleteRaw(name string) {
	delete(c.container, name)
}

func (c *Map[T]) Len() int {
	return len(c.container)
}
