package openapi

import (
	"strconv"
)

// codebeat:disable[TOO_MANY_IVARS]

// Responses Object
type Responses map[string]*Response

// Validate the values of Responses object.
func (responses Responses) Validate() error {
	for status, response := range responses {
		if err := validateStatusCode(status); err != nil {
			return err
		}
		if err := response.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func validateStatusCode(statusStr string) error {
	switch statusStr {
	case "default", "1XX", "2XX", "3XX", "4XX", "5XX":
		return nil
	}
	statusInt, err := strconv.Atoi(statusStr)
	if err != nil {
		return ErrInvalidStatusCode
	}
	if statusInt < 100 || 599 < statusInt {
		return ErrInvalidStatusCode
	}
	return nil
}
