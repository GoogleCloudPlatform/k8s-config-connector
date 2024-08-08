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

package codegen

import (
	"fmt"
	"io"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

type MessageWriter interface {
	WriteMessage(msg protoreflect.MessageDescriptor)
	WriteField(field protoreflect.FieldDescriptor, msg protoreflect.MessageDescriptor, fieldIndex int)
}

type messageWriter struct {
	opts *options
}

type options struct {
	out             io.Writer
	referenceFields map[string]string
}

type MessageWriterOption func(o *options)

func NewMessageWriter(opts ...MessageWriterOption) MessageWriter {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return &messageWriter{opts: o}
}

func WithReferenceFields(rfs map[string]string) MessageWriterOption {
	return func(o *options) {
		o.referenceFields = rfs
	}
}

func WithOutput(out io.Writer) MessageWriterOption {
	return func(o *options) {
		o.out = out
	}
}

func (w *messageWriter) WriteMessage(msg protoreflect.MessageDescriptor) {
	goType := goNameForProtoMessage(msg, msg)

	fmt.Fprintf(w.opts.out, "\n")
	fmt.Fprintf(w.opts.out, "// +kcc:proto=%s\n", msg.FullName())
	fmt.Fprintf(w.opts.out, "type %s struct {\n", goType)
	for i := 0; i < msg.Fields().Len(); i++ {
		field := msg.Fields().Get(i)
		w.WriteField(field, msg, i)
	}
	fmt.Fprintf(w.opts.out, "}\n")
}

func (w *messageWriter) WriteField(field protoreflect.FieldDescriptor, msg protoreflect.MessageDescriptor, fieldIndex int) {
	sourceLocations := msg.ParentFile().SourceLocations().ByDescriptor(field)

	jsonName := getJSONForKRM(field)
	goFieldName := strings.Title(jsonName)
	goType := ""

	if field.IsMap() {
		entryMsg := field.Message()
		keyKind := entryMsg.Fields().ByName("key").Kind()
		valueKind := entryMsg.Fields().ByName("value").Kind()
		if keyKind == protoreflect.StringKind && valueKind == protoreflect.StringKind {
			goType = "map[string]string"
		} else if keyKind == protoreflect.StringKind && valueKind == protoreflect.Int64Kind {
			goType = "map[string]int64"
		} else {
			fmt.Fprintf(w.opts.out, "\n\n\t// TODO: map type %v %v for %v\n\n", keyKind, valueKind, field.Name())
			return
		}
	} else if referencedGoType, isReferenceField := w.isReferenceField(field); isReferenceField {
		goType = fmt.Sprintf("*refs.%s", referencedGoType)
		jsonName = getJSONForReferencedKRM(field)
		goFieldName = strings.Title(jsonName)
	} else {
		switch field.Kind() {
		case protoreflect.MessageKind:
			goType = goNameForProtoMessage(msg, field.Message())

		case protoreflect.EnumKind:
			goType = "string" //string(field.Enum().Name())

		default:
			goType = goTypeForProtoKind(field.Kind())
		}

		if field.Cardinality() == protoreflect.Repeated {
			goType = "[]" + goType
		} else {
			goType = "*" + goType
		}

		// Special case for proto "bytes" type
		if goType == "*[]byte" {
			goType = "[]byte"
		}
	}

	// Blank line between fields for readability
	if fieldIndex != 0 {
		fmt.Fprintf(w.opts.out, "\n")
	}

	if sourceLocations.LeadingComments != "" {
		comment := strings.TrimSpace(sourceLocations.LeadingComments)
		for _, line := range strings.Split(comment, "\n") {
			if strings.TrimSpace(line) == "" {
				fmt.Fprintf(w.opts.out, "\t//\n")
			} else {
				fmt.Fprintf(w.opts.out, "\t// %s\n", line)
			}
		}
	}

	fmt.Fprintf(w.opts.out, "\t%s %s `json:\"%s,omitempty\"`\n",
		goFieldName,
		goType,
		jsonName,
	)
}

func goNameForProtoMessage(parentMessage protoreflect.MessageDescriptor, msg protoreflect.MessageDescriptor) string {
	fullName := string(msg.FullName())
	fullName = strings.TrimPrefix(fullName, string(parentMessage.ParentFile().FullName()))
	fullName = strings.TrimPrefix(fullName, ".")
	fullName = strings.ReplaceAll(fullName, ".", "_")

	// Some special-case values that are not obvious how to map in KRM
	switch fullName {
	case "google_protobuf_Timestamp":
		return "string"
	case "google_protobuf_Duration":
		return "string"
	case "google_protobuf_Int64Value":
		return "int64"
	}
	return fullName
}

func goTypeForProtoKind(kind protoreflect.Kind) string {
	goType := ""
	switch kind {
	case protoreflect.StringKind:
		goType = "string"

	case protoreflect.Int32Kind:
		goType = "int32"

	case protoreflect.Int64Kind:
		goType = "int64"

	case protoreflect.Uint32Kind:
		goType = "uint32"

	case protoreflect.Uint64Kind:
		goType = "uint64"

	case protoreflect.Fixed64Kind:
		goType = "uint64"

	case protoreflect.BoolKind:
		goType = "bool"

	case protoreflect.DoubleKind:
		goType = "float64"

	case protoreflect.FloatKind:
		goType = "float32"

	case protoreflect.BytesKind:
		goType = "[]byte"

	default:
		klog.Fatalf("unhandled kind %q", kind)
	}

	return goType
}

// getJSONForKRM returns the KRM JSON name for the field,
// honoring KRM conventions
func getJSONForKRM(protoField protoreflect.FieldDescriptor) string {
	jsonName := protoField.JSONName()
	if strings.HasSuffix(jsonName, "Id") {
		jsonName = strings.TrimSuffix(jsonName, "Id") + "ID"
	}
	return jsonName
}

func getJSONForReferencedKRM(protoField protoreflect.FieldDescriptor) string {
	// TODO: handle special cases
	return getJSONForKRM(protoField) + "Ref"
}

func (w *messageWriter) isReferenceField(field protoreflect.FieldDescriptor) (string, bool) {
	if w.opts.referenceFields == nil {
		return "", false
	}
	referencedGoType, exist := w.opts.referenceFields[string(field.FullName())]
	return referencedGoType, exist
}
