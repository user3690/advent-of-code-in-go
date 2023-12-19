package util

type FiFoQueue[T comparable] struct {
	items []T
}

func (q *FiFoQueue[T]) Push(item T) {
	q.items = append(q.items, item)
}

func (q *FiFoQueue[T]) Pop() T {
	firstItem := q.items[0]
	q.items = q.items[1:]

	return firstItem
}

func (q *FiFoQueue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *FiFoQueue[T]) Length() int {
	return len(q.items)
}
