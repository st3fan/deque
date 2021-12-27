package deque

type DequeInterface[T any] interface {
	AddFirst(e T) error
	AddLast(e T) error

	PeekFirst() (T, bool)
	PeekLast() (T, bool)

	RemoveFirst() (T, bool)
	RemoveLast() (T, bool)

	Length() int
	Clear()
}

const (
	blockSize   = 64
	blockCenter = ((blockSize - 1) / 2)
)

type block[T any] struct {
	left  *block[T]
	data  [blockSize]T
	right *block[T]
}

func newBlock[T any]() *block[T] {
	return &block[T]{}
}

type Deque[T any] struct {
	leftBlock  *block[T]
	rightBlock *block[T]
	leftIndex  int
	rightIndex int
	size       int
	state      int // TODO What is this
}

// Python also keeps N blocks around

//

func NewDeque[T any]() *Deque[T] {
	b := newBlock[T]()
	return &Deque[T]{
		leftBlock:  b,
		rightBlock: b,
		leftIndex:  blockCenter + 1,
		rightIndex: blockCenter,
		size:       0,
		state:      0,
	}
}

func (d *Deque[T]) needsTrim() bool {
	return false
}

func (d *Deque[T]) AddFirst(e T) error {
	if d.leftIndex == 0 {
		b := newBlock[T]()
		b.right = d.leftBlock
		d.leftBlock.left = b
		d.leftBlock = b
		d.leftIndex = blockSize
	}

	d.size++
	d.leftIndex--
	d.leftBlock.data[d.leftIndex] = e

	if d.needsTrim() {
		d.RemoveLast()
	} else {
		d.state++
	}

	return nil
}

func (d *Deque[T]) AddLast(e T) error {
	if d.rightIndex == blockSize-1 {
		b := newBlock[T]()
		b.left = d.rightBlock
		d.rightBlock.right = b
		d.rightBlock = b
		d.rightIndex = -1
	}

	d.size++
	d.rightIndex++
	d.rightBlock.data[d.rightIndex] = e

	if d.needsTrim() {
		d.RemoveFirst()
	} else {
		d.state++
	}

	return nil
}

func (d *Deque[T]) PeekFirst() (T, bool) {
	if d.size == 0 {
		var empty T
		return empty, false
	}

	return d.leftBlock.data[d.leftIndex], true
}

func (d *Deque[T]) PeekLast() (T, bool) {
	if d.size == 0 {
		var empty T
		return empty, false
	}

	return d.rightBlock.data[d.rightIndex], true
}

func (d *Deque[T]) RemoveFirst() (T, bool) {
	if d.size == 0 {
		var empty T
		return empty, false
	}

	element := d.leftBlock.data[d.leftIndex]

	d.leftIndex++
	d.size--
	d.state++

	if d.leftIndex == blockSize {
		if d.leftIndex == blockSize {
			if d.size != 0 {
				d.leftBlock = d.leftBlock.right
				d.leftIndex = 0
			} else {
				// re-center instead of freeing a block
				d.leftIndex = blockCenter + 1
				d.rightIndex = blockCenter
			}
		}
	}

	return element, true
}

func (d *Deque[T]) RemoveLast() (T, bool) {
	if d.size == 0 {
		var empty T
		return empty, false
	}

	item := d.rightBlock.data[d.rightIndex]
	d.rightIndex--

	d.size--
	d.state++

	if d.rightIndex < 0 {
		if d.size != 0 {
			d.rightBlock = d.rightBlock.left
			d.rightIndex = blockSize - 1
		} else {
			// re-center instead of freeing a block
			d.leftIndex = blockCenter + 1
			d.rightIndex = blockCenter
		}
	}

	return item, true
}

func (d *Deque[T]) Length() int {
	return d.size
}

func (d *Deque[T]) Clear() {
	d.leftBlock = nil
	d.rightBlock = nil
	d.leftIndex = 0
	d.rightIndex = 0
	d.size = 0
	d.state = 0
}
