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
	"strings"

	"github.com/GoogleCloudPlatform/kubectl-ai/gollm"
	"k8s.io/klog/v2"
)

func init() {
	RegisterTool(&EditFile{})
}

type EditFile struct {
	ExistingText string `json:"existing_text"`
	NewText      string `json:"new_text"`
	Filename     string `json:"filename"`
}

type EditFileResults struct {
	Success bool `json:"success"`
}

func (t *EditFile) Run(ctx context.Context, c *Chat, args map[string]any) (any, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	if err := json.Unmarshal(b, t); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", t, err)
	}

	klog.V(2).Infof("%T: %+v", t, t)

	results, err := t.runEditFile(ctx, c.baseDir)
	if err != nil {
		return nil, fmt.Errorf("finding in files: %w", err)
	}
	return results, nil
}

func (t *EditFile) BuildFunctionDefinition() *gollm.FunctionDefinition {
	declaration := &gollm.FunctionDefinition{
		Name: "EditFile",
		Description: `
Make a change to an existing file in the user's workspace, by replacing existing_text with new_text.  This tool only applies the first replacement.

For example, given the following function:

<file>example.go
func a() error {
  return fmt.Errorf("oops")
}

func b() error {
  a()
  return nil
}
</file>

If you wanted to add error handling to the function "b", you could call EditFile with these parameters:

<filename>example.go</filename>
<new_text>
func b() error {
  if err := a(); err != nil {
    return fmt.Errorf("error calling b: %w", err")
  }
  return nil
</new_text>
<existing_text>
func b() error {
  a()
  return nil
</existing_text>
`,
		Parameters: &gollm.Schema{
			Type:     gollm.TypeObject,
			Required: []string{"new_text", "existing_text", "filename"},
			Properties: map[string]*gollm.Schema{
				"existing_text": {
					Type:        gollm.TypeString,
					Description: `The text to find, which will be replaced with the contents of the new_text argument.  Provide all the lines of the existing content you want to replace.`,
				},
				"new_text": {
					Type:        gollm.TypeString,
					Description: `The new content of the file, with a few lines of context. Separate chunks of content with ... on its own line.`,
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

func (t *EditFile) runEditFile(ctx context.Context, baseDir string) (*EditFileResults, error) {
	p := filepath.Join(baseDir, t.Filename)
	fileContents, err := os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("reading file %q: %w", p, err)
	}

	if t.ExistingText == "" {
		return nil, fmt.Errorf("the existing_text argument is requiremnt")
	}

	var f fuzzyFile
	f.SetContents(string(fileContents))

	verbose := true
	if verbose {
		klog.Infof("Before edit:")
		for i, line := range f.lines {
			klog.Infof("%d | %s", i, line)
		}

		klog.Infof("ExistingText:")
		for i, line := range strings.Split(t.ExistingText, "\n") {
			klog.Infof("%d | %s", i, line)
		}

		klog.Infof("NewText:")
		for i, line := range strings.Split(t.NewText, "\n") {
			klog.Infof("%d | %s", i, line)
		}
	}

	if err := f.FindAndReplace(t.ExistingText, t.NewText); err != nil {
		return nil, err
	}

	if verbose {
		klog.Infof("After edit:")
		for i, line := range f.lines {
			klog.Infof("%d | %s", i, line)
		}
	}

	newContents := f.Contents()
	if err := os.WriteFile(p, []byte(newContents), 0644); err != nil {
		return nil, fmt.Errorf("writing file %q: %w", p, err)
	}

	klog.V(2).Infof("wrote %v: %v", p, string(newContents))
	return &EditFileResults{
		Success: true,
	}, nil
}

type fuzzyFile struct {
	lines []string
}

func (f *fuzzyFile) SetContents(contents string) {
	f.lines = strings.Split(contents, "\n")
}

func (f *fuzzyFile) Contents() string {
	return strings.Join(f.lines, "\n")
}

func (f *fuzzyFile) FindAndReplace(find string, replace string) error {
	contents := f.Contents()

	// Check for an exact match first
	ix := strings.Index(contents, find)
	if ix != -1 {
		contents = contents[:ix] + replace + contents[ix+len(find):]
		f.SetContents(contents)
		return nil
	}

	// Look for a fuzzy match, where the lines match ignoring whitespace
	findLines := strings.Split(find, "\n")
	for i := range findLines {
		findLines[i] = strings.TrimSpace(findLines[i])
	}
	for i := range f.lines {
		match := true
		// TODO: Hold on ... this is an exact match comparison!  We could do this with strings.Index!
		for j := 0; j < len(findLines) && i+j < len(f.lines); j++ {
			if strings.TrimSpace(f.lines[i+j]) != findLines[j] {
				match = false
				break
			}
		}
		if match {
			replaceLines := strings.Split(replace, "\n")
			newLines := make([]string, 0, len(f.lines)+len(replaceLines))
			newLines = append(newLines, f.lines[:i]...)
			newLines = append(newLines, replaceLines...)
			newLines = append(newLines, f.lines[i+len(findLines):]...)
			f.lines = newLines
			return nil
		}
	}
	return fmt.Errorf("could not find the `existing_text` string %q", find)
}
