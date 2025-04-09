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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot/ui"
	"github.com/GoogleCloudPlatform/kubectl-ai/gollm"
	"k8s.io/klog/v2"
)

type Chat struct {
	client  gollm.Client
	session gollm.Chat
	baseDir string
	ui      ui.UI

	toolbox *Toolbox
}

type FileInfo struct {
	Path    string
	Content string
}

func NewChat(ctx context.Context, llmClient gollm.Client, model string, baseDir string, contextFiles map[string]*FileInfo, toolbox *Toolbox, ui ui.UI) (*Chat, error) {
	systemPrompt := `
You are a helpful AI coding assistant, expert in the go programming language and in creating kubernetes controllers.

Help the user with their problems.  Do not do more than the user asks; propose minimal changes to solve what the user wants.

If you think the code should be changed, use the tools to apply those changes instead of printing them.  Do not make multiple parallel changes to the same file.
`

	functionDefinitions := toolbox.GetFunctionDefinitions()

	if len(contextFiles) != 0 {
		var sb strings.Builder

		fmt.Fprintf(&sb, "\n")
		fmt.Fprintf(&sb, "Consider the following files:\n")

		for _, file := range contextFiles {
			fmt.Fprintf(&sb, "File %q:\n", file.Path)
			fmt.Fprintf(&sb, "```go\n")
			fmt.Fprintf(&sb, "%s\n", file.Content)
			fmt.Fprintf(&sb, "```\n")
			fmt.Fprintf(&sb, "\n")
		}

		systemPrompt += sb.String()
	}

	session := llmClient.StartChat(systemPrompt, model)

	session.SetFunctionDefinitions(functionDefinitions)

	return &Chat{
		client:  llmClient,
		baseDir: baseDir,
		session: session,
		ui:      ui,
		toolbox: toolbox,
	}, nil

}

func (c *Chat) Close() error {
	return c.client.Close()
}

func (c *Chat) SendMessage(ctx context.Context, userParts ...any) error {
	resp, err := c.session.Send(ctx, userParts...)
	if err != nil {
		return fmt.Errorf("sending message to LLM: %w", err)
	}

	for {
		// Print the usage metadata (includes token count i.e. cost)
		klog.Infof("UsageMetadata: %+v", resp.UsageMetadata())

		if len(resp.Candidates()) == 0 {
			return fmt.Errorf("no response candidates from LLM")
		}

		candidate := resp.Candidates()[0]
		if len(resp.Candidates()) > 1 {
			klog.Warningf("found multiple responses from LLM, only considering the first")
		}

		klog.Infof("processing candidate %+v", candidate)

		var functionResponses []any

		for _, part := range candidate.Parts() {
			if text, ok := part.AsText(); ok {
				s := string(text)
				c.ui.AddLLMOutput(&ui.LLMOutput{Text: s})
				klog.V(2).Infof("TEXT: %+v", s)
			}
			if functionCalls, ok := part.AsFunctionCalls(); ok {
				for _, functionCall := range functionCalls {
					// klog.Infof("functionCall: %+v", functionCall)
					klog.V(2).Infof("functionCall: %+v", functionCall.Name)
					c.ui.AddLLMOutput(&ui.LLMOutput{Text: fmt.Sprintf("functionCall: %+v", functionCall)})
					result, err := c.toolbox.CallFunction(ctx, c, functionCall)
					if err != nil {
						return fmt.Errorf("unexpected error running function: %w", err)
					}
					var response map[string]any
					if result.Error != nil {
						response = make(map[string]any)
						response["result"] = "error"
						response["error"] = fmt.Sprintf("%v", result.Error)
						c.ui.AddLLMOutput(&ui.LLMOutput{Text: fmt.Sprintf("error running function: %v", result.Error)})
					} else {
						b, err := json.Marshal(result.Response)
						if err != nil {
							return fmt.Errorf("converting %T to json: %w", result.Response, err)
						}
						m := make(map[string]any)
						if err := json.Unmarshal(b, &m); err != nil {
							return fmt.Errorf("unmarshaling json: %w", err)
						}
						response = m
					}
					functionResponses = append(functionResponses, gollm.FunctionCallResult{
						Name:   functionCall.Name,
						Result: response,
					})
					b, _ := json.Marshal(response)
					c.ui.AddLLMOutput(&ui.LLMOutput{Text: fmt.Sprintf("sending response: %v", string(b))})
				}
			}
		}

		if len(functionResponses) == 0 {
			return nil
		}

		// Go round again, but this time reply with the function responses
		klog.Infof("functionResponses: %+v", functionResponses)
		resp, err = c.session.Send(ctx, functionResponses...)
		if err != nil {
			return fmt.Errorf("sending message to LLM: %w", err)
		}
	}
}
