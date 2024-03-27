package openapi

import "github.com/nasa9084/go-openapi/oauth"

// codebeat:disable[TOO_MANY_IVARS]

// OAuthFlows Object
type OAuthFlows struct {
	Implicit          *OAuthFlow
	Password          *OAuthFlow
	ClientCredentials *OAuthFlow `yaml:"clientCredentials"`
	AuthorizationCode *OAuthFlow `yaml:"authorizationCode"`
}

// Validate the values of OAuthFlows Object.
func (oauthFlows OAuthFlows) Validate() error {
	if oauthFlows.Implicit != nil {
		oauthFlows.Implicit.SetFlowType(oauth.ImplicitFlow)
		if err := oauthFlows.Implicit.Validate(); err != nil {
			return err
		}
	}
	if oauthFlows.Password != nil {
		oauthFlows.Password.SetFlowType(oauth.PasswordFlow)
		if err := oauthFlows.Password.Validate(); err != nil {
			return err
		}
	}
	if oauthFlows.ClientCredentials != nil {
		oauthFlows.ClientCredentials.SetFlowType(oauth.ClientCredentialsFlow)
		if err := oauthFlows.ClientCredentials.Validate(); err != nil {
			return err
		}
	}
	if oauthFlows.AuthorizationCode != nil {
		oauthFlows.AuthorizationCode.SetFlowType(oauth.AuthorizationCodeFlow)
		if err := oauthFlows.AuthorizationCode.Validate(); err != nil {
			return err
		}
	}
	return nil
}
