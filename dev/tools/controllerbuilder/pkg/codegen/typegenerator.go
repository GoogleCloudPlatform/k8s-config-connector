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
	"os"
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

	// memoizeOutputFields speeds up computations of whether types include output fields
	memoizeHasOutputField map[string]bool
}

type ProtoOverride struct {
	OutputOnly *bool

	// Ignore is set if we should not generate the field
	Ignore bool
}

type MessageDetails struct {
	Message protoreflect.MessageDescriptor

	ReachableFromInput  bool
	ReachableFromOutput bool

	// IsRoot is set if this is a top-level (entrypoint) message;
	// these are considered the roots for reachability analysis.
	IsRoot bool
}

type FieldDetails struct {
	Field   protoreflect.FieldDescriptor
	TypeKey string
}

func NewTypeGenerator(goPackage string, outputBaseDir string, api *protoapi.Proto) *TypeGenerator {
	g := &TypeGenerator{
		goPackage: goPackage,
		api:       api,
		TypeDiscovery: &TypeDiscovery{
			ReachableMessages:     make(map[string]*MessageDetails),
			memoizeHasOutputField: make(map[string]bool),
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

	if err := g.visitTopLevelMessage(messageDescriptor); err != nil {
		return err
	}

	return nil
}

func (g *TypeDiscovery) hasOutputField(message protoreflect.MessageDescriptor) bool {
	key := string(message.FullName())
	if val, found := g.memoizeHasOutputField[key]; found {
		return val
	}
	// Avoid infinite recursion
	g.memoizeHasOutputField[key] = false

	foundOutputField := false
	fields := message.Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		if g.IsOutputOnlyField(field) {
			foundOutputField = true
			break
		}
		msg := field.Message()
		if msg != nil && g.hasOutputField(msg) {
			foundOutputField = true
			break
		}
	}
	g.memoizeHasOutputField[key] = foundOutputField
	return foundOutputField
}

func (g *TypeDiscovery) visitTopLevelMessage(startMessage protoreflect.MessageDescriptor) error {
	klog.Infof("visitTopLevelMessage %q", startMessage.FullName())
	messageDetails := &MessageDetails{
		Message:             startMessage,
		ReachableFromInput:  true,
		ReachableFromOutput: true,
		IsRoot:              true,
	}
	g.ReachableMessages[string(startMessage.FullName())] = messageDetails

	{
		// For spec, a field is reachable if there is a path to it which does not include an output-only field.
		pred := func(fieldPath []protoreflect.FieldDescriptor) bool {
			for _, field := range fieldPath {
				if g.IsOutputOnlyField(field) {
					return false
				}
			}

			return true
		}
		reachableFromInput := make(map[string]protoreflect.MessageDescriptor)
		if err := findReachableTypesFromMessage(startMessage, pred, reachableFromInput); err != nil {
			return err
		}
		for key, message := range reachableFromInput {
			messageDetails := g.ReachableMessages[key]
			if messageDetails == nil {
				messageDetails = &MessageDetails{
					Message: message,
				}
				g.ReachableMessages[key] = messageDetails
			}
			messageDetails.ReachableFromInput = true
		}
	}

	{
		// For status, a field is reachable if there is a path to it that includes an output-only field,
		// or where the referenced type itself has an output only field.
		pred := func(fieldPath []protoreflect.FieldDescriptor) bool {
			for _, field := range fieldPath {
				if g.IsOutputOnlyField(field) {
					return true
				}
			}
			lastField := fieldPath[len(fieldPath)-1]
			messageType := lastField.Message()
			if messageType != nil {
				if g.hasOutputField(messageType) {
					log(4, "Output field reachable from %v", lastField.FullName())
					return true
				}
			}
			return false
		}
		reachableFromOutput := make(map[string]protoreflect.MessageDescriptor)
		if err := findReachableTypesFromMessage(startMessage, pred, reachableFromOutput); err != nil {
			return err
		}
		for key, message := range reachableFromOutput {
			messageDetails := g.ReachableMessages[key]
			if messageDetails == nil {
				messageDetails = &MessageDetails{
					Message: message,
				}
				g.ReachableMessages[key] = messageDetails
			}
			messageDetails.ReachableFromOutput = true
		}
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
	for _, msgDetails := range deduplicateAndSortMessages(g.ReachableMessages) {
		if !msgDetails.ReachableFromInput {
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

	// for _, msgDetails := range deduplicateAndSortMessages(g.OutputMessages()) {
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

	messageDetails := g.MessageDetails(msg)
	if messageDetails == nil {
		klog.Fatalf("message details not found for message %q", msg.FullName())
	}

	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "// %s=%s\n", KCCProtoMessageAnnotationMisc, msg.FullName())
	fmt.Fprintf(out, "type %s struct {\n", goType)
	for i := 0; i < msg.Fields().Len(); i++ {
		field := msg.Fields().Get(i)

		if g.IgnoreField(field) {
			continue
		}

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

	var writeFields []protoreflect.FieldDescriptor

	msgFields := msg.Fields()
	for i := 0; i < msgFields.Len(); i++ {
		field := msgFields.Get(i)
		if g.IgnoreField(field) {
			continue
		}
		if !g.IsOutputOnlyField(field) {
			if messageDetails.ReachableFromInput {
				fieldMessage := field.Message()
				if field.IsMap() {
					fieldMessage = field.MapValue().Message()
				}
				if fieldMessage != nil {
					if !g.hasOutputField(fieldMessage) {
						log(2, "Skipping ObservedState field %s/%s (field message does not have output fields)", msg.Name(), field.Name())
						continue
					}
				} else {
					log(2, "Skipping ObservedState field %s/%s (field message not found and not marked output)", msg.Name(), field.Name())
					continue
				}
			}
		}
		writeFields = append(writeFields, field)
	}

	if len(writeFields) == 0 {
		// klog.Infof("skipping output of output message which has no output fields %q", goType)
		// return
		klog.Infof("writing output message %q that has no output fields", goType)
	}

	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "// %s=%s\n", KCCProtoMessageAnnotationObservedState, msg.FullName())
	fmt.Fprintf(out, "type %s struct {\n", goType)

	for i, field := range writeFields {
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

func log(level int, format string, args ...any) {
	trace := os.Getenv("TRACE")
	if trace != "" {
		s := fmt.Sprintf(format, args...)
		if strings.Contains(s, trace) {
			level = 0
		}
	}
	klog.V(klog.Level(level)).Infof(format, args...)
}

func (g *TypeDiscovery) IsOutputOnlyField(field protoreflect.FieldDescriptor) bool {
	fieldPath := string(field.ContainingMessage().FullName()) + ":" + string(field.Name())

	fieldOverride := g.fieldOverrides[fieldPath]
	if fieldOverride == nil {
		fieldOverride = g.fieldOverrides[string(field.ContainingMessage().FullName())+":*"]
	}
	if fieldOverride != nil {
		if fieldOverride.OutputOnly != nil {
			v := *fieldOverride.OutputOnly
			log(4, "IsOutputOnlyField(%s) => %t (field override)", fieldPath, v)
			return v
		}
	}

	fieldBehaviors := proto.GetExtension(field.Options(), annotations.E_FieldBehavior).([]annotations.FieldBehavior)
	for _, fieldBehavior := range fieldBehaviors {
		if fieldBehavior == annotations.FieldBehavior_OUTPUT_ONLY {
			log(4, "IsOutputOnlyField(%s) => %t (proto annotation)", fieldPath, true)
			return true
		}
	}

	log(4, "IsOutputOnlyField(%s) => %t", fieldPath, false)
	return false
}

func (g *TypeDiscovery) IgnoreField(field protoreflect.FieldDescriptor) bool {
	fieldPath := string(field.ContainingMessage().FullName()) + ":" + string(field.Name())

	fieldOverride := g.fieldOverrides[fieldPath]
	if fieldOverride == nil {
		fieldOverride = g.fieldOverrides[string(field.ContainingMessage().FullName())+":*"]
	}
	if fieldOverride != nil {
		if fieldOverride.Ignore {
			log(4, "IgnoreField(%s) => %t (field override)", fieldPath, true)
			return true
		}
	}

	log(4, "IgnoreField(%s) => %t", fieldPath, false)
	return false
}

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
	case "string", "int64", "bool":
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
// We call this method twice per top level message, once to discover input-only types and once to discover output-only types.
// These correspond to reachedOnPath == InputOnly and reachedOnPath == OutputOnly.
func findReachableTypesFromMessage(message protoreflect.MessageDescriptor, pred func(fieldPath []protoreflect.FieldDescriptor) bool, dest map[string]protoreflect.MessageDescriptor) error {
	for i := 0; i < message.Fields().Len(); i++ {
		field := message.Fields().Get(i)
		fieldPath := []protoreflect.FieldDescriptor{field}
		findReachableTypesFromField(fieldPath, pred, dest)
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
// We call this method twice per reached field, once to discover input-only types and once to discover output-only types.
// These correspond to reachedOnPath == InputOnly and reachedOnPath == OutputOnly.
func findReachableTypesFromField(fieldPath []protoreflect.FieldDescriptor, pred func(fieldPath []protoreflect.FieldDescriptor) bool, dest map[string]protoreflect.MessageDescriptor) {
	if pred(fieldPath) == false {
		return
	}

	field := fieldPath[len(fieldPath)-1]
	// klog.Infof("findReachableTypesFromField(%v, %v)", field.FullName(), reachedOnPath)
	if field.Message() != nil {
		// no need to find dependencies for proto messages that are not mapped to KRM Go struct
		if _, ok := protoMessagesNotMappedToGoStruct[string(field.Message().FullName())]; ok {
			return
		}
	}

	// If the field is a map, descend into the key and value types
	if field.IsMap() {
		mapEntry := field.Message()
		if keyField := mapEntry.Fields().ByName("key"); keyField != nil {
			keyFieldPath := append([]protoreflect.FieldDescriptor{}, fieldPath...)
			keyFieldPath = append(keyFieldPath, keyField)

			findReachableTypesFromField(keyFieldPath, pred, dest)
		}

		if valueField := mapEntry.Fields().ByName("value"); valueField != nil {
			valueFieldPath := append([]protoreflect.FieldDescriptor{}, fieldPath...)
			valueFieldPath = append(valueFieldPath, valueField)

			findReachableTypesFromField(valueFieldPath, pred, dest)
		}
		return
	}

	switch field.Kind() {
	case protoreflect.MessageKind:
		msg := field.Message()
		fqn := string(msg.FullName())
		if _, found := dest[fqn]; found {
			// No need to visit again
			return
		}
		dest[fqn] = msg

		for i := 0; i < msg.Fields().Len(); i++ {
			childField := msg.Fields().Get(i)

			childFieldPath := append([]protoreflect.FieldDescriptor{}, fieldPath...)
			childFieldPath = append(childFieldPath, childField)

			findReachableTypesFromField(childFieldPath, pred, dest)
		}

	case protoreflect.EnumKind:
		// deps[string(field.Enum().FullName())] = true  // Skip enum because enum is mapped to Go string in code generation
	}
}
