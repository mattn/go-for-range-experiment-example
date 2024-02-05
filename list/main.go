package main

import (
	"iter"
)

type List[T any] struct {
	Value T
	next  *List[T]
}

func (l *List[T]) Add(v T) {
	last := l
	for last.next != nil {
		last = last.next
	}
	last.next = &List[T]{
		Value: v,
	}
}

func (l *List[T]) All() iter.Seq[List[T]] {
	return func(yield func(List[T]) bool) {
		for l.next != nil {
			l = l.next
			if !yield(*l) {
				return
			}
		}
	}
}

func main() {
	var list List[int]
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	for v := range list.All() {
		println(v.Value)
	}
}
