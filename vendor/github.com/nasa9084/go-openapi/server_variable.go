package openapi

// codebeat:disable[TOO_MANY_IVARS]

// ServerVariable Object
type ServerVariable struct {
	Enum        []string
	Default     string
	Description string
}

// Validate the values of Server Variable object.
func (sv ServerVariable) Validate() error {
	if sv.Default == "" {
		return ErrRequired{Target: "serverVariable.default"}
	}
	return nil
}
