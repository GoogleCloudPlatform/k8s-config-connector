// Copyright 2025 Google LLC
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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

type TypeGeneratorV2 struct {
	generatorBase

	api          *protoapi.Proto
	goPackage    string
	protoPackage string             // proto package name which contains the resources to be generated
	resources    []ResourceMetadata // top level resources to be generated

	allowOverride bool // allow override of existing types

	// implementation details, internal to the generator
	visitedMessages []protoreflect.MessageDescriptor
	outputMessages  []*OutputMessageDetails
	ignoredFields   map[string]bool // map of fully qualified field names to skip
}

func NewTypeGeneratorV2(goPackage string, api *protoapi.Proto, outputBaseDir string, serviceMetadata *ServiceMetadata) *TypeGeneratorV2 {
	// Create map of ignored fields
	ignoredFields := make(map[string]bool)
	for _, resource := range serviceMetadata.Resources {
		for _, field := range resource.IgnoredFields {
			ignoredFields[field] = true
		}
	}

	g := &TypeGeneratorV2{
		goPackage:     goPackage,
		api:           api,
		protoPackage:  serviceMetadata.Service,
		resources:     serviceMetadata.Resources,
		allowOverride: true, // TODO: this should be a command line flag
		ignoredFields: ignoredFields,
	}
	g.generatorBase.init(outputBaseDir)
	return g
}

func (g *TypeGeneratorV2) VisitProto() error {
	for _, resource := range g.resources {
		resourceProtoFullName := g.protoPackage + "." + resource.ProtoName

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
	}

	return nil
}

func (g *TypeGeneratorV2) visitMessage(message protoreflect.MessageDescriptor) error {
	g.visitedMessages = append(g.visitedMessages, message)

	msgs, err := FindDependenciesForMessage(message, nil)
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

func (g *TypeGeneratorV2) getGeneratedTypesFile() (*generatedFile, error) {
	krmVersion := filepath.Base(g.goPackage)

	k := generatedFileKey{
		GoPackage: g.goPackage,
		FileName:  "types.generated.go",
	}
	out := g.getOutputFile(k)
	out.packageName = krmVersion
	return out, nil
}

func (g *TypeGeneratorV2) WriteVisitedMessages() error {
	out, err := g.getGeneratedTypesFile()
	if err != nil {
		return err
	}

	for _, msg := range deduplicateAndSort(g.visitedMessages) {
		if msg.IsMapEntry() {
			continue
		}

		if !g.allowOverride { // skip on existing types if override is not allowed
			exist, err := g.typeExists(string(msg.FullName()), out.OutputDir())
			if err != nil {
				return err
			}
			if exist {
				klog.Infof("found existing non-generated go type for proto message %q, will override", string(msg.FullName()))
				continue
			}
		}

		// Use a different Go type name for top level messages that correspond to a KRM resource
		var kind string
		for _, resource := range g.resources {
			if string(msg.FullName()) == g.protoPackage+"."+resource.ProtoName {
				kind = resource.Kind
				break
			}
		}

		if kind != "" {
			g.writeTopLevelSpecMessage(&out.body, msg, kind)
		} else {
			g.writeMessage(&out.body, msg)
		}
	}

	// add imports
	out.addImport("refs", "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1")
	out.addImport("commonv1alpha1", "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1")

	return errors.Join(g.errors...)
}

func (g *TypeGeneratorV2) WriteOutputMessages() error {
	out, err := g.getGeneratedTypesFile()
	if err != nil {
		return err
	}

	for _, msgDetails := range deduplicateAndSortOutputMessages(g.outputMessages) {
		msg := msgDetails.Message
		if msg.IsMapEntry() {
			continue
		}

		if !g.allowOverride { // skip on existing types if override is not allowed
			exist, err := g.typeExists(string(msg.FullName()), out.OutputDir())
			if err != nil {
				return err
			}
			if exist {
				klog.Infof("found existing non-generated go type for proto message %q, will override", string(msg.FullName()))
				continue
			}
		}

		g.writeOutputMessage(&out.body, msgDetails)
	}
	return errors.Join(g.errors...)
}

func (g *TypeGeneratorV2) typeExists(protoFullName string, outputDir string) (bool, error) {
	goType, err := g.findTypeDeclarationWithProtoTag(protoFullName, outputDir, true)
	if err != nil {
		return false, fmt.Errorf("looking up go type by proto tag: %w", err)
	}
	if goType != nil {
		klog.Infof("found existing non-generated go type with proto tag %q, will override", protoFullName)
		return true, nil
	}
	return false, nil
}

func (g *TypeGeneratorV2) writeTopLevelSpecMessage(out io.Writer, msg protoreflect.MessageDescriptor, kind string) {
	goType := kind + "Spec"

	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "// %s=%s\n", KCCProtoMessageAnnotation, msg.FullName())
	fmt.Fprintf(out, "type %s struct {\n", goType)

	// Add boilerplate fields
	fmt.Fprintf(out, "\tcommonv1alpha1.CommonSpec `json:\",inline\"`\n\n")

	fmt.Fprintf(out, "\t// +required\n")
	fmt.Fprintf(out, "\tLocation string `json:\"location\"`\n\n")

	fmt.Fprintf(out, "\t// The %s name. If not given, the metadata.name will be used.\n", kind)
	fmt.Fprintf(out, "\tResourceID *string `json:\"resourceID,omitempty\"`\n\n")

	// Write the rest of the fields
	g.writeFields(out, msg)

	fmt.Fprintf(out, "}\n")
}

func (g *TypeGeneratorV2) writeMessage(out io.Writer, msg protoreflect.MessageDescriptor) {
	goType := GoNameForProtoMessage(msg)

	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "// %s=%s\n", KCCProtoMessageAnnotation, msg.FullName())
	fmt.Fprintf(out, "type %s struct {\n", goType)

	g.writeFields(out, msg)

	fmt.Fprintf(out, "}\n")
}

func (g *TypeGeneratorV2) writeOutputMessage(out io.Writer, msgDetails *OutputMessageDetails) {
	msg := msgDetails.Message
	goType := goNameForOutputProtoMessage(msg)

	// Use a different Go type name for top level messages that correspond to a KRM resource
	for _, resource := range g.resources {
		if string(msg.FullName()) == g.protoPackage+"."+resource.ProtoName {
			goType = resource.Kind + "ObservedState"
			break
		}
	}

	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "// %s=%s\n", KCCProtoMessageAnnotation, msg.FullName())
	fmt.Fprintf(out, "type %s struct {\n", goType)

	fieldIndex := 0 // keep track of the index manually since we skip fields that are ignored
	for _, field := range msgDetails.OutputFields {
		if g.ignoredFields[string(field.FullName())] {
			continue
		}

		if !IsFieldBehavior(field, annotations.FieldBehavior_OUTPUT_ONLY) {
			// If field is not explicitly listed as an output, but it appears in OutputMessageDetails,
			// then it must be a parent message that contains a child message with an output.
			WriteField(out, field, msg, fieldIndex, true)
		} else {
			WriteField(out, field, msg, fieldIndex, false)
		}
		fieldIndex++
	}

	fmt.Fprintf(out, "}\n")
}

// writeFields writes configurable fields of a message, skipping ignored fields and output-only fields
func (g *TypeGeneratorV2) writeFields(out io.Writer, msg protoreflect.MessageDescriptor) {
	fieldIndex := 0 // keep track of the index manually since we skip fields that are ignored
	for i := 0; i < msg.Fields().Len(); i++ {
		field := msg.Fields().Get(i)
		if g.ignoredFields[string(field.FullName())] {
			continue
		}
		if !IsFieldBehavior(field, annotations.FieldBehavior_OUTPUT_ONLY) {
			writeFieldV2(out, field, msg, fieldIndex, false)
		}
		fieldIndex++
	}
}

func writeFieldV2(out io.Writer, field protoreflect.FieldDescriptor, msg protoreflect.MessageDescriptor, fieldIndex int, isTransitiveOutput bool) {
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

	// Write field comments and check for required
	if sourceLocations.LeadingComments != "" {
		comment := strings.TrimSpace(sourceLocations.LeadingComments)
		hasRequired := false
		for _, line := range strings.Split(comment, "\n") {
			if strings.TrimSpace(line) == "" {
				fmt.Fprintf(out, "\t//\n")
			} else {
				fmt.Fprintf(out, "\t// %s\n", line)
				if strings.Contains(strings.ToLower(line), "required") {
					hasRequired = true
				}
			}
		}
		// Add +required annotation if needed
		if hasRequired {
			fmt.Fprintf(out, "\t// +required\n")
		}
	}

	fmt.Fprintf(out, "\t// %s=%s\n", KCCProtoFieldAnnotation, field.FullName())
	fmt.Fprintf(out, "\t%s %s `json:\"%s,omitempty\"`\n",
		GoFieldName,
		goType,
		jsonName,
	)
}
