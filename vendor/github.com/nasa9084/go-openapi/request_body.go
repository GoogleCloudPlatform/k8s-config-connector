package openapi

// codebeat:disable[TOO_MANY_IVARS]

// RequestBody Object
type RequestBody struct {
	Description string
	Content     map[string]*MediaType
	Required    bool

	Ref string `yaml:"$ref"`
}

// Validate the values of RequestBody object.
func (requestBody RequestBody) Validate() error {
	if requestBody.Ref != "" {
		return nil // validated in doc.Components
	}
	if requestBody.Content == nil || len(requestBody.Content) == 0 {
		return ErrRequired{Target: "requestBody.content"}
	}
	for _, mediaType := range requestBody.Content {
		if err := mediaType.Validate(); err != nil {
			return err
		}
	}
	return nil
}
