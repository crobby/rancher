package common

import (
	"strings"

	"github.com/sirupsen/logrus"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
)

func TransformToAuthProvider(authConfig map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	if m, ok := authConfig["metadata"].(map[string]interface{}); ok {
		result["id"] = m["name"]
	}
	authConfigObject := &v3.AuthConfig{}
	err := Decode(authConfig, authConfigObject)
	if err != nil {
		logrus.Errorf("unable to decode auth provider: %v", err)
	}
	activeProvider, err := ActiveProviderInConfig(*authConfigObject)
	if err != nil {
		logrus.Errorf("unable to generate valid auth provider: %v", err)
	}

	result["type"] = strings.Replace(activeProvider.Type, "Config", "Provider", -1)
	return result
}
