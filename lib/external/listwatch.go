package external

// import (
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"k8s.io/apimachinery/pkg/runtime"
// 	"k8s.io/apimachinery/pkg/watch"
// 	"k8s.io/client-go/tools/cache"
// )

// type ListerWatcher struct {
// 	lister *Lister
// 	store  cache.Store	
// } 

// // List should return a list type object; the Items field will be extracted, and the
// // ResourceVersion field will be used to start the watch in the right place.
// func (lw *ListerWatcher) List(options metav1.ListOptions) (object runtime.Object, err error) {
// 	return
// }

// // Watch should begin a watch at the specified version.
// func (lw *ListerWatcher) Watch(options metav1.ListOptions) (watch.Interface, error) {
// 	return nil, nil
// }
