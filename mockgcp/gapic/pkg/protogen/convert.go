package protogen

import (
	"context"
	"fmt"
	"path"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/gapic/pkg/openapi"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"k8s.io/klog/v2"
)

type OpenAPIConverter struct {
	doc *openapi.Document

	fileDescriptor *descriptorpb.FileDescriptorProto
}

func NewOpenAPIConverter(doc *openapi.Document) *OpenAPIConverter {
	return &OpenAPIConverter{doc: doc}
}

func (c *OpenAPIConverter) resolveProperty(name string) *openapi.Property {
	prop, found := c.doc.Schemas[name]
	if !found {
		klog.Fatalf("unable to resolve property %q", name)
	}
	return prop
}

func (c *OpenAPIConverter) buildMessageFromOpenAPI(message *openapi.Property) (*descriptorpb.DescriptorProto, error) {
	nextTag := int32(1)
	desc := &descriptorpb.DescriptorProto{}
	desc.Name = PtrTo(message.ID)
	for propertyID, property := range message.Properties {
		propertyID := propertyID
		tag := nextTag
		nextTag++

		field := &descriptorpb.FieldDescriptorProto{
			Number: &tag,
		}
		field.JsonName = PtrTo(propertyID)
		field.Name = PtrTo(ToProtoFieldName(propertyID))

		switch property.Type {
		case "object":
			if property.Ref != "" {
				resolved := c.resolveProperty(property.Ref)
				field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
				field.TypeName = &resolved.ID

			} else if property.AdditionalProperties != nil {
				switch property.AdditionalProperties.Type {
				case "string":
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
						mapMessage.Field = append(mapMessage.Field, &descriptorpb.FieldDescriptorProto{
							Name:   PtrTo("value"),
							Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
							Number: PtrTo[int32](2),
						})

						desc.NestedType = append(desc.NestedType, mapMessage)
					}

					field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
					field.TypeName = &mapType
					field.Label = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()

				default:
					klog.Fatalf("unhandled additionalProperties for object: %+v", property)
				}
			} else {
				klog.Fatalf("expected property.Ref to be set for object: %+v", property)
			}

		case "":
			// TODO: Combine with object?
			if property.Ref != "" {
				resolved := c.resolveProperty(property.Ref)
				field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
				field.TypeName = &resolved.ID
			} else {
				klog.Fatalf("expected property.Ref to be set for empty type: %+v", property)
			}
		case "string":
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()

		case "boolean":
			field.Type = descriptorpb.FieldDescriptorProto_TYPE_BOOL.Enum()

		case "integer":
			switch property.Format {
			case "uint32":
				field.Type = descriptorpb.FieldDescriptorProto_TYPE_UINT32.Enum()
			case "int32":
				field.Type = descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum()
			default:
				klog.Fatalf("unhandled property.Format in %+v", property)
			}
		case "number":
			switch property.Format {
			case "double":
				field.Type = descriptorpb.FieldDescriptorProto_TYPE_DOUBLE.Enum()
			default:
				klog.Fatalf("unhandled property.Format in %+v", property)
			}
		case "array":
			field.Label = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()

			if property.Ref != "" {
				resolved := c.resolveProperty(property.Ref)
				field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
				field.TypeName = &resolved.ID
			} else if property.Items != nil {
				if property.Items.Ref != "" {
					resolved := c.resolveProperty(property.Items.Ref)
					field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
					field.TypeName = &resolved.ID
				} else {
					switch property.Items.Type {
					case "string":
						field.Type = descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()

					default:
						klog.Fatalf("unhandled property.Items for array: %+v", property.Items)
					}
				}
			} else {
				klog.Fatalf("expected property.Ref to be set for array: %+v", property)
			}

		default:
			klog.Fatalf("unsupported property.Type %q %+v", property.Type, property)
		}

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
		b = append(b, byte(c))
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

	service := &descriptorpb.ServiceDescriptorProto{
		Name: PtrTo(ToProtoIdentifier(pluralName)),
	}

	for _, methodName := range sortedKeys(resource.Methods) {
		methodName := methodName
		method := resource.Methods[methodName]

		serviceMethod := &descriptorpb.MethodDescriptorProto{}
		serviceMethod.Options = &descriptorpb.MethodOptions{}

		httpPath := method.FlatPath
		httpPath = "/" + httpPath

		parameterRenames := make(map[string]string)
		parameterRenames[ToProtoFieldName(singularName)] = "name"

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
		case "create":
			serviceMethod.Name = PtrTo("Create" + StartWithUpper(singularName))
		case "get":
			serviceMethod.Name = PtrTo("Get" + StartWithUpper(singularName))
		case "list":
			serviceMethod.Name = PtrTo("List" + StartWithUpper(pluralName))
		case "delete":
			serviceMethod.Name = PtrTo("Delete" + StartWithUpper(singularName))
		case "patch":
			serviceMethod.Name = PtrTo("Patch" + StartWithUpper(singularName))
		case "update":
			serviceMethod.Name = PtrTo("Update" + StartWithUpper(singularName))
		case "testIamPermissions", "getIamPolicy", "setIamPolicy":
			klog.Warningf("skipping method %q", methodName)
			continue
		default:
			klog.Fatalf("unhandled methodName %q", methodName)
		}

		{
			requestTypeName := *serviceMethod.Name + "Request"
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

			for parameterName, parameter := range method.Parameters {

				switch parameter.Location {
				case "path":
					tag := nextTag
					nextTag++

					field := &descriptorpb.FieldDescriptorProto{}
					field.Number = PtrTo(tag)
					//field.JsonName = PtrTo(parameterName)
					field.Name = PtrTo(ToProtoFieldName(parameterName))
					if rename := parameterRenames[field.GetName()]; rename != "" {
						field.Name = PtrTo(rename)
					}
					switch parameter.Type {
					case "string":
						field.Type = descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
					default:
						klog.Fatalf("parameter type not recognized %+v", parameter)
					}

					desc.Field = append(desc.Field, field)
				case "query":
					klog.Warningf("ignoring parameter %+v", parameter)
				default:
					klog.Fatalf("parameter location not recognized %+v", parameter)
				}
			}

			if method.Request != nil {
				if method.Request.Ref == "" {
					klog.Fatalf("unexpected method Request: %+v", method)
				}
				tag := nextTag
				nextTag++

				field := &descriptorpb.FieldDescriptorProto{}
				field.Number = PtrTo(tag)
				//field.JsonName = PtrTo(parameterName)
				field.Name = PtrTo(ToProtoFieldName(singularName))
				field.Type = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
				field.TypeName = PtrTo(method.Request.Ref)

				desc.Field = append(desc.Field, field)

				httpRule.Body = ToProtoFieldName(singularName)
				// serviceMethod.InputType = PtrTo(method.Request.Ref)
			}
		}

		if method.Response != nil {
			if method.Response.Ref == "" {
				klog.Fatalf("unexpected method Response: %+v", method)
			}
			serviceMethod.OutputType = PtrTo(method.Response.Ref)
		} else {
			responseType := *serviceMethod.Name + "Response"
			serviceMethod.OutputType = &responseType

			desc := &descriptorpb.DescriptorProto{}
			desc.Name = &responseType
			c.fileDescriptor.MessageType = append(c.fileDescriptor.MessageType, desc)
		}

		proto.SetExtension(serviceMethod.Options, annotations.E_Http, httpRule)

		service.Method = append(service.Method, serviceMethod)

		klog.Infof("%s/%s => %v", singularName, methodName, prototext.Format(serviceMethod))
	}

	c.fileDescriptor.Service = append(c.fileDescriptor.Service, service)
	return nil
}

func (c *OpenAPIConverter) Convert(ctx context.Context) (*descriptorpb.FileDescriptorProto, error) {

	{
		version := c.doc.Version
		serviceID := c.doc.Name
		prefix := []string{"google", "cloud"}

		name := path.Join(path.Join(prefix...), serviceID, version, serviceID+"pb.proto")
		packageName := strings.Join(prefix, ".") + "." + serviceID + "." + version
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
			desc, err := c.buildMessageFromOpenAPI(message)
			if err != nil {
				return nil, fmt.Errorf("buildMessageFromOpenAPI failed: %w", err)
			}
			c.fileDescriptor.MessageType = append(c.fileDescriptor.MessageType, desc)

			klog.Infof("%s => %+v\n", schemaName, prototext.Format(desc))
		} else {
			klog.Fatalf("unexpected type %q in doc.Schemas", message.Type)
		}
	}

	for _, resourceName := range sortedKeys(c.doc.Resources) {
		resource := c.doc.Resources[resourceName]
		if err := c.buildServiceFromOpenAPI(resourceName, resource); err != nil {
			return nil, fmt.Errorf("buildServiceFromOpenAPI failed: %w", err)
		}

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
