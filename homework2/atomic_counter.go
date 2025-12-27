package main

import "sync/atomic"

type AtomicCounter struct {
	value int64
}

func (ac *AtomicCounter) Increment() {
	// 使用原子操作增加计数器的值
	atomic.AddInt64(&ac.value, 1)
}

func (ac *AtomicCounter) Decrement() {
	// 使用原子操作减少计数器的值
	atomic.AddInt64(&ac.value, -1)
}

func (ac *AtomicCounter) Value() int64 {
	// 使用原子操作获取计数器的值
	return atomic.LoadInt64(&ac.value)
}
