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
	"github.com/rancher/lasso/pkg/controller"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v3.AddToScheme)
}

type Interface interface {
	ADFSConfig() ADFSConfigController
	APIService() APIServiceController
	ActiveDirectoryConfig() ActiveDirectoryConfigController
	ActiveDirectoryProvider() ActiveDirectoryProviderController
	AuthConfig() AuthConfigController
	AuthConfigCommon() AuthConfigCommonController
	AuthProvider() AuthProviderController
	AuthToken() AuthTokenController
	AzureADConfig() AzureADConfigController
	AzureADProvider() AzureADProviderController
	Catalog() CatalogController
	CatalogTemplate() CatalogTemplateController
	CatalogTemplateVersion() CatalogTemplateVersionController
	CloudCredential() CloudCredentialController
	Cluster() ClusterController
	ClusterAlert() ClusterAlertController
	ClusterAlertGroup() ClusterAlertGroupController
	ClusterAlertRule() ClusterAlertRuleController
	ClusterCatalog() ClusterCatalogController
	ClusterLogging() ClusterLoggingController
	ClusterMonitorGraph() ClusterMonitorGraphController
	ClusterRegistrationToken() ClusterRegistrationTokenController
	ClusterRoleTemplateBinding() ClusterRoleTemplateBindingController
	ClusterTemplate() ClusterTemplateController
	ClusterTemplateRevision() ClusterTemplateRevisionController
	ComposeConfig() ComposeConfigController
	DynamicSchema() DynamicSchemaController
	EtcdBackup() EtcdBackupController
	Feature() FeatureController
	FleetWorkspace() FleetWorkspaceController
	FreeIpaConfig() FreeIpaConfigController
	FreeIpaProvider() FreeIpaProviderController
	GithubConfig() GithubConfigController
	GithubProvider() GithubProviderController
	GlobalDns() GlobalDnsController
	GlobalDnsProvider() GlobalDnsProviderController
	GlobalRole() GlobalRoleController
	GlobalRoleBinding() GlobalRoleBindingController
	GoogleOAuthProvider() GoogleOAuthProviderController
	GoogleOauthConfig() GoogleOauthConfigController
	Group() GroupController
	GroupMember() GroupMemberController
	KeyCloakConfig() KeyCloakConfigController
	KeyCloakOIDCConfig() KeyCloakOIDCConfigController
	KontainerDriver() KontainerDriverController
	LdapConfig() LdapConfigController
	LocalConfig() LocalConfigController
	LocalProvider() LocalProviderController
	ManagedChart() ManagedChartController
	MonitorMetric() MonitorMetricController
	MultiClusterApp() MultiClusterAppController
	MultiClusterAppRevision() MultiClusterAppRevisionController
	Node() NodeController
	NodeDriver() NodeDriverController
	NodePool() NodePoolController
	NodeTemplate() NodeTemplateController
	Notifier() NotifierController
	OIDCConfig() OIDCConfigController
	OIDCProvider() OIDCProviderController
	OKTAConfig() OKTAConfigController
	OneProvider() OneProviderController
	OpenLdapConfig() OpenLdapConfigController
	OpenLdapProvider() OpenLdapProviderController
	PingConfig() PingConfigController
	PodSecurityAdmissionConfigurationTemplate() PodSecurityAdmissionConfigurationTemplateController
	PodSecurityPolicyTemplate() PodSecurityPolicyTemplateController
	PodSecurityPolicyTemplateProjectBinding() PodSecurityPolicyTemplateProjectBindingController
	Preference() PreferenceController
	Principal() PrincipalController
	Project() ProjectController
	ProjectAlert() ProjectAlertController
	ProjectAlertGroup() ProjectAlertGroupController
	ProjectAlertRule() ProjectAlertRuleController
	ProjectCatalog() ProjectCatalogController
	ProjectLogging() ProjectLoggingController
	ProjectMonitorGraph() ProjectMonitorGraphController
	ProjectNetworkPolicy() ProjectNetworkPolicyController
	ProjectRoleTemplateBinding() ProjectRoleTemplateBindingController
	RancherUserNotification() RancherUserNotificationController
	RkeAddon() RkeAddonController
	RkeK8sServiceOption() RkeK8sServiceOptionController
	RkeK8sSystemImage() RkeK8sSystemImageController
	RoleTemplate() RoleTemplateController
	SamlConfig() SamlConfigController
	SamlProvider() SamlProviderController
	SamlToken() SamlTokenController
	Setting() SettingController
	ShibbolethConfig() ShibbolethConfigController
	Template() TemplateController
	TemplateContent() TemplateContentController
	TemplateVersion() TemplateVersionController
	Token() TokenController
	User() UserController
	UserAttribute() UserAttributeController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (v *version) ADFSConfig() ADFSConfigController {
	return generic.NewController[*v3.ADFSConfig, *v3.ADFSConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ADFSConfig"}, "adfsconfigs", true, v.controllerFactory)
}

func (v *version) APIService() APIServiceController {
	return generic.NewNonNamespacedController[*v3.APIService, *v3.APIServiceList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "APIService"}, "apiservices", v.controllerFactory)
}

func (v *version) ActiveDirectoryConfig() ActiveDirectoryConfigController {
	return generic.NewController[*v3.ActiveDirectoryConfig, *v3.ActiveDirectoryConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ActiveDirectoryConfig"}, "activedirectoryconfigs", true, v.controllerFactory)
}

func (v *version) ActiveDirectoryProvider() ActiveDirectoryProviderController {
	return generic.NewNonNamespacedController[*v3.ActiveDirectoryProvider, *v3.ActiveDirectoryProviderList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ActiveDirectoryProvider"}, "activedirectoryproviders", v.controllerFactory)
}

func (v *version) AuthConfig() AuthConfigController {
	return generic.NewNonNamespacedController[*v3.AuthConfig, *v3.AuthConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "AuthConfig"}, "authconfigs", v.controllerFactory)
}

func (v *version) AuthConfigCommon() AuthConfigCommonController {
	return generic.NewNonNamespacedController[*v3.AuthConfigCommon, *v3.AuthConfigCommonList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "AuthConfigCommon"}, "authconfigcommons", v.controllerFactory)
}

func (v *version) AuthProvider() AuthProviderController {
	return generic.NewNonNamespacedController[*v3.AuthProvider, *v3.AuthProviderList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "AuthProvider"}, "authproviders", v.controllerFactory)
}

func (v *version) AuthToken() AuthTokenController {
	return generic.NewNonNamespacedController[*v3.AuthToken, *v3.AuthTokenList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "AuthToken"}, "authtokens", v.controllerFactory)
}

func (v *version) AzureADConfig() AzureADConfigController {
	return generic.NewController[*v3.AzureADConfig, *v3.AzureADConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "AzureADConfig"}, "azureadconfigs", true, v.controllerFactory)
}

func (v *version) AzureADProvider() AzureADProviderController {
	return generic.NewNonNamespacedController[*v3.AzureADProvider, *v3.AzureADProviderList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "AzureADProvider"}, "azureadproviders", v.controllerFactory)
}

func (v *version) Catalog() CatalogController {
	return generic.NewNonNamespacedController[*v3.Catalog, *v3.CatalogList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Catalog"}, "catalogs", v.controllerFactory)
}

func (v *version) CatalogTemplate() CatalogTemplateController {
	return generic.NewController[*v3.CatalogTemplate, *v3.CatalogTemplateList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "CatalogTemplate"}, "catalogtemplates", true, v.controllerFactory)
}

func (v *version) CatalogTemplateVersion() CatalogTemplateVersionController {
	return generic.NewController[*v3.CatalogTemplateVersion, *v3.CatalogTemplateVersionList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "CatalogTemplateVersion"}, "catalogtemplateversions", true, v.controllerFactory)
}

func (v *version) CloudCredential() CloudCredentialController {
	return generic.NewController[*v3.CloudCredential, *v3.CloudCredentialList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "CloudCredential"}, "cloudcredentials", true, v.controllerFactory)
}

func (v *version) Cluster() ClusterController {
	return generic.NewNonNamespacedController[*v3.Cluster, *v3.ClusterList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Cluster"}, "clusters", v.controllerFactory)
}

func (v *version) ClusterAlert() ClusterAlertController {
	return generic.NewController[*v3.ClusterAlert, *v3.ClusterAlertList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterAlert"}, "clusteralerts", true, v.controllerFactory)
}

func (v *version) ClusterAlertGroup() ClusterAlertGroupController {
	return generic.NewController[*v3.ClusterAlertGroup, *v3.ClusterAlertGroupList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterAlertGroup"}, "clusteralertgroups", true, v.controllerFactory)
}

func (v *version) ClusterAlertRule() ClusterAlertRuleController {
	return generic.NewController[*v3.ClusterAlertRule, *v3.ClusterAlertRuleList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterAlertRule"}, "clusteralertrules", true, v.controllerFactory)
}

func (v *version) ClusterCatalog() ClusterCatalogController {
	return generic.NewController[*v3.ClusterCatalog, *v3.ClusterCatalogList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterCatalog"}, "clustercatalogs", true, v.controllerFactory)
}

func (v *version) ClusterLogging() ClusterLoggingController {
	return generic.NewController[*v3.ClusterLogging, *v3.ClusterLoggingList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterLogging"}, "clusterloggings", true, v.controllerFactory)
}

func (v *version) ClusterMonitorGraph() ClusterMonitorGraphController {
	return generic.NewController[*v3.ClusterMonitorGraph, *v3.ClusterMonitorGraphList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterMonitorGraph"}, "clustermonitorgraphs", true, v.controllerFactory)
}

func (v *version) ClusterRegistrationToken() ClusterRegistrationTokenController {
	return generic.NewController[*v3.ClusterRegistrationToken, *v3.ClusterRegistrationTokenList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterRegistrationToken"}, "clusterregistrationtokens", true, v.controllerFactory)
}

func (v *version) ClusterRoleTemplateBinding() ClusterRoleTemplateBindingController {
	return generic.NewController[*v3.ClusterRoleTemplateBinding, *v3.ClusterRoleTemplateBindingList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterRoleTemplateBinding"}, "clusterroletemplatebindings", true, v.controllerFactory)
}

func (v *version) ClusterTemplate() ClusterTemplateController {
	return generic.NewController[*v3.ClusterTemplate, *v3.ClusterTemplateList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterTemplate"}, "clustertemplates", true, v.controllerFactory)
}

func (v *version) ClusterTemplateRevision() ClusterTemplateRevisionController {
	return generic.NewController[*v3.ClusterTemplateRevision, *v3.ClusterTemplateRevisionList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterTemplateRevision"}, "clustertemplaterevisions", true, v.controllerFactory)
}

func (v *version) ComposeConfig() ComposeConfigController {
	return generic.NewNonNamespacedController[*v3.ComposeConfig, *v3.ComposeConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ComposeConfig"}, "composeconfigs", v.controllerFactory)
}

func (v *version) DynamicSchema() DynamicSchemaController {
	return generic.NewNonNamespacedController[*v3.DynamicSchema, *v3.DynamicSchemaList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "DynamicSchema"}, "dynamicschemas", v.controllerFactory)
}

func (v *version) EtcdBackup() EtcdBackupController {
	return generic.NewController[*v3.EtcdBackup, *v3.EtcdBackupList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "EtcdBackup"}, "etcdbackups", true, v.controllerFactory)
}

func (v *version) Feature() FeatureController {
	return generic.NewNonNamespacedController[*v3.Feature, *v3.FeatureList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Feature"}, "features", v.controllerFactory)
}

func (v *version) FleetWorkspace() FleetWorkspaceController {
	return generic.NewNonNamespacedController[*v3.FleetWorkspace, *v3.FleetWorkspaceList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "FleetWorkspace"}, "fleetworkspaces", v.controllerFactory)
}

func (v *version) FreeIpaConfig() FreeIpaConfigController {
	return generic.NewController[*v3.FreeIpaConfig, *v3.FreeIpaConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "FreeIpaConfig"}, "freeipaconfigs", true, v.controllerFactory)
}

func (v *version) FreeIpaProvider() FreeIpaProviderController {
	return generic.NewNonNamespacedController[*v3.FreeIpaProvider, *v3.FreeIpaProviderList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "FreeIpaProvider"}, "freeipaproviders", v.controllerFactory)
}

func (v *version) GithubConfig() GithubConfigController {
	return generic.NewNonNamespacedController[*v3.GithubConfig, *v3.GithubConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GithubConfig"}, "githubconfigs", v.controllerFactory)
}

func (v *version) GithubProvider() GithubProviderController {
	return generic.NewNonNamespacedController[*v3.GithubProvider, *v3.GithubProviderList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GithubProvider"}, "githubproviders", v.controllerFactory)
}

func (v *version) GlobalDns() GlobalDnsController {
	return generic.NewController[*v3.GlobalDns, *v3.GlobalDnsList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GlobalDns"}, "globaldnses", true, v.controllerFactory)
}

func (v *version) GlobalDnsProvider() GlobalDnsProviderController {
	return generic.NewController[*v3.GlobalDnsProvider, *v3.GlobalDnsProviderList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GlobalDnsProvider"}, "globaldnsproviders", true, v.controllerFactory)
}

func (v *version) GlobalRole() GlobalRoleController {
	return generic.NewNonNamespacedController[*v3.GlobalRole, *v3.GlobalRoleList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GlobalRole"}, "globalroles", v.controllerFactory)
}

func (v *version) GlobalRoleBinding() GlobalRoleBindingController {
	return generic.NewNonNamespacedController[*v3.GlobalRoleBinding, *v3.GlobalRoleBindingList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GlobalRoleBinding"}, "globalrolebindings", v.controllerFactory)
}

func (v *version) GoogleOAuthProvider() GoogleOAuthProviderController {
	return generic.NewNonNamespacedController[*v3.GoogleOAuthProvider, *v3.GoogleOAuthProviderList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GoogleOAuthProvider"}, "googleoauthproviders", v.controllerFactory)
}

func (v *version) GoogleOauthConfig() GoogleOauthConfigController {
	return generic.NewController[*v3.GoogleOauthConfig, *v3.GoogleOauthConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GoogleOauthConfig"}, "googleoauthconfigs", true, v.controllerFactory)
}

func (v *version) Group() GroupController {
	return generic.NewNonNamespacedController[*v3.Group, *v3.GroupList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Group"}, "groups", v.controllerFactory)
}

func (v *version) GroupMember() GroupMemberController {
	return generic.NewNonNamespacedController[*v3.GroupMember, *v3.GroupMemberList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GroupMember"}, "groupmembers", v.controllerFactory)
}

func (v *version) KeyCloakConfig() KeyCloakConfigController {
	return generic.NewController[*v3.KeyCloakConfig, *v3.KeyCloakConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "KeyCloakConfig"}, "keycloakconfigs", true, v.controllerFactory)
}

func (v *version) KeyCloakOIDCConfig() KeyCloakOIDCConfigController {
	return generic.NewController[*v3.KeyCloakOIDCConfig, *v3.KeyCloakOIDCConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "KeyCloakOIDCConfig"}, "keycloakoidcconfigs", true, v.controllerFactory)
}

func (v *version) KontainerDriver() KontainerDriverController {
	return generic.NewNonNamespacedController[*v3.KontainerDriver, *v3.KontainerDriverList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "KontainerDriver"}, "kontainerdrivers", v.controllerFactory)
}

func (v *version) LdapConfig() LdapConfigController {
	return generic.NewController[*v3.LdapConfig, *v3.LdapConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "LdapConfig"}, "ldapconfigs", true, v.controllerFactory)
}

func (v *version) LocalConfig() LocalConfigController {
	return generic.NewController[*v3.LocalConfig, *v3.LocalConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "LocalConfig"}, "localconfigs", true, v.controllerFactory)
}

func (v *version) LocalProvider() LocalProviderController {
	return generic.NewNonNamespacedController[*v3.LocalProvider, *v3.LocalProviderList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "LocalProvider"}, "localproviders", v.controllerFactory)
}

func (v *version) ManagedChart() ManagedChartController {
	return generic.NewController[*v3.ManagedChart, *v3.ManagedChartList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ManagedChart"}, "managedcharts", true, v.controllerFactory)
}

func (v *version) MonitorMetric() MonitorMetricController {
	return generic.NewController[*v3.MonitorMetric, *v3.MonitorMetricList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "MonitorMetric"}, "monitormetrics", true, v.controllerFactory)
}

func (v *version) MultiClusterApp() MultiClusterAppController {
	return generic.NewController[*v3.MultiClusterApp, *v3.MultiClusterAppList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "MultiClusterApp"}, "multiclusterapps", true, v.controllerFactory)
}

func (v *version) MultiClusterAppRevision() MultiClusterAppRevisionController {
	return generic.NewController[*v3.MultiClusterAppRevision, *v3.MultiClusterAppRevisionList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "MultiClusterAppRevision"}, "multiclusterapprevisions", true, v.controllerFactory)
}

func (v *version) Node() NodeController {
	return generic.NewController[*v3.Node, *v3.NodeList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Node"}, "nodes", true, v.controllerFactory)
}

func (v *version) NodeDriver() NodeDriverController {
	return generic.NewNonNamespacedController[*v3.NodeDriver, *v3.NodeDriverList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "NodeDriver"}, "nodedrivers", v.controllerFactory)
}

func (v *version) NodePool() NodePoolController {
	return generic.NewController[*v3.NodePool, *v3.NodePoolList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "NodePool"}, "nodepools", true, v.controllerFactory)
}

func (v *version) NodeTemplate() NodeTemplateController {
	return generic.NewController[*v3.NodeTemplate, *v3.NodeTemplateList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "NodeTemplate"}, "nodetemplates", true, v.controllerFactory)
}

func (v *version) Notifier() NotifierController {
	return generic.NewController[*v3.Notifier, *v3.NotifierList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Notifier"}, "notifiers", true, v.controllerFactory)
}

func (v *version) OIDCConfig() OIDCConfigController {
	return generic.NewController[*v3.OIDCConfig, *v3.OIDCConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "OIDCConfig"}, "oidcconfigs", true, v.controllerFactory)
}

func (v *version) OIDCProvider() OIDCProviderController {
	return generic.NewNonNamespacedController[*v3.OIDCProvider, *v3.OIDCProviderList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "OIDCProvider"}, "oidcproviders", v.controllerFactory)
}

func (v *version) OKTAConfig() OKTAConfigController {
	return generic.NewController[*v3.OKTAConfig, *v3.OKTAConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "OKTAConfig"}, "oktaconfigs", true, v.controllerFactory)
}

func (v *version) OneProvider() OneProviderController {
	return generic.NewController[*v3.OneProvider, *v3.OneProviderList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "OneProvider"}, "oneproviders", true, v.controllerFactory)
}

func (v *version) OpenLdapConfig() OpenLdapConfigController {
	return generic.NewController[*v3.OpenLdapConfig, *v3.OpenLdapConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "OpenLdapConfig"}, "openldapconfigs", true, v.controllerFactory)
}

func (v *version) OpenLdapProvider() OpenLdapProviderController {
	return generic.NewNonNamespacedController[*v3.OpenLdapProvider, *v3.OpenLdapProviderList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "OpenLdapProvider"}, "openldapproviders", v.controllerFactory)
}

func (v *version) PingConfig() PingConfigController {
	return generic.NewController[*v3.PingConfig, *v3.PingConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "PingConfig"}, "pingconfigs", true, v.controllerFactory)
}

func (v *version) PodSecurityAdmissionConfigurationTemplate() PodSecurityAdmissionConfigurationTemplateController {
	return generic.NewNonNamespacedController[*v3.PodSecurityAdmissionConfigurationTemplate, *v3.PodSecurityAdmissionConfigurationTemplateList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "PodSecurityAdmissionConfigurationTemplate"}, "podsecurityadmissionconfigurationtemplates", v.controllerFactory)
}

func (v *version) PodSecurityPolicyTemplate() PodSecurityPolicyTemplateController {
	return generic.NewNonNamespacedController[*v3.PodSecurityPolicyTemplate, *v3.PodSecurityPolicyTemplateList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "PodSecurityPolicyTemplate"}, "podsecuritypolicytemplates", v.controllerFactory)
}

func (v *version) PodSecurityPolicyTemplateProjectBinding() PodSecurityPolicyTemplateProjectBindingController {
	return generic.NewController[*v3.PodSecurityPolicyTemplateProjectBinding, *v3.PodSecurityPolicyTemplateProjectBindingList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "PodSecurityPolicyTemplateProjectBinding"}, "podsecuritypolicytemplateprojectbindings", true, v.controllerFactory)
}

func (v *version) Preference() PreferenceController {
	return generic.NewController[*v3.Preference, *v3.PreferenceList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Preference"}, "preferences", true, v.controllerFactory)
}

func (v *version) Principal() PrincipalController {
	return generic.NewNonNamespacedController[*v3.Principal, *v3.PrincipalList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Principal"}, "principals", v.controllerFactory)
}

func (v *version) Project() ProjectController {
	return generic.NewController[*v3.Project, *v3.ProjectList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Project"}, "projects", true, v.controllerFactory)
}

func (v *version) ProjectAlert() ProjectAlertController {
	return generic.NewController[*v3.ProjectAlert, *v3.ProjectAlertList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectAlert"}, "projectalerts", true, v.controllerFactory)
}

func (v *version) ProjectAlertGroup() ProjectAlertGroupController {
	return generic.NewController[*v3.ProjectAlertGroup, *v3.ProjectAlertGroupList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectAlertGroup"}, "projectalertgroups", true, v.controllerFactory)
}

func (v *version) ProjectAlertRule() ProjectAlertRuleController {
	return generic.NewController[*v3.ProjectAlertRule, *v3.ProjectAlertRuleList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectAlertRule"}, "projectalertrules", true, v.controllerFactory)
}

func (v *version) ProjectCatalog() ProjectCatalogController {
	return generic.NewController[*v3.ProjectCatalog, *v3.ProjectCatalogList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectCatalog"}, "projectcatalogs", true, v.controllerFactory)
}

func (v *version) ProjectLogging() ProjectLoggingController {
	return generic.NewController[*v3.ProjectLogging, *v3.ProjectLoggingList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectLogging"}, "projectloggings", true, v.controllerFactory)
}

func (v *version) ProjectMonitorGraph() ProjectMonitorGraphController {
	return generic.NewController[*v3.ProjectMonitorGraph, *v3.ProjectMonitorGraphList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectMonitorGraph"}, "projectmonitorgraphs", true, v.controllerFactory)
}

func (v *version) ProjectNetworkPolicy() ProjectNetworkPolicyController {
	return generic.NewController[*v3.ProjectNetworkPolicy, *v3.ProjectNetworkPolicyList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectNetworkPolicy"}, "projectnetworkpolicies", true, v.controllerFactory)
}

func (v *version) ProjectRoleTemplateBinding() ProjectRoleTemplateBindingController {
	return generic.NewController[*v3.ProjectRoleTemplateBinding, *v3.ProjectRoleTemplateBindingList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectRoleTemplateBinding"}, "projectroletemplatebindings", true, v.controllerFactory)
}

func (v *version) RancherUserNotification() RancherUserNotificationController {
	return generic.NewNonNamespacedController[*v3.RancherUserNotification, *v3.RancherUserNotificationList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "RancherUserNotification"}, "rancherusernotifications", v.controllerFactory)
}

func (v *version) RkeAddon() RkeAddonController {
	return generic.NewController[*v3.RkeAddon, *v3.RkeAddonList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "RkeAddon"}, "rkeaddons", true, v.controllerFactory)
}

func (v *version) RkeK8sServiceOption() RkeK8sServiceOptionController {
	return generic.NewController[*v3.RkeK8sServiceOption, *v3.RkeK8sServiceOptionList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "RkeK8sServiceOption"}, "rkek8sserviceoptions", true, v.controllerFactory)
}

func (v *version) RkeK8sSystemImage() RkeK8sSystemImageController {
	return generic.NewController[*v3.RkeK8sSystemImage, *v3.RkeK8sSystemImageList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "RkeK8sSystemImage"}, "rkek8ssystemimages", true, v.controllerFactory)
}

func (v *version) RoleTemplate() RoleTemplateController {
	return generic.NewNonNamespacedController[*v3.RoleTemplate, *v3.RoleTemplateList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "RoleTemplate"}, "roletemplates", v.controllerFactory)
}

func (v *version) SamlConfig() SamlConfigController {
	return generic.NewController[*v3.SamlConfig, *v3.SamlConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "SamlConfig"}, "samlconfigs", true, v.controllerFactory)
}

func (v *version) SamlProvider() SamlProviderController {
	return generic.NewNonNamespacedController[*v3.SamlProvider, *v3.SamlProviderList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "SamlProvider"}, "samlproviders", v.controllerFactory)
}

func (v *version) SamlToken() SamlTokenController {
	return generic.NewNonNamespacedController[*v3.SamlToken, *v3.SamlTokenList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "SamlToken"}, "samltokens", v.controllerFactory)
}

func (v *version) Setting() SettingController {
	return generic.NewNonNamespacedController[*v3.Setting, *v3.SettingList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Setting"}, "settings", v.controllerFactory)
}

func (v *version) ShibbolethConfig() ShibbolethConfigController {
	return generic.NewController[*v3.ShibbolethConfig, *v3.ShibbolethConfigList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ShibbolethConfig"}, "shibbolethconfigs", true, v.controllerFactory)
}

func (v *version) Template() TemplateController {
	return generic.NewNonNamespacedController[*v3.Template, *v3.TemplateList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Template"}, "templates", v.controllerFactory)
}

func (v *version) TemplateContent() TemplateContentController {
	return generic.NewNonNamespacedController[*v3.TemplateContent, *v3.TemplateContentList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "TemplateContent"}, "templatecontents", v.controllerFactory)
}

func (v *version) TemplateVersion() TemplateVersionController {
	return generic.NewNonNamespacedController[*v3.TemplateVersion, *v3.TemplateVersionList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "TemplateVersion"}, "templateversions", v.controllerFactory)
}

func (v *version) Token() TokenController {
	return generic.NewNonNamespacedController[*v3.Token, *v3.TokenList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Token"}, "tokens", v.controllerFactory)
}

func (v *version) User() UserController {
	return generic.NewNonNamespacedController[*v3.User, *v3.UserList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "User"}, "users", v.controllerFactory)
}

func (v *version) UserAttribute() UserAttributeController {
	return generic.NewNonNamespacedController[*v3.UserAttribute, *v3.UserAttributeList](schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "UserAttribute"}, "userattributes", v.controllerFactory)
}
