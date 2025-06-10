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
	"os/exec"
	"strings"

	"github.com/GoogleCloudPlatform/kubectl-ai/gollm"
	"k8s.io/klog/v2"
)

func init() {
	RegisterTool(&RunTerminalCommand{})
}

type RunTerminalCommand struct {
	Command string `json:"command"`
	Args    string `json:"args"`
}

type RunTerminalCommandResults struct {
	Success bool   `json:"success"`
	Output  string `json:"output"`
	Error   string `json:"error"`
}

func (t *RunTerminalCommand) Run(ctx context.Context, c *Chat, args map[string]any) (any, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	if err := json.Unmarshal(b, t); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", t, err)
	}

	klog.V(2).Infof("%T: %+v", t, t)

	result := &RunTerminalCommandResults{}

	klog.V(2).Infof("RunTerminalCommand: %+v", t)

	var cmd *exec.Cmd

	tokens := strings.Split(t.Command, " ")
	switch tokens[0] {
	case "go":
		cmd = exec.CommandContext(ctx, "go", tokens[1:]...)
	case "make":
		cmd = exec.CommandContext(ctx, "make", tokens[1:]...)
	case "gcloud":
		cmd = exec.CommandContext(ctx, "gcloud", tokens[1:]...)
	default:
		cmd = exec.CommandContext(ctx, "bash", "-c", t.Command)
	}
	cmd.Dir = c.baseDir
	cmd.Args = append(tokens, strings.Split(t.Args, " ")...)

	output, err := cmd.CombinedOutput()
	if err != nil {

		if exitError, ok := err.(*exec.ExitError); ok {
			result.Error = fmt.Sprintf("command failed with exit code %d: %s", exitError.ExitCode(), string(exitError.Stderr))
		} else {
			result.Error = fmt.Sprintf("command failed: %s", err.Error())
		}
		result.Success = false
	} else {
		result.Success = true
	}
	if len(output) > 0 {
		result.Output = string(output)
	}
	return result, nil
}

func (t *RunTerminalCommand) BuildFunctionDefinition() *gollm.FunctionDefinition {
	declaration := &gollm.FunctionDefinition{
		Name: "RunTerminalCommand",
		Description: `
		The function is to run a command in the operating system's terminal/command prompt. This is different from a function that adds two numbers. It interacts with the external environment.

		For example, if I want to run a make rule ` + "`" + `make ready-pr` + "`" + `, this function should be called with the ` + "`" + `command` + "`" + ` to be ` + "`" + `make` + "`" + ` and ` + "`" + `args` + "`" + ` to be ` + "`" + `ready-pr` + "`" + `.
		`,
		Parameters: &gollm.Schema{
			Description: "The input is the command to run and its arguments. The command is required and the arguments is optional.",
			Type:        gollm.TypeObject,
			Required:    []string{"command", "args"},
			Properties: map[string]*gollm.Schema{
				"command": {
					Type: gollm.TypeString,
					Description: `
		command is the operating system's terminal command we want to run. It can be golang command, make rules, or bash script.

		For example, make rule should have command ` + "`" + `make ` + "`" + `, golang command should have command ` + "`" + `go test` + "`" + `, or ` + "`" + `go run` + "`" + `, 
		bash script is the default and don't need a command.
		`,
				},
				"args": {
					Type: gollm.TypeString,
					Description: `
		args is the arguments to the command. It is optional. If the command is ` + "`" + `make` + "`" + `, the args should be the make rule.
		If the command is ` + "`" + `go` + "`" + `, the args should be the go command with go arguments.
		If the command is bash script, the args should be the bash script arguments.
		`,
				},
			},
		},
	}
	// TODO: Response?
	return declaration
}
