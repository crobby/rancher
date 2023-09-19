package v3

import (
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"k8s.io/apimachinery/pkg/runtime"
)

type AuthConfigCommonLifecycle interface {
	Create(obj *v3.AuthConfigCommon) (runtime.Object, error)
	Remove(obj *v3.AuthConfigCommon) (runtime.Object, error)
	Updated(obj *v3.AuthConfigCommon) (runtime.Object, error)
}

type authConfigCommonLifecycleAdapter struct {
	lifecycle AuthConfigCommonLifecycle
}

func (w *authConfigCommonLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *authConfigCommonLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *authConfigCommonLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v3.AuthConfigCommon))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *authConfigCommonLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v3.AuthConfigCommon))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *authConfigCommonLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v3.AuthConfigCommon))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewAuthConfigCommonLifecycleAdapter(name string, clusterScoped bool, client AuthConfigCommonInterface, l AuthConfigCommonLifecycle) AuthConfigCommonHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(AuthConfigCommonGroupVersionResource)
	}
	adapter := &authConfigCommonLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v3.AuthConfigCommon) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
