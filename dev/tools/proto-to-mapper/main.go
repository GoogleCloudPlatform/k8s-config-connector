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

package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/proto-to-mapper/gocode"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"k8s.io/klog/v2"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

type GenerateOptions struct {
	OutDir       string
	APIDir       string
	APIGoPackage string
}

func run(ctx context.Context) error {
	var opt GenerateOptions
	opt.OutDir = "generated/"
	inputPath := "build/googleapis.pb"
	opt.APIDir = ""
	flag.StringVar(&opt.OutDir, "out", opt.OutDir, "output directory")
	flag.StringVar(&opt.APIDir, "apis", opt.APIDir, "API directory")
	flag.StringVar(&opt.APIGoPackage, "api-packages", opt.APIGoPackage, "API go package")

	flag.Parse()

	if opt.APIDir == "" {
		return fmt.Errorf("missing --apis flag")
	}
	if opt.APIGoPackage == "" {
		return fmt.Errorf("missing --api-packages flag")
	}

	b, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("reading %q: %w", inputPath, err)
	}

	fds := &descriptorpb.FileDescriptorSet{}
	if err := proto.Unmarshal(b, fds); err != nil {
		return fmt.Errorf("unmarshalling %q: %w", inputPath, err)
	}

	files, err := protodesc.NewFiles(fds)
	if err != nil {
		return fmt.Errorf("building file description: %w", err)
	}

	packages, err := gocode.LoadPackageTree(opt.APIGoPackage, opt.APIDir)
	if err != nil {
		return fmt.Errorf("inspecting go code: %w", err)
	}

	v := &visitor{}
	v.goPackages = make(map[string]*gocode.Package)
	for _, pkg := range packages {
		annotation := pkg.GetAnnotation("+kcc:proto")
		// klog.Infof("got package %v for proto %v", pkg.SourceDir, pkg.Comments)
		if annotation != "" {
			klog.Infof("got package %v for proto %v", pkg.SourceDir, annotation)
			v.goPackages[annotation] = pkg
		}
	}

	v.generatedFiles = make(map[generatedFileKey]*generatedFile)

	var sortedFiles []protoreflect.FileDescriptor
	files.RangeFiles(func(f protoreflect.FileDescriptor) bool {
		sortedFiles = append(sortedFiles, f)
		return true
	})
	sort.Slice(sortedFiles, func(i, j int) bool {
		return sortedFiles[i].FullName() < sortedFiles[j].FullName()
	})

	for _, f := range sortedFiles {
		v.visitFile(f)
	}

	v.writeMapFunctions()

	for _, f := range v.generatedFiles {
		if err := f.Write(opt.OutDir); err != nil {
			return err
		}
	}

	return nil
}

type visitor struct {
	goPackages     map[string]*gocode.Package
	generatedFiles map[generatedFileKey]*generatedFile

	typePairs []typePair
}

type typePair struct {
	ProtoPackage protoreflect.FullName
	KRMType      *gocode.GoStruct
	Proto        protoreflect.MessageDescriptor
}

type generatedFile struct {
	key      generatedFileKey
	contents bytes.Buffer
}

type generatedFileKey struct {
	GoPackage string

	File string
}

func (f *generatedFile) Write(baseDir string) error {
	fullName := f.key.GoPackage
	tokens := strings.Split(fullName, ".")
	dirTokens := []string{baseDir}
	dirTokens = append(dirTokens, tokens...)
	dir := filepath.Join(dirTokens...)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating directory %q: %w", dir, err)
	}

	p := filepath.Join(dir, f.key.File)
	if err := os.WriteFile(p, f.contents.Bytes(), 0644); err != nil {
		return fmt.Errorf("writing %q: %w", p, err)
	}

	return nil
}

// func (v *visitor) krmNameForType(msg protoreflect.MessageDescriptor) string {
// 	goType := v.findKRMStructForProto(msg)
// 	if goType == nil {
// 		// TODO: Deprecate this method - we need to handle the not-mapped case
// 		return protoNameForType(msg)
// 	}
// 	goTypeName := goType.Name
// 	return goTypeName
// }

func splitProtoName(msg protoreflect.MessageDescriptor) (packageName string, messageName string) {
	// TODO: Use ParentFile instead?
	var pos protoreflect.Descriptor = msg
	for {
		switch v := pos.(type) {
		case protoreflect.MessageDescriptor:
			parentMsg, ok := v.Parent().(protoreflect.MessageDescriptor)
			if ok {
				pos = parentMsg
			} else {
				packageName := string(v.Parent().FullName())
				messageName := string(msg.FullName())
				messageName = strings.TrimPrefix(messageName, packageName+".")
				return packageName, messageName
			}
		default:
			klog.Fatalf("unhandled proto message type %T %+v", pos, pos)
		}
	}
}
func (v *visitor) findKRMStructsForProto(msg protoreflect.MessageDescriptor) map[string]*gocode.GoStruct {
	// if strings.HasPrefix(string(msg.FullName()), "google.monitoring.dashboard.v1.") {
	// 	klog.Infof("FOUND %q", msg.FullName())
	// }

	matches := make(map[string]*gocode.GoStruct)

	// protoPackage, _ := splitProtoName(msg)
	// klog.Infof("protoPackage %q for %q", protoPackage, msg.FullName())
	// goPackage := v.goPackages[protoPackage]
	// if goPackage == nil {
	// 	// klog.Infof("no go package for %q for %q", protoPackage, msg.FullName())
	// 	return nil
	// }

	// goType := "Dashboard" + protoNameForType(msg)

	// TODO: Precompute this!
	for _, goPackage := range v.goPackages {
		for _, s := range goPackage.Structs {
			// klog.Infof("PACKAGE %s => %+v", s.Name, s.Comments)
			if len(s.Comments) != 0 {
				for _, c := range s.Comments {
					for _, line := range strings.Split(c, "\n") {
						line = strings.TrimSpace(line)
						if strings.HasPrefix(line, "+kcc:proto=") {
							// klog.Infof("+kcc:proto= %+v", line)
							if line == "+kcc:proto="+string(msg.Name()) {
								matches[s.Name] = s
							}
							if line == "+kcc:proto="+string(msg.FullName()) {
								matches[s.Name] = s
							}
						}
					}
				}
			}
			// if s.Name == goType {
			// 	matches[s.Name] = s
			// }
		}
	}
	if len(matches) == 0 {
		klog.Infof("did not find mapping for %q", msg.FullName())
	}
	return matches
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
	// fullName := string(msg.FullName())
	// fullName = strings.TrimPrefix(fullName, string(msg.ParentFile().FullName()))
	// fullName = strings.TrimPrefix(fullName, ".")
	// fullName = strings.ReplaceAll(fullName, ".", "_")
	oneofKey := ToGoFieldName(field.Name())
	name := protoNameForType(msg) + "_" + oneofKey

	// Special case: check for a collision
	if field.Message() != nil {
		elemTypeName := protoNameForType(field.Message())
		if name == elemTypeName {
			name += "_"
		}
	}
	// name += "_" + elemTypeName
	// 	parentFile := msg.ParentFile()
	// 	for i := 0; i < parentFile.Messages().Len(); i++ {
	// 		m := parentFile.Messages().Get(i)
	// 		if string(m.Name()) == name {
	// 			name += "_"
	// 		}
	// 	}
	return name
}

func (v *visitor) writeTypes(out io.Writer, msg protoreflect.MessageDescriptor) {
	// sourceLocations := msg.ParentFile().SourceLocations().ByDescriptor(msg)
	// klog.Infof("sourceLocations: %+v", sourceLocations)

	goType := protoNameForType(msg)
	// pbType := protoNameForType(msg)

	{
		fmt.Fprintf(out, "// +kcc:proto=%s\n", msg.FullName())
		fmt.Fprintf(out, "type %s struct {\n", goType)
		for i := 0; i < msg.Fields().Len(); i++ {
			field := msg.Fields().Get(i)
			sourceLocations := msg.ParentFile().SourceLocations().ByDescriptor(field)

			goFieldName := strings.Title(field.JSONName())
			jsonName := field.JSONName()
			goType := ""

			switch field.Kind() {
			case protoreflect.MessageKind:
				goType = protoNameForType(field.Message())

			case protoreflect.EnumKind:
				goType = "string" //string(field.Enum().Name())

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

			case protoreflect.BoolKind:
				goType = "bool"

			case protoreflect.DoubleKind:
				goType = "float64"

			case protoreflect.FloatKind:
				goType = "float32"

			case protoreflect.BytesKind:
				goType = "[]byte"

			default:
				klog.Fatalf("unhandled kind %q for field %v", field.Kind(), field)
			}

			if field.Cardinality() == protoreflect.Repeated {
				goType = "[]" + goType
			} else {
				goType = "*" + goType
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
		v.writeTypes(out, m)
	}

}

func (v *visitor) writeMapFunctionsForPair(out io.Writer, pair *typePair) {
	// sourceLocations := msg.ParentFile().SourceLocations().ByDescriptor(msg)
	// klog.Infof("sourceLocations: %+v", sourceLocations)

	// goTypeName := v.krmNameForType(msg)
	// pbTypeName := protoNameForType(msg)

	msg := pair.Proto
	pbTypeName := protoNameForType(msg)

	goType := pair.KRMType
	goTypeName := goType.Name

	goFields := make(map[string]*gocode.StructField)
	for _, f := range goType.Fields {
		goFields[f.Name] = f
	}

	{
		fmt.Fprintf(out, "func %s_FromProto(ctx *MapContext, in *pb.%s) *krm.%s {\n", goTypeName, pbTypeName, goTypeName)
		fmt.Fprintf(out, "\tif in == nil {\n")
		fmt.Fprintf(out, "\t\treturn nil\n")
		fmt.Fprintf(out, "\t}\n")
		fmt.Fprintf(out, "\tout := &krm.%s{}\n", goTypeName)
		for i := 0; i < msg.Fields().Len(); i++ {
			protoField := msg.Fields().Get(i)
			protoFieldName := strings.Title(protoField.JSONName())
			protoAccessor := "Get" + protoFieldName + "()"

			krmFieldName := strings.Title(protoField.JSONName())
			krmField := goFields[krmFieldName]
			if krmField == nil {
				fmt.Fprintf(out, "\t// MISSING: %s\n", krmFieldName)
				continue
			}

			if protoField.Cardinality() == protoreflect.Repeated {
				useSliceFromProto := ""
				useCustomMethod := false

				switch protoField.Kind() {
				case protoreflect.MessageKind:
					krmElemTypeName := krmField.Type
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "*")
					krmElemTypeName = strings.TrimPrefix(krmElemTypeName, "[]")

					functionName := krmElemTypeName + "_FromProto"
					useSliceFromProto = functionName
				case protoreflect.StringKind:
					if krmField.Type != "[]string" {
						useCustomMethod = true
						// useSliceFromProto = fmt.Sprintf("%s_%s_FromProto", goTypeName, protoFieldName)
					}

				case
					protoreflect.FloatKind,
					protoreflect.DoubleKind,
					protoreflect.BoolKind,
					protoreflect.Int64Kind,
					protoreflect.Int32Kind,
					protoreflect.Uint32Kind,
					protoreflect.Uint64Kind,
					protoreflect.BytesKind:
					useSliceFromProto = ""
				default:
					klog.Fatalf("unhandled kind %q for repeated field %v", protoField.Kind(), protoField)
				}

				if useSliceFromProto != "" {
					fmt.Fprintf(out, "\tout.%s = Slice_FromProto(ctx, in.%s, %s)\n",
						krmFieldName,
						krmFieldName,
						useSliceFromProto,
					)
				} else if useCustomMethod {
					methodName := fmt.Sprintf("%s_%s_FromProto", goTypeName, protoFieldName)
					fmt.Fprintf(out, "\tout.%s = %s(ctx, in.%s)\n",
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
					// functionName = "String_" + string(protoField.Message().Name()) + "_ToProto"
					functionName = string(msg.Name()) + "_" + krmFieldName + "_FromProto"
				}
				// krmType := v.findKRMStructForProto(protoField.Message())
				// if krmType == nil {
				// 	fmt.Fprintf(out, "\t// MISSING: type for %s\n", protoField.Message().FullName())
				// 	continue
				// }

				fmt.Fprintf(out, "\tout.%s = %s(ctx, in.%s)\n",
					krmFieldName,
					functionName,
					protoAccessor,
				)
			case protoreflect.EnumKind:
				functionName := "Enum_FromProto"
				fmt.Fprintf(out, "\tout.%s = %s(ctx, in.%s)\n",
					krmFieldName,
					functionName,
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
				protoreflect.BytesKind:
				fmt.Fprintf(out, "\tout.%s = LazyPtr(in.%s)\n",
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

	{
		fmt.Fprintf(out, "func %s_ToProto(ctx *MapContext, in *krm.%s) *pb.%s {\n", goTypeName, goTypeName, pbTypeName)
		fmt.Fprintf(out, "\tif in == nil {\n")
		fmt.Fprintf(out, "\t\treturn nil\n")
		fmt.Fprintf(out, "\t}\n")
		fmt.Fprintf(out, "\tout := &pb.%s{}\n", pbTypeName)
		for i := 0; i < msg.Fields().Len(); i++ {
			protoField := msg.Fields().Get(i)
			krmFieldName := strings.Title(protoField.JSONName())
			krmField := goFields[krmFieldName]
			if krmField == nil {
				fmt.Fprintf(out, "\t// MISSING: %s\n", krmFieldName)
				continue
			}

			protoFieldName := strings.Title(protoField.JSONName())

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

					// case protoreflect.EnumKind:
				// 	protoTypeName := "pb." + protoNameForEnum(protoField.Enum())
				// 	functionName := "Enum_ToProto"
				// 	fmt.Fprintf(out, "\tout.%s = %s[%s](ctx, in.%s)\n",
				// 		krmFieldName,
				// 		functionName,
				// 		protoTypeName,
				// 		goFieldName,
				// 	)
				case protoreflect.StringKind:
					if krmField.Type != "[]string" {
						useCustomMethod = true
						//useSliceToProtoFunction = fmt.Sprintf("%s_%s_ToProto", goTypeName, protoFieldName)
					}

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

				if useSliceToProtoFunction != "" {
					fmt.Fprintf(out, "\tout.%s = Slice_ToProto(ctx, in.%s, %s)\n",
						protoFieldName,
						krmFieldName,
						useSliceToProtoFunction,
					)
				} else if useCustomMethod {
					methodName := fmt.Sprintf("%s_%s_ToProto", goTypeName, protoFieldName)
					fmt.Fprintf(out, "\tout.%s = %s(ctx, in.%s)\n",
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
				// krmType := v.findKRMStructForProto(protoField.Message())
				// if krmType == nil {
				// 	fmt.Fprintf(out, "\t// MISSING: type for %s\n", protoField.Message().FullName())
				// 	continue
				// }

				krmTypeName := krmField.Type
				krmTypeName = strings.TrimPrefix(krmTypeName, "*")

				functionName := krmTypeName + "_ToProto"
				switch krmTypeName {
				case "string":
					// functionName = "String_" + string(protoField.Message().Name()) + "_ToProto"
					functionName = string(msg.Name()) + "_" + krmFieldName + "_ToProto"
				}

				oneof := protoField.ContainingOneof()
				if oneof != nil {
					fmt.Fprintf(out, "\tif oneof := %s(ctx, in.%s); oneof != nil {\n",
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
				fmt.Fprintf(out, "\tout.%s = %s(ctx, in.%s)\n",
					protoFieldName,
					functionName,
					krmFieldName,
				)
			case protoreflect.EnumKind:
				protoTypeName := "pb." + protoNameForEnum(protoField.Enum())
				functionName := "Enum_ToProto"
				fmt.Fprintf(out, "\tout.%s = %s[%s](ctx, in.%s)\n",
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
				protoreflect.BytesKind:

				useCustomMethod := false

				switch protoField.Kind() {
				case protoreflect.StringKind:
					if krmField.Type != "*string" {
						useCustomMethod = true
					}
				}

				oneof := protoField.ContainingOneof()
				if oneof != nil {
					functionName := fmt.Sprintf("%s_%s_ToProto", goTypeName, protoFieldName)
					fmt.Fprintf(out, "\tif oneof := %s(ctx, in.%s); oneof != nil {\n",
						functionName,
						krmFieldName,
					)

					oneofFieldName := ToGoFieldName(oneof.Name())

					// oneofTypeName := protoNameForOneOf(protoField)

					fmt.Fprintf(out, "\t\tout.%s = oneof\n",
						oneofFieldName)
					fmt.Fprintf(out, "\t}\n")

					// oneofFieldName := ToGoFieldName(oneof.Name())

					// oneofTypeName := protoNameForOneOf(protoField)

					// fmt.Fprintf(out, "\t\tout.%s = &pb.%s{%s: oneof}\n",
					// 	oneofFieldName,
					// 	oneofTypeName,
					// 	protoFieldName)
					// fmt.Fprintf(out, "\t}\n")
					// fmt.Fprintf(out, "\tif oneof := %s(ctx, in.%s); oneof != nil {\n",
					// 	functionName,
					// 	krmFieldName,
					// )

					// oneofFieldName := ToGoFieldName(oneof.Name())

					// oneofTypeName := protoNameForOneOf(protoField)

					// fmt.Fprintf(out, "\t\tout.%s = %s(ctx, in.%s)\n",
					// 	oneofFieldName,
					// 	functionName,
					// 	krmFieldName)
					// fmt.Fprintf(out, "\t}\n")
				} else if useCustomMethod {
					methodName := fmt.Sprintf("%s_%s_ToProto", goTypeName, protoFieldName)
					fmt.Fprintf(out, "\tout.%s = %s(ctx, in.%s)\n",
						krmFieldName,
						methodName,
						krmFieldName,
					)
				} else {
					fmt.Fprintf(out, "\tout.%s = ValueOf(in.%s)\n",
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

// func (v *visitor) writeTypes(out io.Writer, msg protoreflect.MessageDescriptor) {
// 	// sourceLocations := msg.ParentFile().SourceLocations().ByDescriptor(msg)
// 	// klog.Infof("sourceLocations: %+v", sourceLocations)

// 	goType := v.krmNameForType(msg)
// 	// pbType := protoNameForType(msg)

// 	{
// 		fmt.Fprintf(out, "type %s struct {\n", goType)
// 		for i := 0; i < msg.Fields().Len(); i++ {
// 			field := msg.Fields().Get(i)
// 			sourceLocations := msg.ParentFile().SourceLocations().ByDescriptor(field)

// 			goFieldName := strings.Title(field.JSONName())
// 			jsonName := field.JSONName()
// 			goType := ""

// 			switch field.Kind() {
// 			case protoreflect.MessageKind:
// 				goType = v.krmNameForType(field.Message())

// 			case protoreflect.EnumKind:
// 				goType = string(field.Enum().Name())

// 			case protoreflect.StringKind:
// 				goType = "string"

// 			case protoreflect.Int32Kind:
// 				goType = "int32"

// 			case protoreflect.Int64Kind:
// 				goType = "int64"

// 			case protoreflect.Uint32Kind:
// 				goType = "uint32"

// 			case protoreflect.Uint64Kind:
// 				goType = "uint64"

// 			case protoreflect.BoolKind:
// 				goType = "bool"

// 			case protoreflect.DoubleKind:
// 				goType = "float64"

// 			case protoreflect.FloatKind:
// 				goType = "float32"

// 			case protoreflect.BytesKind:
// 				goType = "[]byte"

// 			default:
// 				klog.Fatalf("unhandled kind %q for field %v", field.Kind(), field)
// 			}

// 			if field.Cardinality() == protoreflect.Repeated {
// 				goType = "[]" + goType
// 			} else {
// 				goType = "*" + goType
// 			}

// 			if sourceLocations.LeadingComments != "" {
// 				comment := strings.TrimSpace(sourceLocations.LeadingComments)
// 				for _, line := range strings.Split(comment, "\n") {
// 					fmt.Fprintf(out, "\t// %s\n", line)
// 				}
// 			}

// 			fmt.Fprintf(out, "\t%s %s `json:\"%s,omitempty\"`\n",
// 				goFieldName,
// 				goType,
// 				jsonName,
// 			)
// 		}
// 		fmt.Fprintf(out, "}\n")
// 	}

// 	for i := 0; i < msg.Messages().Len(); i++ {
// 		m := msg.Messages().Get(i)
// 		v.writeTypes(out, m)
// 	}

// }

func (v *visitor) visitMessage(msg protoreflect.MessageDescriptor) {
	// klog.Infof("visitMessage %s", msg.FullName())
	// goTypeName := v.krmNameForType(msg)
	// pbTypeName := protoNameForType(msg)

	goTypes := v.findKRMStructsForProto(msg)

	if len(goTypes) == 0 {
		klog.Infof("no krm for %v", msg.FullName())
		return
	}
	for _, goType := range goTypes {
		v.typePairs = append(v.typePairs, typePair{
			ProtoPackage: msg.ParentFile().Package(),
			KRMType:      goType,
			Proto:        msg,
		})
	}

	for _, msg := range sorted(msg.Messages()) {
		v.visitMessage(msg)
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

func (v *visitor) visitFile(f protoreflect.FileDescriptor) {
	// namespace := string(f.FullName())

	for _, msg := range sorted(f.Messages()) {
		v.visitMessage(msg)
	}

	{
		protoPackagePath := string(f.Package())
		protoPackagePath = strings.TrimPrefix(protoPackagePath, "google.")
		goPackage := "apis/" + strings.Join(strings.Split(protoPackagePath, "."), "/")
		k := generatedFileKey{
			GoPackage: goPackage,
			File:      "types.go",
		}
		out := v.generatedFiles[k]
		if out == nil {
			out = &generatedFile{key: k}
			v.generatedFiles[k] = out
		}

		for _, msg := range sorted(f.Messages()) {
			v.writeTypes(&out.contents, msg)
		}
	}

}

func (v *visitor) writeMapFunctions() {
	sort.Slice(v.typePairs, func(i, j int) bool {
		return v.typePairs[i].KRMType.Name < v.typePairs[j].KRMType.Name
	})
	for _, pair := range v.typePairs {
		// klog.Infof("pair %+v", pair)
		// namespace := string(pair.ProtoPackage)
		goPackage := pair.KRMType.GoPackage

		k := generatedFileKey{
			GoPackage: goPackage,
			File:      "mapper.go",
		}
		out := v.generatedFiles[k]
		if out == nil {
			out = &generatedFile{key: k}
			v.generatedFiles[k] = out

			out.contents.WriteString(fmt.Sprintf("package %s\n\n", lastGoComponent(goPackage)))
			out.contents.WriteString(fmt.Sprintf("import pb %q\n\n", "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"))
			out.contents.WriteString(fmt.Sprintf("import krm %q\n\n", "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"))
		}

		v.writeMapFunctionsForPair(&out.contents, &pair)
	}
}

func lastGoComponent(goPackage string) string {
	return filepath.Base(goPackage)
	// components := strings.Split(goPackage, ".")
	// return components[len(components)-1]
}

func computeGoNamespace(protoPackage protoreflect.FullName) string {
	// hack while we figure out the real logic; for now skip the last component, expected to be "v1" or similar
	return string(protoPackage.Parent().Parent().Name())
}

func ToGoFieldName(name protoreflect.Name) string {
	tokens := strings.Split(string(name), "_")
	for i, token := range tokens {
		tokens[i] = strings.Title(token)
	}
	return strings.Join(tokens, "")
}
