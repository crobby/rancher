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
	OneProviderGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "OneProvider",
	}
	OneProviderResource = metav1.APIResource{
		Name:         "oneproviders",
		SingularName: "oneprovider",
		Namespaced:   false,
		Kind:         OneProviderGroupVersionKind.Kind,
	}

	OneProviderGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "oneproviders",
	}
)

func init() {
	resource.Put(OneProviderGroupVersionResource)
}

// Deprecated: use v3.OneProvider instead
type OneProvider = v3.OneProvider

func NewOneProvider(namespace, name string, obj v3.OneProvider) *v3.OneProvider {
	obj.APIVersion, obj.Kind = OneProviderGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type OneProviderHandlerFunc func(key string, obj *v3.OneProvider) (runtime.Object, error)

type OneProviderChangeHandlerFunc func(obj *v3.OneProvider) (runtime.Object, error)

type OneProviderLister interface {
	List(namespace string, selector labels.Selector) (ret []*v3.OneProvider, err error)
	Get(namespace, name string) (*v3.OneProvider, error)
}

type OneProviderController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() OneProviderLister
	AddHandler(ctx context.Context, name string, handler OneProviderHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync OneProviderHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler OneProviderHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler OneProviderHandlerFunc)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, after time.Duration)
}

type OneProviderInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v3.OneProvider) (*v3.OneProvider, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v3.OneProvider, error)
	Get(name string, opts metav1.GetOptions) (*v3.OneProvider, error)
	Update(*v3.OneProvider) (*v3.OneProvider, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*v3.OneProviderList, error)
	ListNamespaced(namespace string, opts metav1.ListOptions) (*v3.OneProviderList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() OneProviderController
	AddHandler(ctx context.Context, name string, sync OneProviderHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync OneProviderHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle OneProviderLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle OneProviderLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync OneProviderHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync OneProviderHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle OneProviderLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle OneProviderLifecycle)
}

type oneProviderLister struct {
	ns         string
	controller *oneProviderController
}

func (l *oneProviderLister) List(namespace string, selector labels.Selector) (ret []*v3.OneProvider, err error) {
	if namespace == "" {
		namespace = l.ns
	}
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v3.OneProvider))
	})
	return
}

func (l *oneProviderLister) Get(namespace, name string) (*v3.OneProvider, error) {
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
			Group:    OneProviderGroupVersionKind.Group,
			Resource: OneProviderGroupVersionResource.Resource,
		}, key)
	}
	return obj.(*v3.OneProvider), nil
}

type oneProviderController struct {
	ns string
	controller.GenericController
}

func (c *oneProviderController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *oneProviderController) Lister() OneProviderLister {
	return &oneProviderLister{
		ns:         c.ns,
		controller: c,
	}
}

func (c *oneProviderController) AddHandler(ctx context.Context, name string, handler OneProviderHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.OneProvider); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *oneProviderController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler OneProviderHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.OneProvider); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *oneProviderController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler OneProviderHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.OneProvider); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *oneProviderController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler OneProviderHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.OneProvider); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type oneProviderFactory struct {
}

func (c oneProviderFactory) Object() runtime.Object {
	return &v3.OneProvider{}
}

func (c oneProviderFactory) List() runtime.Object {
	return &v3.OneProviderList{}
}

func (s *oneProviderClient) Controller() OneProviderController {
	genericController := controller.NewGenericController(s.ns, OneProviderGroupVersionKind.Kind+"Controller",
		s.client.controllerFactory.ForResourceKind(OneProviderGroupVersionResource, OneProviderGroupVersionKind.Kind, false))

	return &oneProviderController{
		ns:                s.ns,
		GenericController: genericController,
	}
}

type oneProviderClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   OneProviderController
}

func (s *oneProviderClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *oneProviderClient) Create(o *v3.OneProvider) (*v3.OneProvider, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v3.OneProvider), err
}

func (s *oneProviderClient) Get(name string, opts metav1.GetOptions) (*v3.OneProvider, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v3.OneProvider), err
}

func (s *oneProviderClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v3.OneProvider, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v3.OneProvider), err
}

func (s *oneProviderClient) Update(o *v3.OneProvider) (*v3.OneProvider, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v3.OneProvider), err
}

func (s *oneProviderClient) UpdateStatus(o *v3.OneProvider) (*v3.OneProvider, error) {
	obj, err := s.objectClient.UpdateStatus(o.Name, o)
	return obj.(*v3.OneProvider), err
}

func (s *oneProviderClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *oneProviderClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *oneProviderClient) List(opts metav1.ListOptions) (*v3.OneProviderList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*v3.OneProviderList), err
}

func (s *oneProviderClient) ListNamespaced(namespace string, opts metav1.ListOptions) (*v3.OneProviderList, error) {
	obj, err := s.objectClient.ListNamespaced(namespace, opts)
	return obj.(*v3.OneProviderList), err
}

func (s *oneProviderClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *oneProviderClient) Patch(o *v3.OneProvider, patchType types.PatchType, data []byte, subresources ...string) (*v3.OneProvider, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v3.OneProvider), err
}

func (s *oneProviderClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *oneProviderClient) AddHandler(ctx context.Context, name string, sync OneProviderHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *oneProviderClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync OneProviderHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *oneProviderClient) AddLifecycle(ctx context.Context, name string, lifecycle OneProviderLifecycle) {
	sync := NewOneProviderLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *oneProviderClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle OneProviderLifecycle) {
	sync := NewOneProviderLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *oneProviderClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync OneProviderHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *oneProviderClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync OneProviderHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *oneProviderClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle OneProviderLifecycle) {
	sync := NewOneProviderLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *oneProviderClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle OneProviderLifecycle) {
	sync := NewOneProviderLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}
