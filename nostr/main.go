package main

import (
	"context"
	"iter"
	"log"

	"github.com/nbd-wtf/go-nostr"
)

func relayFunc(url string) iter.Seq[*nostr.Event] {
	return func(yield func(ev *nostr.Event) bool) {
		ctx := context.Background()
		relay, err := nostr.RelayConnect(ctx, url)
		if err != nil {
			log.Fatal(err)
		}
		defer relay.Close()

		now := nostr.Now()
		filter := nostr.Filters{
			{
				Kinds: []int{nostr.KindTextNote},
				Since: &now,
			},
		}
		sub, err := relay.Subscribe(ctx, filter)
		if err != nil {
			log.Println(err)
			return
		}
		for {
			ev, ok := <-sub.Events
			if !ok || !yield(ev) {
				return
			}
		}
	}
}

func main() {
	for note := range relayFunc("wss://yabu.me") {
		println(note.Content)
	}
}
