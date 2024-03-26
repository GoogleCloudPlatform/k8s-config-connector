package openapi

import (
	"fmt"
	"strings"
)

// ErrFormatInvalid is returned some error caused by string format is occurred.
type ErrFormatInvalid struct {
	Target string
	Format string
}

func (fe ErrFormatInvalid) Error() string {
	if fe.Format == "" {
		return fmt.Sprintf("%s format is invalid", fe.Target)
	}
	return fmt.Sprintf("%s format is invalid: should be %s", fe.Target, fe.Format)
}

// central error variables relating format
var (
	ErrMapKeyFormat      = ErrFormatInvalid{Target: "map key"}
	ErrPathFormat        = ErrFormatInvalid{Target: "path"}
	ErrRuntimeExprFormat = ErrFormatInvalid{Target: "key", Format: "RuntimeExpression"}
)

// ErrRequired is returned when missing some required parameter
type ErrRequired struct {
	Target string
}

func (re ErrRequired) Error() string {
	return fmt.Sprintf("%s is required", re.Target)
}

type errString string

func (ge errString) Error() string {
	return string(ge)
}

const (
	// ErrUnsupportedVersion is returned when the openapi version
	// is unsupported by this package.
	ErrUnsupportedVersion errString = "the OAS version is not supported"
	// ErrInvalidFlowType is returned when the OAuth flow type is invalid
	// or not set to the object.
	ErrInvalidFlowType errString = "invalid flow type"
	// ErrRequiredMustTrue is returned when the value of parameter.required is
	// false when parameter.in is path.
	ErrRequiredMustTrue errString = "required must be true if parameter.in is path"
	// ErrAllowEmptyValueNotValid is returned when allowEmptyValue is specified
	// but parameter.in is not query.
	ErrAllowEmptyValueNotValid errString = "allowEmptyValue is valid only for query parameters"
	// ErrInvalidStatusCode is returned when specified status code is not
	// valid as HTTP status code.
	ErrInvalidStatusCode errString = "status code is invalid"
	// ErrMissingRootDocument is returned when validating securityRequirement
	// object but root document is not set.
	ErrMissingRootDocument errString = "missing root document for security requirement"
)

type errTooManyContentEntry struct {
	target string
}

func (etme errTooManyContentEntry) Error() string {
	return fmt.Sprintf("%s.content must only contain one entry", etme.target)
}

var (
	// ErrTooManyHeaderContent is returned when the length of header.content
	// is more than 2.
	ErrTooManyHeaderContent = errTooManyContentEntry{target: "header"}
	// ErrTooManyParameterContent is returned when the length of parameter.content
	// is more than 2.
	ErrTooManyParameterContent = errTooManyContentEntry{target: "parameter"}
)

type errDuplicated struct {
	target string
}

func (de errDuplicated) Error() string {
	return fmt.Sprintf("some %s are duplicated", de.target)
}

var (
	// ErrOperationIDDuplicated is returned when some operation ids are
	// duplicated but operation ids cannot be duplicated.
	ErrOperationIDDuplicated = errDuplicated{target: "operation ids"}
	// ErrParameterDuplicated is returned when some parameters are duplicated
	// but cannot be duplicated.
	ErrParameterDuplicated = errDuplicated{target: "parameters"}
	// ErrPathsDuplicated is returned when some paths are duplicated.
	ErrPathsDuplicated = errDuplicated{target: "paths"}
)

// ErrNotDeclared is returned when the securityScheme name is
// not defined in components object in the document.
type ErrNotDeclared struct {
	Name string
}

func (snde ErrNotDeclared) Error() string {
	return fmt.Sprintf("%s is not declared in components.securitySchemes", snde.Name)
}

// ErrMustEmpty returned when the securityRequirement is not
// empty but must be empty.
type ErrMustEmpty struct {
	Type string
}

func (srmee ErrMustEmpty) Error() string {
	return fmt.Sprintf("securityRequirements for %s type must be empty", srmee.Type)
}

// ErrMustOneOf is returned some value must be one of given list, but not one.
type ErrMustOneOf struct {
	Object      string
	ValidValues []string
}

func (ooe ErrMustOneOf) Error() string {
	return fmt.Sprintf("%s must be one of: %s", ooe.Object, strings.Join(ooe.ValidValues, ", "))
}
