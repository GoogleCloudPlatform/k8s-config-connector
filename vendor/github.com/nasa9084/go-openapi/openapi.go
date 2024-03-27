package openapi

import (
	"bytes"
	"io"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// LoadFile OpenAPI Specification v3.0 spec file.
func LoadFile(filename string) (*Document, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, f); err != nil {
		return nil, err
	}
	b := buf.Bytes()

	if err := f.Close(); err != nil {
		panic(err)
	}

	return Load(b)
}

// Load OpenAPI Specification v3.0 spec.
func Load(b []byte) (*Document, error) {
	doc := &Document{}
	if err := yaml.Unmarshal(b, doc); err != nil {
		return nil, err
	}
	// If the servers property is not provided, or is an empty array, the default value would be a Server Object with a url value of /.
	// see: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#oasObject
	if doc.Servers == nil || len(doc.Servers) == 0 {
		doc.Servers = []*Server{&Server{URL: "/"}}
	}
	for i := range doc.Security {
		doc.Security[i].setDocument(doc)
	}
	for _, pi := range doc.Paths {
		for _, op := range pi.Operations() {
			for _, sr := range op.Security {
				sr.setDocument(doc)
			}
		}
	}
	return doc, nil
}
