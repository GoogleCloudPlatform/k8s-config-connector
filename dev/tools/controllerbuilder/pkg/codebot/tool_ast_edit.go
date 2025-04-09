// Copyright 2025 Google LLC
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

package codebot

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/GoogleCloudPlatform/kubectl-ai/gollm"
	"k8s.io/klog/v2"
)

func init() {
	// ASTEdit is currently not enabled, EditFile seems to do OK with ReadFile
	// RegisterTool(&ASTEdit{})
}

type ASTEdit struct {
	Node     string `json:"node"`
	Action   string `json:"action"`
	Filename string `json:"filename"`
}

type ASTEditResults struct {
	Success bool `json:"success"`
}

func (t *ASTEdit) Run(ctx context.Context, c *Chat, args map[string]any) (*ASTEditResults, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	if err := json.Unmarshal(b, t); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", t, err)
	}

	klog.V(2).Infof("%T: %+v", t, t)

	result := &ASTEditResults{}

	klog.V(2).Infof("ASTEdit: %+v", t)

	p := filepath.Join(c.baseDir, t.Filename)
	fileContentsBytes, err := os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("reading file %q: %w", p, err)
	}

	fileContents := string(fileContentsBytes)

	if t.Action == "" {
		return nil, fmt.Errorf("the action argument is requiremnt")
	}

	if t.Node == "" {
		return nil, fmt.Errorf("the node argument is requiremnt")
	}

	switch t.Action {
	case "ADD":
		fileContents = fileContents + "\n" + t.Node

	case "REPLACE":
		klog.Fatalf("REPLACE node not yet implemented: %+v", t)

	case "DELETE":
		klog.Fatalf("DELETE node not yet implemented: %+v", t)

	default:
		return nil, fmt.Errorf("the action %q is not known; it should be ADD, REPLACE or DELETE", t.Action)
	}
	if err := os.WriteFile(p, []byte(fileContents), 0644); err != nil {
		return nil, fmt.Errorf("writing file %q: %w", p, err)
	}

	result.Success = true
	return result, nil
}

func (t *ASTEdit) BuildFunctionDefinition() *gollm.FunctionDefinition {
	declaration := &gollm.FunctionDefinition{
		Name: "ast_edit",
		Description: `
		Makes changes to an existing file in a way that understands the syntax of the file (the AST, or Abstract Syntax Tree).
		`,
		Parameters: &gollm.Schema{
			Type:     gollm.TypeObject,
			Required: []string{"filename", "action", "node"},
			Properties: map[string]*gollm.Schema{
				"action": {
					Type: gollm.TypeString,
					Description: `
		The action to perform; must be one of ADD, REPLACE, DELETE.
	
		For example, to update a method, use REPLACE here and provided the updated method in node.`,
				},
				"node": {
					Type: gollm.TypeString,
					Description: `
		The AST node to ADD, REPLACE or DELETE.  This should be a top level node in the AST of the language you're using, for example a function declaration or a type declaration.
		`,
				},
				"filename": {
					Type:        gollm.TypeString,
					Description: "The path to the file you want to change",
				},
			},
		},
	}
	// TODO: Response?
	return declaration
}
