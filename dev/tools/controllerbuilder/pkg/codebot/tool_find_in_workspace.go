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

	"github.com/GoogleCloudPlatform/kubectl-ai/gollm"
	"k8s.io/klog/v2"
)

// func init() {
// 	RegisterTool(&FindInWorkspace{})
// }

type FindInWorkspace struct {
	FindText string `json:"find_text"`
}

type FindInWorkspaceResult struct {
	Matches []*Match `json:"matches"`
	Result  string   `json:"result"`
}

func (t *FindInWorkspace) Run(ctx context.Context, c *Chat, args map[string]any) (any, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	if err := json.Unmarshal(b, t); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", t, err)
	}

	result := &FindInWorkspaceResult{}

	klog.V(2).Infof("%T: %+v", t, t)

	matches, err := t.findInFiles(ctx, c.baseDir)
	if err != nil {
		return nil, fmt.Errorf("finding in files: %w", err)
	}

	result.Matches = matches
	result.Result = "success"

	return result, nil
}

func (t *FindInWorkspace) BuildFunctionDefinition() *gollm.FunctionDefinition {
	declaration := &gollm.FunctionDefinition{
		Name: "FindInWorkspace",
		Description: `
Search the code for a particular string.  This returns matches from the workspace, including the filename and a few lines of context for each match.
`,
		Parameters: &gollm.Schema{
			Type:     gollm.TypeObject,
			Required: []string{"find_text"},
			Properties: map[string]*gollm.Schema{
				"find_text": {
					Type: gollm.TypeString,
					Description: `
Find files in the workspace that include the specified string.
`,
				},
			},
		},
	}
	// TODO: Response?
	return declaration
}

type Match struct {
	Filename     string `json:"filename"`
	MatchingLine string `json:"matching_line"`
	Context      string `json:"context"`
}

func (t *FindInWorkspace) findInFiles(ctx context.Context, baseDir string) ([]*Match, error) {
	var matches []*Match
	if err := filepath.WalkDir(baseDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		fileContents, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading file %q: %w", path, err)
		}
		lines := bytes.Split(fileContents, []byte("\n"))
		for i, line := range lines {
			if bytes.Contains(line, []byte(t.FindText)) {
				var context bytes.Buffer
				start := max(0, i-2)
				end := min(len(lines), i+3)
				for j := start; j < end; j++ {
					fmt.Fprintf(&context, "%d: %s\n", j+1, lines[j])
				}
				relativePath, err := filepath.Rel(baseDir, path)
				if err != nil {
					return fmt.Errorf("getting relative path for %q: %w", path, err)
				}
				match := &Match{
					Filename:     relativePath,
					MatchingLine: string(line),
					Context:      context.String(),
				}
				matches = append(matches, match)
				return nil
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return matches, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
