package openapi

// codebeat:disable[TOO_MANY_IVARS]

// Parameter Object
type Parameter struct {
	Name            string
	In              InType
	Description     string
	Required        bool
	Deprecated      string
	AllowEmptyValue bool `yaml:"allowEmptyValue"`

	Style         string
	Explode       bool
	AllowReserved bool `yaml:"allowReserved"`
	Schema        *Schema
	Example       interface{}
	Examples      map[string]*Example

	Content map[string]*MediaType

	Ref string `yaml:"$ref"`
}

// Validate the values of Parameter object.
// This function DOES NOT check whether the name field correspond to the associated path or not.
func (parameter Parameter) Validate() error {
	if err := parameter.validateRequiredObjects(); err != nil {
		return err
	}
	switch parameter.In {
	case InQuery, InHeader, InPath, InCookie:
	default:
		return ErrMustOneOf{Object: "parameter.in", ValidValues: ParameterInList}
	}
	if parameter.In == InPath && !parameter.Required {
		return ErrRequiredMustTrue
	}
	if parameter.In != InQuery && parameter.AllowEmptyValue {
		return ErrAllowEmptyValueNotValid
	}
	if len(parameter.Content) > 1 {
		return ErrTooManyParameterContent
	}

	return validateAll(parameter.reduceValidaters())
}

func (parameter Parameter) validateRequiredObjects() error {
	if parameter.Name == "" {
		return ErrRequired{Target: "parameter.name"}
	}
	if parameter.In == "" {
		return ErrRequired{Target: "parameter.in"}
	}
	return nil
}

func (parameter Parameter) reduceValidaters() []validater {
	validaters := []validater{}
	if parameter.Schema != nil {
		validaters = append(validaters, parameter.Schema)
	}
	if v, ok := parameter.Example.(validater); ok {
		validaters = append(validaters, v)
	}

	// example has no validation

	for _, mediaType := range parameter.Content {
		validaters = append(validaters, mediaType)
	}
	return validaters
}
