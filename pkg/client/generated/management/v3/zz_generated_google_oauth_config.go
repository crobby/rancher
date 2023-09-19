package client

const (
	GoogleOauthConfigType                              = "googleOauthConfig"
	GoogleOauthConfigFieldAdminEmail                   = "adminEmail"
	GoogleOauthConfigFieldAnnotations                  = "annotations"
	GoogleOauthConfigFieldCommon                       = "common"
	GoogleOauthConfigFieldCreated                      = "created"
	GoogleOauthConfigFieldCreatorID                    = "creatorId"
	GoogleOauthConfigFieldHostname                     = "hostname"
	GoogleOauthConfigFieldLabels                       = "labels"
	GoogleOauthConfigFieldName                         = "name"
	GoogleOauthConfigFieldNestedGroupMembershipEnabled = "nestedGroupMembershipEnabled"
	GoogleOauthConfigFieldOauthCredential              = "oauthCredential"
	GoogleOauthConfigFieldOwnerReferences              = "ownerReferences"
	GoogleOauthConfigFieldRemoved                      = "removed"
	GoogleOauthConfigFieldServiceAccountCredential     = "serviceAccountCredential"
	GoogleOauthConfigFieldUUID                         = "uuid"
	GoogleOauthConfigFieldUserInfoEndpoint             = "userInfoEndpoint"
)

type GoogleOauthConfig struct {
	AdminEmail                   string            `json:"adminEmail,omitempty" yaml:"adminEmail,omitempty"`
	Annotations                  map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Common                       *AuthConfigCommon `json:"common,omitempty" yaml:"common,omitempty"`
	Created                      string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                    string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Hostname                     string            `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	Labels                       map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                         string            `json:"name,omitempty" yaml:"name,omitempty"`
	NestedGroupMembershipEnabled bool              `json:"nestedGroupMembershipEnabled,omitempty" yaml:"nestedGroupMembershipEnabled,omitempty"`
	OauthCredential              string            `json:"oauthCredential,omitempty" yaml:"oauthCredential,omitempty"`
	OwnerReferences              []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed                      string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	ServiceAccountCredential     string            `json:"serviceAccountCredential,omitempty" yaml:"serviceAccountCredential,omitempty"`
	UUID                         string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	UserInfoEndpoint             string            `json:"userInfoEndpoint,omitempty" yaml:"userInfoEndpoint,omitempty"`
}
