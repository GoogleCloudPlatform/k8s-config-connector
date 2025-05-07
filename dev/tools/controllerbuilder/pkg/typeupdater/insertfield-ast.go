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
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"k8s.io/klog/v2"
)

type target struct {
	goName string
	endPos int
}

func IsFieldBehavior(field protoreflect.FieldDescriptor, fieldBehavior annotations.FieldBehavior) bool {
	d := field.Options()
	fieldBehaviors := proto.GetExtension(d, annotations.E_FieldBehavior).([]annotations.FieldBehavior)
	for _, f := range fieldBehaviors {
		if f == fieldBehavior {
			return true
		}
	}
	return false
}

func (u *FieldInserter) insertGoField() error {
	klog.Infof("inserting the generated Go code for field %s", u.newField.proto.Name())

	targetComment := fmt.Sprintf("+kcc:proto=%s", u.newField.parent.FullName())
	outputOnly := IsFieldBehavior(u.newField.proto, annotations.FieldBehavior_OUTPUT_ONLY)

	filepath.WalkDir(u.opts.APIDirectory, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}

		// read the file bytes
		srcBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// parse the file
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		// use a CommentMap to associate comments with nodes
		docMap := gocode.NewDocMap(fset, file)

		// find the target Go struct and the ending position in the source
		// there are 2 cases considered.
		// - case 1, there is only 1 matching target.
		// - case 2, there are two matching targets (Spec and ObservedState).
		var targets []target
		ast.Inspect(file, func(n ast.Node) bool {
			ts, ok := n.(*ast.TypeSpec)
			if !ok {
				return true
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				return true
			}

			comments := docMap[ts]
			if !commentContains(comments, targetComment) {
				return true
			}

			if len(st.Fields.List) == 0 {
				return true // empty struct? this should not happen
			}

			klog.Infof("found potential target Go struct %s", ts.Name.Name)
			targets = append(targets, target{
				goName: ts.Name.Name,
				endPos: int(fset.Position(st.End()).Offset),
			})
			return true // continue searching for potential target Go struct
		})

		var chosenTarget *target
		if len(targets) == 0 { // no target, continue to next file
			return nil
		} else if len(targets) == 1 { // case 1, one matching Go struct
			chosenTarget = &targets[0]
		} else if len(targets) == 2 { // case 2, Spec/ObservedState pair
			for _, t := range targets {
				if !outputOnly && strings.HasSuffix(t.goName, "Spec") ||
					outputOnly && strings.HasSuffix(t.goName, "ObservedState") {
					chosenTarget = &t
					break
				}
			}
		}

		if chosenTarget != nil { // target Go struct was found, modify the source bytes
			var newSrcBytes []byte
			// TODO: use the same field ordering as in proto message?
			newSrcBytes = append(newSrcBytes, srcBytes[:chosenTarget.endPos-1]...) // up to before '}'
			newSrcBytes = append(newSrcBytes, u.newField.generatedContent...)      // insert new field
			newSrcBytes = append(newSrcBytes, srcBytes[chosenTarget.endPos-1:]...) // include the '}'

			if err := os.WriteFile(path, newSrcBytes, d.Type()); err != nil {
				return err
			}
			klog.Infof("modified file %s", path)
		}
		return nil
	})

	return nil
}
