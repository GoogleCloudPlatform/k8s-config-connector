package openapi

import (
	"encoding/json"
	"sort"
)

// codebeat:disable[TOO_MANY_IVARS]

// SecurityRequirement Object
type SecurityRequirement struct {
	document *Document
	mp       map[string][]string
}

// UnmarshalJSON implements json.Unmarshaler.
func (secReq *SecurityRequirement) UnmarshalJSON(data []byte) error {
	v := map[string][]string{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	secReq.mp = v
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (secReq *SecurityRequirement) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return unmarshal(&secReq.mp)
}

// Get returns required security schemes. If there is not given name,
// this function returns nil.
func (secReq SecurityRequirement) Get(name string) []string {
	val, ok := secReq.mp[name]
	if !ok {
		return nil
	}
	return val
}

// Names returns the keys of security requirements.
// The returned slice is sorted.
func (secReq SecurityRequirement) Names() []string {
	ret := make([]string, 0, len(secReq.mp))
	for k := range secReq.mp {
		ret = append(ret, k)
	}
	sort.Strings(ret)
	return ret
}

// Validate the values of SecurityRequirement object.
func (secReq SecurityRequirement) Validate() error {
	if len(secReq.mp) == 0 {
		return nil
	}
	if secReq.document == nil {
		return ErrMissingRootDocument
	}
	components := secReq.document.Components
	if components == nil {
		return ErrRequired{Target: "components object in parent document"}
	}
	return secReq.validateScopes()
}

func (secReq SecurityRequirement) validateScopes() error {
	for name, scopes := range secReq.mp {
		secScheme, ok := secReq.document.Components.SecuritySchemes[name]
		if !ok {
			return ErrNotDeclared{Name: name}
		}
		if secScheme.Type != OAuth2Type {
			if len(scopes) != 0 {
				return ErrMustEmpty{Type: string(secScheme.Type)}
			}
			continue
		}
		for _, scope := range scopes {
			_, implicit := secScheme.Flows.Implicit.Scopes[scope]
			_, password := secScheme.Flows.Password.Scopes[scope]
			_, cc := secScheme.Flows.ClientCredentials.Scopes[scope]
			_, ac := secScheme.Flows.AuthorizationCode.Scopes[scope]
			if !implicit && !password && !cc && !ac {
				return ErrNotDeclared{Name: scope}
			}
		}
	}
	return nil
}

func (secReq *SecurityRequirement) setDocument(doc *Document) {
	secReq.document = doc
}
