package util

import "testing"

func TestNewFiFoQueue(t *testing.T) {
	type crucible struct {
		heatLoss int
	}

	q := FiFoQueue[crucible]{}

	q.Push(crucible{heatLoss: 10})
	q.Push(crucible{heatLoss: 20})
	q.Push(crucible{heatLoss: 1})

	if q.Length() != 3 {
		t.Fatalf("wrong priority queue size, expected %d but got %d\n", 8, q.Length())
	}

	item := q.Pop()
	if item.heatLoss != 10 {
		t.Fatalf("got wrong item of queue, expected value %d but got %d\n", 10, item.heatLoss)
	}

	item = q.Pop()
	if item.heatLoss != 20 {
		t.Fatalf("got wrong item of queue, expected value %d but got %d\n", 20, item.heatLoss)
	}

	item = q.Pop()
	if item.heatLoss != 1 {
		t.Fatalf("got wrong item of queue, expected value %d but got %d\n", 1, item.heatLoss)
	}

	if !q.IsEmpty() {
		t.Fatalf("queue not empty, got %d items\n", q.Length())
	}
}
