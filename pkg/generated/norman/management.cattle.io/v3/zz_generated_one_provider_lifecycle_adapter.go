package v3

import (
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"k8s.io/apimachinery/pkg/runtime"
)

type OneProviderLifecycle interface {
	Create(obj *v3.OneProvider) (runtime.Object, error)
	Remove(obj *v3.OneProvider) (runtime.Object, error)
	Updated(obj *v3.OneProvider) (runtime.Object, error)
}

type oneProviderLifecycleAdapter struct {
	lifecycle OneProviderLifecycle
}

func (w *oneProviderLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *oneProviderLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *oneProviderLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v3.OneProvider))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *oneProviderLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v3.OneProvider))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *oneProviderLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v3.OneProvider))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewOneProviderLifecycleAdapter(name string, clusterScoped bool, client OneProviderInterface, l OneProviderLifecycle) OneProviderHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(OneProviderGroupVersionResource)
	}
	adapter := &oneProviderLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v3.OneProvider) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
