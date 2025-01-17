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

	"cloud.google.com/go/vertexai/genai"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot/ui"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/llm"
	"k8s.io/klog/v2"
)

type Chat struct {
	client  *genai.Client
	session *genai.ChatSession
	baseDir string
	ui      ui.UI
}

type FileInfo struct {
	Path    string
	Content string
}

func NewChat(ctx context.Context, baseDir string, contextFiles map[string]*FileInfo, ui ui.UI) (*Chat, error) {
	gemini, err := llm.BuildGeminiClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("initializing gemini: %w", err)
	}

	model := gemini.GenerativeModel("gemini-2.0-flash-exp")
	// model := gemini.GenerativeModel("gemini-1.5-pro-002")

	// Some values that are recommended by aistudio
	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"

	systemPrompt := `
You are a helpful AI coding assistant, expert in the go programming language and in creating kubernetes controllers.

Help the user with their problems.  Do not do more than the user asks; propose minimal changes to solve what the user wants.

If you think the code should be changed, use the tools to apply those changes instead of printing them.  Do not make multiple parallel changes to the same file.
`

	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(systemPrompt),
		},
	}

	var functionDeclarations []*genai.FunctionDeclaration

	functionDeclarations = append(functionDeclarations, &genai.FunctionDeclaration{
		Name: "create_file",
		Description: `
Create a new file in the user's workspace, with the contents specified by contents.  This tool will fail if the file already exists.
	`,
		Parameters: &genai.Schema{
			Type:     genai.TypeObject,
			Required: []string{"contents", "filename"},
			Properties: map[string]*genai.Schema{
				"contents": {
					Type: genai.TypeString,
					Description: `
The text that should be the contents of the new file.
`,
				},
				"filename": {
					Type:        genai.TypeString,
					Description: "The path to the file you want to create",
				},
			},
		},
	})

	for _, tool := range GetTools() {
		functionDeclarations = append(functionDeclarations, tool.BuildFunctionDefinition())
	}

	// 	functionDeclarations = append(functionDeclarations, &genai.FunctionDeclaration{
	// 		Name: "ast_edit",
	// 		Description: `
	// Makes changes to an existing file in a way that understands the syntax of the file (the AST, or Abstract Syntax Tree).
	// `,
	// 		Parameters: &genai.Schema{
	// 			Type:     genai.TypeObject,
	// 			Required: []string{"filename", "action", "node"},
	// 			Properties: map[string]*genai.Schema{
	// 				"action": {
	// 					Type: genai.TypeString,
	// 					Description: `
	// The action to perform; must be one of ADD, REPLACE, DELETE.

	// For example, to update a method, use REPLACE here and provided the updated method in node.`,
	// 				},
	// 				"node": {
	// 					Type: genai.TypeString,
	// 					Description: `
	// The AST node to ADD, REPLACE or DELETE.  This should be a top level node in the AST of the language you're using, for example a function declaration or a type declaration.
	// `,
	// 				},
	// 				"filename": {
	// 					Type:        genai.TypeString,
	// 					Description: "The path to the file you want to change",
	// 				},
	// 			},
	// 		},
	// 	})

	model.Tools = append(model.Tools, &genai.Tool{
		FunctionDeclarations: functionDeclarations,
	})

	if model.ToolConfig == nil {
		model.ToolConfig = &genai.ToolConfig{}
	}
	if model.ToolConfig.FunctionCallingConfig == nil {
		model.ToolConfig.FunctionCallingConfig = &genai.FunctionCallingConfig{}
	}
	model.ToolConfig.FunctionCallingConfig.Mode = genai.FunctionCallingAuto

	if len(contextFiles) != 0 {
		model.SystemInstruction.Parts = append(model.SystemInstruction.Parts, genai.Text(`
Consider the following files:
`))
	}
	for _, file := range contextFiles {
		var sb strings.Builder

		fmt.Fprintf(&sb, "File %q:\n", file.Path)
		fmt.Fprintf(&sb, "```go\n")
		fmt.Fprintf(&sb, "%s\n", file.Content)
		fmt.Fprintf(&sb, "```\n")

		model.SystemInstruction.Parts = append(model.SystemInstruction.Parts, genai.Text(sb.String()))
	}

	session := model.StartChat()

	return &Chat{
		client:  gemini,
		baseDir: baseDir,
		session: session,
		ui:      ui,
	}, nil

}

func (c *Chat) Close() error {
	return c.client.Close()
}

func (c *Chat) SendMessage(ctx context.Context, userParts ...genai.Part) error {
	for {
		resp, err := c.session.SendMessage(ctx, userParts...)
		if err != nil {
			return fmt.Errorf("sending message to LLM: %w", err)
		}
		// Print the usage metadata (includes token count i.e. cost)
		klog.V(2).Infof("UsageMetadata: %+v", resp.UsageMetadata)

		if len(resp.Candidates) == 0 {
			return fmt.Errorf("no candidates returned")
		}

		if len(resp.Candidates) > 1 {
			// TODO: It is probably worth evaluating all the candidates
			klog.Warningf("LLM returned multiple candidates; ignoring all but the first")
		}

		candidate := resp.Candidates[0]
		klog.V(2).Infof("processing candidate %+v", candidate)
		content := candidate.Content

		var functionResponses []genai.Part

		for _, part := range content.Parts {
			if text, ok := part.(genai.Text); ok {
				s := string(text)
				c.ui.AddLLMOutput(&ui.LLMOutput{Text: s})
				klog.V(2).Infof("TEXT: %+v", s)
			} else if functionCall, ok := part.(genai.FunctionCall); ok {
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
				functionResponses = append(functionResponses, genai.FunctionResponse{
					Name:     functionCall.Name,
					Response: response,
				})
				b, _ := json.Marshal(response)
				c.ui.AddLLMOutput(&ui.LLMOutput{Text: fmt.Sprintf("sending response: %v", string(b))})

			} else {
				klog.Infof("UNKNOWN: %T %+v", part, part)
			}
		}

		if len(functionResponses) == 0 {
			return nil
		}

		// functionResponseContent := genai.Content{
		// 	Role:  "model",
		// 	Parts: functionResponses,
		// }
		userParts = functionResponses
		// c.session.History = append(c.session.History, functionResponseContent)
	}
}
