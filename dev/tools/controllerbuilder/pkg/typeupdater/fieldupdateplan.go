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

package typeupdater

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog"
)

// FieldUpdatePlan represents a planned update to a specific Go field
type FieldUpdatePlan struct {
	filepath string // path to the file containing the field

	structName string     // name of the Go struct containing the field
	fieldName  string     // name of the Go field to update
	fieldInfo  *fieldInfo // original field info for reference

	protoParentName string                       // fully qualified name of the proto parent message
	protoName       string                       // fully qualified name of the proto field
	protoField      protoreflect.FieldDescriptor // proto field descriptor

	content []byte // generated field content
}

func (s *ProtoPackageSyncer) createFieldUpdatePlan(msgInfo messageInfo, fieldInfo *fieldInfo, msgDesc protoreflect.MessageDescriptor) (*FieldUpdatePlan, error) {
	if fieldInfo.IsVirtual {
		if s.opts.LegacyMode {
			// HACK: infer the proto name for legacy go fields without proto name annotation
			fieldInfo.ProtoName = fmt.Sprintf("%s.%s", msgInfo.ProtoName, getProtoFieldName(fieldInfo.GoName))
			klog.Infof("Inferring proto name for legacy field %s in struct %s: %s", fieldInfo.GoName, msgInfo.GoName, fieldInfo.ProtoName)
		} else {
			klog.Infof("Skipping virtual field %s in %s", fieldInfo.GoName, msgInfo.GoName)
			return nil, nil
		}
	}

	// 1. find the proto field
	name := getProtoFieldName(fieldInfo.ProtoName) // e.g. "google.cloud.bigquery.datatransfer.v1.TransferConfig.DisplayName" -> "display_name"
	protoField := msgDesc.Fields().ByName(protoreflect.Name(name))
	if protoField == nil {
		klog.Warningf("proto field %s (full name: %s) not found in message %s", name, fieldInfo.ProtoName, msgInfo.ProtoName)
		return nil, nil
	}

	// 2. generate Go structs for the field
	var buf bytes.Buffer

	// 2.1 special annotations such as "// +required" are manually added to the generated code, we need to preserve them
	specialAnnotations := getSpecialAnnotations(fieldInfo.Comments)
	if len(specialAnnotations) > 0 {
		for _, annotation := range specialAnnotations {
			fmt.Fprintf(&buf, "\t// %s\n", annotation)
		}
	}

	// 2.2 regenerate the field content based on the proto field descriptor
	if fieldInfo.IsReference { // For reference fields, preserve original comments and reference type
		return nil, nil // skip generating reference fields for now since we don't plan to update them
		/* for _, comment := range fieldInfo.Comments {
		       fmt.Fprintf(&buf, "\t// %s\n", comment)
		   }
		   jsonName := codegen.GetJSONForKRM(protoField)
		   fmt.Fprintf(&buf, "\t%s *refv1beta1.%s `json:\"%s,omitempty\"`\n",
		       fieldInfo.GoName,
		       fieldInfo.RefType,
		       jsonName) */
	} else if fieldInfo.IsIgnored { // for ignored fields, generate only the field declaration without comments
		goType, err := codegen.GoTypeForField(protoField, false) // TODO: add support for transitive output fields
		if err != nil {
			return nil, fmt.Errorf("determining Go type for ignored field %s (proto: %s): %w", fieldInfo.GoName, fieldInfo.ProtoName, err)
		}
		jsonName := codegen.GetJSONForKRM(protoField)
		fmt.Fprintf(&buf, "\t%s %s `json:\"%s,omitempty\"`\n",
			fieldInfo.GoName,
			goType,
			jsonName)
	} else { // for regular fields, generate complete field with comments
		codegen.WriteField(&buf, protoField, msgDesc, 0, false) // HACK: use fieldIndex=0 to avoid generating a leading blank line on comments
		// TODO: add support for transitive output fields when writing the field
	}

	// 3. create the update plan to record every information we need to update the field
	plan := &FieldUpdatePlan{
		filepath:        msgInfo.FilePath,
		structName:      msgInfo.GoName,
		fieldName:       fieldInfo.GoName,
		fieldInfo:       fieldInfo,
		protoParentName: msgInfo.ProtoName,
		protoName:       fieldInfo.ProtoName,
		protoField:      protoField,
		content:         buf.Bytes(),
	}

	return plan, nil
}

func (s *ProtoPackageSyncer) applyFieldUpdatePlan(plan FieldUpdatePlan) error {
	content, err := os.ReadFile(plan.filepath)
	if err != nil {
		return fmt.Errorf("reading file %s: %w", plan.filepath, err)
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, plan.filepath, content, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("parsing file %s: %w", plan.filepath, err)
	}

	docMap := gocode.NewDocMap(fset, file)

	// find the target struct and field by matching the proto name
	targetMessageAnnotation := fmt.Sprintf("%s=%s", codegen.KCCProtoMessageAnnotationMisc, plan.protoParentName)
	targetFieldAnnotation := fmt.Sprintf("%s=%s", codegen.KCCProtoFieldAnnotation, plan.protoName)
	var fieldNode *ast.Field
	var found bool
	ast.Inspect(file, func(n ast.Node) bool {
		if found {
			return false
		}

		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}
		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			return false
		}
		if !commentContains(docMap[ts], targetMessageAnnotation) { // match by fully qualified proto name annotation
			return true
		}

		// find the target field
		for _, field := range st.Fields.List {
			fieldComments := docMap[field]
			if commentContains(fieldComments, targetFieldAnnotation) ||
				(s.opts.LegacyMode && len(field.Names) > 0 && field.Names[0].Name == plan.fieldName) { // HACK: match the field name for legacy Go fields without proper proto name annotation
				fieldNode = field
				found = true
				return false
			}
		}
		return true
	})

	if !found {
		return fmt.Errorf("field %s not found in struct %s", plan.fieldName, plan.structName)
	}

	// get the start position (accounting for doc comments if they exist)
	var startPos token.Pos
	var hasDoc bool
	if doc := docMap[fieldNode]; doc != nil {
		startPos = doc.Pos()
		hasDoc = true
	} else {
		startPos = fieldNode.Pos()
	}
	start := fset.Position(startPos)
	end := fset.Position(fieldNode.End())

	if hasDoc { // HACK: remove the leading tab ("\t") from the original field content
		start.Offset--
	}

	// replace the field content
	newContent := make([]byte, 0, len(content)+len(plan.content))
	newContent = append(newContent, content[:start.Offset]...)
	newContent = append(newContent, plan.content...)
	newContent = append(newContent, content[end.Offset:]...)

	if err := os.WriteFile(plan.filepath, newContent, 0644); err != nil {
		return fmt.Errorf("writing file %s: %w", plan.filepath, err)
	}

	return nil
}

func printUpdatePlans(plans []FieldUpdatePlan) {
	klog.Infof("Field update plans:")
	for _, plan := range plans {
		klog.Infof("- File: %s", plan.filepath)
		klog.Infof("  Struct: %s", plan.structName)
		klog.Infof("  Field: %s", plan.fieldName)
		klog.Infof("  Proto: %s", plan.protoName)
		klog.Infof("  IsReference: %v", plan.fieldInfo.IsReference)
		klog.Infof("  IsIgnored: %v", plan.fieldInfo.IsIgnored)
		klog.Infof("  Content: %s", string(plan.content))
	}
}
