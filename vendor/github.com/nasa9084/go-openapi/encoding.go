package openapi

// codebeat:disable[TOO_MANY_IVARS]

// Encoding Object
type Encoding struct {
	ContentType   string `yaml:"contentType"`
	Headers       map[string]*Header
	Style         string
	Explode       bool
	AllowReserved bool `yaml:"allowReserved"`
}

// Validate the values of Encoding object.
func (encoding Encoding) Validate() error {
	for _, header := range encoding.Headers {
		if err := header.Validate(); err != nil {
			return err
		}
	}
	return nil
}
