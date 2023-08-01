package openapi

// codebeat:disable[TOO_MANY_IVARS]

// XML Object
type XML struct {
	Name      string
	Namespace string
	Prefix    string
	Attribute bool
	Wrapped   bool
}

// Validate the values of XML object.
func (xml XML) Validate() error {
	return mustURL("xml.namespace", xml.Namespace)
}
