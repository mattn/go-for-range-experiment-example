package main

import (
	"fmt"
	"time"
)

func TimeFrom(s time.Duration, d time.Duration) func(func(time.Time) bool) {
	return func(yield func(time.Time) bool) {
		start := time.Now().Add(s)
		for start.Before(time.Now()) {
			if !yield(start) {
				return
			}
			start = start.Add(d)
		}
	}
}

func main() {
	for v := range TimeFrom(-3*time.Minute, time.Minute) {
		fmt.Println(v)
	}
}
