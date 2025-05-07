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

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/llm"
	"k8s.io/klog/v2"
)

func init() {
	RegisterTool(&CreateFile{})
}

type CreateFile struct {
	Contents  string `json:"contents"`
	Filename  string `json:"filename"`
	Overwrite bool   `json:"overwrite"`
}

type CreateFileResults struct {
	Success bool `json:"success"`
}

func (t *CreateFile) Run(ctx context.Context, c *Chat, args map[string]any) (any, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	if err := json.Unmarshal(b, t); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", t, err)
	}

	klog.V(2).Infof("%T: %+v", t, t)

	result := &CreateFileResults{}

	klog.V(2).Infof("CreateFile: %+v", t)

	p := filepath.Join(c.baseDir, t.Filename)
	if !t.Overwrite {
		if _, err := os.Stat(p); err == nil {
			return nil, fmt.Errorf("file %q already exists", t.Filename)
		}
	}
	err = os.MkdirAll(filepath.Dir(p), 0755)
	if err != nil {
		return nil, fmt.Errorf("creating dir %s: %w", filepath.Dir(p), err)
	}
	f, err := os.Create(p)
	if err != nil {
		return nil, fmt.Errorf("creating file %s: %w", p, err)
	}
	defer f.Close()

	if t.Contents == "" {
		return nil, fmt.Errorf("the contents argument is requiremnt")
	}
	if _, err := f.WriteString(t.Contents); err != nil {
		return nil, fmt.Errorf("writing file %q: %w", p, err)
	}

	result.Success = true
	return result, nil
}

func (t *CreateFile) BuildFunctionDefinition() *llm.FunctionDefinition {
	declaration := &llm.FunctionDefinition{
		Name:        "CreateFile",
		Description: `Create a new file in the user's workspace, with the contents specified by contents.  This tool will fail if the file already exists.`,
		Parameters: &llm.Schema{
			Type:     llm.TypeObject,
			Required: []string{"contents", "filename"},
			Properties: map[string]*llm.Schema{
				"contents": {
					Type:        llm.TypeString,
					Description: `The text that should be the contents of the new file.`,
				},
				"filename": {
					Type:        llm.TypeString,
					Description: "The path to the file you want to create",
				},
				"overwrite": {
					Type:        llm.TypeBoolean,
					Description: "Whether to overwrite the file if it already exists",
				},
			},
		},
	}
	// TODO: Response?
	return declaration
}
