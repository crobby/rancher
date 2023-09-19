package client

const (
	LocalConfigType                 = "localConfig"
	LocalConfigFieldAnnotations     = "annotations"
	LocalConfigFieldCommon          = "common"
	LocalConfigFieldCreated         = "created"
	LocalConfigFieldCreatorID       = "creatorId"
	LocalConfigFieldLabels          = "labels"
	LocalConfigFieldName            = "name"
	LocalConfigFieldOwnerReferences = "ownerReferences"
	LocalConfigFieldRemoved         = "removed"
	LocalConfigFieldUUID            = "uuid"
)

type LocalConfig struct {
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Common          *AuthConfigCommon `json:"common,omitempty" yaml:"common,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
