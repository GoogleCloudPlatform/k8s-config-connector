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
	precomputedMappings map[protoreflect.FullName][]*gocode.GoStruct

	generatedFileAnnotation *annotations.FileAnnotation

	multiversion bool

	importedPackages map[string]importedPackage
}

type importedPackage struct {
	alias     string
	goPackage string
}

func NewMapperGenerator(goPathForMessage OutputFunc, outputBaseDir string, generatedFileAnnotation *annotations.FileAnnotation, multiversion bool) *MapperGenerator {
	g := &MapperGenerator{
		goPathForMessage:        goPathForMessage,
		precomputedMappings:     make(map[protoreflect.FullName][]*gocode.GoStruct),
		generatedFileAnnotation: generatedFileAnnotation,
		multiversion:            multiversion,
		importedPackages:        make(map[string]importedPackage),
	}
	g.generatorBase.init(outputBaseDir)
	return g
}

func (v *MapperGenerator) AddGoImportAlias(goPackage string, alias string) string {
	v.importedPackages[goPackage] = importedPackage{
		alias:     alias,
		goPackage: goPackage,
	}
	return alias
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
							v.precomputedMappings[protoName] = append(v.precomputedMappings[protoName], s)
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

func (v *MapperGenerator) findKRMStructsForProto(msg protoreflect.MessageDescriptor) []*gocode.GoStruct {
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
	if len(goTypes) == 0 {
		klog.Infof("no go types found for proto %v", msg.FullName())
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

func (v *MapperGenerator) GenerateMappers(goImports map[string]string) error {
	sort.Slice(v.typePairs, func(i, j int) bool {
		if v.typePairs[i].KRMType.Name == v.typePairs[j].KRMType.Name {
			return v.typePairs[i].KRMType.GoPackage < v.typePairs[j].KRMType.GoPackage
		}
		return v.typePairs[i].KRMType.Name < v.typePairs[j].KRMType.Name
	})

	for alias, pkgName := range goImports {
		v.importedPackages[pkgName] = importedPackage{
			alias:     alias,
			goPackage: pkgName,
		}
	}

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
		out.goPackage = lastGoComponent(goPackage)

		out.fileAnnotation = v.generatedFileAnnotation

		{
			out.addImport("refs", "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1")
			out.addImport("", "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct")
		}

		v.writeMapFunctionsForPair(&out.body, out.OutputDir(), &pair)

		for _, importedPackage := range v.importedPackages {
			goPackage := importedPackage.goPackage
			out.addImport(importedPackage.alias, goPackage)
		}
	}

	return nil
}
func (v *MapperGenerator) writeMapFunctionsForPair(out io.Writer, srcDir string, pair *typePair) {
	klog.V(2).InfoS("writeMapFunctionsForPair", "pair.Proto.FullName", pair.Proto.FullName(), "pair.KRMType.Name", pair.KRMType.Name)
	msg := pair.Proto
	pbTypeName := protoNameForType(msg)
	pbTypeGoImport := v.goPackageForProto(msg.ParentFile())

	goType := pair.KRMType
	goTypeName := goType.Name

	goFields := make(map[string]*gocode.StructField)
	for _, f := range goType.Fields {
		goFields[f.Name] = f
	}

	krmImportName := v.getGoImportAlias(pair.KRMType.GoPackage)

	versionSpecifier := ""
	if v.multiversion {
		versionSpecifier = "_" + lastGoComponent(pair.KRMType.GoPackage)
	}

	if v.findFuncDeclaration(goTypeName+versionSpecifier+"_FromProto", srcDir, true) == nil {
		fmt.Fprintf(out, "func %s_FromProto(mapCtx *direct.MapContext, in *%s.%s) *%s.%s {\n", goTypeName+versionSpecifier, pbTypeGoImport, pbTypeName, krmImportName, goTypeName)
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
				if krmFieldRef := goFields[krmFieldName+"Ref"]; krmFieldRef != nil {
					fmt.Fprintf(out, "\tif in.%s != \"\" {\n", protoAccessor)
					qualifiedTypeName := strings.TrimPrefix(krmFieldRef.Type, "*")
					tokens := strings.SplitN(qualifiedTypeName, ".", 2)
					if len(tokens) > 1 {
						alias := v.getGoImportAlias(krmFieldRef.GoPackage)
						qualifiedTypeName = alias + "." + tokens[1]
					} else if len(tokens) == 1 {
						// In same package
						alias := v.getGoImportAlias(pair.KRMType.GoPackage)
						qualifiedTypeName = alias + "." + tokens[0]
					}
					fmt.Fprintf(out, "\t\tout.%v = &%v{External: in.%v}\n", krmFieldRef.Name, qualifiedTypeName, protoAccessor)
					fmt.Fprintf(out, "\t}\n")
					continue
				}

				if krmFieldRefs := goFields[strings.TrimSuffix(krmFieldName, "s")+"Refs"]; krmFieldRefs != nil {
					template := `
						if v := in.{protoAccessor}; len(v) != 0 {
							for i := range v {
								out.{KRMField} = append(out.{KRMField}, &{RefType}{External: v[i]})
							}
						}
					`

					qualifiedTypeName := strings.TrimPrefix(krmFieldRefs.Type, "*")
					tokens := strings.SplitN(qualifiedTypeName, ".", 2)
					if len(tokens) > 1 {
						alias := v.getGoImportAlias(krmFieldRefs.GoPackage)
						klog.Infof("getGetImportAlias(%q) => %q", krmFieldRefs.GoPackage, alias)
						qualifiedTypeName = alias + "." + tokens[1]
					} else if len(tokens) == 1 {
						// In same package
						alias := v.getGoImportAlias(pair.KRMType.GoPackage)
						qualifiedTypeName = alias + "." + tokens[0]
					}

					s := template
					s = strings.ReplaceAll(s, "{protoAccessor}", protoAccessor)
					s = strings.ReplaceAll(s, "{KRMField}", krmFieldRefs.Name)
					s = strings.ReplaceAll(s, "{RefType}", qualifiedTypeName)

					fmt.Fprintf(out, "%s\n", s)

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

			isKRMFieldSlice := strings.HasPrefix(krmField.Type, "[]")
			isProtoFieldSlice := protoField.Cardinality() == protoreflect.Repeated

			if isProtoFieldSlice && !isKRMFieldSlice && !protoField.IsMap() { // proto slice -> krm single
				var fromProtoElemFunc string
				switch protoField.Kind() {
				case protoreflect.MessageKind:
					krmElemTypeName := strings.TrimPrefix(krmField.Type, "*")
					fromProtoElemFunc = krmElemTypeName + versionSpecifier + "_FromProto"
					if _, ok := protoMessagesNotMappedToGoStruct[string(protoField.Message().FullName())]; ok {
						fromProtoElemFunc = krmFromProtoFunctionName(protoField, krmField.Name)
					}
				case protoreflect.EnumKind:
					fromProtoElemFunc = "direct. "
				}

				fmt.Fprintf(out, "\tif len(in.%s) > 0 {\n", protoAccessor)
				if fromProtoElemFunc != "" {
					fmt.Fprintf(out, "\t\tout.%s = %s(mapCtx, in.%s[0])\n", krmFieldName, fromProtoElemFunc, protoAccessor)
				} else {
					if protoField.Kind() == protoreflect.BytesKind {
						fmt.Fprintf(out, "\t\tout.%s = in.%s[0]\n", krmFieldName, protoAccessor)
					} else {
						fmt.Fprintf(out, "\t\tout.%s = direct.LazyPtr(in.%s[0])\n", krmFieldName, protoAccessor)
					}
				}
				fmt.Fprintf(out, "\t}\n")
				continue
			}

			if !isProtoFieldSlice && isKRMFieldSlice { // proto single -> krm slice
				krmSliceElemType := strings.TrimPrefix(krmField.Type, "[]")

				switch protoField.Kind() {
				case protoreflect.MessageKind:
					krmElemTypeName := strings.TrimPrefix(krmSliceElemType, "*")
					functionName := krmElemTypeName + versionSpecifier + "_FromProto"
					if _, ok := protoMessagesNotMappedToGoStruct[string(protoField.Message().FullName())]; ok {
						functionName = krmFromProtoFunctionName(protoField, krmField.Name)
					}
					fmt.Fprintf(out, "\tif v := in.%s; v != nil {\n", protoAccessor)
					if strings.HasPrefix(krmSliceElemType, "*") {
						fmt.Fprintf(out, "\t\tout.%s = []*%s.%s{%s(mapCtx, v)}\n", krmFieldName, krmImportName, strings.TrimPrefix(krmSliceElemType, "*"), functionName)
					} else {
						fmt.Fprintf(out, "\t\tout.%s = []%s.%s{%s(mapCtx, v)}\n", krmFieldName, krmImportName, krmSliceElemType, functionName)
					}
					fmt.Fprintf(out, "\t}\n")

				case protoreflect.EnumKind:
					functionName := "direct.Enum_FromProto"
					fmt.Fprintf(out, "\tout.%s = []%s.%s{%s(mapCtx, in.%s)}\n", krmFieldName, krmImportName, krmSliceElemType, functionName, protoAccessor)

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
					if protoIsPointerInGo(protoField) {
						fmt.Fprintf(out, "\tif v := in.%s; v != nil {\n", protoFieldName)
						fmt.Fprintf(out, "\t\tout.%s = []%s.%s{v}\n", krmFieldName, krmImportName, krmSliceElemType)
						fmt.Fprintf(out, "\t}\n")
					} else {
						fmt.Fprintf(out, "\tout.%s = []%s.%s{direct.LazyPtr(in.%s)}\n", krmFieldName, krmImportName, krmSliceElemType, protoAccessor)
					}

				default:
					klog.Fatalf("unhandled kind %q for field %v", protoField.Kind(), protoField)
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

					functionName := krmElemTypeName + versionSpecifier + "_FromProto"
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

				functionName := krmTypeName + versionSpecifier + "_FromProto"
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
				protoreflect.Fixed64Kind,
				protoreflect.BytesKind:
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

			default:
				klog.Fatalf("unhandled kind %q for field %v", protoField.Kind(), protoField)
			}
		}
		fmt.Fprintf(out, "\treturn out\n")
		fmt.Fprintf(out, "}\n")
	} else {
		klog.Infof("found existing non-generated mapping function %q, won't generate", goTypeName+"_FromProto")
	}

	if v.findFuncDeclaration(goTypeName+versionSpecifier+"_ToProto", srcDir, true) == nil {
		fmt.Fprintf(out, "func %s_ToProto(mapCtx *direct.MapContext, in *%s.%s) *%s.%s {\n", goTypeName+versionSpecifier, krmImportName, goTypeName, pbTypeGoImport, pbTypeName)
		fmt.Fprintf(out, "\tif in == nil {\n")
		fmt.Fprintf(out, "\t\treturn nil\n")
		fmt.Fprintf(out, "\t}\n")
		fmt.Fprintf(out, "\tout := &%s.%s{}\n", pbTypeGoImport, pbTypeName)
		for i := 0; i < msg.Fields().Len(); i++ {
			protoField := msg.Fields().Get(i)
			protoFieldName := protoNameForField(protoField)
			protoFieldPackage := v.goPackageForProto(protoField.ParentFile())

			krmFieldName := goFieldName(protoField)
			krmField := goFields[krmFieldName]
			if krmField == nil {
				// Support refs
				if krmFieldRef := goFields[krmFieldName+"Ref"]; krmFieldRef != nil {
					fmt.Fprintf(out, "\tif in.%s != nil {\n", krmFieldRef.Name)
					fmt.Fprintf(out, "\t\tout.%s = in.%s.External\n", protoFieldName, krmFieldRef.Name)
					fmt.Fprintf(out, "\t}\n")
					continue
				}

				if krmFieldRefs := goFields[strings.TrimSuffix(krmFieldName, "s")+"Refs"]; krmFieldRefs != nil {
					template := `
						if v := in.{KRMField}; len(v) != 0 {
							for i := range v {
								out.{protoFieldName} = append(out.{protoFieldName}, v[i].External)
							}
						}
						`

					s := template
					s = strings.ReplaceAll(s, "{protoFieldName}", protoFieldName)
					s = strings.ReplaceAll(s, "{KRMField}", krmFieldRefs.Name)

					fmt.Fprintf(out, "%s\n", s)

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

			isKRMFieldSlice := strings.HasPrefix(krmField.Type, "[]")
			isProtoFieldSlice := protoField.Cardinality() == protoreflect.Repeated

			if !isProtoFieldSlice && isKRMFieldSlice { // proto single <- krm slice
				krmElemType := strings.TrimPrefix(krmField.Type, "[]")
				krmElemTypeName := strings.TrimPrefix(krmElemType, "*")

				fmt.Fprintf(out, "\tif len(in.%s) > 0 && in.%s[0] != nil {\n", krmFieldName, krmFieldName)

				switch protoField.Kind() {
				case protoreflect.MessageKind:
					functionName := krmElemTypeName + versionSpecifier + "_ToProto"
					if _, ok := protoMessagesNotMappedToGoStruct[string(protoField.Message().FullName())]; ok {
						functionName = krmToProtoFunctionName(protoField, krmField.Name)
					}
					fmt.Fprintf(out, "\t\tout.%s = %s(mapCtx, in.%s[0])\n", protoFieldName, functionName, krmFieldName)

				case protoreflect.EnumKind:
					protoTypeName := v.goPackageForProto(protoField.Enum().ParentFile()) + "." + protoNameForEnum(protoField.Enum())
					functionName := "direct.Enum_ToProto"
					fmt.Fprintf(out, "\t\tout.%s = %s[%s](mapCtx, *in.%s[0])\n", protoFieldName, functionName, protoTypeName, krmFieldName)

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
					if protoIsPointerInGo(protoField) {
						fmt.Fprintf(out, "\t\tout.%s = in.%s[0]\n", protoFieldName, krmFieldName)
					} else {
						fmt.Fprintf(out, "\t\tout.%s = direct.ValueOf(in.%s[0])\n", protoFieldName, krmFieldName)
					}

				default:
					klog.Fatalf("unhandled kind %q for field %v", protoField.Kind(), protoField)
				}
				fmt.Fprintf(out, "\t}\n")
				continue
			}

			if isProtoFieldSlice && !isKRMFieldSlice && !protoField.IsMap() { // proto slice <- krm single
				var toProtoElemFunc string
				krmElemTypeName := strings.TrimPrefix(krmField.Type, "*")

				var protoSliceElemType string

				switch protoField.Kind() {
				case protoreflect.MessageKind:
					functionName := krmElemTypeName + versionSpecifier + "_ToProto"
					if _, ok := protoMessagesNotMappedToGoStruct[string(protoField.Message().FullName())]; ok {
						functionName = krmToProtoFunctionName(protoField, krmField.Name)
					}
					toProtoElemFunc = functionName
					protoSliceElemType = "*" + v.goPackageForProto(protoField.Message().ParentFile()) + "." + protoNameForType(protoField.Message())

				case protoreflect.EnumKind:
					protoTypeName := v.goPackageForProto(protoField.Enum().ParentFile()) + "." + protoNameForEnum(protoField.Enum())
					toProtoElemFunc = fmt.Sprintf("direct.Enum_ToProto[%s]", protoTypeName)
					protoSliceElemType = v.goPackageForProto(protoField.Message().ParentFile()) + "." + protoNameForEnum(protoField.Enum())

				default:
					protoSliceElemType = goTypeForProtoKind(protoField.Kind())
				}

				fmt.Fprintf(out, "\tif v := in.%s; v != nil {\n", krmFieldName)
				if toProtoElemFunc != "" {
					if protoField.Kind() == protoreflect.EnumKind {
						fmt.Fprintf(out, "\t\tout.%s = []%s{%s(mapCtx, *v)}\n", protoFieldName, protoSliceElemType, toProtoElemFunc)
					} else {
						fmt.Fprintf(out, "\t\tout.%s = []%s{%s(mapCtx, v)}\n", protoFieldName, protoSliceElemType, toProtoElemFunc)
					}
				} else {
					if protoField.Kind() == protoreflect.BytesKind {
						fmt.Fprintf(out, "\t\tout.%s = [][]byte{in.%s}\n", protoFieldName, krmFieldName)
					} else {
						fmt.Fprintf(out, "\t\tout.%s = []%s{direct.ValueOf(in.%s)}\n", protoFieldName, protoSliceElemType, krmFieldName)
					}
				}
				fmt.Fprintf(out, "\t}\n")
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

					functionName := krmElemTypeName + versionSpecifier + "_ToProto"
					useSliceToProtoFunction = functionName

				case protoreflect.StringKind:
					if krmField.Type != "[]string" {
						useCustomMethod = fmt.Sprintf("%s_%s_ToProto", goTypeName, protoFieldName)
					}

				case protoreflect.EnumKind:
					krmElemTypeName := krmField.Type
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "*")
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "[]")

					protoTypeName := v.goPackageForProto(protoField.Enum().ParentFile()) + "." + protoNameForEnum(protoField.Enum())
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

				functionName := krmTypeName + versionSpecifier + "_ToProto"
				switch krmTypeName {
				case "string":
					functionName = string(msg.Name()) + "_" + krmFieldName + "_ToProto"
				}

				// special handling for proto messages that mapped to KRM string
				if _, ok := protoMessagesNotMappedToGoStruct[string(protoField.Message().FullName())]; ok {
					functionName = krmToProtoFunctionName(protoField, krmField.Name)
				}

				oneof := protoField.ContainingOneof()
				if oneof != nil && !protoField.HasOptionalKeyword() {
					fmt.Fprintf(out, "\tif oneof := %s(mapCtx, in.%s); oneof != nil {\n",
						functionName,
						krmFieldName,
					)

					oneofFieldName := ToGoFieldName(oneof.Name())

					oneofTypeName := protoNameForOneOf(protoField)

					fmt.Fprintf(out, "\t\tout.%s = &%s.%s{%s: oneof}\n",
						oneofFieldName,
						protoFieldPackage,
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
				protoTypeName := v.goPackageForProto(protoField.Enum().ParentFile()) + "." + protoNameForEnum(protoField.Enum())
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

	// Generate ToProto helpers for oneof fields that are not messages
	for i := 0; i < msg.Fields().Len(); i++ {
		protoField := msg.Fields().Get(i)
		oneof := protoField.ContainingOneof()
		if oneof == nil || protoField.HasOptionalKeyword() || protoField.Kind() == protoreflect.MessageKind {
			continue
		}

		protoFieldName := protoNameForField(protoField)
		functionName := fmt.Sprintf("%s_%s_ToProto", goTypeName, protoFieldName)

		if v.findFuncDeclaration(functionName, srcDir, true) != nil {
			continue
		}

		krmFieldName := goFieldName(protoField)
		krmField, ok := goFields[krmFieldName]
		if !ok {
			// This can happen if the field is not in the KRM struct (e.g. output-only).
			// We should not generate a helper if there is no corresponding KRM field.
			continue
		}
		krmFieldType := krmField.Type

		oneofWrapperTypeName := v.goPackageForProto(protoField.ParentFile()) + "." + protoNameForOneOf(protoField)

		fmt.Fprintf(out, "func %s(mapCtx *direct.MapContext, in %s) *%s {\n", functionName, krmFieldType, oneofWrapperTypeName)
		fmt.Fprintf(out, "\tif in == nil {\n")
		fmt.Fprintf(out, "\t\treturn nil\n")
		fmt.Fprintf(out, "\t}\n")

		switch protoField.Kind() {
		case protoreflect.EnumKind:
			protoEnumTypeName := v.goPackageForProto(protoField.Enum().ParentFile()) + "." + protoNameForEnum(protoField.Enum())
			fmt.Fprintf(out, "\treturn &%s{%s: direct.Enum_ToProto[%s](mapCtx, in)}\n", oneofWrapperTypeName, protoFieldName, protoEnumTypeName)

		case protoreflect.BoolKind:
			fmt.Fprintf(out, "\treturn &%s{%s: *in}\n", oneofWrapperTypeName, protoFieldName)

		case protoreflect.StringKind,
			protoreflect.FloatKind,
			protoreflect.DoubleKind,
			protoreflect.Int64Kind,
			protoreflect.Int32Kind,
			protoreflect.Uint32Kind,
			protoreflect.Uint64Kind,
			protoreflect.Fixed64Kind,
			protoreflect.BytesKind:

			fmt.Fprintf(out, "\treturn &%s{%s: *in}\n", oneofWrapperTypeName, protoFieldName)
		}
		fmt.Fprintf(out, "}\n")
	}
}

func protoNameForType(msg protoreflect.MessageDescriptor) string {
	fullName := string(msg.FullName())
	namespace := string(msg.ParentFile().Package())
	if namespace != "" {
		namespace = namespace + "."
	}
	fullName = strings.TrimPrefix(fullName, namespace)
	fullName = strings.ReplaceAll(fullName, ".", "_")
	return fullName
}

func protoNameForEnum(msg protoreflect.EnumDescriptor) string {
	fullName := string(msg.FullName())
	namespace := string(msg.ParentFile().Package())
	if namespace != "" {
		namespace = namespace + "."
	}
	fullName = strings.TrimPrefix(fullName, namespace)
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
		return "direct.Struct_FromProto"
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
		return "direct.Struct_ToProto"
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

// getGoImportAlias generates an import alias for go package
func (o *MapperGenerator) getGoImportAlias(goPackage string) string {
	existing, found := o.importedPackages[goPackage]
	if found {
		return existing.alias
	}

	importAlias := lastComponent(goPackage)

	// Disambiguate in a way that preserves compatibility with the existing code
	if strings.Contains(goPackage, "k8s-config-connector/apis/refs/") {
		importAlias = "refs" + importAlias
	} else if _, suffix, found := strings.Cut(goPackage, "k8s-config-connector/apis/"); found {
		tokens := strings.Split(suffix, "/")
		if len(tokens) == 2 {
			importAlias = "krm" + tokens[0] + tokens[1]
		} else {
			klog.Fatalf("cannot determine import alias for %s", goPackage)
		}
	}

	o.importedPackages[goPackage] = importedPackage{
		alias:     importAlias,
		goPackage: goPackage,
	}
	return importAlias
}

// GoPackageForProto returns the full go package import path for a proto message
// For the import alias, see goPackageForProto
func GoPackageForProto(parentFile protoreflect.FileDescriptor) string {
	fileOptions := parentFile.Options().(*descriptorpb.FileOptions)
	protoGoPackage := fileOptions.GetGoPackage()
	if ix := strings.Index(protoGoPackage, ";"); ix != -1 {
		protoGoPackage = protoGoPackage[:ix]
	}
	return protoGoPackage
}

// goPackageForProto returns the go package import alias for a proto message
func (o *MapperGenerator) goPackageForProto(parentFile protoreflect.FileDescriptor) string {
	protoGoPackage := GoPackageForProto(parentFile)
	existing, found := o.importedPackages[protoGoPackage]
	if found {
		return existing.alias
	}

	importAlias := strings.TrimSuffix(lastComponent(protoGoPackage), "pb") + "pb"

	o.AddGoImportAlias(protoGoPackage, importAlias)
	return importAlias
}

func lastComponent(s string) string {
	ix := strings.LastIndex(s, "/")
	return s[ix+1:]
}
