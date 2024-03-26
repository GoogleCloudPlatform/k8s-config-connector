package openapi

// codebeat:disable[TOO_MANY_IVARS]

// Example Object
type Example struct {
	Summary       string
	Description   string
	Value         interface{}
	ExternalValue interface{} `yaml:"externalValue"`

	Ref string `yaml:"$ref"`
}
