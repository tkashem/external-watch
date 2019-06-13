package external

import (
	"log"
	"errors"
	"time"

	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/tkashem/external-watch/lib/appregistry"
)

// Decoder implements the watch.Decoder interface.
type Decoder struct {
	items chan watch.Event
}

func (d *Decoder) Decode() (eventType watch.EventType, object runtime.Object, err error ) {
	item, ok := <-d.items
	if !ok {
		err = errors.New("no more updates")
		return
	}

	eventType = item.Type
	object = item.Object
	return
}

func (d *Decoder) Close() {
	
}

func (d *Decoder) run() {
	defer func() {
		close(d.items)
	}()

	source := &OperatorSource{
		Endpoint: "https://quay.io/cnr",
		Namespace: "akashem",
	}

	lister := &Lister{}
	reflector := &Reflector{
		lister: lister,
		store: map[string]*appregistry.RegistryMetadata{},
		source: source,
	}

	log.Printf("watching external operator source - %s", source)

	log.Print("listing all")
	events, err := reflector.List()
	if err != nil {
		log.Printf("list error - %v", err)
		return
	}

	for _, event := range events {
		d.items<- *event
	}

	log.Print("watching for updates")
	for {
		<-time.After(5 * time.Second)
		// log.Printf("polling for changes")

		events, err := reflector.Watch()
		if err != nil {
			log.Printf("list error - %v", err)
			return
		}
	
		for _, event := range events {
			d.items<- *event
		}		
	}
}