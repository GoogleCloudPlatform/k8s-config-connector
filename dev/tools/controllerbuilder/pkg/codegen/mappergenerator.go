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
	"strings"
	"unicode"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"
	protoapi "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"k8s.io/klog/v2"
)

type MapperGenerator struct {
	generatorBase
	goPathForMessage OutputFunc

	goPackages []*gocode.Package

	typePairs           []typePair
	precomputedMappings map[protoreflect.FullName]map[string]*gocode.GoStruct
}

func NewMapperGenerator(goPathForMessage OutputFunc, outputBaseDir string) *MapperGenerator {
	g := &MapperGenerator{
		goPathForMessage:    goPathForMessage,
		precomputedMappings: make(map[protoreflect.FullName]map[string]*gocode.GoStruct),
	}
	g.generatorBase.init(outputBaseDir)
	return g
}

type OutputFunc func(msg protoreflect.MessageDescriptor) (goPath string, shouldWrite bool)

func (v *MapperGenerator) VisitGoCode(goPackage string, basePath string) error {
	packages, err := gocode.LoadPackageTree(goPackage, basePath)
	if err != nil {
		return fmt.Errorf("inspecting go code: %w", err)
	}

	for _, pkg := range packages {
		// annotation := pkg.GetAnnotation("+kcc:proto")
		// klog.Infof("got package %v for proto %v", pkg.SourceDir, pkg.Comments)
		// if annotation != "" {
		// 	klog.Infof("got package %v for proto %v", pkg.SourceDir, annotation)
		v.goPackages = append(v.goPackages, pkg)
		// v.goPackages[annotation] = pkg
		// }

		// Populate the precomputedMappings
		for _, s := range pkg.Structs {
			if len(s.Comments) != 0 {
				for _, c := range s.Comments {
					for _, line := range strings.Split(c, "\n") {
						line = strings.TrimSpace(line)
						if strings.HasPrefix(line, "+kcc:proto=") {
							protoName := protoreflect.FullName(strings.TrimPrefix(line, "+kcc:proto="))
							if _, ok := v.precomputedMappings[protoName]; !ok {
								v.precomputedMappings[protoName] = make(map[string]*gocode.GoStruct)
							}
							v.precomputedMappings[protoName][s.Name] = s
						}
					}
				}
			}
		}
	}

	return nil
}

func (v *MapperGenerator) VisitProto(api *protoapi.Proto) error {
	sortedFiles := api.SortedFiles()
	for _, f := range sortedFiles {
		v.visitFile(f)
	}
	return nil
}

func (g *MapperGenerator) visitFile(f protoreflect.FileDescriptor) {
	for _, msg := range sortIntoMessageSlice(f.Messages()) {
		g.visitMessage(msg)
	}
}

func (v *MapperGenerator) findKRMStructsForProto(msg protoreflect.MessageDescriptor) map[string]*gocode.GoStruct {
	// Use precomputed mappings
	if matches, found := v.precomputedMappings[msg.FullName()]; found {
		return matches
	}
	klog.V(2).Infof("did not find mapping for %q", msg.FullName())
	return nil
}

func (v *MapperGenerator) visitMessage(msg protoreflect.MessageDescriptor) {
	if _, visit := v.goPathForMessage(msg); !visit {
		return
	}
	goTypes := v.findKRMStructsForProto(msg)

	if goTypes == nil || len(goTypes) == 0 {
		klog.Infof("no krm for %v", msg.FullName())
		return
	}
	parentFile := msg.ParentFile()
	fileOptions := parentFile.Options().(*descriptorpb.FileOptions)
	protoGoPackage := fileOptions.GetGoPackage()
	if ix := strings.Index(protoGoPackage, ";"); ix != -1 {
		protoGoPackage = protoGoPackage[:ix]
	}

	// Some exceptions in our proto mapping
	// TODO: Move to flag?  How many of these are there?
	switch protoGoPackage {
	case "cloud.google.com/go/networkconnectivity/apiv1/networkconnectivitypb":
		protoGoPackage = "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	}

	for _, goType := range goTypes {
		v.typePairs = append(v.typePairs, typePair{
			ProtoPackage:   msg.ParentFile().Package(),
			ProtoGoPackage: protoGoPackage,
			KRMType:        goType,
			Proto:          msg,
		})
	}

	for _, msg := range sortIntoMessageSlice(msg.Messages()) {
		v.visitMessage(msg)
	}
}

type typePair struct {
	ProtoPackage   protoreflect.FullName
	ProtoGoPackage string
	KRMType        *gocode.GoStruct
	Proto          protoreflect.MessageDescriptor
}

func lastGoComponent(goPackage string) string {
	return filepath.Base(goPackage)
}

func (v *MapperGenerator) GenerateMappers() error {
	sort.Slice(v.typePairs, func(i, j int) bool {
		return v.typePairs[i].KRMType.Name < v.typePairs[j].KRMType.Name
	})
	for _, pair := range v.typePairs {
		goPackage, shouldVisit := v.goPathForMessage(pair.Proto)
		if !shouldVisit {
			continue
		}

		k := generatedFileKey{
			GoPackage: goPackage,
			FileName:  "mapper.generated.go",
		}
		out := v.getOutputFile(k)
		if out.contents.Len() == 0 {
			pbPackage := pair.ProtoGoPackage
			krmPackage := pair.KRMType.GoPackage

			// TODO: Can we figure out a better way here?
			switch pbPackage {
			case "cloud.google.com/go/bigtable/admin/apiv2/adminpb":
				pbPackage = "google.golang.org/genproto/googleapis/bigtable/admin/v2"
			}

			out.contents.WriteString(fmt.Sprintf("package %s\n\n", lastGoComponent(goPackage)))
			out.contents.WriteString("import (\n")
			out.contents.WriteString(fmt.Sprintf("\trefs %q\n", "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"))
			out.contents.WriteString(fmt.Sprintf("\tpb %q\n", pbPackage))
			out.contents.WriteString(fmt.Sprintf("\tkrm %q\n", krmPackage))
			out.contents.WriteString(fmt.Sprintf("\t%q\n", "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"))
			out.contents.WriteString(")\n")
		}

		v.writeMapFunctionsForPair(&out.contents, out.OutputDir(), &pair)
	}

	return nil
}

func (v *MapperGenerator) writeMapFunctionsForPair(out io.Writer, srcDir string, pair *typePair) {
	msg := pair.Proto
	pbTypeName := protoNameForType(msg)

	goType := pair.KRMType
	goTypeName := goType.Name

	goFields := make(map[string]*gocode.StructField)
	for _, f := range goType.Fields {
		goFields[f.Name] = f
	}

	if v.findFuncDeclaration(goTypeName+"_FromProto", srcDir, true) == nil {
		fmt.Fprintf(out, "func %s_FromProto(mapCtx *direct.MapContext, in *pb.%s) *krm.%s {\n", goTypeName, pbTypeName, goTypeName)
		fmt.Fprintf(out, "\tif in == nil {\n")
		fmt.Fprintf(out, "\t\treturn nil\n")
		fmt.Fprintf(out, "\t}\n")
		fmt.Fprintf(out, "\tout := &krm.%s{}\n", goTypeName)
		for i := 0; i < msg.Fields().Len(); i++ {
			protoField := msg.Fields().Get(i)
			protoFieldName := buildGoProtoFieldName(protoField)
			protoAccessor := "Get" + protoFieldName + "()"

			krmFieldName := goFieldName(protoField)
			krmField := goFields[krmFieldName]
			if krmField == nil {
				// Support refs
				krmFieldRef := goFields[krmFieldName+"Ref"]
				if krmFieldRef != nil {
					fmt.Fprintf(out, "\tif in.%s != \"\" {\n", protoAccessor)
					fmt.Fprintf(out, "\t	out.%v = &refs.%v{External: in.%v}\n", krmFieldRef.Name, strings.TrimPrefix(krmFieldRef.Type, "*refs."), protoAccessor)
					fmt.Fprintf(out, "\t}\n")
					continue
				}

				if !v.fieldExistInCounterpartStruct(goType, krmFieldName) && !v.fieldExistInCounterpartStruct(goType, krmFieldName+"Ref") { // special handling for Spec and observedState structs which map to the same proto message.
					fmt.Fprintf(out, "\t// MISSING: %s\n", krmFieldName)
				}
				continue
			}

			if protoField.Cardinality() == protoreflect.Repeated {
				useSliceFromProtoFunction := ""
				useCustomMethod := false

				switch protoField.Kind() {
				case protoreflect.MessageKind:
					krmElemTypeName := krmField.Type
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "*")
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "[]")

					functionName := krmElemTypeName + "_FromProto"
					useSliceFromProtoFunction = functionName
				case protoreflect.StringKind:
					if krmField.Type != "[]string" {
						useCustomMethod = true
						// useSliceFromProto = fmt.Sprintf("%s_%s_FromProto", goTypeName, protoFieldName)
					}
				case protoreflect.EnumKind:
					krmElemTypeName := krmField.Type
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "*")
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "[]")

					functionName := "Enum_FromProto"
					useSliceFromProtoFunction = fmt.Sprintf("%s(mapCtx, in.%s)",
						functionName,
						krmFieldName,
					)

				case
					protoreflect.FloatKind,
					protoreflect.DoubleKind,
					protoreflect.BoolKind,
					protoreflect.Int64Kind,
					protoreflect.Int32Kind,
					protoreflect.Uint32Kind,
					protoreflect.Uint64Kind,
					protoreflect.BytesKind:
					useSliceFromProtoFunction = ""
				default:
					klog.Fatalf("unhandled kind %q for repeated field %v", protoField.Kind(), protoField)
				}

				if protoField.IsMap() {
					entryMsg := protoField.Message()
					keyKind := entryMsg.Fields().ByName("key").Kind()
					valueKind := entryMsg.Fields().ByName("value").Kind()
					if keyKind == protoreflect.StringKind && valueKind == protoreflect.StringKind {
						useSliceFromProtoFunction = ""
					} else if keyKind == protoreflect.StringKind && valueKind == protoreflect.Int64Kind {
						useSliceFromProtoFunction = ""
					} else {
						fmt.Fprintf(out, "\t// TODO: map type %v %v for field %v\n", keyKind, valueKind, krmFieldName)
						continue
					}
				}

				if useSliceFromProtoFunction != "" {
					fmt.Fprintf(out, "\tout.%s = direct.Slice_FromProto(mapCtx, in.%s, %s)\n",
						krmFieldName,
						krmFieldName,
						useSliceFromProtoFunction,
					)
				} else if useCustomMethod {
					methodName := fmt.Sprintf("%s_%s_FromProto", goTypeName, protoFieldName)
					fmt.Fprintf(out, "\tout.%s = %s(mapCtx, in.%s)\n",
						krmFieldName,
						methodName,
						krmFieldName,
					)
				} else {
					fmt.Fprintf(out, "\tout.%s = in.%s\n",
						krmFieldName,
						krmFieldName,
					)
				}
				continue
			}

			switch protoField.Kind() {
			case protoreflect.MessageKind:
				krmTypeName := krmField.Type
				krmTypeName = strings.TrimPrefix(krmTypeName, "*")

				functionName := krmTypeName + "_FromProto"
				switch krmTypeName {
				case "string":
					functionName = string(msg.Name()) + "_" + krmFieldName + "_FromProto"
				}

				// special handling for proto messages that mapped to KRM string
				if _, ok := protoMessagesNotMappedToGoStruct[string(protoField.Message().FullName())]; ok {
					functionName = krmFromProtoFunctionName(protoField, krmField.Name)
				}

				fmt.Fprintf(out, "\tout.%s = %s(mapCtx, in.%s)\n",
					krmFieldName,
					functionName,
					protoAccessor,
				)
			case protoreflect.EnumKind:
				functionName := "direct.Enum_FromProto"
				// Not needed if we use the accessor:
				// protoTypeName := "pb." + protoNameForEnum(protoField.Enum())
				// if protoIsPointerInGo(protoField) {
				// 	functionName = "EnumPtr_FromProto[" + protoTypeName + "]"
				// }
				fmt.Fprintf(out, "\tout.%s = %s(mapCtx, in.%s)\n",
					krmFieldName,
					functionName,
					protoAccessor,
				)
			case protoreflect.StringKind,
				protoreflect.FloatKind,
				protoreflect.DoubleKind,
				protoreflect.BoolKind,
				protoreflect.Int64Kind,
				protoreflect.Int32Kind,
				protoreflect.Uint32Kind,
				protoreflect.Uint64Kind,
				protoreflect.Fixed64Kind:
				if protoIsPointerInGo(protoField) {
					fmt.Fprintf(out, "\tout.%s = in.%s\n",
						krmFieldName,
						protoFieldName,
					)
				} else {
					fmt.Fprintf(out, "\tout.%s = direct.LazyPtr(in.%s)\n",
						krmFieldName,
						protoAccessor,
					)
				}

			case protoreflect.BytesKind:
				fmt.Fprintf(out, "\tout.%s = in.%s\n",
					krmFieldName,
					protoAccessor,
				)

			default:
				klog.Fatalf("unhandled kind %q for field %v", protoField.Kind(), protoField)
			}
		}
		fmt.Fprintf(out, "\treturn out\n")
		fmt.Fprintf(out, "}\n")
	}

	if v.findFuncDeclaration(goTypeName+"_ToProto", srcDir, true) == nil {
		fmt.Fprintf(out, "func %s_ToProto(mapCtx *direct.MapContext, in *krm.%s) *pb.%s {\n", goTypeName, goTypeName, pbTypeName)
		fmt.Fprintf(out, "\tif in == nil {\n")
		fmt.Fprintf(out, "\t\treturn nil\n")
		fmt.Fprintf(out, "\t}\n")
		fmt.Fprintf(out, "\tout := &pb.%s{}\n", pbTypeName)
		for i := 0; i < msg.Fields().Len(); i++ {
			protoField := msg.Fields().Get(i)
			protoFieldName := strings.Title(protoField.JSONName())

			krmFieldName := goFieldName(protoField)
			krmField := goFields[krmFieldName]
			if krmField == nil {
				// Support refs
				krmFieldRef := goFields[krmFieldName+"Ref"]
				if krmFieldRef != nil {
					fmt.Fprintf(out, "\tif in.%s != nil {\n", krmFieldRef.Name)
					fmt.Fprintf(out, "\t	out.%v = in.%v.External\n", protoFieldName, krmFieldRef.Name)
					fmt.Fprintf(out, "\t}\n")
					continue
				}

				if !v.fieldExistInCounterpartStruct(goType, krmFieldName) && !v.fieldExistInCounterpartStruct(goType, krmFieldName+"Ref") { // special handling for spec and state structs which map to the same proto message.
					fmt.Fprintf(out, "\t// MISSING: %s\n", krmFieldName)
				}
				continue
			}

			// protoFieldName := buildGoProtoFieldName(protoField)

			if protoField.Cardinality() == protoreflect.Repeated {
				useSliceToProtoFunction := ""
				useCustomMethod := false

				switch protoField.Kind() {
				case protoreflect.MessageKind:
					krmElemTypeName := krmField.Type
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "*")
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "[]")

					functionName := krmElemTypeName + "_ToProto"
					useSliceToProtoFunction = functionName

				case protoreflect.StringKind:
					if krmField.Type != "[]string" {
						useCustomMethod = true
						//useSliceToProtoFunction = fmt.Sprintf("%s_%s_ToProto", goTypeName, protoFieldName)
					}

				case protoreflect.EnumKind:
					krmElemTypeName := krmField.Type
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "*")
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "[]")

					protoTypeName := "pb." + protoNameForEnum(protoField.Enum())
					functionName := "direct.Enum_ToProto"
					useSliceToProtoFunction = fmt.Sprintf("%s[%s](mapCtx, in.%s)",
						functionName,
						protoTypeName,
						krmFieldName,
					)

				case protoreflect.FloatKind,
					protoreflect.DoubleKind,
					protoreflect.BoolKind,
					protoreflect.Int64Kind,
					protoreflect.Int32Kind,
					protoreflect.Uint32Kind,
					protoreflect.Uint64Kind,
					protoreflect.BytesKind:

					useSliceToProtoFunction = ""
				default:
					klog.Fatalf("unhandled kind %q for repeated field %v", protoField.Kind(), protoField)
				}

				if protoField.IsMap() {
					entryMsg := protoField.Message()
					keyKind := entryMsg.Fields().ByName("key").Kind()
					valueKind := entryMsg.Fields().ByName("value").Kind()
					if keyKind == protoreflect.StringKind && valueKind == protoreflect.StringKind {
						useSliceToProtoFunction = ""
					} else if keyKind == protoreflect.StringKind && valueKind == protoreflect.Int64Kind {
						useSliceToProtoFunction = ""
					} else {
						fmt.Fprintf(out, "\t// TODO: map type %v %v for field %v\n", keyKind, valueKind, krmFieldName)
						continue
					}
				}

				if useSliceToProtoFunction != "" {
					fmt.Fprintf(out, "\tout.%s = direct.Slice_ToProto(mapCtx, in.%s, %s)\n",
						protoFieldName,
						krmFieldName,
						useSliceToProtoFunction,
					)
				} else if useCustomMethod {
					methodName := fmt.Sprintf("%s_%s_ToProto", goTypeName, protoFieldName)
					fmt.Fprintf(out, "\tout.%s = %s(mapCtx, in.%s)\n",
						krmFieldName,
						methodName,
						krmFieldName,
					)
				} else {
					fmt.Fprintf(out, "\tout.%s = in.%s\n",
						protoFieldName,
						krmFieldName,
					)
				}
				continue
			}

			switch protoField.Kind() {
			case protoreflect.MessageKind:
				krmTypeName := krmField.Type
				krmTypeName = strings.TrimPrefix(krmTypeName, "*")

				functionName := krmTypeName + "_ToProto"
				switch krmTypeName {
				case "string":
					functionName = string(msg.Name()) + "_" + krmFieldName + "_ToProto"
				}

				// special handling for proto messages that mapped to KRM string
				if _, ok := protoMessagesNotMappedToGoStruct[string(protoField.Message().FullName())]; ok {
					functionName = krmToProtoFunctionName(protoField, krmField.Name)
				}

				oneof := protoField.ContainingOneof()
				if oneof != nil {
					fmt.Fprintf(out, "\tif oneof := %s(mapCtx, in.%s); oneof != nil {\n",
						functionName,
						krmFieldName,
					)

					oneofFieldName := ToGoFieldName(oneof.Name())

					oneofTypeName := protoNameForOneOf(protoField)

					fmt.Fprintf(out, "\t\tout.%s = &pb.%s{%s: oneof}\n",
						oneofFieldName,
						oneofTypeName,
						protoFieldName)
					fmt.Fprintf(out, "\t}\n")
					continue
				}
				fmt.Fprintf(out, "\tout.%s = %s(mapCtx, in.%s)\n",
					protoFieldName,
					functionName,
					krmFieldName,
				)
			case protoreflect.EnumKind:
				protoTypeName := "pb." + protoNameForEnum(protoField.Enum())
				functionName := "direct.Enum_ToProto"
				if protoIsPointerInGo(protoField) {
					functionName = "EnumPtr_ToProto[" + protoTypeName + "]"
				}

				oneof := protoField.ContainingOneof()
				if oneof != nil {
					// These are very rare and irregular; just require a custom method
					functionName := fmt.Sprintf("%s_%s_ToProto", goTypeName, protoFieldName)

					fmt.Fprintf(out, "\tif oneof := %s(mapCtx, in.%s); oneof != nil {\n",
						functionName,
						krmFieldName,
					)

					oneofFieldName := ToGoFieldName(oneof.Name())

					fmt.Fprintf(out, "\t\tout.%s = oneof\n",
						oneofFieldName)
					fmt.Fprintf(out, "\t}\n")
					continue
				}

				fmt.Fprintf(out, "\tout.%s = %s[%s](mapCtx, in.%s)\n",
					protoFieldName,
					functionName,
					protoTypeName,
					krmFieldName,
				)
			case protoreflect.StringKind,
				protoreflect.FloatKind,
				protoreflect.DoubleKind,
				protoreflect.BoolKind,
				protoreflect.Int64Kind,
				protoreflect.Int32Kind,
				protoreflect.Uint32Kind,
				protoreflect.Uint64Kind,
				protoreflect.Fixed64Kind,
				protoreflect.BytesKind:

				useCustomMethod := false

				switch protoField.Kind() {
				case protoreflect.StringKind:
					if krmField.Type != "*string" {
						useCustomMethod = true
					}
				}

				oneof := protoField.ContainingOneof()
				if protoField.HasOptionalKeyword() {
					fmt.Fprintf(out, "\tout.%s = in.%s\n",
						protoFieldName,
						krmFieldName,
					)
				} else if oneof != nil {
					functionName := fmt.Sprintf("%s_%s_ToProto", goTypeName, protoFieldName)
					fmt.Fprintf(out, "\tif oneof := %s(mapCtx, in.%s); oneof != nil {\n",
						functionName,
						krmFieldName,
					)

					oneofFieldName := ToGoFieldName(oneof.Name())

					fmt.Fprintf(out, "\t\tout.%s = oneof\n",
						oneofFieldName)
					fmt.Fprintf(out, "\t}\n")
				} else if useCustomMethod {
					methodName := fmt.Sprintf("%s_%s_ToProto", goTypeName, protoFieldName)
					fmt.Fprintf(out, "\tout.%s = %s(mapCtx, in.%s)\n",
						krmFieldName,
						methodName,
						krmFieldName,
					)
				} else if protoField.Kind() == protoreflect.BytesKind {
					fmt.Fprintf(out, "\tout.%s = in.%s\n",
						protoFieldName,
						krmFieldName,
					)
				} else {
					fmt.Fprintf(out, "\tout.%s = direct.ValueOf(in.%s)\n",
						protoFieldName,
						krmFieldName,
					)
				}

			default:
				klog.Fatalf("unhandled kind %q for field %v", protoField.Kind(), protoField)
			}

		}
		fmt.Fprintf(out, "\treturn out\n")
		fmt.Fprintf(out, "}\n")
	}

}

func protoNameForType(msg protoreflect.MessageDescriptor) string {
	fullName := string(msg.FullName())
	fullName = strings.TrimPrefix(fullName, string(msg.ParentFile().FullName()))
	fullName = strings.TrimPrefix(fullName, ".")
	fullName = strings.ReplaceAll(fullName, ".", "_")
	return fullName
}

func protoNameForEnum(msg protoreflect.EnumDescriptor) string {
	fullName := string(msg.FullName())
	fullName = strings.TrimPrefix(fullName, string(msg.ParentFile().FullName()))
	fullName = strings.TrimPrefix(fullName, ".")
	fullName = strings.ReplaceAll(fullName, ".", "_")
	return fullName
}

func protoNameForOneOf(field protoreflect.FieldDescriptor) string {
	msg := field.Parent().(protoreflect.MessageDescriptor)
	oneofKey := ToGoFieldName(field.Name())
	name := protoNameForType(msg) + "_" + oneofKey

	// Special case: check for a collision
	if field.Message() != nil {
		elemTypeName := protoNameForType(field.Message())
		if name == elemTypeName {
			name += "_"
		}
	}
	if field.Enum() != nil {
		elemTypeName := protoNameForEnum(field.Enum())
		if name == elemTypeName {
			name += "_"
		}
	}
	return name
}

func ToGoFieldName(name protoreflect.Name) string {
	tokens := strings.Split(string(name), "_")
	for i, token := range tokens {
		tokens[i] = strings.Title(token)
	}
	return strings.Join(tokens, "")
}

// protoIsPointerInGo returns if the field is going to be represented as a pointer in go.
// Most proto3 fields are not pointers, but a few are.
func protoIsPointerInGo(field protoreflect.FieldDescriptor) bool {
	switch field.Kind() {
	case protoreflect.EnumKind:
		if field.HasOptionalKeyword() {
			return true
		}
		return false

	case protoreflect.StringKind,
		protoreflect.FloatKind,
		protoreflect.DoubleKind,
		protoreflect.BoolKind,
		protoreflect.Int64Kind,
		protoreflect.Int32Kind,
		protoreflect.Uint32Kind,
		protoreflect.Uint64Kind,
		protoreflect.Fixed64Kind,
		protoreflect.BytesKind:
		return field.HasOptionalKeyword()

	default:
		klog.Fatalf("protoIsPointerInGo not implemented for %v", field)
	}
	return false
}

func sortIntoMessageSlice(messages protoreflect.MessageDescriptors) []protoreflect.MessageDescriptor {
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

func krmFromProtoFunctionName(protoField protoreflect.FieldDescriptor, krmFieldName string) string {
	fullname := string(protoField.Message().FullName())
	switch fullname {
	case "google.protobuf.Timestamp":
		return "direct.StringTimestamp_FromProto"
	case "google.protobuf.Struct":
		return krmFieldName + "_FromProto"
	case "google.protobuf.Duration":
		return "direct.StringDuration_FromProto"
	}
	klog.Fatalf("unhandled case in krmFromProtoFunctionName for proto field %s", fullname)
	return ""
}

func krmToProtoFunctionName(protoField protoreflect.FieldDescriptor, krmFieldName string) string {
	fullname := string(protoField.Message().FullName())
	switch fullname {
	case "google.protobuf.Timestamp":
		return "direct.StringTimestamp_ToProto"
	case "google.protobuf.Struct":
		return krmFieldName + "_ToProto"
	case "google.protobuf.Duration":
		return "direct.StringDuration_ToProto"
	}
	klog.Fatalf("unhandled case in krmToProtoFunctionName for proto field %s", fullname)
	return ""
}

func (v *MapperGenerator) fieldExistInCounterpartStruct(goType *gocode.GoStruct, krmFieldName string) bool {
	counterpartTypeName := getCounterpartTypeName(goType.Name)
	if counterpartTypeName == "" {
		return false
	}

	for _, pair := range v.typePairs {
		if pair.KRMType.Name == counterpartTypeName {
			return fieldExistInStruct(pair.KRMType, krmFieldName)
		}
	}

	return false
}

func getCounterpartTypeName(goTypeName string) string {
	switch {
	case strings.HasSuffix(goTypeName, "Spec"):
		return strings.TrimSuffix(goTypeName, "Spec") + "ObservedState"
	case strings.HasSuffix(goTypeName, "ObservedState"):
		return strings.TrimSuffix(goTypeName, "ObservedState") + "Spec"
	default:
		return ""
	}
}

func fieldExistInStruct(goType *gocode.GoStruct, fieldName string) bool {
	for _, field := range goType.Fields {
		if field.Name == fieldName {
			return true
		}
	}
	return false
}

func buildGoProtoFieldName(fd protoreflect.FieldDescriptor) string {
	s := fd.JSONName()
	var out strings.Builder
	var previous rune
	previous = '_' // Should start with upper case
	for _, r := range s {
		switch previous {
		case '_', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			out.WriteRune(unicode.ToUpper(r))
		default:
			out.WriteRune(r)
		}
		previous = r
	}
	return out.String()
}
