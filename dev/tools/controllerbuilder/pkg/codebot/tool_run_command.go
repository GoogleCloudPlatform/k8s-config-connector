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

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/llm"
	"k8s.io/klog/v2"
)

func init() {
	RegisterTool(&RunShellCommand{})
}

type RunShellCommand struct {
	Command string `json:"shell_command"`
}

func (t *RunShellCommand) Run(ctx context.Context, c *Chat, llmArgs map[string]any) (map[string]any, error) {
	b, err := json.Marshal(llmArgs)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	if err := json.Unmarshal(b, t); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", t, err)
	}

	result := make(map[string]any)

	klog.V(2).Infof("%T: %+v", t, t)

	args := strings.Fields(t.Command)
	if len(args) == 0 {
		return nil, fmt.Errorf("shell_command is required")
	}
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	cmd.Dir = c.baseDir
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		switch err := err.(type) {
		case *exec.ExitError:
			result["exit_code"] = err.ExitCode()
		default:
			return nil, fmt.Errorf("unexpected error running command: %w", err)
		}
	} else {
		result["exit_code"] = 0
	}
	result["stdout"] = stdout.String()
	result["stderr"] = stderr.String()

	return result, nil
}

func (t *RunShellCommand) BuildFunctionDefinition() *llm.FunctionDefinition {
	declaration := &llm.FunctionDefinition{
		Name:        "RunShellCommand",
		Description: `Run a shell command in the workspace.  This returns the stdout, stderr and exit code from running the command.`,
		Parameters: &llm.Schema{
			Type:     llm.TypeObject,
			Required: []string{"shell_command"},
			Properties: map[string]*llm.Schema{
				"shell_command": {
					Type:        llm.TypeString,
					Description: `The shell command to run in the workspace.`,
				},
			},
		},
	}
	// TODO: Response?
	return declaration
}
