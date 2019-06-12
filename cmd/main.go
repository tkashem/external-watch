package main

import (
	"fmt"

	"github.com/tkashem/external-watch/lib/external"
)

func main() {
	fmt.Println("watching external resources")

	watcher := external.NewWatch()

	ch := watcher.ResultChan()
	for event := range ch {
		fmt.Println("event=%v", event)
	}
}