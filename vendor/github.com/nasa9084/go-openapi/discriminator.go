package openapi

// codebeat:disable[TOO_MANY_IVARS]

// Discriminator Object
type Discriminator struct {
	PropertyName string `yaml:"propertyName"`
	Mapping      map[string]string
}

// Validate the values of Descriminator object.
func (discriminator Discriminator) Validate() error {
	if discriminator.PropertyName == "" {
		return ErrRequired{Target: "discriminator.propertyName"}
	}
	return nil
}
