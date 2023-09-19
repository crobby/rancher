/*
Copyright 2023 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v3

import (
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/generic"
)

// OpenLdapConfigController interface for managing OpenLdapConfig resources.
type OpenLdapConfigController interface {
	generic.ControllerInterface[*v3.OpenLdapConfig, *v3.OpenLdapConfigList]
}

// OpenLdapConfigClient interface for managing OpenLdapConfig resources in Kubernetes.
type OpenLdapConfigClient interface {
	generic.ClientInterface[*v3.OpenLdapConfig, *v3.OpenLdapConfigList]
}

// OpenLdapConfigCache interface for retrieving OpenLdapConfig resources in memory.
type OpenLdapConfigCache interface {
	generic.CacheInterface[*v3.OpenLdapConfig]
}
