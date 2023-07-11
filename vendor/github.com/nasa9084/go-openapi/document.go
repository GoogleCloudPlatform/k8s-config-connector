package openapi

import (
	"sort"
	"strconv"
	"strings"
)

// codebeat:disable[TOO_MANY_IVARS]

// Document represents a OpenAPI Specification document.
type Document struct {
	Version      string `yaml:"openapi"`
	Info         *Info
	Servers      []*Server
	Paths        Paths
	Components   *Components
	Security     []*SecurityRequirement
	Tags         []*Tag
	ExternalDocs *ExternalDocumentation `yaml:"externalDocs"`
}

// Validate the values of spec.
func (doc Document) Validate() error {
	if err := doc.validateRequiredFields(); err != nil {
		return err
	}
	if err := doc.validateOASVersion(); err != nil {
		return err
	}
	return doc.validateFields()
}

func (doc Document) validateOASVersion() error {
	splited := strings.FieldsFunc(doc.Version, func(r rune) bool { return r == '.' })
	if len(splited) != 3 {
		return ErrFormatInvalid{Target: "openapi version", Format: "X.Y.Z"}
	}
	major, err := strconv.Atoi(splited[0])
	if err != nil {
		return ErrFormatInvalid{Target: "major part of openapi version"}
	}
	minor, err := strconv.Atoi(splited[1])
	if err != nil {
		return ErrFormatInvalid{Target: "minor part of openapi version"}
	}
	_, err = strconv.Atoi(splited[2])
	if err != nil {
		return ErrFormatInvalid{Target: "patch part of openapi version"}
	}
	if major == 3 && 0 <= minor {
		return nil
	}
	return ErrUnsupportedVersion
}

func (doc Document) validateRequiredFields() error {
	if doc.Version == "" {
		return ErrRequired{Target: "openapi"}
	}
	if doc.Info == nil {
		return ErrRequired{Target: "info"}
	}
	if doc.Paths == nil {
		return ErrRequired{Target: "paths"}
	}
	return nil
}

func (doc Document) validateFields() error {
	var validaters []validater
	validaters = append(validaters, doc.Info)
	for _, s := range doc.Servers {
		validaters = append(validaters, s)
	}
	validaters = append(validaters, doc.Paths)
	if doc.Components != nil {
		validaters = append(validaters, doc.Components)
	}
	for _, securityRequirement := range doc.Security {
		validaters = append(validaters, securityRequirement)
	}
	for _, t := range doc.Tags {
		validaters = append(validaters, t)
	}
	if doc.ExternalDocs != nil {
		validaters = append(validaters, doc.ExternalDocs)
	}
	return validateAll(validaters)
}

type WalkFunc func(doc *Document, method, path string, pathItem *PathItem, op *Operation) error

func (doc *Document) Walk(walkFn WalkFunc) error {
	var paths []string
	for path := range doc.Paths {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	for _, path := range paths {
		pathItem := doc.Paths[path]
		var methods []string
		for method := range pathItem.Operations() {
			methods = append(methods, method)
		}
		sort.Strings(methods)

		for _, method := range methods {
			operation := pathItem.GetOperationByMethod(method)
			if err := walkFn(doc, method, path, pathItem, operation); err != nil {
				return err
			}
		}
	}
	return nil
}
