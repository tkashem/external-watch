package external

import (
	"errors"

	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apimachinery/pkg/runtime"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	d.items<- watch.Event{
		Type: watch.Added,
		Object: &ExternalSource{
			ObjectMeta: metav1.ObjectMeta{
				Name: "redhat-operators",
				Namespace: "redhat-operators",
			},
			Spec: ExternalSourceSpec{
				Endpoint: "https:quay.io",
			},
		},
	}

	close(d.items)
}