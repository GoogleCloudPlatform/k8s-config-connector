package openapi

import (
	"net/url"
	"regexp"
)

var (
	emailRegexp  = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	mapKeyRegexp = regexp.MustCompile("^[a-zA-Z0-9\\.\\-_]+$")
)

type validater interface {
	Validate() error
}

func validateAll(vs []validater) error {
	for _, v := range vs {
		if v == nil {
			continue
		}
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func mustURL(name, urlStr string) error {
	if urlStr == "" {
		return ErrRequired{Target: name}
	}
	if _, err := url.ParseRequestURI(urlStr); err != nil {
		return ErrFormatInvalid{Target: name, Format: "URL"}
	}
	return nil
}
