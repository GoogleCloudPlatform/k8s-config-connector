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
	refenenceFields     map[string]string // map of proto message name to referenced go struct name
}

func NewTypeGenerator(goPathForMessage OutputFunc, outputBaseDir string) *TypeGenerator {
	g := &TypeGenerator{
		goPackageForMessage: goPathForMessage,
		refenenceFields:     make(map[string]string),
	}
	g.generatorBase.init(outputBaseDir)
	return g
}

type OutputFunc func(msg protoreflect.MessageDescriptor) (goPath string, shouldWrite bool)

func (g *TypeGenerator) WithReferenceFields(refs []string) *TypeGenerator {
	for _, ref := range refs {
		parts := strings.Split(ref, ":")
		g.refenenceFields[parts[0]] = parts[1]
	}
	return g
}

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

		opts := []MessageWriterOption{
			WithReferenceFields(g.refenenceFields),
			WithOutput(w),
		}
		NewMessageWriter(opts...).WriteMessage(msg)
	}
}

func sorted(messages []protoreflect.MessageDescriptor) []protoreflect.MessageDescriptor {
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].FullName() < messages[j].FullName()
	})
	return messages
}
