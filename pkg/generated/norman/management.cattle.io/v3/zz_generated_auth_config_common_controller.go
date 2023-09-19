package v3

import (
	"context"
	"time"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"github.com/rancher/norman/resource"
	"github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	AuthConfigCommonGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "AuthConfigCommon",
	}
	AuthConfigCommonResource = metav1.APIResource{
		Name:         "authconfigcommons",
		SingularName: "authconfigcommon",
		Namespaced:   false,
		Kind:         AuthConfigCommonGroupVersionKind.Kind,
	}

	AuthConfigCommonGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "authconfigcommons",
	}
)

func init() {
	resource.Put(AuthConfigCommonGroupVersionResource)
}

// Deprecated: use v3.AuthConfigCommon instead
type AuthConfigCommon = v3.AuthConfigCommon

func NewAuthConfigCommon(namespace, name string, obj v3.AuthConfigCommon) *v3.AuthConfigCommon {
	obj.APIVersion, obj.Kind = AuthConfigCommonGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type AuthConfigCommonHandlerFunc func(key string, obj *v3.AuthConfigCommon) (runtime.Object, error)

type AuthConfigCommonChangeHandlerFunc func(obj *v3.AuthConfigCommon) (runtime.Object, error)

type AuthConfigCommonLister interface {
	List(namespace string, selector labels.Selector) (ret []*v3.AuthConfigCommon, err error)
	Get(namespace, name string) (*v3.AuthConfigCommon, error)
}

type AuthConfigCommonController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() AuthConfigCommonLister
	AddHandler(ctx context.Context, name string, handler AuthConfigCommonHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync AuthConfigCommonHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler AuthConfigCommonHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler AuthConfigCommonHandlerFunc)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, after time.Duration)
}

type AuthConfigCommonInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v3.AuthConfigCommon) (*v3.AuthConfigCommon, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v3.AuthConfigCommon, error)
	Get(name string, opts metav1.GetOptions) (*v3.AuthConfigCommon, error)
	Update(*v3.AuthConfigCommon) (*v3.AuthConfigCommon, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*v3.AuthConfigCommonList, error)
	ListNamespaced(namespace string, opts metav1.ListOptions) (*v3.AuthConfigCommonList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() AuthConfigCommonController
	AddHandler(ctx context.Context, name string, sync AuthConfigCommonHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync AuthConfigCommonHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle AuthConfigCommonLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle AuthConfigCommonLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync AuthConfigCommonHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync AuthConfigCommonHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle AuthConfigCommonLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle AuthConfigCommonLifecycle)
}

type authConfigCommonLister struct {
	ns         string
	controller *authConfigCommonController
}

func (l *authConfigCommonLister) List(namespace string, selector labels.Selector) (ret []*v3.AuthConfigCommon, err error) {
	if namespace == "" {
		namespace = l.ns
	}
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v3.AuthConfigCommon))
	})
	return
}

func (l *authConfigCommonLister) Get(namespace, name string) (*v3.AuthConfigCommon, error) {
	var key string
	if namespace != "" {
		key = namespace + "/" + name
	} else {
		key = name
	}
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    AuthConfigCommonGroupVersionKind.Group,
			Resource: AuthConfigCommonGroupVersionResource.Resource,
		}, key)
	}
	return obj.(*v3.AuthConfigCommon), nil
}

type authConfigCommonController struct {
	ns string
	controller.GenericController
}

func (c *authConfigCommonController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *authConfigCommonController) Lister() AuthConfigCommonLister {
	return &authConfigCommonLister{
		ns:         c.ns,
		controller: c,
	}
}

func (c *authConfigCommonController) AddHandler(ctx context.Context, name string, handler AuthConfigCommonHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.AuthConfigCommon); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *authConfigCommonController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler AuthConfigCommonHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.AuthConfigCommon); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *authConfigCommonController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler AuthConfigCommonHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.AuthConfigCommon); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *authConfigCommonController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler AuthConfigCommonHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.AuthConfigCommon); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type authConfigCommonFactory struct {
}

func (c authConfigCommonFactory) Object() runtime.Object {
	return &v3.AuthConfigCommon{}
}

func (c authConfigCommonFactory) List() runtime.Object {
	return &v3.AuthConfigCommonList{}
}

func (s *authConfigCommonClient) Controller() AuthConfigCommonController {
	genericController := controller.NewGenericController(s.ns, AuthConfigCommonGroupVersionKind.Kind+"Controller",
		s.client.controllerFactory.ForResourceKind(AuthConfigCommonGroupVersionResource, AuthConfigCommonGroupVersionKind.Kind, false))

	return &authConfigCommonController{
		ns:                s.ns,
		GenericController: genericController,
	}
}

type authConfigCommonClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   AuthConfigCommonController
}

func (s *authConfigCommonClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *authConfigCommonClient) Create(o *v3.AuthConfigCommon) (*v3.AuthConfigCommon, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v3.AuthConfigCommon), err
}

func (s *authConfigCommonClient) Get(name string, opts metav1.GetOptions) (*v3.AuthConfigCommon, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v3.AuthConfigCommon), err
}

func (s *authConfigCommonClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v3.AuthConfigCommon, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v3.AuthConfigCommon), err
}

func (s *authConfigCommonClient) Update(o *v3.AuthConfigCommon) (*v3.AuthConfigCommon, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v3.AuthConfigCommon), err
}

func (s *authConfigCommonClient) UpdateStatus(o *v3.AuthConfigCommon) (*v3.AuthConfigCommon, error) {
	obj, err := s.objectClient.UpdateStatus(o.Name, o)
	return obj.(*v3.AuthConfigCommon), err
}

func (s *authConfigCommonClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *authConfigCommonClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *authConfigCommonClient) List(opts metav1.ListOptions) (*v3.AuthConfigCommonList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*v3.AuthConfigCommonList), err
}

func (s *authConfigCommonClient) ListNamespaced(namespace string, opts metav1.ListOptions) (*v3.AuthConfigCommonList, error) {
	obj, err := s.objectClient.ListNamespaced(namespace, opts)
	return obj.(*v3.AuthConfigCommonList), err
}

func (s *authConfigCommonClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *authConfigCommonClient) Patch(o *v3.AuthConfigCommon, patchType types.PatchType, data []byte, subresources ...string) (*v3.AuthConfigCommon, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v3.AuthConfigCommon), err
}

func (s *authConfigCommonClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *authConfigCommonClient) AddHandler(ctx context.Context, name string, sync AuthConfigCommonHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *authConfigCommonClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync AuthConfigCommonHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *authConfigCommonClient) AddLifecycle(ctx context.Context, name string, lifecycle AuthConfigCommonLifecycle) {
	sync := NewAuthConfigCommonLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *authConfigCommonClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle AuthConfigCommonLifecycle) {
	sync := NewAuthConfigCommonLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *authConfigCommonClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync AuthConfigCommonHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *authConfigCommonClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync AuthConfigCommonHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *authConfigCommonClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle AuthConfigCommonLifecycle) {
	sync := NewAuthConfigCommonLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *authConfigCommonClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle AuthConfigCommonLifecycle) {
	sync := NewAuthConfigCommonLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}
