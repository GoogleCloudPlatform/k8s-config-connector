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
	"time"

	protoapi "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"

	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

type TypeGenerator struct {
	generatorBase
	goPathForMessage OutputFunc
}

func NewTypeGenerator(goPathForMessage OutputFunc) *TypeGenerator {
	g := &TypeGenerator{
		goPathForMessage: goPathForMessage,
	}
	g.generatorBase.init()
	return g
}

type OutputFunc func(msg protoreflect.MessageDescriptor) (goPath string, shouldWrite bool)

func (v *TypeGenerator) VisitProto(api *protoapi.Proto) error {
	sortedFiles := api.SortedFiles()
	for _, f := range sortedFiles {
		v.visitFile(f)
	}
	return nil
}

func (g *TypeGenerator) visitFile(f protoreflect.FileDescriptor) {
	for _, msg := range sorted(f.Messages()) {
		g.visitMessage(msg)
	}

	{

		for _, msg := range sorted(f.Messages()) {
			if msg.IsMapEntry() {
				continue
			}

			goPath, ok := g.goPathForMessage(msg)
			if !ok {
				continue
			}

			krmVersion := filepath.Base(goPath)

			k := generatedFileKey{
				GoPackagePath: goPath,
				File:          "types.generated.go",
			}
			out := g.getOutputFile(k)

			w := &out.contents

			if out.contents.Len() == 0 {
				writeCopyright(w, time.Now().Year())

				fmt.Fprintf(w, "package %s\n", krmVersion)
			}

			g.writeTypes(w, msg)
		}
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

func (v *TypeGenerator) visitMessage(msg protoreflect.MessageDescriptor) {
	if _, visit := v.goPathForMessage(msg); !visit {
		return
	}
	// goTypes := v.findKRMStructsForProto(msg)

	// if len(goTypes) == 0 {
	// 	klog.V(2).Infof("no krm for %v", msg.FullName())
	// 	return
	// }
	// for _, goType := range goTypes {
	// 	v.typePairs = append(v.typePairs, typePair{
	// 		ProtoPackage: msg.ParentFile().Package(),
	// 		KRMType:      goType,
	// 		Proto:        msg,
	// 	})
	// }

	for _, msg := range sorted(msg.Messages()) {
		v.visitMessage(msg)
	}
}

func (v *TypeGenerator) writeTypes(out io.Writer, msg protoreflect.MessageDescriptor) {
	goType := goNameForProtoMessage(msg, msg)

	{
		fmt.Fprintf(out, "// +kcc:proto=%s\n", msg.FullName())
		fmt.Fprintf(out, "type %s struct {\n", goType)
		for i := 0; i < msg.Fields().Len(); i++ {
			field := msg.Fields().Get(i)
			sourceLocations := msg.ParentFile().SourceLocations().ByDescriptor(field)

			goFieldName := strings.Title(field.JSONName())
			jsonName := field.JSONName()
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
					fmt.Fprintf(out, "// TODO: map type %v %v\n", keyKind, valueKind)
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
			}

			// Blank line between fields for readability
			if i != 0 {
				fmt.Fprintf(out, "\n")
			}

			if sourceLocations.LeadingComments != "" {
				comment := strings.TrimSpace(sourceLocations.LeadingComments)
				for _, line := range strings.Split(comment, "\n") {
					fmt.Fprintf(out, "    // %s\n", line)
				}
			}

			fmt.Fprintf(out, "    %s %s `json:\"%s,omitempty\"`\n",
				goFieldName,
				goType,
				jsonName,
			)
		}
		fmt.Fprintf(out, "}\n")
	}

	for i := 0; i < msg.Messages().Len(); i++ {
		m := msg.Messages().Get(i)
		if m.IsMapEntry() {
			continue
		}
		v.writeTypes(out, m)
	}

}

func sorted(messages protoreflect.MessageDescriptors) []protoreflect.MessageDescriptor {
	var out []protoreflect.MessageDescriptor
	for i := 0; i < messages.Len(); i++ {
		m := messages.Get(i)
		out = append(out, m)
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].FullName() < out[j].FullName()
	})
	return out
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
