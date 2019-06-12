package external

import (
	"k8s.io/apimachinery/pkg/watch"
)

func NewWatch() watch.Interface {
	decoder := &Decoder{
		items: make(chan watch.Event, 0),
	}
	reporter := &Reporter{}
	watcher := watch.NewStreamWatcher(decoder, reporter)

	go decoder.run()
	return watcher
}


