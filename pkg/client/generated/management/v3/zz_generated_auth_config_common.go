package client

import (
	"github.com/rancher/norman/types"
)

const (
	AuthConfigCommonType                     = "authConfigCommon"
	AuthConfigCommonFieldAccessMode          = "accessMode"
	AuthConfigCommonFieldAllowedPrincipalIDs = "allowedPrincipalIds"
	AuthConfigCommonFieldAnnotations         = "annotations"
	AuthConfigCommonFieldCreated             = "created"
	AuthConfigCommonFieldCreatorID           = "creatorId"
	AuthConfigCommonFieldEnabled             = "enabled"
	AuthConfigCommonFieldLabels              = "labels"
	AuthConfigCommonFieldName                = "name"
	AuthConfigCommonFieldOwnerReferences     = "ownerReferences"
	AuthConfigCommonFieldRemoved             = "removed"
	AuthConfigCommonFieldStatus              = "status"
	AuthConfigCommonFieldType                = "type"
	AuthConfigCommonFieldUUID                = "uuid"
)

type AuthConfigCommon struct {
	types.Resource
	AccessMode          string            `json:"accessMode,omitempty" yaml:"accessMode,omitempty"`
	AllowedPrincipalIDs []string          `json:"allowedPrincipalIds,omitempty" yaml:"allowedPrincipalIds,omitempty"`
	Annotations         map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created             string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID           string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Enabled             bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Labels              map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences     []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed             string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	Status              *AuthConfigStatus `json:"status,omitempty" yaml:"status,omitempty"`
	Type                string            `json:"type,omitempty" yaml:"type,omitempty"`
	UUID                string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type AuthConfigCommonCollection struct {
	types.Collection
	Data   []AuthConfigCommon `json:"data,omitempty"`
	client *AuthConfigCommonClient
}

type AuthConfigCommonClient struct {
	apiClient *Client
}

type AuthConfigCommonOperations interface {
	List(opts *types.ListOpts) (*AuthConfigCommonCollection, error)
	ListAll(opts *types.ListOpts) (*AuthConfigCommonCollection, error)
	Create(opts *AuthConfigCommon) (*AuthConfigCommon, error)
	Update(existing *AuthConfigCommon, updates interface{}) (*AuthConfigCommon, error)
	Replace(existing *AuthConfigCommon) (*AuthConfigCommon, error)
	ByID(id string) (*AuthConfigCommon, error)
	Delete(container *AuthConfigCommon) error
}

func newAuthConfigCommonClient(apiClient *Client) *AuthConfigCommonClient {
	return &AuthConfigCommonClient{
		apiClient: apiClient,
	}
}

func (c *AuthConfigCommonClient) Create(container *AuthConfigCommon) (*AuthConfigCommon, error) {
	resp := &AuthConfigCommon{}
	err := c.apiClient.Ops.DoCreate(AuthConfigCommonType, container, resp)
	return resp, err
}

func (c *AuthConfigCommonClient) Update(existing *AuthConfigCommon, updates interface{}) (*AuthConfigCommon, error) {
	resp := &AuthConfigCommon{}
	err := c.apiClient.Ops.DoUpdate(AuthConfigCommonType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *AuthConfigCommonClient) Replace(obj *AuthConfigCommon) (*AuthConfigCommon, error) {
	resp := &AuthConfigCommon{}
	err := c.apiClient.Ops.DoReplace(AuthConfigCommonType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *AuthConfigCommonClient) List(opts *types.ListOpts) (*AuthConfigCommonCollection, error) {
	resp := &AuthConfigCommonCollection{}
	err := c.apiClient.Ops.DoList(AuthConfigCommonType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *AuthConfigCommonClient) ListAll(opts *types.ListOpts) (*AuthConfigCommonCollection, error) {
	resp := &AuthConfigCommonCollection{}
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

func (cc *AuthConfigCommonCollection) Next() (*AuthConfigCommonCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &AuthConfigCommonCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *AuthConfigCommonClient) ByID(id string) (*AuthConfigCommon, error) {
	resp := &AuthConfigCommon{}
	err := c.apiClient.Ops.DoByID(AuthConfigCommonType, id, resp)
	return resp, err
}

func (c *AuthConfigCommonClient) Delete(container *AuthConfigCommon) error {
	return c.apiClient.Ops.DoResourceDelete(AuthConfigCommonType, &container.Resource)
}
