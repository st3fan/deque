package deque_test

import (
	"testing"

	"github.com/st3fan/deque"
)

func Test_SynchronizedDequeImplementsDeque(t *testing.T) {
	var _ deque.DequeInterface[int] = &deque.SynchronizedDeque[int]{}
}
