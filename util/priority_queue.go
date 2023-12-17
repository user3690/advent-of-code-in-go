package util

type PriorityQueue[T comparable] struct {
	items      []T
	comparator func(itemA, itemB T) int
}

func NewPriorityQueue[T comparable](comparator func(itemA, itemB T) int) PriorityQueue[T] {
	return PriorityQueue[T]{
		comparator: comparator,
	}
}

func (pq *PriorityQueue[T]) Push(item T) {
	pq.items = append(pq.items, item)
	pq.up()
}

func (pq *PriorityQueue[T]) Pop() (item T) {
	firstElement := pq.items[0]
	LastElementIndex := len(pq.items) - 1

	pq.items[0], pq.items[LastElementIndex] = pq.items[LastElementIndex], pq.items[0]
	pq.items = pq.items[:LastElementIndex]
	pq.down()

	return firstElement
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.items) == 0
}

func (pq *PriorityQueue[T]) Length() int {
	return len(pq.items)
}

func (pq *PriorityQueue[T]) up() {
	index := len(pq.items) - 1

	for index > 0 {
		parentIndex := (index - 1) / 2
		if pq.comparator(pq.items[index], pq.items[parentIndex]) < 0 {
			pq.items[index], pq.items[parentIndex] = pq.items[parentIndex], pq.items[index]
			index = parentIndex
		} else {
			break
		}
	}
}

func (pq *PriorityQueue[T]) down() {
	var index int

	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		if leftChild < len(pq.items) && pq.comparator(pq.items[leftChild], pq.items[smallest]) < 0 {
			smallest = leftChild
		}

		if rightChild < len(pq.items) && pq.comparator(pq.items[rightChild], pq.items[smallest]) < 0 {
			smallest = rightChild
		}

		if smallest != index {
			pq.items[index], pq.items[smallest] = pq.items[smallest], pq.items[index]
			index = smallest
		} else {
			break
		}
	}
}
