package protogen

import (
	"context"
	"fmt"
	"path"
	"sort"
	"strings"
	"unicode"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/gapic/pkg/openapi"
	"google.golang.org/protobuf/encoding/prototext"
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

func ToProtoFieldName(s string) string {
	var b []byte
	for _, c := range s {
		if unicode.IsUpper(c) {
			b = append(b, '_')
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

func (c *OpenAPIConverter) buildServiceFromOpenAPI(resourceName string, resource *openapi.Resource) error {

	resourceName = ToProtoIdentifier(resourceName)

	service := &descriptorpb.ServiceDescriptorProto{
		Name: PtrTo(resourceName),
	}

	for _, methodName := range sortedKeys(resource.Methods) {
		methodName := methodName
		method := resource.Methods[methodName]

		serviceMethod := &descriptorpb.MethodDescriptorProto{}
		switch methodName {
		case "create":
			serviceMethod.Name = PtrTo("Create" + resourceName)
		case "get":
			serviceMethod.Name = PtrTo("Get" + resourceName)
		case "list":
			serviceMethod.Name = PtrTo("List" + resourceName)
		case "delete":
			serviceMethod.Name = PtrTo("Delete" + resourceName)
		case "patch":
			serviceMethod.Name = PtrTo("Patch" + resourceName)
		case "update":
			serviceMethod.Name = PtrTo("Update" + resourceName)
		case "testIamPermissions", "getIamPolicy", "setIamPolicy":
			klog.Warningf("skipping method %q", methodName)
			continue
		default:
			klog.Fatalf("unhandled methodName %q", methodName)
		}
		if method.Request != nil {
			if method.Request.Ref == "" {
				klog.Fatalf("unexpected method Request: %+v", method)
			}
			serviceMethod.InputType = PtrTo(method.Request.Ref)
		} else {
			requestType := *serviceMethod.Name + "Request"
			serviceMethod.InputType = &requestType
			desc := &descriptorpb.DescriptorProto{}
			desc.Name = &requestType
			c.fileDescriptor.MessageType = append(c.fileDescriptor.MessageType, desc)
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
		service.Method = append(service.Method, serviceMethod)

		klog.Infof("%s/%s => %v", resourceName, methodName, prototext.Format(serviceMethod))
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
