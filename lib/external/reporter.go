package external

import (
	"k8s.io/apimachinery/pkg/runtime"
)

type Reporter struct {
}

func (f *Reporter) AsObject(err error) runtime.Object {
	return runtime.Unstructured(nil)
}