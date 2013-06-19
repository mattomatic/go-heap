package heap

type Item interface {
	Less(than Item) bool
}

type Heap struct {
	data []Item
	min  bool
}

// Create a new min heap
func NewMinHeap() *Heap {
	return &Heap{make([]Item, 0), true}
}

// Create a new max heap
func NewMaxHeap() *Heap {
	return &Heap{make([]Item, 0), false}
}

// Get the number of elements in the heap
func (h *Heap) Size() int {
	return len(h.data)
}

// Insert an element into the heap
func (h *Heap) Insert(x Item) {
	h.data = append(h.data, x)
	h.siftUp(h.Size() - 1)
}

// Get the smallest element in the heap
func (h *Heap) Peek() Item {
	return h.data[0]
}

// Remove the smallest element from the heap
func (h *Heap) Pop() Item {
	x := h.data[0]
	h.swap(0, h.Size()-1)
	h.trim()
	h.siftDown(0)
	return x
}

func (h *Heap) siftUp(c int) {
	p := parent(c)

	if !h.invariantHolds(p, c) {
		h.swap(p, c)
		h.siftUp(p)
	}
}

func (h *Heap) siftDown(p int) {
	c1 := lhsChild(p)
	c2 := rhsChild(p)

	if !(h.invariantHolds(p, c1) && h.invariantHolds(p, c2)) {
		if c2 >= h.Size() || h.less(c1, c2) {
			h.swap(p, c1)
			h.siftDown(c1)
		} else {
			h.swap(p, c2)
			h.siftDown(c2)
		}
	}
}

func (h *Heap) invariantHolds(p int, c int) bool {
	// all children must be greater than or equal to their parents
	// <==>
	// all children must not be less than their parents
	// note: if there are no children, then the invariant holds
	return c >= h.Size() || !h.less(c, p)
}

func parent(x int) int {
	if x == 0 {
		return 0
	}

	return halve(x - 1)
}

func lhsChild(x int) int {
	return double(x) + 1
}

func rhsChild(x int) int {
	return double(x) + 2
}

func double(x int) int {
	return x << 1
}

func halve(x int) int {
	return x >> 1
}

func (h *Heap) swap(x int, y int) {
	t := h.data[x]
	h.data[x] = h.data[y]
	h.data[y] = t
}

func (h *Heap) less(x int, y int) bool {
	a := h.data[x]
	b := h.data[y]

	if h.min {
		return a.Less(b)
	} else {
		return greaterThan(a, b)
	}
}

func greaterThan(x Item, y Item) bool {
	return !x.Less(y) && !equalTo(x, y)
}

func equalTo(x Item, y Item) bool {
	return !x.Less(y) && !y.Less(x)
}

func (h *Heap) trim() {
	h.data = h.data[:len(h.data)-1]
}
