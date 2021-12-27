package deque

import "sync"

type SynchronizedDeque[T any] struct {
	mutex sync.Mutex
	deque DequeInterface[T]
}

func NewSynchronizedDeque[T any](deque DequeInterface[T]) *SynchronizedDeque[T] {
	return &SynchronizedDeque[T]{deque: deque}
}

func (d *SynchronizedDeque[T]) AddFirst(e T) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.deque.AddFirst(e)
}

func (d *SynchronizedDeque[T]) AddLast(e T) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.deque.AddLast(e)
}

func (d *SynchronizedDeque[T]) PeekFirst() (T, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.deque.PeekFirst()
}

func (d *SynchronizedDeque[T]) PeekLast() (T, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.deque.PeekLast()
}

func (d *SynchronizedDeque[T]) RemoveFirst() (T, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.deque.RemoveFirst()
}

func (d *SynchronizedDeque[T]) RemoveLast() (T, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.deque.RemoveLast()
}

func (d *SynchronizedDeque[T]) Length() int {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.deque.Length()
}

func (d *SynchronizedDeque[T]) Clear() {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.deque.Clear()
}
