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
	"fmt"
	"strings"

	"cloud.google.com/go/vertexai/genai"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/llm"
	"k8s.io/klog"
)

type Chat struct {
	client  *genai.Client
	session *genai.ChatSession
	baseDir string
}

type FileInfo struct {
	Path    string
	Content string
}

func NewChat(ctx context.Context, baseDir string, contextFiles map[string]*FileInfo) (*Chat, error) {
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

	// edit_file seems less robust than ast_edit; it keeps only doing a single-line find...
	// TODO: Try giving an example?
	functionDeclarations = append(functionDeclarations, &genai.FunctionDeclaration{
		Name: "edit_file",
		Description: `
Make a change to an existing file in the user's workspace, by replacing existing_text with new_text.  This tool only applies the first replacement.
	`,
		Parameters: &genai.Schema{
			Type:     genai.TypeObject,
			Required: []string{"existing_text", "new_text", "filename"},
			Properties: map[string]*genai.Schema{
				"existing_text": {
					Type: genai.TypeString,
					Description: `
The text to find, which will be replaced with the contents of the new_text argument.  Provide all the lines of the existing content you want to replace.
`,
				},
				"new_text": {
					Type: genai.TypeString,
					Description: `
The text that should replace the contents of the existing_text argument.
`,
				},
				"filename": {
					Type:        genai.TypeString,
					Description: "The path to the file you want to change",
				},
			},
		},
	})

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
	}, nil

}

func (c *Chat) Close() error {
	return c.client.Close()
}

func (c *Chat) SendMessage(ctx context.Context, userParts ...genai.Part) error {
	resp, err := c.session.SendMessage(ctx, userParts...)
	if err != nil {
		return fmt.Errorf("sending message to LLM: %w", err)
	}
	// Print the usage metadata (includes token count i.e. cost)
	klog.Infof("UsageMetadata: %+v", resp.UsageMetadata)

	for _, candidate := range resp.Candidates {
		klog.Infof("processing candidate %+v", candidate)
		content := candidate.Content

		var functionResponses []genai.Part

		for _, part := range content.Parts {
			if text, ok := part.(genai.Text); ok {
				s := string(text)
				klog.Infof("TEXT: %+v", s)

				// diffs don't seem to apply very reliably (????)
				// if strings.Contains(s, "```diff") {
				// 	inDiff := false
				// 	var diff bytes.Buffer
				// 	for _, line := range strings.Split(s, "\n") {
				// 		if strings.HasPrefix(line, "```diff") {
				// 			inDiff = true
				// 		} else if strings.HasPrefix(line, "```") {
				// 			inDiff = false
				// 		} else if inDiff {
				// 			diff.WriteString(line)
				// 			diff.WriteString("\n")
				// 		}
				// 	}

				// 	cmd := exec.CommandContext(ctx, "patch", "-p1")
				// 	cmd.Stdin = &diff
				// 	cmd.Stdout = os.Stdout
				// 	cmd.Stderr = os.Stderr
				// 	if err := cmd.Run(); err != nil {
				// 		fmt.Fprintf(os.Stderr, "PATCH DID NOT APPLY: %v\n", err)
				// 	} else {
				// 		fmt.Fprintf(os.Stdout, "Applied patch :-)\n")
				// 	}
				// }
			} else if functionCall, ok := part.(genai.FunctionCall); ok {
				// klog.Infof("functionCall: %+v", functionCall)
				klog.Infof("functionCall: %+v", functionCall.Name)
				response, err := c.runFunctionCall(ctx, functionCall)
				if err != nil {
					return fmt.Errorf("unexpected error running function: %w", err)
				}
				functionResponses = append(functionResponses, genai.FunctionResponse{
					Name:     functionCall.Name,
					Response: response,
				})

			} else {
				klog.Infof("UNKNOWN: %T %+v", part, part)
			}
		}

		if len(functionResponses) != 0 {
			c.session.History = append(c.session.History, &genai.Content{
				Role:  "model",
				Parts: functionResponses,
			})
		}

		break
	}
	return nil
}
