package client

const (
	KeyCloakOIDCConfigType                    = "keyCloakOIDCConfig"
	KeyCloakOIDCConfigFieldAnnotations        = "annotations"
	KeyCloakOIDCConfigFieldAuthEndpoint       = "authEndpoint"
	KeyCloakOIDCConfigFieldCertificate        = "certificate"
	KeyCloakOIDCConfigFieldClientID           = "clientId"
	KeyCloakOIDCConfigFieldClientSecret       = "clientSecret"
	KeyCloakOIDCConfigFieldCommon             = "common"
	KeyCloakOIDCConfigFieldCreated            = "created"
	KeyCloakOIDCConfigFieldCreatorID          = "creatorId"
	KeyCloakOIDCConfigFieldGroupSearchEnabled = "groupSearchEnabled"
	KeyCloakOIDCConfigFieldIssuer             = "issuer"
	KeyCloakOIDCConfigFieldLabels             = "labels"
	KeyCloakOIDCConfigFieldName               = "name"
	KeyCloakOIDCConfigFieldOwnerReferences    = "ownerReferences"
	KeyCloakOIDCConfigFieldPrivateKey         = "privateKey"
	KeyCloakOIDCConfigFieldRancherURL         = "rancherUrl"
	KeyCloakOIDCConfigFieldRemoved            = "removed"
	KeyCloakOIDCConfigFieldScopes             = "scope"
	KeyCloakOIDCConfigFieldUUID               = "uuid"
)

type KeyCloakOIDCConfig struct {
	Annotations        map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AuthEndpoint       string            `json:"authEndpoint,omitempty" yaml:"authEndpoint,omitempty"`
	Certificate        string            `json:"certificate,omitempty" yaml:"certificate,omitempty"`
	ClientID           string            `json:"clientId,omitempty" yaml:"clientId,omitempty"`
	ClientSecret       string            `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`
	Common             *AuthConfigCommon `json:"common,omitempty" yaml:"common,omitempty"`
	Created            string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID          string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	GroupSearchEnabled *bool             `json:"groupSearchEnabled,omitempty" yaml:"groupSearchEnabled,omitempty"`
	Issuer             string            `json:"issuer,omitempty" yaml:"issuer,omitempty"`
	Labels             map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name               string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences    []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	PrivateKey         string            `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`
	RancherURL         string            `json:"rancherUrl,omitempty" yaml:"rancherUrl,omitempty"`
	Removed            string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	Scopes             string            `json:"scope,omitempty" yaml:"scope,omitempty"`
	UUID               string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
