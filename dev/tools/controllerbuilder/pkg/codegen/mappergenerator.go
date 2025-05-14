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

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/annotations"
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

	generatedFileAnnotation *annotations.FileAnnotation
}

func NewMapperGenerator(goPathForMessage OutputFunc, outputBaseDir string, generatedFileAnnotation *annotations.FileAnnotation) *MapperGenerator {
	g := &MapperGenerator{
		goPathForMessage:        goPathForMessage,
		precomputedMappings:     make(map[protoreflect.FullName]map[string]*gocode.GoStruct),
		generatedFileAnnotation: generatedFileAnnotation,
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
						if proto, ok := GetProtoMessageFromAnnotation(line); ok {
							protoName := protoreflect.FullName(proto)
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
	case "cloud.google.com/go/bigquery/apiv2/bigquerypb":
		protoGoPackage = "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/v2"
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
		out.packageName = lastGoComponent(goPackage)

		out.fileAnnotation = v.generatedFileAnnotation

		{
			pbPackage := pair.ProtoGoPackage
			krmPackage := pair.KRMType.GoPackage

			// TODO: Can we figure out a better way here?
			switch pbPackage {
			case "cloud.google.com/go/bigtable/admin/apiv2/adminpb":
				pbPackage = "google.golang.org/genproto/googleapis/bigtable/admin/v2"
			}

			out.addImport("refs", "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1")
			out.addImport("pb", pbPackage)
			out.addImport(getKRMImportName(krmPackage), krmPackage)
			out.addImport("", "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct")
		}

		v.writeMapFunctionsForPair(&out.body, out.OutputDir(), &pair)
	}

	return nil
}

func (v *MapperGenerator) writeMapFunctionsForPair(out io.Writer, srcDir string, pair *typePair) {
	klog.V(2).InfoS("writeMapFunctionsForPair", "pair.Proto.FullName", pair.Proto.FullName(), "pair.KRMType.Name", pair.KRMType.Name)
	msg := pair.Proto
	pbTypeName := protoNameForType(msg)

	goType := pair.KRMType
	goTypeName := goType.Name

	goFields := make(map[string]*gocode.StructField)
	for _, f := range goType.Fields {
		goFields[f.Name] = f
	}

	krmImportName := getKRMImportName(pair.KRMType.GoPackage)

	if v.findFuncDeclaration(goTypeName+"_FromProto", srcDir, true) == nil {
		fmt.Fprintf(out, "func %s_FromProto(mapCtx *direct.MapContext, in *pb.%s) *%s.%s {\n", goTypeName, pbTypeName, krmImportName, goTypeName)
		fmt.Fprintf(out, "\tif in == nil {\n")
		fmt.Fprintf(out, "\t\treturn nil\n")
		fmt.Fprintf(out, "\t}\n")
		fmt.Fprintf(out, "\tout := &%s.%s{}\n", krmImportName, goTypeName)
		for i := 0; i < msg.Fields().Len(); i++ {
			protoField := msg.Fields().Get(i)
			protoFieldName := protoNameForField(protoField)
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

				if !v.fieldExistInCounterpartStruct(goType, krmFieldName) && !v.fieldExistInCounterpartStruct(goType, krmFieldName+"Ref") { // special handling for Spec and ObservedState structs which map to the same proto message.
					fmt.Fprintf(out, "\t// MISSING: %s\n", krmFieldName)
					for k := range goFields {
						if strings.EqualFold(k, krmFieldName) {
							fmt.Fprintf(out, "\t// (near miss): %q vs %q\n", krmFieldName, k)
						}
					}
				}
				continue
			}

			if protoField.Cardinality() == protoreflect.Repeated {
				useSliceFromProtoFunction := ""
				useCustomMethod := ""

				switch protoField.Kind() {
				case protoreflect.MessageKind:
					krmElemTypeName := krmField.Type
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "*")
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "[]")

					functionName := krmElemTypeName + "_FromProto"
					useSliceFromProtoFunction = functionName
				case protoreflect.StringKind:
					if krmField.Type != "[]string" {
						useCustomMethod = fmt.Sprintf("%s_%s_FromProto", goTypeName, protoFieldName)
					}
				case protoreflect.EnumKind:
					krmElemTypeName := krmField.Type
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "*")
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "[]")

					useCustomMethod = "direct.EnumSlice_FromProto"

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
				} else if useCustomMethod != "" {
					fmt.Fprintf(out, "\tout.%s = %s(mapCtx, in.%s)\n",
						krmFieldName,
						useCustomMethod,
						protoFieldName,
					)
				} else {
					fmt.Fprintf(out, "\tout.%s = in.%s\n",
						krmFieldName,
						protoFieldName,
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
	} else {
		klog.Infof("found existing non-generated mapping function %q, won't generate", goTypeName+"_FromProto")
	}

	if v.findFuncDeclaration(goTypeName+"_ToProto", srcDir, true) == nil {
		fmt.Fprintf(out, "func %s_ToProto(mapCtx *direct.MapContext, in *%s.%s) *pb.%s {\n", goTypeName, krmImportName, goTypeName, pbTypeName)
		fmt.Fprintf(out, "\tif in == nil {\n")
		fmt.Fprintf(out, "\t\treturn nil\n")
		fmt.Fprintf(out, "\t}\n")
		fmt.Fprintf(out, "\tout := &pb.%s{}\n", pbTypeName)
		for i := 0; i < msg.Fields().Len(); i++ {
			protoField := msg.Fields().Get(i)
			protoFieldName := protoNameForField(protoField)

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

				if !v.fieldExistInCounterpartStruct(goType, krmFieldName) && !v.fieldExistInCounterpartStruct(goType, krmFieldName+"Ref") { // special handling for spec and observedState structs which map to the same proto message.
					fmt.Fprintf(out, "\t// MISSING: %s\n", krmFieldName)
					for k := range goFields {
						if strings.EqualFold(k, krmFieldName) {
							fmt.Fprintf(out, "\t// (near miss): %q vs %q\n", krmFieldName, k)
						}
					}
				}
				continue
			}

			if protoField.Cardinality() == protoreflect.Repeated {
				useSliceToProtoFunction := ""
				useCustomMethod := ""

				switch protoField.Kind() {
				case protoreflect.MessageKind:
					krmElemTypeName := krmField.Type
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "*")
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "[]")

					functionName := krmElemTypeName + "_ToProto"
					useSliceToProtoFunction = functionName

				case protoreflect.StringKind:
					if krmField.Type != "[]string" {
						useCustomMethod = fmt.Sprintf("%s_%s_ToProto", goTypeName, protoFieldName)
					}

				case protoreflect.EnumKind:
					krmElemTypeName := krmField.Type
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "*")
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "[]")

					protoTypeName := "pb." + protoNameForEnum(protoField.Enum())
					useCustomMethod = fmt.Sprintf("direct.EnumSlice_ToProto[%s]", protoTypeName)

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
				} else if useCustomMethod != "" {
					fmt.Fprintf(out, "\tout.%s = %s(mapCtx, in.%s)\n",
						krmFieldName,
						useCustomMethod,
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

				useCustomMethod := ""

				switch protoField.Kind() {
				case protoreflect.StringKind:
					if krmField.Type != "*string" {
						useCustomMethod = fmt.Sprintf("%s_%s_ToProto", goTypeName, protoFieldName)
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
				} else if useCustomMethod != "" {
					fmt.Fprintf(out, "\tout.%s = %s(mapCtx, in.%s)\n",
						krmFieldName,
						useCustomMethod,
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
	} else {
		klog.Infof("found existing non-generated mapping function %q, won't generate", goTypeName+"_ToProto")
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

func protoNameForField(protoField protoreflect.FieldDescriptor) string {
	s := strings.Title(protoField.JSONName())
	return s
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
	case "google.protobuf.Int64Value":
		return "direct.Int64Value_FromProto"
	case "google.protobuf.StringValue":
		return "direct.StringValue_FromProto"
	case "google.protobuf.BoolValue":
		return "direct.BoolValue_FromProto"
	case "google.protobuf.FloatValue":
		return "direct.FloatValue_FromProto"
	case "google.protobuf.DoubleValue":
		return "direct.DoubleValue_FromProto"
	case "google.protobuf.Int32Value":
		return "direct.Int32Value_FromProto"
	case "google.protobuf.UInt32Value":
		return "direct.UInt32Value_FromProto"
	case "google.protobuf.UInt64Value":
		return "direct.UInt64Value_FromProto"
	case "google.protobuf.BytesValue":
		return "direct.BytesValue_FromProto"
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
	case "google.protobuf.Int64Value":
		return "direct.Int64Value_ToProto"
	case "google.protobuf.StringValue":
		return "direct.StringValue_ToProto"
	case "google.protobuf.BoolValue":
		return "direct.BoolValue_ToProto"
	case "google.protobuf.FloatValue":
		return "direct.FloatValue_ToProto"
	case "google.protobuf.DoubleValue":
		return "direct.DoubleValue_ToProto"
	case "google.protobuf.Int32Value":
		return "direct.Int32Value_ToProto"
	case "google.protobuf.UInt32Value":
		return "direct.UInt32Value_ToProto"
	case "google.protobuf.UInt64Value":
		return "direct.UInt64Value_ToProto"
	case "google.protobuf.BytesValue":
		return "direct.BytesValue_ToProto"
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

// getKRMImportName generates an import alias for KRM package based on its version
// e.g. "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1" -> "krmv1alpha1"
func getKRMImportName(krmPackage string) string {
	parts := strings.Split(krmPackage, "/")
	version := parts[len(parts)-1]
	return "krm" + version
}
