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

	"github.com/GoogleCloudPlatform/kubectl-ai/gollm"
	"k8s.io/klog/v2"
)

func init() {
	RegisterTool(&RunShellCommand{})
}

type RunShellCommand struct {
	Command string `json:"shell_command"`
}

type RunShellCommandResults struct {
	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
	ExitCode int    `json:"exit_code"`
}

func (t *RunShellCommand) Run(ctx context.Context, c *Chat, llmArgs map[string]any) (any, error) {
	b, err := json.Marshal(llmArgs)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	if err := json.Unmarshal(b, t); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", t, err)
	}

	result := &RunShellCommandResults{}

	klog.V(2).Infof("%T: %+v", t, t)

	// args := strings.Fields(t.Command)
	// if len(args) == 0 {
	// 	return nil, fmt.Errorf("shell_command is required")
	// }

	args := []string{"/bin/bash", "-c", t.Command}
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	cmd.Dir = c.baseDir
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		switch err := err.(type) {
		case *exec.ExitError:
			result.ExitCode = err.ExitCode()
		default:
			return nil, fmt.Errorf("unexpected error running command: %w", err)
		}
	} else {
		result.ExitCode = 0
	}
	result.Stdout = stdout.String()
	result.Stderr = stderr.String()

	return result, nil
}

func (t *RunShellCommand) BuildFunctionDefinition() *gollm.FunctionDefinition {
	declaration := &gollm.FunctionDefinition{
		Name:        "RunShellCommand",
		Description: `Run a shell command in the workspace.  This returns the stdout, stderr and exit code from running the command.`,
		Parameters: &gollm.Schema{
			Type:     gollm.TypeObject,
			Required: []string{"shell_command"},
			Properties: map[string]*gollm.Schema{
				"shell_command": {
					Type:        gollm.TypeString,
					Description: `The shell command to run in the workspace.`,
				},
			},
		},
	}
	// TODO: Response?
	return declaration
}
