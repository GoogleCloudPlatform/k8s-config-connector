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
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"k8s.io/apimachinery/pkg/util/sets"

	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

// Some special-case values that are not obvious how to map in KRM
var protoMessagesNotMappedToGoStruct = map[string]string{
	"google.protobuf.Timestamp":   "string",
	"google.protobuf.Duration":    "string",
	"google.protobuf.Int64Value":  "int64",
	"google.protobuf.StringValue": "string",
	"google.protobuf.Struct":      "map[string]string",
}

type TypeGenerator struct {
	generatorBase
	api             *protoapi.Proto
	goPackage       string
	visitedMessages []protoreflect.MessageDescriptor
}

func NewTypeGenerator(goPackage string, outputBaseDir string, api *protoapi.Proto) *TypeGenerator {
	g := &TypeGenerator{
		goPackage: goPackage,
		api:       api,
	}
	g.generatorBase.init(outputBaseDir)
	return g
}

func (g *TypeGenerator) VisitProto(resourceProtoFullName string) error {

	descriptor, err := g.api.Files().FindDescriptorByName(protoreflect.FullName(resourceProtoFullName))
	if err != nil {
		return fmt.Errorf("failed to find the proto message %s: %w", resourceProtoFullName, err)
	}
	messageDescriptor, ok := descriptor.(protoreflect.MessageDescriptor)
	if !ok {
		return fmt.Errorf("unexpected descriptor type: %T", descriptor)
	}

	if err := g.visitMessage(messageDescriptor); err != nil {
		return err
	}

	return nil
}

func (g *TypeGenerator) visitMessage(messageDescriptor protoreflect.MessageDescriptor) error {
	//klog.Infof("found message %q", messageDescriptor.FullName())

	g.visitedMessages = append(g.visitedMessages, messageDescriptor)

	msgs, err := findDependenciesForMessage(messageDescriptor)
	if err != nil {
		return err
	}
	g.visitedMessages = append(g.visitedMessages, msgs...)

	return nil
}

func writeCopyright(w io.Writer, year int) {
	s := `// Copyright {{.Year}} Google LLC
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

`
	s = strings.ReplaceAll(s, "{{.Year}}", strconv.Itoa(year))
	if _, err := w.Write([]byte(s)); err != nil {
		klog.Fatalf("writing copyright: %v", err)
	}
}

func (g *TypeGenerator) WriteVisitedMessages() error {
	for _, msg := range deduplicateAndSort(g.visitedMessages) {
		if msg.IsMapEntry() {
			continue
		}

		krmVersion := filepath.Base(g.goPackage)

		k := generatedFileKey{
			GoPackage: g.goPackage,
			FileName:  "types.generated.go",
		}
		out := g.getOutputFile(k)

		goTypeName := goNameForProtoMessage(msg)
		skipGenerated := true
		goType, err := g.findTypeDeclaration(goTypeName, out.OutputDir(), skipGenerated)
		if err != nil {
			return fmt.Errorf("looking up go type: %w", err)
		}
		if goType != nil {
			klog.Infof("found existing non-generated go type %q, won't generate", goTypeName)
			continue
		}

		goType, err = g.findTypeDeclarationWithProtoTag(string(msg.FullName()), out.OutputDir(), skipGenerated)
		if err != nil {
			return fmt.Errorf("looking up go type by proto tag: %w", err)
		}
		if goType != nil {
			klog.Infof("found existing non-generated go type with proto tag %q, won't generate", msg.FullName())
			continue
		}

		out.packageName = krmVersion

		WriteMessage(&out.body, msg)
	}
	return errors.Join(g.errors...)
}

func WriteMessage(out io.Writer, msg protoreflect.MessageDescriptor) {
	goType := goNameForProtoMessage(msg)

	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "// +kcc:proto=%s\n", msg.FullName())
	fmt.Fprintf(out, "type %s struct {\n", goType)
	for i := 0; i < msg.Fields().Len(); i++ {
		field := msg.Fields().Get(i)
		WriteField(out, field, msg, i)
	}
	fmt.Fprintf(out, "}\n")
}

func WriteField(out io.Writer, field protoreflect.FieldDescriptor, msg protoreflect.MessageDescriptor, fieldIndex int) {
	sourceLocations := msg.ParentFile().SourceLocations().ByDescriptor(field)

	jsonName := getJSONForKRM(field)
	goFieldName := goFieldName(field)
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
			fmt.Fprintf(out, "\n\t// TODO: map type %v %v for %v\n\n", keyKind, valueKind, field.Name())
			return
		}
	} else {
		switch field.Kind() {
		case protoreflect.MessageKind:
			goType = goNameForProtoMessage(field.Message())

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
		// Special case for proto "google.protobuf.Struct" type
		if goType == "*map[string]string" {
			goType = "map[string]string"
		}
	}

	// Blank line between fields for readability
	if fieldIndex != 0 {
		fmt.Fprintf(out, "\n")
	}

	if sourceLocations.LeadingComments != "" {
		comment := strings.TrimSpace(sourceLocations.LeadingComments)
		for _, line := range strings.Split(comment, "\n") {
			if strings.TrimSpace(line) == "" {
				fmt.Fprintf(out, "\t//\n")
			} else {
				fmt.Fprintf(out, "\t// %s\n", line)
			}
		}
	}

	fmt.Fprintf(out, "\t%s %s `json:\"%s,omitempty\"`\n",
		goFieldName,
		goType,
		jsonName,
	)
}

func deduplicateAndSort(messages []protoreflect.MessageDescriptor) []protoreflect.MessageDescriptor {
	m := make(map[string]protoreflect.MessageDescriptor)
	for _, msg := range messages {
		key := string(msg.FullName())
		m[key] = msg
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	messages = []protoreflect.MessageDescriptor{}
	for _, key := range keys {
		messages = append(messages, m[key])
	}
	return messages
}

func goNameForProtoMessage(msg protoreflect.MessageDescriptor) string {
	fullName := string(msg.FullName())

	// Some special-case values that are not obvious how to map in KRM
	if goType, ok := protoMessagesNotMappedToGoStruct[fullName]; ok {
		return goType
	}

	fullName = strings.TrimPrefix(fullName, string(msg.ParentFile().FullName()))
	fullName = strings.TrimPrefix(fullName, ".")
	fullName = strings.ReplaceAll(fullName, ".", "")
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
	tokens := strings.Split(string(protoField.Name()), "_")
	for i, token := range tokens {
		if i == 0 {
			// Do not capitalize first token
			continue
		}
		if isAcronym(token) {
			token = strings.ToUpper(token)
		} else {
			token = strings.Title(token)
		}
		tokens[i] = token
	}
	return strings.Join(tokens, "")
}

// goFieldName returns the KRM go name for the field,
// honoring KRM conventions
func goFieldName(protoField protoreflect.FieldDescriptor) string {
	tokens := strings.Split(string(protoField.Name()), "_")
	for i, token := range tokens {
		if isAcronym(token) {
			token = strings.ToUpper(token)
		} else {
			token = strings.Title(token)
		}
		tokens[i] = token
	}
	return strings.Join(tokens, "")
}

func isAcronym(s string) bool {
	switch s {
	case "id":
		return true
	case "html", "url":
		return true
	case "http", "https", "ssh":
		return true
	default:
		return false
	}
}

// findDependenciesForMessage recursively explores the dependent proto messages of the given message.
func findDependenciesForMessage(message protoreflect.MessageDescriptor) ([]protoreflect.MessageDescriptor, error) {
	msgs := make(map[string]protoreflect.MessageDescriptor)
	for i := 0; i < message.Fields().Len(); i++ {
		field := message.Fields().Get(i)
		FindDependenciesForField(field, msgs, nil) // TODO: explicity set ignored fields when generating Go types
	}

	RemoveNotMappedToGoStruct(msgs)

	res := []protoreflect.MessageDescriptor{}
	for _, msg := range msgs {
		res = append(res, msg)
	}
	return res, nil
}

// FindDependenciesForField recursively explores the dependent proto messages of the given field.
func FindDependenciesForField(field protoreflect.FieldDescriptor, deps map[string]protoreflect.MessageDescriptor, ignoredFields sets.String) {
	if ignoredFields.Has(string(field.FullName())) {
		return
	}

	if field.Message() != nil { // no need to find dependencies for proto messages that are not mapped to KRM Go struct
		if _, ok := protoMessagesNotMappedToGoStruct[string(field.Message().FullName())]; ok {
			return
		}
	}

	if field.IsMap() {
		mapEntry := field.Message()
		if keyField := mapEntry.Fields().ByName("key"); keyField != nil {
			FindDependenciesForField(keyField, deps, ignoredFields)
		}
		if valueField := mapEntry.Fields().ByName("value"); valueField != nil {
			FindDependenciesForField(valueField, deps, ignoredFields)
		}
	} else {
		switch field.Kind() {
		case protoreflect.MessageKind:
			msg := field.Message()
			fqn := string(msg.FullName())
			if _, ok := deps[fqn]; !ok {
				deps[fqn] = msg
				for i := 0; i < msg.Fields().Len(); i++ {
					field := msg.Fields().Get(i)
					FindDependenciesForField(field, deps, ignoredFields)
				}
			}
		case protoreflect.EnumKind:
			// deps[string(field.Enum().FullName())] = true  // Skip enum because enum is mapped to Go string in code generation
		}
	}
}

func RemoveNotMappedToGoStruct(msgs map[string]protoreflect.MessageDescriptor) {
	for msg := range protoMessagesNotMappedToGoStruct {
		delete(msgs, msg)
	}
}
