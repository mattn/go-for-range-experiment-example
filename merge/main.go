package main

import (
	"fmt"
	"iter"
	"slices"
)

func merge[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		nexts := make([]func() (T, bool), len(seqs))
		stops := make([]func(), len(seqs))
		for i, seq := range seqs {
			nexts[i], stops[i] = iter.Pull(seq)
			defer stops[i]()
		}
		for len(nexts) > 0 {
			v, ok := nexts[0]()
			if !ok {
				nexts = nexts[1:]
				continue
			}
			if !yield(v) {
				break
			}
			nexts = append(nexts[1:], nexts[0])
		}
	}
}

func main() {
	a1 := []int{1, 4, 7}
	a2 := []int{2, 5, 8, 10}
	a3 := []int{3, 6, 9}

	for v := range merge(slices.Values(a1), slices.Values(a2), slices.Values(a3)) {
		fmt.Println(v)
	}
}
