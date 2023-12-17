package util

import "testing"

func TestNewPriorityQueue(t *testing.T) {
	type crucible struct {
		heatLoss int
	}

	pq := NewPriorityQueue(func(itemA, itemB crucible) int { return itemA.heatLoss - itemB.heatLoss })

	pq.Push(crucible{heatLoss: 10})
	pq.Push(crucible{heatLoss: 20})
	pq.Push(crucible{heatLoss: 1})
	pq.Push(crucible{heatLoss: 10})
	pq.Push(crucible{heatLoss: 6})
	pq.Push(crucible{heatLoss: 30})
	pq.Push(crucible{heatLoss: 2})
	pq.Push(crucible{heatLoss: 15})

	if pq.Length() != 8 {
		t.Fatalf("wrong priority queue size, expected %d but got %d\n", 8, pq.Length())
	}

	item := pq.Pop()
	if item.heatLoss != 1 {
		t.Fatalf("got wrong item of queue, expected value %d but got %d\n", 1, item.heatLoss)
	}

	item = pq.Pop()
	if item.heatLoss != 2 {
		t.Fatalf("got wrong item of queue, expected value %d but got %d\n", 2, item.heatLoss)
	}

	item = pq.Pop()
	if item.heatLoss != 6 {
		t.Fatalf("got wrong item of queue, expected value %d but got %d\n", 6, item.heatLoss)
	}

	item = pq.Pop()
	if item.heatLoss != 10 {
		t.Fatalf("got wrong item of queue, expected value %d but got %d\n", 10, item.heatLoss)
	}

	item = pq.Pop()
	if item.heatLoss != 10 {
		t.Fatalf("got wrong item of queue, expected value %d but got %d\n", 10, item.heatLoss)
	}

	item = pq.Pop()
	if item.heatLoss != 15 {
		t.Fatalf("got wrong item of queue, expected value %d but got %d\n", 15, item.heatLoss)
	}

	item = pq.Pop()
	if item.heatLoss != 20 {
		t.Fatalf("got wrong item of queue, expected value %d but got %d\n", 20, item.heatLoss)
	}

	item = pq.Pop()
	if item.heatLoss != 30 {
		t.Fatalf("got wrong item of queue, expected value %d but got %d\n", 30, item.heatLoss)
	}
}
