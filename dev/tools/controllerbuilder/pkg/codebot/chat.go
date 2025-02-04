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

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot/repocontext"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot/rules"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot/ui"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/llm"
	"k8s.io/klog/v2"
)

type Chat struct {
	client  llm.Client
	session llm.Chat
	baseDir string
	ui      ui.UI
}

type Options struct {
	BaseDir string
	ContextFiles map[string]*repocontext.FileInfo
	MockGCPFiles map[string]string
	APIsFiles map[string]string

	Rules []rules.Rule // user defined rules for the system prompt
}

func NewChat(ctx context.Context, llmClient llm.Client, ui ui.UI, opts *Options) (*Chat, error) {

	systemPrompt := `
You are a helpful AI coding assistant, expert in the go programming language and in creating kubernetes controllers.

Help the user with their problems.  Do not do more than the user asks; propose minimal changes to solve what the user wants.

If you think the code should be changed, use the tools to apply those changes instead of printing them.  Do not make multiple parallel changes to the same file.
`

	var functionDefinitions []*llm.FunctionDefinition

	for _, tool := range GetTools() {
		functionDefinitions = append(functionDefinitions, tool.BuildFunctionDefinition())
	}

	if len(opts.ContextFiles) != 0 {
		var sb strings.Builder

		fmt.Fprintf(&sb, "\n")
		fmt.Fprintf(&sb, "Consider the following files:\n")

		for path, file := range opts.ContextFiles {
			fmt.Fprintf(&sb, "File %q:\n", path)
			fmt.Fprintf(&sb, "```go\n")
			fmt.Fprintf(&sb, "%s\n", file.Content)
			fmt.Fprintf(&sb, "```\n")
			fmt.Fprintf(&sb, "\n")
		}

		systemPrompt += sb.String()
	}

	if len(opts.MockGCPFiles) > 0 {
		var sb strings.Builder

		fmt.Fprintf(&sb, "\n")
		fmt.Fprintf(&sb, "For mockGCP file generation in particular use the following files as good examples for code organization and existing patterns that you will follow:\n")
		for _, fileContent := range opts.MockGCPFiles {
			// todo acpana separate out by files and filepath
			// fmt.Fprintf(&sb, "File %q:\n", path)
			fmt.Fprintf(&sb, "%s\n", fileContent)
			fmt.Fprintf(&sb, "\n")
		}

		systemPrompt += sb.String()
	}

	// separeating APIs context load for custom instruction sets
	if len(opts.APIsFiles) > 0 {
		var sb strings.Builder

		fmt.Fprintf(&sb, "\n")
		fmt.Fprintf(&sb, "For API file generation in particular use the following files as good examples for code organization and existing patterns that you will follow:\n")
		for _, fileContent := range opts.APIsFiles {
			// todo acpana separate out by files and filepath
			// fmt.Fprintf(&sb, "File %q:\n", path)
			fmt.Fprintf(&sb, "%s\n", fileContent)
			fmt.Fprintf(&sb, "\n")
		}
		fmt.Fprintf(&sb, "DO NOT GENERATE deepcopy files.")

		systemPrompt += sb.String()
	}

	if len(opts.Rules) > 0 {
		systemPrompt = rules.ApplyRules(systemPrompt, opts.Rules)
	}

	session := llmClient.StartChat(systemPrompt)

	session.SetFunctionDefinitions(functionDefinitions)

	return &Chat{
		client:  llmClient,
		baseDir: opts.BaseDir,
		session: session,
		ui:      ui,
	}, nil

}

func (c *Chat) Close() error {
	return c.client.Close()
}

func (c *Chat) SendMessage(ctx context.Context, userParts ...string) error {
	resp, err := c.session.SendMessage(ctx, userParts...)
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

		var functionResponses []llm.FunctionCallResult

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
					result, err := c.runFunctionCall(ctx, functionCall)
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
					functionResponses = append(functionResponses, llm.FunctionCallResult{
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
		resp, err = c.session.SendFunctionResults(ctx, functionResponses)
		if err != nil {
			return fmt.Errorf("sending message to LLM: %w", err)
		}
	}
}
