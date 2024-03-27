package openapi

import (
	"net/url"
)

// codebeat:disable[TOO_MANY_IVARS]

// Info Object
type Info struct {
	Title          string
	Description    string
	TermsOfService string `yaml:"termsOfService"`
	Contact        *Contact
	License        *License
	Version        string
}

// Validate the values of Info object.
func (info Info) Validate() error {
	if err := info.validateRequiredFields(); err != nil {
		return err
	}
	return info.validateFields()
}

func (info Info) validateRequiredFields() error {
	if info.Title == "" {
		return ErrRequired{Target: "info.title"}
	}
	if info.Version == "" {
		return ErrRequired{Target: "info.version"}
	}
	return nil
}

func (info Info) validateFields() error {
	if info.TermsOfService != "" {
		if _, err := url.ParseRequestURI(info.TermsOfService); err != nil {
			return ErrFormatInvalid{Target: "info.termsOfService", Format: "URL"}
		}
	}

	var validaters []validater
	if info.Contact != nil {
		validaters = append(validaters, info.Contact)
	}
	if info.License != nil {
		validaters = append(validaters, info.License)
	}
	return validateAll(validaters)
}
