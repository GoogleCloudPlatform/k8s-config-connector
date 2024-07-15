// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protogen

import (
	"context"
	"fmt"
	"path"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/tools/gapic/pkg/openapi"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"k8s.io/klog/v2"
)

type OpenAPIConverter struct {
	doc *openapi.Document

	imports map[string]bool

	fileDescriptor *descriptorpb.FileDescriptorProto

	// Comments holds field and message level comments
	Comments

	// protoPackageName is the name of the proto package we are generating.
	protoPackageName string
}

// Comments holds all the comments for our messages / fields
type Comments struct {
	comments map[string]*Comment
}

func (c *Comments) SetComment(key string, comment string) {
	c.comments[key] = &Comment{Text: comment}

}
func (c *Comments) GetComment(key string) string {
	o, found := c.comments[key]
	if !found {
		return ""
	}
	return o.Text
}

// Comment is a comment on a message/field
type Comment struct {
	Text string
}

func NewOpenAPIConverter(protoPackageName string, doc *openapi.Document) *OpenAPIConverter {
	return &OpenAPIConverter{
		doc:              doc,
		protoPackageName: protoPackageName,
		imports:          make(map[string]bool),
		Comments: Comments{
			comments: make(map[string]*Comment),
		},
	}
}

func (c *OpenAPIConverter) addImport(s string) {
	c.imports[s] = true
}

func (c *OpenAPIConverter) buildMessageFromOpenAPI(message *openapi.Property) (*descriptorpb.DescriptorProto, error) {
	nextTag := int32(1)
	desc := &descriptorpb.DescriptorProto{}
	desc.Name = PtrTo(message.ID)
	c.SetComment(c.protoPackageName+"."+desc.GetName(), message.Description)
	for _, entry := range message.Properties.Entries() {
		propertyID := entry.Key
		property := entry.Value

		tag := nextTag
		nextTag++

		field := &descriptorpb.FieldDescriptorProto{
			Number: &tag,
		}
		field.JsonName = PtrTo(propertyID)
		field.Name = PtrTo(ToProtoFieldName(propertyID))

		switch property.Type {
		case "string", "boolean", "integer", "number":
			c.setPrimitiveType(property, field)

		case "object":
			if property.Ref != "" {
				typeName := c.resolveMessageType(property.Ref)
				field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
				field.TypeName = &typeName

			} else if property.AdditionalProperties != nil {
				valueField := &descriptorpb.FieldDescriptorProto{
					Name:   PtrTo("value"),
					Number: PtrTo[int32](2),
				}
				switch property.AdditionalProperties.Type {
				case "string", "boolean", "integer", "number":
					c.setPrimitiveType(property.AdditionalProperties, valueField)

				case "any":
					valueField.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
					valueField.TypeName = PtrTo("google.protobuf.Any")
					c.addImport("google/protobuf/any.proto")

				case "":
					if property.AdditionalProperties.Ref != "" {
						valueMessage := c.resolveMessageType(property.AdditionalProperties.Ref)

						valueField.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
						valueField.TypeName = PtrTo(valueMessage)
					}
				}

				if valueField.Type == nil {
					klog.Fatalf("unhandled additionalProperties for object: %+v", property)
				}

				mapType := MapEntryName(propertyID)

				{
					mapMessage := &descriptorpb.DescriptorProto{}
					mapMessage.Name = &mapType
					mapMessage.Options = &descriptorpb.MessageOptions{
						MapEntry: PtrTo(true),
					}

					mapMessage.Field = append(mapMessage.Field, &descriptorpb.FieldDescriptorProto{
						Name:   PtrTo("key"),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
						Number: PtrTo[int32](1),
					})
					mapMessage.Field = append(mapMessage.Field, valueField)

					desc.NestedType = append(desc.NestedType, mapMessage)
				}

				field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
				field.TypeName = &mapType
				field.Label = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()
			} else if len(property.Properties.Entries()) != 0 {
				// An inline definition
				inlineMessage, err := c.buildMessageFromOpenAPI(property)
				if err != nil {
					return nil, err
				}
				typeName := message.ID + StartWithUpper(propertyID)
				inlineMessage.Name = &typeName
				c.fileDescriptor.MessageType = append(c.fileDescriptor.MessageType, inlineMessage)

				field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
				field.TypeName = &typeName
			} else {
				klog.Fatalf("expected property.Ref to be set for object: %+v", property)
			}

		case "":
			// TODO: Combine with object?
			if property.Ref != "" {
				typeName := c.resolveMessageType(property.Ref)
				field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
				field.TypeName = &typeName
			} else {
				klog.Fatalf("expected property.Ref to be set for empty type: %+v", property)
			}

		case "array":
			field.Label = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()

			if property.Ref != "" {
				typeName := c.resolveMessageType(property.Ref)
				field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
				field.TypeName = &typeName
			} else if property.Items != nil {
				if property.Items.Ref != "" {
					typeName := c.resolveMessageType(property.Items.Ref)
					field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
					field.TypeName = &typeName
				} else {
					switch property.Items.Type {
					case "string", "boolean", "integer", "number":
						c.setPrimitiveType(property.Items, field)

					case "object":
						if property.Items.AdditionalProperties != nil {
							switch property.Items.AdditionalProperties.Type {
							case "any":
								field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
								field.TypeName = PtrTo("google.protobuf.Any")
								c.addImport("google/protobuf/any.proto")
							default:
								klog.Fatalf("unhandled property.Format in %+v", property)
							}
						} else if len(property.Items.Properties.Entries()) != 0 {
							// An inline definition
							inlineMessage, err := c.buildMessageFromOpenAPI(property.Items)
							if err != nil {
								return nil, err
							}
							typeName := message.ID + StartWithUpper(propertyID)
							inlineMessage.Name = &typeName
							c.fileDescriptor.MessageType = append(c.fileDescriptor.MessageType, inlineMessage)

							field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
							field.TypeName = &typeName
						} else {
							klog.Fatalf("unhandled property.Format in %+v", property)
						}

					case "any":
						field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
						field.TypeName = PtrTo("google.protobuf.Any")
						c.addImport("google/protobuf/any.proto")

					case "array":
						if property.Items.Items != nil {
							switch property.Items.Items.Type {
							case "any":
								field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
								field.TypeName = PtrTo("google.protobuf.Any")
								c.addImport("google/protobuf/any.proto")
							default:
								klog.Fatalf("unhandled property.Items.Items.Type in %+v", property)
							}
						} else {
							klog.Fatalf("unhandled property.Items in %+v", property)
						}

					default:
						klog.Fatalf("unhandled property.Items for array: %+v", property)
					}
				}
			} else {
				klog.Fatalf("expected property.Ref to be set for array: %+v", property)
			}

		case "any":
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
			field.TypeName = PtrTo("google.protobuf.Any")
			c.addImport("google/protobuf/any.proto")

		default:
			klog.Fatalf("unsupported property.Type %q %+v", property.Type, property)
		}

		c.SetComment(c.protoPackageName+"."+desc.GetName()+"."+field.GetName(), property.Description)
		desc.Field = append(desc.Field, field)
	}

	return desc, nil
}

func PtrTo[T any](val T) *T {
	return &val
}

func ToProtoIdentifier(s string) string {
	var b []byte
	upperNext := true
	for _, c := range s {
		switch {
		case c == '_':
			upperNext = true
		case upperNext:
			b = append(b, byte(unicode.ToUpper(c)))
			upperNext = false
		default:
			b = append(b, byte(c))
		}
	}
	return string(b)
}

func StartWithUpper(s string) string {
	var b []byte
	for i, c := range s {
		if i == 0 {
			b = append(b, byte(unicode.ToUpper(c)))
		} else {
			b = append(b, byte(c))
		}
	}
	return string(b)
}

func ToProtoFieldName(s string) string {
	var b []byte
	for i, c := range s {
		if unicode.IsUpper(c) {
			if i != 0 {
				b = append(b, '_')
			}
			c = unicode.ToLower(c)
		}
		switch c {
		case '.':
			b = append(b, '_')

		default:
			b = append(b, byte(c))
		}
	}
	return string(b)
}

// See https://github.com/protocolbuffers/protobuf-go/blob/v1.32.0/internal/strs/strings.go#L125
func MapEntryName(s string) string {
	var b []byte
	upperNext := true
	for _, c := range s {
		switch {
		case c == '_':
			upperNext = true
		case upperNext:
			b = append(b, byte(unicode.ToUpper(c)))
			upperNext = false
		default:
			b = append(b, byte(c))
		}
	}
	b = append(b, "Entry"...)
	return string(b)
}

func (c *OpenAPIConverter) findMessageDescriptor(name string) *descriptorpb.DescriptorProto {
	for _, md := range c.fileDescriptor.MessageType {
		if md.GetName() == name {
			return md
		}
	}
	return nil
}

func (c *OpenAPIConverter) buildServiceFromOpenAPI(pluralName string, resource *openapi.Resource) error {
	singularName := pluralName
	if strings.HasSuffix(singularName, "ies") {
		singularName = strings.TrimSuffix(singularName, "ies") + "y"
	} else if strings.HasSuffix(singularName, "s") {
		singularName = strings.TrimSuffix(singularName, "s")
	}

	serviceName := ToProtoIdentifier(pluralName)
	if !strings.HasSuffix(serviceName, "Server") {
		serviceName += "Server"
	}

	service := &descriptorpb.ServiceDescriptorProto{
		Name: PtrTo(serviceName),
	}

	for _, methodName := range sortedKeys(resource.Methods) {
		methodName := methodName
		method := resource.Methods[methodName]

		serviceMethod := &descriptorpb.MethodDescriptorProto{}
		serviceMethod.Options = &descriptorpb.MethodOptions{}

		httpPath := method.Path
		if httpPath == "" {
			httpPath = method.FlatPath
		}
		basePath := c.doc.BasePath
		if !strings.HasPrefix(basePath, "/") {
			basePath = "/" + basePath
		}
		httpPath = basePath + httpPath

		parameterRenames := make(map[string]string)
		if ToProtoFieldName(singularName) != "project" { // Otherwise doesn't line up with parameters
			parameterRenames[ToProtoFieldName(singularName)] = "name"
		}

		// Map path parameters from json to proto syntax
		{
			r := regexp.MustCompile(`{[^}]+}`)
			httpPath = r.ReplaceAllStringFunc(httpPath, func(match string) string {
				match = strings.TrimPrefix(match, "{")
				match = strings.TrimSuffix(match, "}")
				id := ToProtoFieldName(match)
				if rename := parameterRenames[id]; rename != "" {
					id = rename
				}
				return "{" + id + "}"
			})
		}

		// Replace {+project_id} => {project_id} (and other parameters)
		{
			r := regexp.MustCompile(`{\+[^}]+}`)
			httpPath = r.ReplaceAllStringFunc(httpPath, func(match string) string {
				match = strings.TrimPrefix(match, "{+")
				match = strings.TrimSuffix(match, "}")

				expansion := "*"
				if parameter := method.Parameters.Get(match); parameter != nil {
					if parameter.Pattern != "" {
						pattern := parameter.Pattern
						pattern = strings.TrimPrefix(pattern, "^")
						pattern = strings.TrimSuffix(pattern, "$")
						pattern = strings.ReplaceAll(pattern, "[^/]+", "*")
						expansion = pattern
					}
				}

				return "{" + match + "=" + expansion + "}"
			})
		}

		httpRule := &annotations.HttpRule{}
		switch method.HTTPMethod {
		case "PUT":
			httpRule.Pattern = &annotations.HttpRule_Put{
				Put: httpPath,
			}
		case "POST":
			httpRule.Pattern = &annotations.HttpRule_Post{
				Post: httpPath,
			}
		case "GET":
			httpRule.Pattern = &annotations.HttpRule_Get{
				Get: httpPath,
			}
		case "DELETE":
			httpRule.Pattern = &annotations.HttpRule_Delete{
				Delete: httpPath,
			}
		case "PATCH":
			httpRule.Pattern = &annotations.HttpRule_Patch{
				Patch: httpPath,
			}
		default:
			klog.Fatalf("unhandled HTTPMethod %q in %+v", method.HTTPMethod, method)
		}
		switch methodName {
		case "create", "add":
			serviceMethod.Name = PtrTo(StartWithUpper(methodName) + StartWithUpper(singularName))
		case "get":
			serviceMethod.Name = PtrTo("Get" + StartWithUpper(singularName))
		case "list":
			serviceMethod.Name = PtrTo("List" + StartWithUpper(pluralName))
		case "delete", "remove":
			serviceMethod.Name = PtrTo("Delete" + StartWithUpper(singularName))
		case "patch":
			serviceMethod.Name = PtrTo("Patch" + StartWithUpper(singularName))
		case "update":
			serviceMethod.Name = PtrTo("Update" + StartWithUpper(singularName))
		case "testIamPermissions", "getIamPolicy", "setIamPolicy":
			klog.Warningf("skipping method %q", methodName)
			continue

			// TODO: Just toUpper it?
		// case "cancel":
		// 	serviceMethod.Name = PtrTo("Cancel" + StartWithUpper(singularName))
		// case "addSubnetwork":
		// 	serviceMethod.Name = PtrTo("AddSubnetwork" + StartWithUpper(singularName))
		// case "disableVpcServiceControls":
		// 	serviceMethod.Name = PtrTo("DisableVpcServiceControls" + StartWithUpper(singularName))
		// case "enableVpcServiceControls":
		// 	serviceMethod.Name = PtrTo("EnableVpcServiceControls" + StartWithUpper(singularName))
		// case "searchRange":
		// 	serviceMethod.Name = PtrTo("SearchRange" + StartWithUpper(singularName))
		// case "validate":
		// 	serviceMethod.Name = PtrTo("Validate" + StartWithUpper(singularName))

		default:
			serviceMethod.Name = PtrTo(StartWithUpper(methodName) + StartWithUpper(singularName))
			// klog.Fatalf("unhandled methodName %q", methodName)
		}

		{
			requestTypeName := *serviceMethod.Name + "Request"
			if method.Request != nil {
				if method.Request.Ref == requestTypeName {
					requestTypeName = *serviceMethod.Name + "ServiceRequest"
				}
			}

			nextTag := int32(1)

			desc := c.findMessageDescriptor(requestTypeName)
			if desc == nil {
				desc = &descriptorpb.DescriptorProto{}
				desc.Name = PtrTo(requestTypeName)
				c.fileDescriptor.MessageType = append(c.fileDescriptor.MessageType, desc)
			} else {
				for _, fd := range desc.Field {
					if fd.GetNumber() >= nextTag {
						nextTag++
					}
				}
			}
			serviceMethod.InputType = &requestTypeName

			for _, parameterEntry := range method.Parameters.Entries() {
				parameterName := parameterEntry.Key
				parameter := parameterEntry.Value

				addParameter := false

				fieldName := ToProtoFieldName(parameterName)
				if rename := parameterRenames[fieldName]; rename != "" {
					fieldName = rename
				}

				switch parameter.Location {
				case "path":
					addParameter = true
				case "query":
					addParameter = true
					// Only if field not explicitly declared
					for _, f := range desc.Field {
						if f.GetName() == fieldName {
							addParameter = false
						}
					}

				default:
					klog.Fatalf("parameter location not recognized %+v", parameter)
				}

				if addParameter {
					tag := nextTag
					nextTag++

					field := &descriptorpb.FieldDescriptorProto{}
					field.Number = PtrTo(tag)
					field.Name = PtrTo(fieldName)
					// TODO: Merge parameter and property?
					switch parameter.Type {
					case "string":
						field.Type = descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
					case "boolean":
						field.Type = descriptorpb.FieldDescriptorProto_TYPE_BOOL.Enum()
					case "integer":
						field.Type = descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum()
					default:
						klog.Fatalf("parameter type not recognized %+v", parameter)
					}

					desc.Field = append(desc.Field, field)
				}
			}

			if method.Request != nil {
				if method.Request.Ref == "" {
					klog.Fatalf("unexpected method Request: %+v", method)
				}

				tag := nextTag
				nextTag++

				bodyFieldName := ToProtoFieldName(singularName)
				for _, f := range desc.Field {
					if f.GetName() == bodyFieldName {
						bodyFieldName = bodyFieldName + "_body"
						break
					}
				}
				field := &descriptorpb.FieldDescriptorProto{}
				field.Number = PtrTo(tag)
				//field.JsonName = PtrTo(parameterName)
				field.Name = PtrTo(bodyFieldName)
				field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
				field.TypeName = PtrTo(method.Request.Ref)

				desc.Field = append(desc.Field, field)

				httpRule.Body = bodyFieldName
				// serviceMethod.InputType = PtrTo(method.Request.Ref)
			}
		}

		if method.Response != nil {
			if method.Response.Ref == "" {
				klog.Fatalf("unexpected method Response: %+v", method)
			}
			serviceMethod.OutputType = PtrTo(c.resolveMessageType(method.Response.Ref))
		} else {
			serviceMethod.OutputType = PtrTo("google.protobuf.Empty")
			c.addImport("google/protobuf/empty.proto")
		}

		proto.SetExtension(serviceMethod.Options, annotations.E_Http, httpRule)

		c.SetComment(c.protoPackageName+"."+service.GetName()+"."+serviceMethod.GetName(), method.Description)

		service.Method = append(service.Method, serviceMethod)

		klog.V(4).Infof("%s/%s => %v", singularName, methodName, prototext.Format(serviceMethod))
	}

	c.fileDescriptor.Service = append(c.fileDescriptor.Service, service)
	return nil
}

func (c *OpenAPIConverter) Convert(ctx context.Context) (*descriptorpb.FileDescriptorProto, error) {

	{
		tokens := strings.Split(c.protoPackageName, ".")
		if len(tokens) < 2 {
			return nil, fmt.Errorf("unexpected proto package name %q (expected more tokens)", c.protoPackageName)
		}
		name := path.Join(tokens...) + ".pb.proto"
		packageName := c.protoPackageName
		serviceID := tokens[len(tokens)-2]
		version := tokens[len(tokens)-1]
		goPackageName := path.Join("cloud.google.com/go", serviceID, "api"+version, serviceID+"pb") + ";" + serviceID + "pb"
		c.fileDescriptor = &descriptorpb.FileDescriptorProto{
			Name:    PtrTo(name),
			Package: PtrTo(packageName),
			Options: &descriptorpb.FileOptions{
				GoPackage: PtrTo(goPackageName),
			},
		}
	}

	for _, schemaName := range sortedKeys(c.doc.Schemas) {
		message := c.doc.Schemas[schemaName]

		if message.Type == "object" {
			if !c.isWellKnown(message.ID) {
				desc, err := c.buildMessageFromOpenAPI(message)
				if err != nil {
					return nil, fmt.Errorf("buildMessageFromOpenAPI failed: %w", err)
				}
				c.fileDescriptor.MessageType = append(c.fileDescriptor.MessageType, desc)

				klog.V(4).Infof("%s => %+v\n", schemaName, prototext.Format(desc))
			} else {
				klog.Infof("skipping well known message %q", schemaName)
			}
		} else if message.Type == "any" {
			klog.Warningf("skipping schema with type any: %q", message.ID)
		} else {
			klog.Fatalf("unexpected type %q in doc.Schemas", message.Type)
		}
	}

	if err := c.visitResources(ctx, "", c.doc.Resources); err != nil {
		return nil, err
	}

	for k := range c.imports {
		c.fileDescriptor.Dependency = append(c.fileDescriptor.Dependency, k)
	}

	return c.fileDescriptor, nil
}

func sortedKeys[V any](m map[string]V) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func (c *OpenAPIConverter) visitResources(ctx context.Context, namePrefix string, resources map[string]*openapi.Resource) error {
	for _, resourceName := range sortedKeys(resources) {
		if resourceName == "operations" {
			klog.Warningf("skipping operations resource")
			continue
		}
		resource := resources[resourceName]
		if err := c.buildServiceFromOpenAPI(namePrefix+StartWithUpper(resourceName), resource); err != nil {
			return err
		}

		if err := c.visitResources(ctx, namePrefix+StartWithUpper(resourceName), resource.Resources); err != nil {
			return err
		}
	}

	return nil
}

func (c *OpenAPIConverter) resolveMessageType(ref string) string {
	_, found := c.doc.Schemas[ref]
	if !found {
		klog.Fatalf("unable to resolve property ref %q", ref)
	}

	switch ref {
	case "Operation", "GoogleLongrunningOperation":
		c.addImport("google/longrunning/operations.proto")
		return "google.longrunning.Operation"

	default:
		return ref
	}
}

func (c *OpenAPIConverter) isWellKnown(ref string) bool {
	if ref == "" {
		return false
	}
	resolved := c.resolveMessageType(ref)
	switch resolved {
	case "google.longrunning.Operation":
		return true
	}

	return false
}

func (c *OpenAPIConverter) setPrimitiveType(property *openapi.Property, field *descriptorpb.FieldDescriptorProto) {

	switch property.Type {

	case "string":
		switch property.Format {
		case "int64":
			// JSON can't represent the full 64 bit range.
			// 64 bit protos are represented in JSON (and thus OpenAPI) as string.
			// But, thankfully, we get a hint that this happened.
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_INT64.Enum()
		case "uint64":
			// Same story as int64
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_UINT64.Enum()
		case "byte":
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_BYTES.Enum()
		case "":
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
		case "date-time", "google-datetime":
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
			field.TypeName = PtrTo("google.protobuf.Timestamp")
			c.addImport("google/protobuf/timestamp.proto")

		case "date":
			// TODO: use google.type.Date?
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()

		case "google-duration":
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
			field.TypeName = PtrTo("google.protobuf.Duration")
			c.addImport("google/protobuf/duration.proto")

		case "google-fieldmask":
			// TODO: use fieldmask type?
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
		default:
			klog.Fatalf("unhandled property.Format %q in %+v", property.Format, property)
		}

	case "integer":
		switch property.Format {
		case "uint32":
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_UINT32.Enum()
		case "int32":
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum()
		default:
			klog.Fatalf("unhandled property.Format %q in %+v", property.Format, property)
		}

	case "number":
		switch property.Format {
		case "double":
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_DOUBLE.Enum()
		case "float":
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_FLOAT.Enum()
		default:
			klog.Fatalf("unhandled property.Format %q in %+v", property.Format, property)
		}

	case "boolean":
		switch property.Format {
		case "":
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_BOOL.Enum()
		default:
			klog.Fatalf("unhandled property.Format %q in %+v", property.Format, property)
		}

	default:
		klog.Fatalf("unhandled primitive property type %q in %+v", property.Type, property)
	}
}
