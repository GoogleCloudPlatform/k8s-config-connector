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
	RegisterTool(&ReadFile{})
}

type ReadFile struct {
	Filename string `json:"filename"`
}

type ReadFileResponse struct {
	Contents string `json:"contents"`
}

func (t *ReadFile) Run(ctx context.Context, c *Chat, args map[string]any) (any, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	if err := json.Unmarshal(b, t); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", t, err)
	}

	klog.V(2).Infof("%T: %+v", t, t)

	p := filepath.Join(c.baseDir, t.Filename)
	fileContents, err := os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("reading file %q: %w", p, err)
	}

	return &ReadFileResponse{
		Contents: string(fileContents),
	}, nil
}

func (t *ReadFile) BuildFunctionDefinition() *gollm.FunctionDefinition {
	declaration := &gollm.FunctionDefinition{
		Name: "ReadFile",
		Description: `
Reads the contents of a file in the workspace.  This returns the full contents of the given project file.
`,
		Parameters: &gollm.Schema{
			Type:     gollm.TypeObject,
			Required: []string{"filename"},
			Properties: map[string]*gollm.Schema{
				"filename": {
					Type: gollm.TypeString,
					Description: `
The path to the file in the workspace you want to read.
`,
				},
			},
		},
	}
	// TODO: Response?
	return declaration
}
