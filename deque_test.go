package deque_test

import (
	"testing"

	"github.com/st3fan/deque"
)

func Test_DequeImplementsDequeInterface(t *testing.T) {
	var _ deque.DequeInterface[int] = &deque.Deque[int]{}
}

func Test_AddRemoveLast(t *testing.T) {
	d := deque.NewDeque[int]()

	for i := 0; i < 2500; i++ {
		d.AddLast(i)
		if d.Length() != (i + 1) {
			t.Fail()
		}
	}

	for i := 0; i < 2500; i++ {
		if e, ok := d.RemoveLast(); !ok || e != (2500-i-1) {
			t.Fail()
		}
		if d.Length() != (2500 - i - 1) {
			t.Fail()
		}
	}
}

func Test_AddRemoveFirst(t *testing.T) {
	d := deque.NewDeque[int]()

	for i := 0; i < 2500; i++ {
		d.AddFirst(i)
		if d.Length() != (i + 1) {
			t.Fail()
		}
	}

	for i := 0; i < 2500; i++ {
		if e, ok := d.RemoveFirst(); !ok || e != (2500-i-1) {
			t.Fail()
		}
		if d.Length() != (2500 - i - 1) {
			t.Fail()
		}
	}
}

func Test_AddRemove(t *testing.T) {
	d := deque.NewDeque[int]()

	for i := 0; i < 2500; i++ {
		if i%2 == 0 {
			d.AddFirst(i)
		} else {
			d.AddLast(i)
		}
		if d.Length() != (i + 1) {
			t.Fail()
		}
	}

	// for i := 0; i < 2500; i++ {
	// 	if e, ok := d.RemoveFirst(); ok != nil || e != (2500-i-1) {
	// 		t.Fail()
	// 	}
	// 	if d.Length() != (2500 - i - 1) {
	// 		t.Fail()
	// 	}
	// }
}

func Test_LeftAndRight(t *testing.T) {
	d := deque.NewDeque[string]()

	d.AddFirst("a")
	d.AddLast("b")
	d.AddLast("c")

	if e, ok := d.RemoveLast(); !ok || e != "c" {
		t.Fail()
	}
	if e, ok := d.RemoveLast(); !ok || e != "b" {
		t.Fail()
	}
	if e, ok := d.RemoveLast(); !ok || e != "a" {
		t.Fail()
	}
}

func Test_Clear(t *testing.T) {
	d := deque.NewDeque[int]()
	d.AddFirst(42)
	if d.Length() != 1 {
		t.Fail()
	}
	d.Clear()
	if d.Length() != 0 {
		t.Fail()
	}
}

func Benchmark_AddRemoveOne(b *testing.B) {
	b.StopTimer()
	d := deque.NewDeque[int]()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		d.AddFirst(i)
		d.RemoveFirst()
	}
}

func Benchmark_AddRemoveManyFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		d := deque.NewDeque[int]()
		b.StartTimer()
		for n := 0; n < 40; n++ {
			d.AddFirst(n)
		}
	}
}

func Benchmark_AddRemoveManyLast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		d := deque.NewDeque[int]()
		b.StartTimer()
		for n := 0; n < 40; n++ {
			d.AddLast(n)
		}
	}
}

func Benchmark_AddRemoveManyLastWithStruct(b *testing.B) {
	type thing struct {
		value int
	}
	b.StopTimer()
	d := deque.NewDeque[*thing]()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d.AddLast(&thing{value: 42})
		d.RemoveLast()
	}
}

func Benchmark_AddRemoveManyLastWithPrimitive(b *testing.B) {
	b.StopTimer()
	d := deque.NewDeque[int]()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d.AddLast(42)
		d.RemoveLast()
	}
}
