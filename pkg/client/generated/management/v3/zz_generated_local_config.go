package client

const (
	LocalConfigType        = "localConfig"
	LocalConfigFieldCommon = "common"
)

type LocalConfig struct {
	Common *AuthConfigCommon `json:"common,omitempty" yaml:"common,omitempty"`
}
