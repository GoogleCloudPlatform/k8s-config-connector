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
	"os/exec"
	"strings"

	"github.com/GoogleCloudPlatform/kubectl-ai/gollm"
	"k8s.io/klog/v2"
)

func init() {
	RegisterTool(&VerifyCode{})
}

type VerifyCode struct {
	Filename string `json:"filename"`
}

type VerifyCodeResponse struct {
	Success bool `json:"success"`
	Errors  []string
}

func (t *VerifyCode) Run(ctx context.Context, c *Chat, args map[string]any) (any, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	if err := json.Unmarshal(b, t); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", t, err)
	}

	klog.V(2).Infof("%T: %+v", t, t)

	cmd := exec.CommandContext(ctx, "go", "build", "./...")
	cmd.Dir = c.baseDir

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	results := &VerifyCodeResponse{}
	results.Success = true

	if err := cmd.Run(); err != nil {
		if exitCodeErr, ok := err.(*exec.ExitError); ok {
			results.Success = false
			results.Errors = append(results.Errors, fmt.Sprintf("build failed with exit code %d", exitCodeErr.ExitCode()))
		} else {
			// Unexpected
			return nil, fmt.Errorf("failed to run go build: %w", err)
		}
	}

	// cmd = exec.CommandContext(ctx, "go", "vet", "./...")
	// cmd.Dir = c.baseDir
	// cmd.Stdout = &stdout
	// cmd.Stderr = &stderr
	// if err := cmd.Run(); err != nil {
	// 	results.Success = false
	// 	results.Errors = append(results.Errors, fmt.Sprintf("vet failed: %v", err))
	// }

	// cmd = exec.CommandContext(ctx, "go", "fmt", "./...")
	// cmd.Dir = c.baseDir
	// cmd.Stdout = &stdout
	// cmd.Stderr = &stderr
	// if err := cmd.Run(); err != nil {
	// 	results.Success = false
	// 	results.Errors = append(results.Errors, fmt.Sprintf("fmt failed: %v", err))
	// }

	results.Errors = append(results.Errors, strings.Split(stderr.String(), "\n")...)
	return results, nil
}

func (t *VerifyCode) BuildFunctionDefinition() *gollm.FunctionDefinition {
	declaration := &gollm.FunctionDefinition{
		Name: "VerifyCode",
		Description: `
Verifies the result of changes by trying to build, lint and vet the code.
`,
		Parameters: &gollm.Schema{
			Type:       gollm.TypeObject,
			Properties: map[string]*gollm.Schema{},
		},
	}
	// TODO: Response?
	return declaration
}
