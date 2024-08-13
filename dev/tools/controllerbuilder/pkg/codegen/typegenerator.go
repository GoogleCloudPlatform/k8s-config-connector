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
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	protoapi "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"

	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

type TypeGenerator struct {
	generatorBase
	goPackageForMessage OutputFunc
	visitedMessages     []protoreflect.MessageDescriptor
}

func NewTypeGenerator(goPathForMessage OutputFunc, outputBaseDir string) *TypeGenerator {
	g := &TypeGenerator{
		goPackageForMessage: goPathForMessage,
	}
	g.generatorBase.init(outputBaseDir)
	return g
}

type OutputFunc func(msg protoreflect.MessageDescriptor) (goPath string, shouldWrite bool)

func (g *TypeGenerator) VisitProto(api *protoapi.Proto) error {
	seen := make(map[string]bool) // Track seen message names
	sortedFiles := api.SortedFiles()
	for _, f := range sortedFiles {
		g.visitFile(f, seen)
	}

	g.writeVisitedMessages()
	return nil
}

func (g *TypeGenerator) visitFile(f protoreflect.FileDescriptor, seen map[string]bool) {
	// visit and collect all messages (including nested)
	for i := 0; i < f.Messages().Len(); i++ {
		msg := f.Messages().Get(i)
		g.visitMessage(msg, seen)
	}
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

func (g *TypeGenerator) visitMessage(msg protoreflect.MessageDescriptor, seen map[string]bool) {
	if _, visit := g.goPackageForMessage(msg); !visit {
		return
	}

	if _, ok := seen[string(msg.FullName())]; ok {
		return
	}
	seen[string(msg.FullName())] = true

	g.visitedMessages = append(g.visitedMessages, msg)

	for i := 0; i < msg.Messages().Len(); i++ {
		nestedMsg := msg.Messages().Get(i)
		g.visitMessage(nestedMsg, seen)
	}
}

func (g *TypeGenerator) writeVisitedMessages() {
	for _, msg := range sorted(g.visitedMessages) {
		if msg.IsMapEntry() {
			continue
		}

		goPackage, ok := g.goPackageForMessage(msg)
		if !ok {
			continue
		}

		krmVersion := filepath.Base(goPackage)

		k := generatedFileKey{
			GoPackage: goPackage,
			FileName:  "types.generated.go",
		}
		out := g.getOutputFile(k)

		goTypeName := goNameForProtoMessage(msg, msg)
		skipGenerated := true
		goType, err := g.findTypeDeclaration(goTypeName, out.OutputDir(), skipGenerated)
		if err != nil {
			g.Errorf("looking up go type: %w", err)
			continue
		}
		if goType != nil {
			klog.Infof("found existing non-generated go type %q, won't generate", goTypeName)
			continue
		}

		goType, err = g.findTypeDeclarationWithProtoTag(string(msg.FullName()), out.OutputDir(), skipGenerated)
		if err != nil {
			g.Errorf("looking up go type by proto tag: %w", err)
			continue
		}
		if goType != nil {
			klog.Infof("found existing non-generated go type with proto tag %q, won't generate", msg.FullName())
			continue
		}

		w := &out.contents

		if out.contents.Len() == 0 {
			fmt.Fprintf(w, "package %s\n", krmVersion)
		}
		WriteMessage(w, msg)
	}
}

func WriteMessage(out io.Writer, msg protoreflect.MessageDescriptor) {
	goType := goNameForProtoMessage(msg, msg)

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

func sorted(messages []protoreflect.MessageDescriptor) []protoreflect.MessageDescriptor {
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].FullName() < messages[j].FullName()
	})
	return messages
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
