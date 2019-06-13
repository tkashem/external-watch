package main

import (
	"log"

	"github.com/tkashem/external-watch/lib/external"
)

func main() {
	watcher := external.NewWatch()

	ch := watcher.ResultChan()
	for event := range ch {
		registry, ok := event.Object.(*external.OperatorRegistry)
		if !ok {
			log.Printf("wrong type of object, event type=%s", event.Type)
			continue
		}
		
		log.Printf("new watch event - type=%s object=%s", event.Type, registry.Spec.String())
	}
}