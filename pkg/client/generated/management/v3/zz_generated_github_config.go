package client

const (
	GithubConfigType                     = "githubConfig"
	GithubConfigFieldAdditionalClientIDs = "additionalClientIds"
	GithubConfigFieldAnnotations         = "annotations"
	GithubConfigFieldClientID            = "clientId"
	GithubConfigFieldClientSecret        = "clientSecret"
	GithubConfigFieldCommon              = "common"
	GithubConfigFieldCreated             = "created"
	GithubConfigFieldCreatorID           = "creatorId"
	GithubConfigFieldHostname            = "hostname"
	GithubConfigFieldHostnameToClientID  = "hostnameToClientId"
	GithubConfigFieldLabels              = "labels"
	GithubConfigFieldName                = "name"
	GithubConfigFieldOwnerReferences     = "ownerReferences"
	GithubConfigFieldRemoved             = "removed"
	GithubConfigFieldTLS                 = "tls"
	GithubConfigFieldUUID                = "uuid"
)

type GithubConfig struct {
	AdditionalClientIDs map[string]string `json:"additionalClientIds,omitempty" yaml:"additionalClientIds,omitempty"`
	Annotations         map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ClientID            string            `json:"clientId,omitempty" yaml:"clientId,omitempty"`
	ClientSecret        string            `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`
	Common              *AuthConfigCommon `json:"common,omitempty" yaml:"common,omitempty"`
	Created             string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID           string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Hostname            string            `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	HostnameToClientID  map[string]string `json:"hostnameToClientId,omitempty" yaml:"hostnameToClientId,omitempty"`
	Labels              map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences     []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed             string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	TLS                 bool              `json:"tls,omitempty" yaml:"tls,omitempty"`
	UUID                string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
