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

func ofChan[T any](c <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range c {
			if !yield(v) {
				return
			}
		}
	}
}

func main() {
	a1 := []int{1, 5, 9}
	a2 := []int{2, 6, 10, 13}
	a3 := []int{3, 7, 11}
	a4 := make(chan int)

	go func() {
		defer close(a4)
		a4 <- 4
		a4 <- 8
		a4 <- 12
		a4 <- 14
	}()

	for v := range merge(slices.Values(a1), slices.Values(a2), slices.Values(a3), ofChan(a4)) {
		fmt.Println(v)
	}
}
