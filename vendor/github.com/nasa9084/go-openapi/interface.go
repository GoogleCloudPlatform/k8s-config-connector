package openapi

// InType represents where the parameter or the securityScheme is in.
type InType string

// InTypes
const (
	InQuery  InType = "query"
	InHeader InType = "header"
	InPath   InType = "path"
	InCookie InType = "cookie"
)

// InType Lists for ErrOneOf
var (
	ParameterInList      = []string{string(InQuery), string(InHeader), string(InPath), string(InCookie)}
	SecuritySchemeInList = []string{string(InQuery), string(InHeader), string(InCookie)}
)
