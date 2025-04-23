// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package typeupdater

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"
)

// messageInfo contains information about a Go struct parsed from existing types files.
// This struct is used to keep track of existing information about a message in the
// generated and human-edited code.
type messageInfo struct {
	GoName    string                // The Go struct name
	ProtoName string                // The proto message name from +kcc:proto annotation
	IsVirtual bool                  // KRM-specific messages that don't map to proto
	Comments  []string              // Original comments
	Fields    map[string]*fieldInfo // Map of field name to field info
	FilePath  string                // The file path where this Go struct was located
}

// fieldInfo contains information about a field in a Go struct parsed from existing types files.
// This struct is used to keep track of existing information about a field in the
// generated and human-edited code.
type fieldInfo struct {
	GoName      string   // Field name in Go
	ProtoName   string   // The fully qualified proto field name from +kcc:proto:field annotation
	IsVirtual   bool     // KRM-specific fields that don't map to proto
	IsIgnored   bool     // Field explicitly marked as not implemented
	IsReference bool     // Is this a reference field?
	RefType     string   // What type of reference (ProjectRef, etc)
	Comments    []string // Preserve original comments for reference fields
}

func extractMessageInfoFromGoFiles(dir string) (map[string]messageInfo, error) {
	messages := make(map[string]messageInfo)

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}

		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		docMap := gocode.NewDocMap(fset, file)

		ast.Inspect(file, func(n ast.Node) bool {
			ts, ok := n.(*ast.TypeSpec)
			if !ok {
				return true
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				return true
			}

			msgInfo := newMessageInfo(ts.Name.Name, path)
			msgInfo.parseComments(ts, docMap)

			// parse fields within the message
			for _, field := range st.Fields.List {
				if len(field.Names) == 0 {
					continue
				}
				fieldInfo := newFieldInfo(field.Names[0].Name)
				fieldInfo.parseComments(field, docMap)
				msgInfo.Fields[fieldInfo.GoName] = fieldInfo
			}

			messages[msgInfo.GoName] = msgInfo
			return true
		})
		return nil
	})

	return messages, err
}

func newMessageInfo(name, filePath string) messageInfo {
	return messageInfo{
		GoName:   name,
		FilePath: filePath,
		Fields:   make(map[string]*fieldInfo),
	}
}

func (info *messageInfo) parseComments(ts *ast.TypeSpec, docMap map[ast.Node]*ast.CommentGroup) {
	info.IsVirtual = true

	if comments := docMap[ts]; comments != nil {
		info.Comments = make([]string, 0, len(comments.List))
		for _, c := range comments.List {
			text := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
			info.Comments = append(info.Comments, text)

			// check for proto annotation
			for _, annotation := range []string{
				codegen.KCCProtoMessageAnnotationMisc,
				codegen.KCCProtoMessageAnnotationSpec,
				codegen.KCCProtoMessageAnnotationObservedState,
			} {
				if strings.HasPrefix(text, annotation+"=") {
					protoName := strings.TrimSpace(strings.TrimPrefix(text, annotation+"="))
					info.ProtoName = protoName
					info.IsVirtual = false
				}
			}
		}
	}
}

func newFieldInfo(name string) *fieldInfo {
	return &fieldInfo{
		GoName: name,
	}
}

func (info *fieldInfo) parseComments(field *ast.Field, docMap map[ast.Node]*ast.CommentGroup) {
	info.IsVirtual = true

	// check if field is a reference field
	if expr, ok := field.Type.(*ast.StarExpr); ok {
		if sel, ok := expr.X.(*ast.SelectorExpr); ok {
			if ident, ok := sel.X.(*ast.Ident); ok {
				if ident.Name == "refv1beta1" { // HACK: this is a hack to identify reference fields
					info.IsReference = true
					info.RefType = sel.Sel.Name
				}
			}
		}
	}

	// parse comments to find kcc codegen annotations
	if comments := docMap[field]; comments != nil {
		info.Comments = make([]string, 0, len(comments.List))
		for _, c := range comments.List {
			text := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
			info.Comments = append(info.Comments, text)

			if strings.HasPrefix(text, codegen.KCCProtoFieldAnnotation+"=") {
				protoName := strings.TrimSpace(strings.TrimPrefix(text, codegen.KCCProtoFieldAnnotation+"="))
				info.ProtoName = protoName
				info.IsVirtual = false
			}
			if strings.Contains(text, "NOTYET") || strings.Contains(text, codegen.KCCProtoIgnoreAnnotation) {
				info.IsIgnored = true
			}
		}
	}
}

// getSpecialAnnotations extracts special annotations like +required from comment group
// These annotations are manually added to the generated code, we need to preserve them.
func getSpecialAnnotations(comments []string) []string {
	if comments == nil {
		return nil
	}

	var annotations []string
	for _, c := range comments {
		if strings.Contains(c, "+genclient") ||
			strings.Contains(c, "+k8s") ||
			strings.Contains(c, "+kubebuilder") ||
			strings.Contains(c, "+required") ||
			strings.Contains(c, "+optional") ||
			strings.Contains(c, "Immutable") {
			annotations = append(annotations, c)
		}
	}
	return annotations
}
