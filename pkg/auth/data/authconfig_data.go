package data

import (
	"github.com/sirupsen/logrus"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	apimgmtv3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/rancher/pkg/auth/providers/activedirectory"
	"github.com/rancher/rancher/pkg/auth/providers/azure"
	localprovider "github.com/rancher/rancher/pkg/auth/providers/local"
	client "github.com/rancher/rancher/pkg/client/generated/management/v3"
	"github.com/rancher/rancher/pkg/controllers/management/auth"
	v3 "github.com/rancher/rancher/pkg/generated/norman/management.cattle.io/v3"
	"github.com/rancher/rancher/pkg/types/config"
)

func AuthConfigs(management *config.ManagementContext) error {
	//	if err := addAuthConfig(github.Name, client.GithubConfigType, false, management); err != nil {
	//		return err
	//	}

	if err := addAuthConfig(activedirectory.Name, client.ActiveDirectoryConfigType, false, management); err != nil {
		return err
	}

	//if err := addAuthConfig(azure.Name, client.AzureADConfigType, false, management); err != nil {
	//	return err
	//}
	//
	//if err := addAuthConfig(ldap.OpenLdapName, client.OpenLdapConfigType, false, management); err != nil {
	//	return err
	//}
	//
	//if err := addAuthConfig(ldap.FreeIpaName, client.FreeIpaConfigType, false, management); err != nil {
	//	return err
	//}
	//
	//if err := addAuthConfig(saml.PingName, client.PingConfigType, false, management); err != nil {
	//	return err
	//}
	//
	//if err := addAuthConfig(saml.ADFSName, client.ADFSConfigType, false, management); err != nil {
	//	return err
	//}
	//
	//if err := addAuthConfig(saml.KeyCloakName, client.KeyCloakConfigType, false, management); err != nil {
	//	return err
	//}
	//
	//if err := addAuthConfig(saml.OKTAName, client.OKTAConfigType, false, management); err != nil {
	//	return err
	//}
	//
	//if err := addAuthConfig(saml.ShibbolethName, client.ShibbolethConfigType, false, management); err != nil {
	//	return err
	//}
	//
	//if err := addAuthConfig(googleoauth.Name, client.GoogleOauthConfigType, false, management); err != nil {
	//	return err
	//}
	//
	//if err := addAuthConfig(oidc.Name, client.OIDCConfigType, false, management); err != nil {
	//	return err
	//}
	//
	//if err := addAuthConfig(keycloakoidc.Name, client.KeyCloakOIDCConfigType, false, management); err != nil {
	//	return err
	//}

	return addAuthConfig(localprovider.Name, client.LocalConfigType, true, management)
}

func addAuthConfig(name, aType string, enabled bool, management *config.ManagementContext) error {
	annotations := make(map[string]string)
	if name == azure.Name {
		annotations[azure.GraphEndpointMigratedAnnotation] = "true"
	}
	annotations[auth.CleanupAnnotation] = auth.CleanupRancherLocked

	authConfigObject := &v3.AuthConfig{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:        name,
			Annotations: annotations,
		},
	}
	if aType == client.LocalConfigType {
		authConfigObject.Local = apimgmtv3.LocalConfig{
			Common: apimgmtv3.AuthConfigCommon{
				Enabled: true,
				Type:    "local",
			},
		}
	} else {
		authConfigObject.ActiveDirectory = apimgmtv3.ActiveDirectoryConfig{
			Common: apimgmtv3.AuthConfigCommon{
				Enabled: false,
				Type:    "activedirectory",
			},
		}
	}

	created, err := management.Management.AuthConfigs("").Create(authConfigObject)

	logrus.Infof("the authconfig created was: %v", created)
	if err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	}

	return nil
}
