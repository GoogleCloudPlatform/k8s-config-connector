package openapi

// codebeat:disable[TOO_MANY_IVARS]

// SecurityScheme Object
type SecurityScheme struct {
	Type             SecuritySchemeType
	Description      string
	Name             string
	In               InType
	Scheme           string
	BearerFormat     string `yaml:"bearerFormat"`
	Flows            *OAuthFlows
	OpenIDConnectURL string `yaml:"openIdConnectUrl"`

	Ref string `yaml:"$ref"`
}

// SecuritySchemeType represents a securityScheme.type value.
type SecuritySchemeType string

// SecuritySchemeTypes
const (
	APIKeyType        SecuritySchemeType = "apiKey"
	HTTPType          SecuritySchemeType = "http"
	OAuth2Type        SecuritySchemeType = "oauth2"
	OpenIDConnectType SecuritySchemeType = "openIdConnect"
)

// SecuritySchemeTypeList is a list of valid values of securityScheme.Type.
var SecuritySchemeTypeList = []string{string(APIKeyType), string(HTTPType), string(OAuth2Type), string(OpenIDConnectType)}

// Validate the values of SecurityScheme object.
func (secScheme SecurityScheme) Validate() error {
	switch secScheme.Type {
	case "":
		return ErrRequired{Target: "securityScheme.type"}
	case APIKeyType:
		return secScheme.validateFieldForAPIKey()
	case HTTPType:
		return secScheme.validateFieldForHTTP()
	case OAuth2Type:
		return secScheme.validateFieldForOAuth2()
	case OpenIDConnectType:
		return secScheme.validateFieldForOpenIDConnect()
	}
	return ErrMustOneOf{Object: "securityScheme.type", ValidValues: SecuritySchemeTypeList}
}

func (secScheme SecurityScheme) validateFieldForAPIKey() error {
	if secScheme.Name == "" {
		return ErrRequired{"securityScheme.name"}
	}
	if secScheme.In == "" {
		return ErrRequired{"securityScheme.in"}
	}
	if secScheme.In != InQuery && secScheme.In != InHeader && secScheme.In != InCookie {
		return ErrMustOneOf{Object: "securityScheme.in", ValidValues: SecuritySchemeInList}
	}
	return nil
}

func (secScheme SecurityScheme) validateFieldForHTTP() error {
	if secScheme.Scheme == "" {
		return ErrRequired{Target: "securityScheme.scheme"}
	}
	return nil
}

func (secScheme SecurityScheme) validateFieldForOAuth2() error {
	if secScheme.Flows == nil {
		return ErrRequired{Target: "securityScheme.flows"}
	}
	return secScheme.Flows.Validate()
}

func (secScheme SecurityScheme) validateFieldForOpenIDConnect() error {
	return mustURL("securityScheme.openIdConnectUrl", secScheme.OpenIDConnectURL)
}
