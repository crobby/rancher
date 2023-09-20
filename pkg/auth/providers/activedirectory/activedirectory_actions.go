package activedirectory

import (
	"fmt"
	"strings"

	"github.com/rancher/norman/api/handler"
	"github.com/rancher/norman/httperror"
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/convert"
	"github.com/sirupsen/logrus"

	v32 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/rancher/pkg/auth/providers/common"
	client "github.com/rancher/rancher/pkg/client/generated/management/v3"
	v3 "github.com/rancher/rancher/pkg/generated/norman/management.cattle.io/v3"
	managementschema "github.com/rancher/rancher/pkg/schemas/management.cattle.io/v3"
)

func (p *adProvider) formatter(apiContext *types.APIContext, resource *types.RawResource) {
	common.AddCommonActions(apiContext, resource)
	resource.AddAction(apiContext, "testAndApply")
}

func (p *adProvider) actionHandler(actionName string, action *types.Action, request *types.APIContext) error {
	handled, err := common.HandleCommonAction(actionName, action, request, Name, p.authConfigs)
	if err != nil {
		return err
	}
	if handled {
		return nil
	}

	if actionName == "testAndApply" {
		return p.testAndApply(actionName, action, request)
	}

	return httperror.NewAPIError(httperror.ActionNotAvailable, "")
}

func (p *adProvider) testAndApply(actionName string, action *types.Action, request *types.APIContext) error {
	input, err := handler.ParseAndValidateActionBody(request, request.Schemas.Schema(&managementschema.Version,
		client.ActiveDirectoryTestAndApplyInputType))
	if err != nil {
		return err
	}
	configApplyInput := &v32.ActiveDirectoryTestAndApplyInput{}
	if err := common.Decode(input, configApplyInput); err != nil {
		return httperror.NewAPIError(httperror.InvalidBodyContent,
			fmt.Sprintf("Failed to parse body: %v", err))
	}

	config := &configApplyInput.ActiveDirectoryConfig

	login := &v32.BasicLogin{
		Username: configApplyInput.Username,
		Password: configApplyInput.Password,
	}

	if config.ServiceAccountPassword != "" {
		value, err := common.ReadFromSecret(p.secrets, config.ServiceAccountPassword,
			strings.ToLower(client.ActiveDirectoryConfigFieldServiceAccountPassword))
		if err != nil {
			return err
		}
		config.ServiceAccountPassword = value
	}

	caPool, err := newCAPool(config.Certificate)
	if err != nil {
		return err
	}

	if len(config.Servers) < 1 {
		return httperror.NewAPIError(httperror.InvalidBodyContent, "must supply a server")
	}

	userPrincipal, groupPrincipals, err := p.loginUser(login, config, caPool, true)
	if err != nil {
		return err
	}

	//if this works, save adConfig CR adding enabled flag
	config.Common.Enabled = configApplyInput.Enabled
	err = p.saveActiveDirectoryConfig(config)
	if err != nil {
		return httperror.NewAPIError(httperror.ServerError, fmt.Sprintf("Failed to save activedirectory config: %v", err))
	}

	user, err := p.userMGR.SetPrincipalOnCurrentUser(request, userPrincipal)
	if err != nil {
		return err
	}

	userExtraInfo := p.GetUserExtraAttributes(userPrincipal)

	return p.tokenMGR.CreateTokenAndSetCookie(user.Name, userPrincipal, groupPrincipals, "", 0, "Token via AD Configuration", request, userExtraInfo)
}

func (p *adProvider) saveActiveDirectoryConfig(config *v32.ActiveDirectoryConfig) error {
	//TODO Rethink this whole function

	//storedConfig, _, err := p.getActiveDirectoryConfig()

	field := strings.ToLower(client.ActiveDirectoryConfigFieldServiceAccountPassword)
	if err := common.CreateOrUpdateSecrets(p.secrets, config.ServiceAccountPassword, field, strings.ToLower(convert.ToString(config.Common.Type))); err != nil {
		return err
	}

	authConfig := &v32.AuthConfig{}
	authConfig.APIVersion = "management.cattle.io/v3"
	authConfig.Kind = v3.AuthConfigGroupVersionKind.Kind
	authConfig.ActiveDirectory.Common.Type = client.ActiveDirectoryConfigType
	config.ServiceAccountPassword = common.GetFullSecretName(config.Common.Type, field)
	authConfig.ActiveDirectory = *config

	logrus.Debugf("updating activeDirectoryConfig")
	_, err := p.authConfigs.Update(authConfig)
	if err != nil {
		return err
	}
	return nil
}
