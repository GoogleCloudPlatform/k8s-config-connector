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
	"os"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"

	"k8s.io/klog/v2"
)

type protoStruct struct {
	name string    // fully qualified name of the proto message
	end  token.Pos // the ending position of the corresponding Go struct in file
}

func (u *TypeUpdater) insertGoMessages() error {
	// find the file containing existing Go types
	goTypesFile, err := findGoTypesFile(u.opts.apiDirectory)
	if err != nil {
		return err
	}

	// parse the file
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, goTypesFile, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	// use a CommentMap to associate comments with nodes
	docMap := gocode.NewDocMap(fset, file)

	// collect existing structs with `+kcc:proto` comments
	existingStructs := []protoStruct{}
	ast.Inspect(file, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}
		_, ok = ts.Type.(*ast.StructType)
		if !ok {
			return true
		}

		comments := docMap[ts]
		name, err := protoNameFromComment(comments)
		if err != nil {
			return true // not a valid comment for proto name, continue
		}
		existingStructs = append(existingStructs,
			protoStruct{
				name: name,
				end:  n.End(),
			})
		return true
	})

	// read the existing go types file
	srcBytes, err := os.ReadFile(goTypesFile)
	if err != nil {
		return err
	}

	// update the go types file
	// maintain the ordering of the go types based on alphabetical order of the corresponding proto message name
	newSrcBytes := srcBytes
	for _, newStruct := range u.generatedGoStructs {
		// find the correct insertion point
		insertIndex := 0
		insertPos := 0
		for i, existing := range existingStructs {
			if newStruct.name < existing.name {
				insertIndex = i
				break
			}
			insertPos = int(fset.Position(existing.end).Offset) // update insertPos to the end of the current struct
		}

		// insert the new struct at the calculated position
		newSrcBytes = append(newSrcBytes[:insertPos], append(newStruct.content, newSrcBytes[insertPos:]...)...)
		insertPos += len(newStruct.content) // update to the end of the newly inserted struct

		// update the end positions of all subsequent structs
		for j := insertIndex; j < len(existingStructs); j++ {
			existingStructs[j].end += token.Pos(len(newStruct.content))
		}

		// update the existing structs list to include the newly inserted struct
		existingStructs = append(existingStructs[:insertIndex],
			append([]protoStruct{{name: newStruct.name, end: token.Pos(insertPos)}},
				existingStructs[insertIndex:]...)...)

		klog.Infof("inserted the generated Go struct for proto message %s", newStruct.name)
	}

	if err := os.WriteFile(goTypesFile, newSrcBytes, 0644); err != nil {
		return err
	}

	return nil
}

// findGoTypesFile walk the directory to find the target file
func findGoTypesFile(directory string) (string, error) {
	var goTypesFile string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == "types.generated.go" {
			goTypesFile = path
			return filepath.SkipDir // found the target, no need to continue walking
		}

		// if no "types.generated.go", check for a file ending with "_types.go"
		if goTypesFile == "" && filepath.Ext(info.Name()) == ".go" {
			if matched, _ := filepath.Match("*_types.go", info.Name()); matched {
				goTypesFile = path
			}
		}

		return nil
	})

	if err != nil {
		return "", err
	}
	if goTypesFile == "" {
		return "", fmt.Errorf("did not found Go types file in directory %s", directory)
	}
	return goTypesFile, nil
}

func protoNameFromComment(cg *ast.CommentGroup) (string, error) {
	if cg == nil {
		return "", fmt.Errorf("empty comment group")
	}
	for _, c := range cg.List {
		trimmed := strings.TrimPrefix(c.Text, "//")
		trimmed = strings.TrimSpace(trimmed)
		if !strings.HasPrefix(trimmed, kccProtoPrefix) {
			continue
		}
		return strings.TrimSpace(strings.TrimPrefix(trimmed, kccProtoPrefix)), nil // found the comment with proto name
	}
	return "", fmt.Errorf("not found")
}
