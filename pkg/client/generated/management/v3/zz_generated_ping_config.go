package client

const (
	PingConfigType                    = "pingConfig"
	PingConfigFieldAnnotations        = "annotations"
	PingConfigFieldCommon             = "common"
	PingConfigFieldCreated            = "created"
	PingConfigFieldCreatorID          = "creatorId"
	PingConfigFieldDisplayNameField   = "displayNameField"
	PingConfigFieldEntityID           = "entityID"
	PingConfigFieldGroupsField        = "groupsField"
	PingConfigFieldIDPMetadataContent = "idpMetadataContent"
	PingConfigFieldLabels             = "labels"
	PingConfigFieldName               = "name"
	PingConfigFieldOwnerReferences    = "ownerReferences"
	PingConfigFieldRancherAPIHost     = "rancherApiHost"
	PingConfigFieldRemoved            = "removed"
	PingConfigFieldSpCert             = "spCert"
	PingConfigFieldSpKey              = "spKey"
	PingConfigFieldUIDField           = "uidField"
	PingConfigFieldUUID               = "uuid"
	PingConfigFieldUserNameField      = "userNameField"
)

type PingConfig struct {
	Annotations        map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Common             *AuthConfigCommon `json:"common,omitempty" yaml:"common,omitempty"`
	Created            string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID          string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	DisplayNameField   string            `json:"displayNameField,omitempty" yaml:"displayNameField,omitempty"`
	EntityID           string            `json:"entityID,omitempty" yaml:"entityID,omitempty"`
	GroupsField        string            `json:"groupsField,omitempty" yaml:"groupsField,omitempty"`
	IDPMetadataContent string            `json:"idpMetadataContent,omitempty" yaml:"idpMetadataContent,omitempty"`
	Labels             map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name               string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences    []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	RancherAPIHost     string            `json:"rancherApiHost,omitempty" yaml:"rancherApiHost,omitempty"`
	Removed            string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	SpCert             string            `json:"spCert,omitempty" yaml:"spCert,omitempty"`
	SpKey              string            `json:"spKey,omitempty" yaml:"spKey,omitempty"`
	UIDField           string            `json:"uidField,omitempty" yaml:"uidField,omitempty"`
	UUID               string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	UserNameField      string            `json:"userNameField,omitempty" yaml:"userNameField,omitempty"`
}
