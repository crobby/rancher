package client

const (
	AuthConfigCommonType                     = "authConfigCommon"
	AuthConfigCommonFieldAccessMode          = "accessMode"
	AuthConfigCommonFieldAllowedPrincipalIDs = "allowedPrincipalIds"
	AuthConfigCommonFieldEnabled             = "enabled"
	AuthConfigCommonFieldStatus              = "status"
	AuthConfigCommonFieldType                = "type"
)

type AuthConfigCommon struct {
	AccessMode          string            `json:"accessMode,omitempty" yaml:"accessMode,omitempty"`
	AllowedPrincipalIDs []string          `json:"allowedPrincipalIds,omitempty" yaml:"allowedPrincipalIds,omitempty"`
	Enabled             bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Status              *AuthConfigStatus `json:"status,omitempty" yaml:"status,omitempty"`
	Type                string            `json:"type,omitempty" yaml:"type,omitempty"`
}
