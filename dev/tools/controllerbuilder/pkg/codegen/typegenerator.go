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

	codegenannotations "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/annotations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

type TypeGenerator struct {
	generatorBase
	api                     *protoapi.Proto
	goPackage               string
	visitedMessages         []protoreflect.MessageDescriptor
	outputMessages          []*OutputMessageDetails
	generatedFileAnnotation *codegenannotations.FileAnnotation
}

type OutputMessageDetails struct {
	Message      protoreflect.MessageDescriptor
	OutputFields []protoreflect.FieldDescriptor
}

func NewTypeGenerator(goPackage string, outputBaseDir string, api *protoapi.Proto) *TypeGenerator {
	g := &TypeGenerator{
		goPackage: goPackage,
		api:       api,
	}
	g.generatorBase.init(outputBaseDir)
	return g
}

// WithGeneratedFileAnnotation sets the generated file annotation
func (g *TypeGenerator) WithGeneratedFileAnnotation(generatedFileAnnotation *codegenannotations.FileAnnotation) *TypeGenerator {
	g.generatedFileAnnotation = generatedFileAnnotation
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

func (g *TypeGenerator) visitMessage(message protoreflect.MessageDescriptor) error {
	//klog.Infof("found message %q", messageDescriptor.FullName())

	g.visitedMessages = append(g.visitedMessages, message)

	msgs, err := FindDependenciesForMessage(message, nil) // TODO: explicitly set ignored fields when generating Go types
	if err != nil {
		return err
	}
	g.visitedMessages = append(g.visitedMessages, msgs...)

	outputMessages, err := findOutputsForMessage(message)
	if err != nil {
		return err
	}
	g.outputMessages = append(g.outputMessages, outputMessages...)

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

		goTypeName := GoNameForProtoMessage(msg)
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

func (g *TypeGenerator) WriteOutputMessages() error {
	for _, msgDetails := range deduplicateAndSortOutputMessages(g.outputMessages) {
		msg := msgDetails.Message
		if msg.IsMapEntry() {
			continue
		}

		krmVersion := filepath.Base(g.goPackage)

		k := generatedFileKey{
			GoPackage: g.goPackage,
			FileName:  "types.generated.go",
		}
		out := g.getOutputFile(k)

		out.fileAnnotation = g.generatedFileAnnotation

		goTypeName := goNameForOutputProtoMessage(msg)
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

		WriteOutputMessage(&out.body, msgDetails)
	}
	return errors.Join(g.errors...)
}

func WriteMessage(out io.Writer, msg protoreflect.MessageDescriptor) {
	goType := GoNameForProtoMessage(msg)

	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "// %s=%s\n", KCCProtoMessageAnnotation, msg.FullName())
	fmt.Fprintf(out, "type %s struct {\n", goType)
	for i := 0; i < msg.Fields().Len(); i++ {
		field := msg.Fields().Get(i)
		if !IsFieldBehavior(field, annotations.FieldBehavior_OUTPUT_ONLY) {
			// Only write non-output fields.
			WriteField(out, field, msg, i, false)
		}
	}
	fmt.Fprintf(out, "}\n")
}

func WriteOutputMessage(out io.Writer, msgDetails *OutputMessageDetails) {
	msg := msgDetails.Message
	goType := goNameForOutputProtoMessage(msg)

	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "// %s=%s\n", KCCProtoMessageAnnotation, msg.FullName())
	fmt.Fprintf(out, "type %s struct {\n", goType)
	for i, field := range msgDetails.OutputFields {
		if !IsFieldBehavior(field, annotations.FieldBehavior_OUTPUT_ONLY) {
			// If field is not explicitly listed as an output, but it appears in OutputMessageDetails,
			// then it must be a parent message that contains a child message with an output.
			WriteField(out, field, msg, i, true)
		} else {
			WriteField(out, field, msg, i, false)
		}
	}
	fmt.Fprintf(out, "}\n")
}

func GoTypeForField(field protoreflect.FieldDescriptor, isTransitiveOutput bool) (string, error) {
	if field.IsMap() {
		entryMsg := field.Message()
		keyKind := entryMsg.Fields().ByName("key").Kind()
		valueKind := entryMsg.Fields().ByName("value").Kind()
		if keyKind == protoreflect.StringKind && valueKind == protoreflect.StringKind {
			return "map[string]string", nil
		} else if keyKind == protoreflect.StringKind && valueKind == protoreflect.Int64Kind {
			return "map[string]int64", nil
		} else {
			return "", fmt.Errorf("unsupported map type with key %v and value %v", keyKind, valueKind)
		}
	}

	var goType string
	switch field.Kind() {
	case protoreflect.MessageKind:
		if isTransitiveOutput {
			goType = goNameForOutputProtoMessage(field.Message())
		} else {
			goType = GoNameForProtoMessage(field.Message())
		}
	case protoreflect.EnumKind:
		goType = "string"
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

	return goType, nil
}

func WriteField(out io.Writer, field protoreflect.FieldDescriptor, msg protoreflect.MessageDescriptor, fieldIndex int, isTransitiveOutput bool) {
	sourceLocations := msg.ParentFile().SourceLocations().ByDescriptor(field)

	jsonName := GetJSONForKRM(field)
	GoFieldName := goFieldName(field)

	goType, err := GoTypeForField(field, isTransitiveOutput)
	if err != nil {
		fmt.Fprintf(out, "\n\t// TODO: %v\n\n", err)
		return
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

	fmt.Fprintf(out, "\t// %s=%s\n", KCCProtoFieldAnnotation, field.FullName())
	fmt.Fprintf(out, "\t%s %s `json:\"%s,omitempty\"`\n",
		GoFieldName,
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

func deduplicateAndSortOutputMessages(messages []*OutputMessageDetails) []*OutputMessageDetails {
	m := make(map[string]*OutputMessageDetails)
	for _, msg := range messages {
		key := string(msg.Message.FullName())
		m[key] = msg
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	messages = []*OutputMessageDetails{}
	for _, key := range keys {
		messages = append(messages, m[key])
	}
	return messages
}

func GoNameForProtoMessage(msg protoreflect.MessageDescriptor) string {
	fullName := string(msg.FullName())

	// Some special-case values that are not obvious how to map in KRM
	if goType, ok := protoMessagesNotMappedToGoStruct[fullName]; ok {
		return goType
	}

	fullName = strings.TrimPrefix(fullName, string(msg.ParentFile().FullName()))
	fullName = strings.TrimPrefix(fullName, ".")
	// Ensure acronyms in type names are also handled.
	parts := strings.Split(fullName, ".")
	for i, part := range parts {
		partInSnakeCase := text.AsSnakeCase(part)
		tokens := strings.Split(partInSnakeCase, "_")
		for j, token := range tokens {
			if IsAcronym(token) {
				token = strings.ToUpper(token)
			} else {
				token = strings.Title(token)
			}
			tokens[j] = token
		}
		parts[i] = strings.Join(tokens, "")
	}
	return strings.Join(parts, "_")
}

func goNameForOutputProtoMessage(msg protoreflect.MessageDescriptor) string {
	return GoNameForProtoMessage(msg) + "ObservedState"
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

// GetJSONForKRM returns the KRM JSON name for the field,
// honoring KRM conventions
func GetJSONForKRM(protoField protoreflect.FieldDescriptor) string {
	tokens := strings.Split(string(protoField.Name()), "_")
	for i, token := range tokens {
		if i == 0 {
			// Do not capitalize first token
			continue
		}
		if IsAcronym(token) {
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
		if IsAcronym(token) {
			token = strings.ToUpper(token)
		} else {
			token = strings.Title(token)
		}
		tokens[i] = token
	}
	return strings.Join(tokens, "")
}

// FindDependenciesForMessage recursively explores the dependent proto messages of the given message.
func FindDependenciesForMessage(message protoreflect.MessageDescriptor, ignoredFields sets.String) ([]protoreflect.MessageDescriptor, error) {
	msgs := make(map[string]protoreflect.MessageDescriptor)
	for i := 0; i < message.Fields().Len(); i++ {
		field := message.Fields().Get(i)
		FindDependenciesForField(field, msgs, ignoredFields)
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

// findOutputsForMessage recursively explores a given message and its dependencies and assembles a
// list of messages that contain output-only fields and the output-only fields within them.
func findOutputsForMessage(message protoreflect.MessageDescriptor) ([]*OutputMessageDetails, error) {
	outputDeps := make(map[string]*OutputMessageDetails)

	seen := make(map[string]bool)
	for i := 0; i < message.Fields().Len(); i++ {
		field := message.Fields().Get(i)
		// TODO: explicitly set ignored fields when generating Go types
		if hasOutput := FindOutputsForField(field, seen, outputDeps, nil); hasOutput {
			fqn := string(message.FullName())
			if _, ok := outputDeps[fqn]; !ok {
				outputDeps[fqn] = &OutputMessageDetails{
					Message: message,
				}
			}
			outputDeps[fqn].OutputFields = append(outputDeps[fqn].OutputFields, field)
		}
		// TODO: explicitly set ignored fields when generating Go types
		FindOutputsForField(field, seen, outputDeps, nil)
	}

	res := []*OutputMessageDetails{}
	for _, msgDetails := range outputDeps {
		res = append(res, msgDetails)
	}
	return res, nil
}

// FindDependenciesForField recursively explores the dependent proto messages of the given field.
// It returns true if the current field or any of its nested field is marked as an output field, false otherwise.
func FindOutputsForField(field protoreflect.FieldDescriptor, seen map[string]bool, outputDeps map[string]*OutputMessageDetails, ignoredFields sets.String) bool {
	if ignoredFields.Has(string(field.FullName())) {
		return false
	}

	isOutput := false
	if IsFieldBehavior(field, annotations.FieldBehavior_OUTPUT_ONLY) {
		isOutput = true
	}

	if field.Message() != nil {
		// There is no need to recurse for proto messages that are not mapped to KRM Go struct.
		if _, ok := protoMessagesNotMappedToGoStruct[string(field.Message().FullName())]; ok {
			return isOutput
		}
	}

	if field.IsMap() {
		// Map outputs are not supported.
	} else {
		switch field.Kind() {
		case protoreflect.MessageKind:
			msg := field.Message()
			fqn := string(msg.FullName())
			if _, ok := seen[fqn]; !ok {
				seen[fqn] = true
				for i := 0; i < msg.Fields().Len(); i++ {
					field := msg.Fields().Get(i)
					if FindOutputsForField(field, seen, outputDeps, ignoredFields) {
						if _, ok := outputDeps[fqn]; !ok {
							outputDeps[fqn] = &OutputMessageDetails{
								Message: msg,
							}
						}
						outputDeps[fqn].OutputFields = append(outputDeps[fqn].OutputFields, field)
						isOutput = true
					}
				}
			}
		case protoreflect.EnumKind:
			// There is no need to recurse for enum messages since they are mapped to Go string.
		}
	}

	return isOutput
}

func IsFieldBehavior(field protoreflect.FieldDescriptor, fieldBehavior annotations.FieldBehavior) bool {
	d := field.Options()
	fieldBehaviors := proto.GetExtension(d, annotations.E_FieldBehavior).([]annotations.FieldBehavior)
	for _, f := range fieldBehaviors {
		if f == fieldBehavior {
			return true
		}
	}
	return false
}
