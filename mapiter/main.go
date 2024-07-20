package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	m.Store("alice", 11)
	m.Store("bob", 12)
	m.Store("cindy", 13)
	for key, val := range m.Range {
		fmt.Println(key, val)
	}
}
