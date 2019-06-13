package external

import (
	"fmt"

	"k8s.io/apimachinery/pkg/watch"
	"github.com/tkashem/external-watch/lib/appregistry"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Reflector struct {
	lister *Lister
	// store  cache.Store
	source *OperatorSource	

	store map[string]*appregistry.RegistryMetadata 
}

func (r *Reflector) List() (events []*watch.Event, err error) {
	items, err := r.lister.List(r.source)
	if err != nil {
		return
	}

	events = make([]*watch.Event, 0)
	for _, item := range items {
		r.store[item.ID()] = item
		
		event := watch.Event{
			Type: watch.Added,
			Object: &OperatorRegistry{
				ObjectMeta: metav1.ObjectMeta{
					Name: "redhat-operators",
					Namespace: "redhat-operators",
				},
				Spec: *item,
			},
		}
		
		events = append(events, &event)
	}

	return
}


func (r *Reflector) Watch() (events []*watch.Event, err error ) {
	items, err := r.lister.List(r.source)
	if err != nil {
		return
	}

	return r.updates(items)
}

func (r *Reflector) updates(items []*appregistry.RegistryMetadata) (events []*watch.Event, err error) {
	events = make([]*watch.Event, 0)
	for _, remote := range items {
		if remote.Release == "" {
			err = fmt.Errorf("Release not specified for repository [%s]", remote.ID())
			return
		}

		object := &OperatorRegistry{
			ObjectMeta: metav1.ObjectMeta{
				Name: "redhat-operators",
				Namespace: "redhat-operators",
			},
			Spec: *remote,
		}

		local, exists := r.store[remote.ID()]
		if !exists {
			// This is a new repository that has been pushed.
			event := watch.Event{
				Type: watch.Added,
				Object: object,
			}
			events = append(events, &event)
			r.store[ remote.ID() ] = remote
			continue
		}

		if local.Release != remote.Release {
			// The repository has gone through an update.
			event := watch.Event{
				Type: watch.Modified,
				Object: object,
			}
			events = append(events, &event)
			r.store[ remote.ID() ] = remote
		}
	}

	itemsMap := map[string]*appregistry.RegistryMetadata{}
	for _, remote := range items {
		itemsMap[remote.ID()] = remote
	}

	for key, value := range r.store {
		_, exists := itemsMap[key]
		if !exists {
			object := &OperatorRegistry{
				ObjectMeta: metav1.ObjectMeta{
					Name: "redhat-operators",
					Namespace: "redhat-operators",
				},
				Spec: *value,
			}

			// The repository has gone through an update.
			event := watch.Event{
				Type: watch.Deleted,
				Object: object,
			}
			events = append(events, &event)
			delete(r.store, key)
		}
	}

	return
}
