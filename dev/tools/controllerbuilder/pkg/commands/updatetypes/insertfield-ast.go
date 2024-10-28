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

package updatetypes

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

	"k8s.io/klog/v2"
)

func (u *TypeUpdater) insertGoField() error {
	klog.Infof("inserting the generated Go code for field %s", u.newField.field.Name())

	targetComment := fmt.Sprintf("+kcc:proto=%s", u.generatedGoField.parentMessage)

	filepath.WalkDir(u.opts.apiDirectory, func(path string, d fs.DirEntry, err error) error {
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

		// find the target Go struct and its ending position in the source
		var endPos int
		ast.Inspect(file, func(n ast.Node) bool {
			if endPos != 0 {
				return false // already found the target
			}

			ts, ok := n.(*ast.TypeSpec)
			if !ok {
				return true
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				return true
			}

			comments := docMap[ts]
			if !isTargetStruct(comments, targetComment) {
				return true
			}

			if len(st.Fields.List) == 0 {
				return true // empty struct? this should not happen
			}

			klog.Infof("found target Go struct %s", ts.Name.Name)

			endPos = int(fset.Position(st.End()).Offset)
			return false // stop searching, we found the target Go struct
		})

		// if the target Go struct was found, modify the source bytes
		if endPos != 0 {
			var newSrcBytes []byte
			// TODO: use the same field ordering as in proto message
			newSrcBytes = append(newSrcBytes, srcBytes[:endPos-1]...)        // up to before '}'
			newSrcBytes = append(newSrcBytes, u.generatedGoField.content...) // insert new field
			newSrcBytes = append(newSrcBytes, srcBytes[endPos-1:]...)        // include the '}'

			if err := os.WriteFile(path, newSrcBytes, d.Type()); err != nil {
				return err
			}
			klog.Infof("modified file %s", path)
		}
		return nil
	})

	return nil
}

func isTargetStruct(cg *ast.CommentGroup, target string) bool {
	if cg == nil {
		return false
	}
	for _, c := range cg.List {
		trimmed := strings.TrimPrefix(c.Text, "//")
		trimmed = strings.TrimSpace(trimmed)
		if trimmed == target {
			return true
		}
	}
	return false
}
