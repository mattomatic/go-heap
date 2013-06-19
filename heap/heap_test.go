package heap

import (
	"math/rand"
	"testing"
)

func TestCreate(t *testing.T) {
	h := NewMinHeap()

	if h == nil {
		t.Error()
	}

	h2 := NewMaxHeap()

	if h2 == nil {
		t.Error()
	}
}

func TestLess(t *testing.T) {
	a := Int(1)
	b := Int(1)

	if a.Less(b) {
		t.Error()
	}

	c := Int(2)

	if c.Less(a) {
		t.Error()
	}
}

func TestInsert(t *testing.T) {
	h := NewMinHeap()
	h.Insert(Int(2))
	h.Insert(Int(3))
	h.Insert(Int(1))
	h.Insert(Int(4))

	if h.Peek() != Int(1) {
		t.Error()
	}
}

func TestRemove(t *testing.T) {
	h := NewMinHeap()
	h.Insert(Int(2))
	h.Insert(Int(3))
	h.Insert(Int(1))
	h.Insert(Int(4))
	h.Pop()
	h.Pop()
	h.Pop()
	h.Pop()
}

func TestRandom(t *testing.T) {
	h := NewMinHeap()
	for i := 0; i < 500; i++ {
		x := Int(rand.Int())
		h.Insert(x)
	}

	x := h.Peek()
	for h.Size() > 0 {
		if h.Peek().Less(x) {
			t.Error()
		}

		x = h.Pop()
	}
}

func TestRandomMax(t *testing.T) {
	h := NewMaxHeap()
	for i := 0; i < 500; i++ {
		h.Insert(Int(rand.Intn(100)))
	}

	x := h.Peek()
	for h.Size() > 0 {
		if greaterThan(h.Peek(), x) {
			t.Error()
		}

		x = h.Pop()
	}
}

func TestParentChild(t *testing.T) {
	for i := 0; i < 50; i++ {
		c1 := lhsChild(i)
		c2 := rhsChild(i)

		if c1 == c2 {
			t.Error()
		}

		if parent(c1) != i || parent(c2) != i {
			t.Error()
		}
	}
}

func BenchmarkDouble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		double(i)
	}
}

func BenchmarkHalve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		halve(i)
	}
}
