package heap

import (
	"testing"
	"math/rand"
)

func TestCreate(t *testing.T) {
	h := New()

	if h == nil {
		t.Error()
	}
}

func TestInvariant(t *testing.T) {
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
	h := New()
	h.Insert(Int(2))
	h.Insert(Int(3))
	h.Insert(Int(1))
	h.Insert(Int(4))
	
	if h.Min() != Int(1) {
	    t.Error()
	}	
}

func TestRemove(t *testing.T) {
	h := New()
	h.Insert(Int(2))
	h.Insert(Int(3))
	h.Insert(Int(1))
	h.Insert(Int(4))
	h.Remove()
	h.Remove()
	h.Remove()
	h.Remove()
}

func TestRandom(t *testing.T) {
    h := New()
    for i:=0; i<500; i++ {
        x := Int(rand.Int())
        h.Insert(x)
    }
    
    x := h.Min()
    for h.Size() > 0 {
        if h.Min().Less(x) {
            t.Error()
        }
        
        x = h.Remove()
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