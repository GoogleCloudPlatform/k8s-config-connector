package openapi

import (
	"strconv"
)

// codebeat:disable[TOO_MANY_IVARS]

// Operation Object
type Operation struct {
	Tags         []string
	Summary      string
	Description  string
	ExternalDocs *ExternalDocumentation `yaml:"externalDocs"`
	OperationID  string                 `yaml:"operationId"`
	Parameters   []*Parameter
	RequestBody  *RequestBody `yaml:"requestBody"`
	Responses    Responses
	Callbacks    map[string]*Callback
	Deprecated   bool
	Security     []*SecurityRequirement
	Servers      []*Server
}

// SuccessResponse returns a success response object.
// If there are 2 or more success responses (like created and ok),
// it's not sure which is returned.
// If only match the default response or 2XX response, returned status code will be 0.
func (operation *Operation) SuccessResponse() (*Response, int, bool) {
	if operation == nil || operation.Responses == nil {
		return nil, -1, false
	}
	var defaultResponse *Response
	for statusStr, resp := range operation.Responses {
		switch statusStr {
		case "default", "2XX":
			defaultResponse = resp
		case "1XX", "3XX", "4XX", "5XX":
			continue
		}
		statusInt, err := strconv.Atoi(statusStr)
		if err != nil {
			continue
		}
		if statusInt/100 == 2 {
			if resp == nil {
				continue
			}
			return resp, statusInt, true
		}
	}
	return defaultResponse, 0, (defaultResponse != nil)
}

// Validate the values of Operation object.
func (operation Operation) Validate() error {
	validaters := []validater{}
	if operation.ExternalDocs != nil {
		validaters = append(validaters, operation.ExternalDocs)
	}
	if hasDuplicatedParameter(operation.Parameters) {
		return ErrParameterDuplicated
	}
	if operation.RequestBody != nil {
		validaters = append(validaters, operation.RequestBody)
	}
	if operation.Responses == nil {
		return ErrRequired{Target: "operation.responses"}
	}
	validaters = append(validaters, operation.Responses)
	for _, callback := range operation.Callbacks {
		validaters = append(validaters, callback)
	}
	for _, security := range operation.Security {
		validaters = append(validaters, security)
	}
	for _, server := range operation.Servers {
		validaters = append(validaters, server)
	}
	return validateAll(validaters)
}
