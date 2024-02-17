package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"os/exec"
	"time"
)

func OutputLines(p string, args ...string) func(func(string) bool) {
	return func(yield func(string) bool) {
		cmd := exec.Command(p, args...)
		r, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer r.Close()
		err = cmd.Start()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}
	}
}

func ZipPull[V1, V2 any](
	left func(yield func(v V1) bool),
	right func(yield func(v V2) bool),
) func(yield func(l V1, r V2) bool) {
	return func(yield func(l V1, r V2) bool) {
		nextL, stopL := iter.Pull(left)
		nextR, stopR := iter.Pull(right)
		defer stopL()
		defer stopR()

		for {
			l, lOk := nextL()
			r, rOk := nextR()

			if !lOk || !rOk {
				return
			}
			if !yield(l, r) {
				return
			}
		}
	}
}

func main() {
	tic := OutputLines("yes", "tic")
	toc := OutputLines("yes", "toc")
	for l, r := range ZipPull(tic, toc) {
		fmt.Println(l, r)
		time.Sleep(time.Second)
	}
}
