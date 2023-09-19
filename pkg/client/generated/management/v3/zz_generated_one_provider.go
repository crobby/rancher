package client

import (
	"github.com/rancher/norman/types"
)

const (
	OneProviderType                 = "oneProvider"
	OneProviderFieldAnnotations     = "annotations"
	OneProviderFieldCreated         = "created"
	OneProviderFieldCreatorID       = "creatorId"
	OneProviderFieldLabels          = "labels"
	OneProviderFieldLocal           = "inline"
	OneProviderFieldName            = "name"
	OneProviderFieldOwnerReferences = "ownerReferences"
	OneProviderFieldRemoved         = "removed"
	OneProviderFieldUUID            = "uuid"
)

type OneProvider struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Local           *LocalConfig      `json:"inline,omitempty" yaml:"inline,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type OneProviderCollection struct {
	types.Collection
	Data   []OneProvider `json:"data,omitempty"`
	client *OneProviderClient
}

type OneProviderClient struct {
	apiClient *Client
}

type OneProviderOperations interface {
	List(opts *types.ListOpts) (*OneProviderCollection, error)
	ListAll(opts *types.ListOpts) (*OneProviderCollection, error)
	Create(opts *OneProvider) (*OneProvider, error)
	Update(existing *OneProvider, updates interface{}) (*OneProvider, error)
	Replace(existing *OneProvider) (*OneProvider, error)
	ByID(id string) (*OneProvider, error)
	Delete(container *OneProvider) error
}

func newOneProviderClient(apiClient *Client) *OneProviderClient {
	return &OneProviderClient{
		apiClient: apiClient,
	}
}

func (c *OneProviderClient) Create(container *OneProvider) (*OneProvider, error) {
	resp := &OneProvider{}
	err := c.apiClient.Ops.DoCreate(OneProviderType, container, resp)
	return resp, err
}

func (c *OneProviderClient) Update(existing *OneProvider, updates interface{}) (*OneProvider, error) {
	resp := &OneProvider{}
	err := c.apiClient.Ops.DoUpdate(OneProviderType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *OneProviderClient) Replace(obj *OneProvider) (*OneProvider, error) {
	resp := &OneProvider{}
	err := c.apiClient.Ops.DoReplace(OneProviderType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *OneProviderClient) List(opts *types.ListOpts) (*OneProviderCollection, error) {
	resp := &OneProviderCollection{}
	err := c.apiClient.Ops.DoList(OneProviderType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *OneProviderClient) ListAll(opts *types.ListOpts) (*OneProviderCollection, error) {
	resp := &OneProviderCollection{}
	resp, err := c.List(opts)
	if err != nil {
		return resp, err
	}
	data := resp.Data
	for next, err := resp.Next(); next != nil && err == nil; next, err = next.Next() {
		data = append(data, next.Data...)
		resp = next
		resp.Data = data
	}
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (cc *OneProviderCollection) Next() (*OneProviderCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &OneProviderCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *OneProviderClient) ByID(id string) (*OneProvider, error) {
	resp := &OneProvider{}
	err := c.apiClient.Ops.DoByID(OneProviderType, id, resp)
	return resp, err
}

func (c *OneProviderClient) Delete(container *OneProvider) error {
	return c.apiClient.Ops.DoResourceDelete(OneProviderType, &container.Resource)
}
