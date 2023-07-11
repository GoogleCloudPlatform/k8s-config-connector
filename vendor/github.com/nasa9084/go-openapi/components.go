package openapi

// codebeat:disable[TOO_MANY_IVARS]

// Components Object
type Components struct {
	Schemas         map[string]*Schema
	Responses       map[string]*Response
	Parameters      map[string]*Parameter
	Examples        map[string]*Example
	RequestBodies   map[string]*RequestBody `yaml:"requestBodies"`
	Headers         map[string]*Header
	SecuritySchemes map[string]*SecurityScheme `yaml:"securitySchemes"`
	Links           map[string]*Link
	Callbacks       map[string]*Callback
}

// Validate the values of Components object.
func (components Components) Validate() error {
	if err := validateComponentKeys(components); err != nil {
		return err
	}
	validaters := reduceComponentObjects(components)
	return validateAll(validaters)
}

func validateComponentKeys(components Components) error {
	keys := reduceComponentKeys(components)
	for _, k := range keys {
		if !mapKeyRegexp.MatchString(k) {
			return ErrMapKeyFormat
		}
	}
	return nil
}

func reduceComponentKeys(components Components) []string {
	keys := []string{}
	for k := range components.Schemas {
		keys = append(keys, k)
	}
	for k := range components.Responses {
		keys = append(keys, k)
	}
	for k := range components.Parameters {
		keys = append(keys, k)
	}
	for k := range components.Examples {
		keys = append(keys, k)
	}
	for k := range components.RequestBodies {
		keys = append(keys, k)
	}
	for k := range components.Headers {
		keys = append(keys, k)
	}
	for k := range components.SecuritySchemes {
		keys = append(keys, k)
	}
	for k := range components.Links {
		keys = append(keys, k)
	}
	for k := range components.Callbacks {
		keys = append(keys, k)
	}
	return keys
}

func reduceComponentObjects(components Components) []validater {
	validaters := []validater{}
	for _, schema := range components.Schemas {
		validaters = append(validaters, schema)
	}
	for _, response := range components.Responses {
		validaters = append(validaters, response)
	}
	for _, parameter := range components.Parameters {
		validaters = append(validaters, parameter)
	}

	// example has no validation

	for _, reqBody := range components.RequestBodies {
		validaters = append(validaters, reqBody)
	}
	for _, header := range components.Headers {
		validaters = append(validaters, header)
	}
	for _, secScheme := range components.SecuritySchemes {
		validaters = append(validaters, secScheme)
	}
	for _, link := range components.Links {
		validaters = append(validaters, link)
	}
	for _, callback := range components.Callbacks {
		validaters = append(validaters, callback)
	}
	return validaters
}
