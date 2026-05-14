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
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"

	codegenannotations "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/annotations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

type TypeGenerator struct {
	generatorBase
	api                     *protoapi.Proto
	goPackage               string
	generatedFileAnnotation *codegenannotations.FileAnnotation
	includeSkippedOutput    bool

	*TypeDiscovery
}

type TypeDiscovery struct {
	ReachableMessages map[string]*MessageDetails
	fieldOverrides    map[string]*ProtoOverride
}

type ProtoOverride struct {
	OutputOnly *bool
}

type MessageDetails struct {
	Message protoreflect.MessageDescriptor

	ReachableFromInput  bool
	ReachableFromOutput bool
	// InputFields  []protoreflect.FieldDescriptor
	// OutputFields []protoreflect.FieldDescriptor
}

func (m *MessageDetails) AllFields() []protoreflect.FieldDescriptor {
	var out []protoreflect.FieldDescriptor
	fields := m.Message.Fields()
	for i := 0; i < fields.Len(); i++ {
		out = append(out, fields.Get(i))
	}
	return out
}

func NewTypeGenerator(goPackage string, outputBaseDir string, api *protoapi.Proto) *TypeGenerator {
	g := &TypeGenerator{
		goPackage: goPackage,
		api:       api,
		TypeDiscovery: &TypeDiscovery{
			ReachableMessages: make(map[string]*MessageDetails),
		},
	}
	g.generatorBase.init(outputBaseDir)
	return g
}

// WithProtoOverrides sets the field overrides from config file
func (g *TypeDiscovery) WithProtoOverrides(fieldOverrides map[string]*ProtoOverride) {
	g.fieldOverrides = fieldOverrides
}

// WithGeneratedFileAnnotation sets the generated file annotation
func (g *TypeGenerator) WithGeneratedFileAnnotation(generatedFileAnnotation *codegenannotations.FileAnnotation) *TypeGenerator {
	g.generatedFileAnnotation = generatedFileAnnotation
	return g
}

// WithIncludeSkippedOutput sets whether to output skipped types as commented-out code
func (g *TypeGenerator) WithIncludeSkippedOutput(includeSkippedOutput bool) *TypeGenerator {
	g.includeSkippedOutput = includeSkippedOutput
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

func (g *TypeDiscovery) visitMessage(startMessage protoreflect.MessageDescriptor) error {
	//klog.Infof("found message %q", messageDescriptor.FullName())
	g.ReachableMessages[string(startMessage.FullName())] = &MessageDetails{
		Message:             startMessage,
		ReachableFromInput:  true,
		ReachableFromOutput: true,
	}

	if err := g.findReachableTypesFromMessage(startMessage, g.ReachableMessages, MixedInputOutput); err != nil {
		return err
	}

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

func (g *TypeGenerator) WriteReachableMessages() error {
	// for k, msgDetails := range g.ReachableMessages {
	// 	klog.Infof("visit %q => %q", k, msgDetails.Message.FullName())

	// }

	for _, msgDetails := range deduplicateAndSortMessages(g.ReachableMessages) {
		// klog.Infof("write msg %q", msgDetails.Message.FullName())
		if !msgDetails.ReachableFromInput {
			// klog.Infof("skipping msg %q because not reachable from input", msgDetails.Message.FullName())
			continue
		}

		if _, ok := protoMessagesNotMappedToGoStruct[string(msgDetails.Message.FullName())]; ok {
			// klog.Infof("skipping msg %q because not mapped to go struct", msgDetails.Message.FullName())
			continue
		}

		msg := msgDetails.Message
		if msg.IsMapEntry() {
			// klog.Infof("skipping msg %q because it is a map entry", msgDetails.Message.FullName())
			continue
		}

		k := generatedFileKey{
			GoPackage: g.goPackage,
			FileName:  "types.generated.go",
		}
		out := g.getOutputFile(k)

		out.goPackage = lastGoComponent(g.goPackage)

		out.fileAnnotation = g.generatedFileAnnotation

		goTypeName := GoNameForProtoMessage(msg)
		skipGenerated := true
		goType, err := g.findTypeDeclaration(goTypeName, out.OutputDir(), skipGenerated)
		if err != nil {
			return fmt.Errorf("looking up go type: %w", err)
		}
		if goType != nil {
			klog.Infof("found existing non-generated go type %q, won't generate", goTypeName)
			if g.includeSkippedOutput {
				g.WriteSpecMessageAsComment(&out.body, msg, fmt.Sprintf("found existing non-generated go type %q, skipping", goTypeName))
			}
			continue
		}

		goType, err = g.findTypeDeclarationWithProtoTag(string(msg.FullName()), out.OutputDir(), skipGenerated, []string{KCCProtoMessageAnnotationMisc, KCCProtoMessageAnnotationSpec})
		if err != nil {
			return fmt.Errorf("looking up go type by proto tag: %w", err)
		}
		if goType != nil {
			klog.Infof("found existing non-generated go type with proto tag %q, won't generate", msg.FullName())
			if g.includeSkippedOutput {
				g.WriteSpecMessageAsComment(&out.body, msg, fmt.Sprintf("found existing non-generated go type with proto tag %q, skipping", msg.FullName()))
			}
			continue
		}

		g.WriteSpecMessage(&out.body, msg)
	}

	for _, msgDetails := range deduplicateAndSortMessages(g.ReachableMessages) {
		if !msgDetails.ReachableFromOutput {
			continue
		}

		if _, ok := protoMessagesNotMappedToGoStruct[string(msgDetails.Message.FullName())]; ok {
			continue
		}

		msg := msgDetails.Message
		if msg.IsMapEntry() {
			continue
		}

		k := generatedFileKey{
			GoPackage: g.goPackage,
			FileName:  "types.generated.go",
		}
		out := g.getOutputFile(k)
		out.goPackage = lastGoComponent(g.goPackage)

		out.fileAnnotation = g.generatedFileAnnotation

		goTypeName := g.goNameForOutputProtoMessage(msg)
		skipGenerated := true
		goType, err := g.findTypeDeclaration(goTypeName, out.OutputDir(), skipGenerated)
		if err != nil {
			return fmt.Errorf("looking up go type: %w", err)
		}
		if goType != nil {
			klog.V(1).Infof("found existing non-generated go type %q, won't generate", goTypeName)
			if g.includeSkippedOutput {
				g.WriteObservedStateMessageAsComment(&out.body, msgDetails, fmt.Sprintf("found existing non-generated go type %q, skipping", goTypeName))
			}
			continue
		}

		goType, err = g.findTypeDeclarationWithProtoTag(string(msg.FullName()), out.OutputDir(), skipGenerated, []string{KCCProtoMessageAnnotationObservedState})
		if err != nil {
			return fmt.Errorf("looking up go type by proto tag: %w", err)
		}
		if goType != nil {
			klog.V(1).Infof("found existing non-generated go type with proto tag %q, won't generate", msg.FullName())
			if g.includeSkippedOutput {
				g.WriteObservedStateMessageAsComment(&out.body, msgDetails, fmt.Sprintf("found existing non-generated go type with proto tag %q, skipping", msg.FullName()))
			}
			continue
		}

		g.WriteObservedStateMessage(&out.body, msgDetails)
	}
	return errors.Join(g.errors...)
}

func (g *TypeGenerator) WriteSpecMessageAsComment(out io.Writer, msg protoreflect.MessageDescriptor, reason string) {
	var b bytes.Buffer
	g.WriteSpecMessage(&b, msg)
	fmt.Fprintf(out, "\n/* %s\n", reason)
	fmt.Fprintf(out, "%s", strings.ReplaceAll(b.String(), "*/", "* /"))
	fmt.Fprintf(out, "*/\n")
}

func (g *TypeGenerator) WriteObservedStateMessageAsComment(out io.Writer, msgDetails *MessageDetails, reason string) {
	var b bytes.Buffer
	g.WriteObservedStateMessage(&b, msgDetails)
	fmt.Fprintf(out, "\n/* %s\n", reason)
	fmt.Fprintf(out, "%s", strings.ReplaceAll(b.String(), "*/", "* /"))
	fmt.Fprintf(out, "*/\n")
}

func (g *TypeGenerator) WriteSpecMessage(out io.Writer, msg protoreflect.MessageDescriptor) {
	goType := GoNameForProtoMessage(msg)
	klog.Infof("writing input message %q", goType)

	messageDetails := g.MessageDetails(msg)
	if messageDetails == nil {
		klog.Fatalf("message details not found for message %q", msg.FullName())
	}

	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "// %s=%s\n", KCCProtoMessageAnnotationMisc, msg.FullName())
	fmt.Fprintf(out, "type %s struct {\n", goType)
	for i := 0; i < msg.Fields().Len(); i++ {
		field := msg.Fields().Get(i)

		if g.IsOutputOnlyField(field) {
			continue
		}

		g.WriteField(out, field, msg, i, false)
	}
	fmt.Fprintf(out, "}\n")
}

func (g *TypeGenerator) WriteObservedStateMessage(out io.Writer, messageDetails *MessageDetails) {
	msg := messageDetails.Message
	goType := g.goNameForOutputProtoMessage(msg)

	var fields []protoreflect.FieldDescriptor
	for _, field := range messageDetails.AllFields() {
		if messageDetails.ReachableFromInput {
			if !g.IsOutputOnlyField(field) {
				continue
			}
		}
		fields = append(fields, field)
	}

	if len(fields) == 0 {
		// klog.Infof("skipping output of output message which has no output fields %q", goType)
		// return
		klog.Infof("writing output message %q that has no output fields", goType)
	}

	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "// %s=%s\n", KCCProtoMessageAnnotationObservedState, msg.FullName())
	fmt.Fprintf(out, "type %s struct {\n", goType)

	for i, field := range fields {
		g.WriteField(out, field, msg, i, true)
	}
	fmt.Fprintf(out, "}\n")
}

func (g *TypeDiscovery) MessageDetails(message protoreflect.MessageDescriptor) *MessageDetails {
	key := string(message.FullName())
	if val, ok := g.ReachableMessages[key]; ok {
		return val
	}
	klog.Warningf("message details not found for message %q", key)
	return nil
}

func (g *TypeDiscovery) IsOutputOnlyField(field protoreflect.FieldDescriptor) bool {
	fieldPath := string(field.ContainingMessage().FullName()) + ":" + string(field.Name())
	klog.V(4).Infof("IsOutputOnlyField: %s", fieldPath)
	fieldOverride := g.fieldOverrides[fieldPath]
	if fieldOverride == nil {
		fieldOverride = g.fieldOverrides[string(field.ContainingMessage().FullName())+":*"]
	}
	if fieldOverride != nil {
		if fieldOverride.OutputOnly != nil {
			return *fieldOverride.OutputOnly
		}
	}

	fieldBehaviors := proto.GetExtension(field.Options(), annotations.E_FieldBehavior).([]annotations.FieldBehavior)
	for _, fieldBehavior := range fieldBehaviors {
		if fieldBehavior == annotations.FieldBehavior_OUTPUT_ONLY {
			return true
		}
	}
	return false
}

// // MessageType determines whether the message is output only, input only, or mixed
// func (g *TypeGenerator) MessageType(message protoreflect.MessageDescriptor) InputOutputType {
// 	key := string(message.FullName()) + ":*"
// 	messageType, ok := g.messageTypes[key]
// 	if ok {
// 		return messageType
// 	}

// 	hasInputFields := false
// 	hasOutputFields := false

// 	fields := message.Fields()

// 	for i := 0; i < fields.Len(); i++ {
// 		field := fields.Get(i)
// 		if g.IsOutputOnlyField(field) {
// 			hasOutputFields = true
// 		} else {
// 			hasInputFields = true
// 		}

// 		switch field.Kind() {
// 		case protoreflect.MessageKind:
// 			fieldMessage := field.Message()
// 			switch g.MessageType(fieldMessage) {
// 			case InputOnly:
// 				hasInputFields = true
// 			case OutputOnly:
// 				hasOutputFields = true
// 			case MixedInputOutput:
// 				hasInputFields = true
// 				hasOutputFields = true
// 			}
// 		}
// 	}

// 	if hasInputFields && !hasOutputFields {
// 		messageType = InputOnly
// 	} else if !hasInputFields && hasOutputFields {
// 		messageType = OutputOnly
// 	} else {
// 		messageType = MixedInputOutput
// 	}
// 	g.messageTypes[key] = messageType
// 	return messageType
// }

func (g *TypeGenerator) GoTypeForField(field protoreflect.FieldDescriptor, isTransitiveOutput bool) (string, error) {
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
			goType = g.goNameForOutputProtoMessage(field.Message())
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
	if goType == "*apiextensionsv1.JSON" {
		goType = "apiextensionsv1.JSON"
	}

	return goType, nil
}

func (g *TypeGenerator) WriteField(out io.Writer, field protoreflect.FieldDescriptor, msg protoreflect.MessageDescriptor, fieldIndex int, isTransitiveOutput bool) {
	sourceLocations := msg.ParentFile().SourceLocations().ByDescriptor(field)

	jsonName := GetJSONForKRM(field)
	GoFieldName := goFieldName(field)

	goType, err := g.GoTypeForField(field, isTransitiveOutput)
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

func deduplicateAndSortMessages(messages map[string]*MessageDetails) []*MessageDetails {
	var keys []string
	for key := range messages {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var out []*MessageDetails
	for _, key := range keys {
		out = append(out, messages[key])
	}
	return out
}

// AsSnakeCase returns the given string converted to lowercase snake_case. If the input is already snake_case, no
// change is made. Any transitions in the input from lowercase to uppercase are interpreted as camelCase-style word
// transitions, and are replaced with an underscore.
func AsSnakeCase(s string) string {
	res := regexp.MustCompile("(.)([A-Z][a-z]+)").ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(res, "${1}_${2}"))
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
		partInSnakeCase := AsSnakeCase(part)
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

func (g *TypeGenerator) goNameForOutputProtoMessage(msg protoreflect.MessageDescriptor) string {
	base := GoNameForProtoMessage(msg)

	// Some well-known types never get a suffix
	switch base {
	case "string", "int64":
		return base

	}

	useBaseNameWhereNoConflict := false
	if useBaseNameWhereNoConflict {
		// Use the base name for simple types that only appear in output.
		// This is just for (human) readability.
		messageDetails := g.MessageDetails(msg)
		if messageDetails == nil {
			klog.Fatalf("no message details for %q / %q", msg.FullName(), base)
		}
		if messageDetails.ReachableFromOutput && !messageDetails.ReachableFromInput {
			return base
		}
	}

	return base + "ObservedState"
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

// findReachableTypesFromMessage recursively explores the dependent proto messages of the given message.
func (g *TypeDiscovery) findReachableTypesFromMessage(message protoreflect.MessageDescriptor, dest map[string]*MessageDetails, reachedOnPath InputOutputType) error {
	for i := 0; i < message.Fields().Len(); i++ {
		field := message.Fields().Get(i)
		g.findReachableTypesFromField(field, dest, reachedOnPath)
	}

	return nil
}

type InputOutputType string

const (
	InputOnly        InputOutputType = "InputOnly"
	OutputOnly       InputOutputType = "OutputOnly"
	MixedInputOutput InputOutputType = "Mixed"
)

// findReachableTypesFromField recursively explores the dependent proto messages of the given field.
func (g *TypeDiscovery) findReachableTypesFromField(field protoreflect.FieldDescriptor, dest map[string]*MessageDetails, reachedOnPath InputOutputType) {
	// klog.Infof("findReachableTypesFromField(%v, %v)", field.FullName(), reachedOnPath)
	if field.Message() != nil {
		// no need to find dependencies for proto messages that are not mapped to KRM Go struct
		if _, ok := protoMessagesNotMappedToGoStruct[string(field.Message().FullName())]; ok {
			return
		}
	}

	if g.IsOutputOnlyField(field) {
		if reachedOnPath == InputOnly {
			return
		}
		if reachedOnPath == MixedInputOutput {
			reachedOnPath = OutputOnly
		}
	}

	if field.IsMap() {
		mapEntry := field.Message()
		if keyField := mapEntry.Fields().ByName("key"); keyField != nil {
			g.findReachableTypesFromField(keyField, dest, reachedOnPath)
		}
		if valueField := mapEntry.Fields().ByName("value"); valueField != nil {
			g.findReachableTypesFromField(valueField, dest, reachedOnPath)
		}
	} else {
		switch field.Kind() {
		case protoreflect.MessageKind:
			msg := field.Message()
			fqn := string(msg.FullName())
			visit := false
			messageDetails := dest[fqn]
			if messageDetails == nil {
				// klog.Infof("adding new msg to dest: %q", fqn)
				messageDetails = &MessageDetails{
					Message: msg,
				}
				dest[fqn] = messageDetails
				visit = true
			}

			if reachedOnPath != OutputOnly && !messageDetails.ReachableFromInput {
				// klog.Infof("marked %q as reachable from input", fqn)
				messageDetails.ReachableFromInput = true
				visit = true
			}

			if reachedOnPath != InputOnly && !messageDetails.ReachableFromOutput {
				// klog.Infof("marked %q as reachable from output", fqn)
				messageDetails.ReachableFromOutput = true
				visit = true
			}

			if visit {
				for i := 0; i < msg.Fields().Len(); i++ {
					field := msg.Fields().Get(i)
					g.findReachableTypesFromField(field, dest, reachedOnPath)
				}
			}
		case protoreflect.EnumKind:
			// deps[string(field.Enum().FullName())] = true  // Skip enum because enum is mapped to Go string in code generation
		}
	}
}

// func RemoveNotMappedToGoStruct(msgs map[string]protoreflect.MessageDescriptor) {
// 	for msg := range protoMessagesNotMappedToGoStruct {
// 		delete(msgs, msg)
// 	}
// }

// // findOutputsForMessage recursively explores a given message and its dependencies and assembles a
// // list of messages that contain output-only fields and the output-only fields within them.
// func (g *TypeGenerator) findOutputsForMessage(message protoreflect.MessageDescriptor) ([]*OutputMessageDetails, error) {
// 	outputDeps := make(map[string]*OutputMessageDetails)

// 	seen := make(map[string]bool)
// 	for i := 0; i < message.Fields().Len(); i++ {
// 		field := message.Fields().Get(i)
// 		// TODO: explicitly set ignored fields when generating Go types
// 		if hasOutput := g.FindOutputsForField(field, seen, outputDeps, nil); hasOutput {
// 			fqn := string(message.FullName())
// 			if _, ok := outputDeps[fqn]; !ok {
// 				outputDeps[fqn] = &OutputMessageDetails{
// 					Message: message,
// 				}
// 			}
// 			outputDeps[fqn].OutputFields = append(outputDeps[fqn].OutputFields, field)
// 		}
// 		// TODO: explicitly set ignored fields when generating Go types
// 		g.FindOutputsForField(field, seen, outputDeps, nil)
// 	}

// 	res := []*OutputMessageDetails{}
// 	for _, msgDetails := range outputDeps {
// 		res = append(res, msgDetails)
// 	}
// 	return res, nil
// }

// // FindDependenciesForField recursively explores the dependent proto messages of the given field.
// // It returns true if the current field or any of its nested field is marked as an output field, false otherwise.
// func (g *TypeGenerator) FindOutputsForField(field protoreflect.FieldDescriptor, seen map[string]bool, outputDeps map[string]*OutputMessageDetails, ignoredFields sets.String) bool {
// 	if ignoredFields.Has(string(field.FullName())) {
// 		return false
// 	}

// 	isOutput := false
// 	if g.IsOutputOnlyField(field) {
// 		isOutput = true
// 	}

// 	if field.Message() != nil {
// 		// There is no need to recurse for proto messages that are not mapped to KRM Go struct.
// 		if _, ok := protoMessagesNotMappedToGoStruct[string(field.Message().FullName())]; ok {
// 			return isOutput
// 		}
// 	}

// 	if field.IsMap() {
// 		// Map outputs are not supported.
// 	} else {
// 		switch field.Kind() {
// 		case protoreflect.MessageKind:
// 			msg := field.Message()
// 			fqn := string(msg.FullName())
// 			if _, ok := seen[fqn]; !ok {
// 				seen[fqn] = true
// 				for i := 0; i < msg.Fields().Len(); i++ {
// 					field := msg.Fields().Get(i)
// 					if g.FindOutputsForField(field, seen, outputDeps, ignoredFields) {
// 						if _, ok := outputDeps[fqn]; !ok {
// 							outputDeps[fqn] = &OutputMessageDetails{
// 								Message: msg,
// 							}
// 						}
// 						outputDeps[fqn].OutputFields = append(outputDeps[fqn].OutputFields, field)
// 						isOutput = true
// 					}
// 				}
// 			}
// 		case protoreflect.EnumKind:
// 			// There is no need to recurse for enum messages since they are mapped to Go string.
// 		}
// 	}

// 	return isOutput
// }
