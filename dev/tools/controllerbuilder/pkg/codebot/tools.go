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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"cloud.google.com/go/vertexai/genai"
	"k8s.io/klog/v2"
)

func (c *Chat) runFunctionCall(ctx context.Context, functionCall genai.FunctionCall) (map[string]any, error) {
	log := klog.FromContext(ctx)
	switch functionCall.Name {
	case "edit_file":
		result, err := c.runEditFile(ctx, functionCall.Args)
		if err != nil {
			result = make(map[string]any)
			result["result"] = "error"
			result["error"] = fmt.Sprintf("%v", err)
			log.Info("unable to apply edit_file", result, result)
			return result, nil
		}
		return result, nil
	case "create_file":
		result, err := c.runCreateFile(ctx, functionCall.Args)
		if err != nil {
			result = make(map[string]any)
			result["result"] = "error"
			result["error"] = fmt.Sprintf("%v", err)
			log.Info("unable to apply create_file", result, result)
			return result, nil
		}
		return result, nil
	case "ast_edit":
		result, err := c.runASTEdit(ctx, functionCall.Args)
		if err != nil {
			result = make(map[string]any)
			result["result"] = "error"
			result["error"] = fmt.Sprintf("%v", err)
			log.Info("unable to apply ast_edit", result, result)
			return result, nil
		}
		return result, nil
	default:
		// TODO: Fatal or return an error?
		return nil, fmt.Errorf("unknown function %q", functionCall.Name)
	}
}

type EditFile struct {
	Find     string `json:"existing_text"`
	Replace  string `json:"new_text"`
	Filename string `json:"filename"`
}

func (c *Chat) runEditFile(ctx context.Context, args map[string]any) (map[string]any, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	var editFile EditFile
	if err := json.Unmarshal(b, &editFile); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", &editFile, err)
	}

	result := make(map[string]any)

	klog.Infof("EditFile: %+v", editFile)

	p := filepath.Join(c.baseDir, editFile.Filename)
	fileContents, err := os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("reading file %q: %w", p, err)
	}

	if editFile.Find == "" {
		return nil, fmt.Errorf("the find argument is requiremnt")
	}

	ix := bytes.Index(fileContents, []byte(editFile.Find))
	if ix == -1 {
		return nil, fmt.Errorf("could not find the `find` string %q in the file %q", editFile.Find, p)
	}

	newContents := bytes.Replace(fileContents, []byte(editFile.Find), []byte(editFile.Replace), 1)
	if err := os.WriteFile(p, newContents, 0644); err != nil {
		return nil, fmt.Errorf("writing file %q: %w", p, err)
	}

	result["result"] = "success"
	return result, nil
}

type CreateFile struct {
	Contents string `json:"contents"`
	Filename string `json:"filename"`
}

func (c *Chat) runCreateFile(ctx context.Context, args map[string]any) (map[string]any, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	var createFile CreateFile
	if err := json.Unmarshal(b, &createFile); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", &createFile, err)
	}

	result := make(map[string]any)

	klog.Infof("CreateFile: %+v", createFile)

	p := filepath.Join(c.baseDir, createFile.Filename)
	if _, err := os.Stat(p); err == nil {
		return nil, fmt.Errorf("file %q already exists", createFile.Filename)
	}

	if createFile.Contents == "" {
		return nil, fmt.Errorf("the contents argument is requiremnt")
	}

	newContents := []byte(createFile.Contents)
	if err := os.WriteFile(p, newContents, 0644); err != nil {
		return nil, fmt.Errorf("writing file %q: %w", p, err)
	}

	result["result"] = "success"
	return result, nil
}

type ASTEdit struct {
	Node     string `json:"node"`
	Action   string `json:"action"`
	Filename string `json:"filename"`
}

func (c *Chat) runASTEdit(ctx context.Context, args map[string]any) (map[string]any, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	var astEdit ASTEdit
	if err := json.Unmarshal(b, &astEdit); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", &astEdit, err)
	}

	result := make(map[string]any)

	klog.Infof("ASTEdit: %+v", astEdit)

	p := filepath.Join(c.baseDir, astEdit.Filename)
	fileContentsBytes, err := os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("reading file %q: %w", p, err)
	}

	fileContents := string(fileContentsBytes)

	if astEdit.Action == "" {
		return nil, fmt.Errorf("the action argument is requiremnt")
	}

	if astEdit.Node == "" {
		return nil, fmt.Errorf("the node argument is requiremnt")
	}

	switch astEdit.Action {
	case "ADD":
		fileContents = fileContents + "\n" + astEdit.Node

	case "REPLACE":
		klog.Fatalf("REPLACE node not yet implemented: %+v", astEdit)

	case "DELETE":
		klog.Fatalf("DELETE node not yet implemented: %+v", astEdit)

	default:
		return nil, fmt.Errorf("the action %q is not known; it should be ADD, REPLACE or DELETE", astEdit.Action)
	}
	if err := os.WriteFile(p, []byte(fileContents), 0644); err != nil {
		return nil, fmt.Errorf("writing file %q: %w", p, err)
	}

	result["result"] = "success"
	return result, nil
}
