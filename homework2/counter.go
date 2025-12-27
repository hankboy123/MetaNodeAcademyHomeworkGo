package main

import "sync"

type Counter struct {
	value int
}

var mu sync.Mutex

func (c *Counter) Increment() {
	mu.Lock()
	c.value++
	defer mu.Unlock()
}

func (c *Counter) Decrement() {
	mu.Lock()
	c.value--
	defer mu.Unlock()
}

func (c *Counter) Value() int {
	return c.value
}
