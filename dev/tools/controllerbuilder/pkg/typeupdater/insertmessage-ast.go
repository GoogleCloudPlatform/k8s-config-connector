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
	"os"
	"path/filepath"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"

	"k8s.io/klog/v2"
)

type goStruct struct {
	name  string // fully qualified name of the proto message
	start int    // byte offset of the start of this struct
	end   int    // byte offset of the end of this struct
}

func (u *FieldInserter) insertGoMessages() error {
	if len(u.dependentMessages) == 0 {
		return nil
	}

	// find the file containing existing Go types
	goTypesFile, err := findGoTypesFile(u.opts.APIDirectory)
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
	existingStructs := []goStruct{}
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
			goStruct{
				name:  name,
				start: fset.Position(n.Pos()).Offset,
				end:   fset.Position(n.End()).Offset,
			})
		return true
	})
	if len(existingStructs) == 0 {
		return fmt.Errorf("no Go struct was found in file %s", goTypesFile)
	}

	// read the existing go types file
	srcBytes, err := os.ReadFile(goTypesFile)
	if err != nil {
		return err
	}

	// update the go types file
	// maintain the ordering of the go types based on alphabetical order of the corresponding proto message name
	newSrcBytes := srcBytes
	for _, msg := range u.dependentMessages {
		targetName := string(msg.proto.FullName())

		// find the correct insertion point
		insertIndex := 0
		insertPos := existingStructs[0].start
		for i, existing := range existingStructs {
			if targetName < existing.name {
				// found the first struct that should come after the new one
				insertIndex = i
				break
			}
			insertPos = existing.end // update insertPos to the end of the current struct
		}

		if insertIndex == len(existingStructs) { //  append at the end
			insertPos = existingStructs[insertIndex-1].end
		}

		// insert the new struct to file
		newSrcBytes = append(newSrcBytes[:insertPos],
			append(msg.generatedContent, newSrcBytes[insertPos:]...)...)

		// update existing structs list
		newStruct := goStruct{
			name:  targetName,
			start: insertPos,
			end:   insertPos + len(msg.generatedContent),
		}
		existingStructs = append(existingStructs[:insertIndex],
			append([]goStruct{newStruct}, existingStructs[insertIndex:]...)...)
		// update the start and end positions of all subsequent structs
		for j := insertIndex + 1; j < len(existingStructs); j++ {
			existingStructs[j].start += len(msg.generatedContent)
			existingStructs[j].end += len(msg.generatedContent)
		}

		klog.Infof("inserted the generated Go struct for proto message %s", targetName)
	}

	if err := os.WriteFile(goTypesFile, newSrcBytes, 0644); err != nil {
		return err
	}
	klog.Infof("modified file %s", goTypesFile)

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
		return "", fmt.Errorf("did not find Go types file in directory %s", directory)
	}
	return goTypesFile, nil
}

func protoNameFromComment(cg *ast.CommentGroup) (string, error) {
	if cg == nil {
		return "", fmt.Errorf("empty comment group")
	}
	for _, c := range cg.List {
		proto, ok := codegen.GetProtoMessageFromAnnotation(c.Text)
		if ok {
			return proto, nil // found the comment with proto name
		}
	}
	return "", fmt.Errorf("not found")
}
