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

package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/vertexai/genai"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/llm"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/toolbot"
	"k8s.io/klog/v2"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type FileInfo struct {
	Path    string
	Content string
}

type Options struct {
	ProtoDir string
}

func run(ctx context.Context) error {
	var o Options

	klog.InitFlags(nil)

	flag.StringVar(&o.ProtoDir, "proto-dir", o.ProtoDir, "base directory for checkout of proto API definitions")
	flag.Parse()

	gemini, err := llm.BuildGeminiClient(ctx)
	if err != nil {
		return fmt.Errorf("initializing gemini: %w", err)
	}
	defer gemini.Close()

	if o.ProtoDir == "" {
		return fmt.Errorf("proto-dir is required")
	}
	protoEnhancer, err := toolbot.NewEnhanceWithProtoDefinition(o.ProtoDir)
	if err != nil {
		return fmt.Errorf("loading proto definitions: %w", err)
	}

	files := flag.Args()

	contextFiles := make(map[string]FileInfo)
	for _, file := range files {
		b, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("reading file %q: %w", file, err)
		}
		contextFiles[file] = FileInfo{
			Path:    file,
			Content: string(b),
		}
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
Make a change to an existing file, by replacing existing text with new text.  This function only applies the first replacement.
	`,
		Parameters: &genai.Schema{
			Type:     genai.TypeObject,
			Required: []string{"find", "replace", "filename"},
			Properties: map[string]*genai.Schema{
				"find": {
					Type: genai.TypeString,
					Description: `
The text to find, which will be replaced with the contents of the replace argument.
`,
				},
				"replace": {
					Type: genai.TypeString,
					Description: `
The text to use to replace the contents of the find argument.
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
		Name: "ast_edit",
		Description: `
Makes changes to an existing file in a way that understands the syntax of the file (the AST, or Abstract Syntax Tree).
`,
		Parameters: &genai.Schema{
			Type:     genai.TypeObject,
			Required: []string{"filename", "action", "node"},
			Properties: map[string]*genai.Schema{
				"action": {
					Type: genai.TypeString,
					Description: `
The action to perform; must be one of ADD, REPLACE, DELETE.

For example, to update a method, use REPLACE here and provided the updated method in node.`,
				},
				"node": {
					Type: genai.TypeString,
					Description: `
The AST node to ADD, REPLACE or DELETE.  This should be a top level node in the AST of the language you're using, for example a function declaration or a type declaration.
`,
				},
				"filename": {
					Type:        genai.TypeString,
					Description: "The path to the file you want to change",
				},
			},
		},
	})

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

	chatSession := model.StartChat()

	reader := bufio.NewReader(os.Stdin)
	for {
		var userParts []genai.Part
		fmt.Printf(">>> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("reading from stdin: %w", err)
		}
		// fmt.Println(text)

		var additionalContext strings.Builder

		tokens := strings.Split(text, " ")
		for i, token := range tokens {
			if strings.HasPrefix(token, "@proto.service:") {
				tokens[i] = ""
				v := strings.TrimPrefix(token, "@proto.service:")
				dataPoint := &toolbot.DataPoint{}
				dataPoint.SetInput("proto.service", v)
				if err := protoEnhancer.EnhanceDataPoint(ctx, dataPoint); err != nil {
					return fmt.Errorf("error getting proto service definition: %w", err)
				}
				def := dataPoint.Input["proto.service.definition"]
				if def == "" {
					return fmt.Errorf("proto service definition for %q was empty", v)
				}
				fmt.Fprintf(&additionalContext, "Protobuf service definition for %s:\n", v)
				fmt.Fprintf(&additionalContext, "```proto")
				fmt.Fprintf(&additionalContext, "%v", def)
				fmt.Fprintf(&additionalContext, "```")
				fmt.Fprintf(&additionalContext, "---\n")
			}
			if strings.HasPrefix(token, "@proto.message:") {
				tokens[i] = ""
				v := strings.TrimPrefix(token, "@proto.message:")
				dataPoint := &toolbot.DataPoint{}
				dataPoint.SetInput("proto.message", v)
				if err := protoEnhancer.EnhanceDataPoint(ctx, dataPoint); err != nil {
					return fmt.Errorf("error getting proto message definition: %w", err)
				}
				def := dataPoint.Input["proto.message.definition"]
				if def == "" {
					return fmt.Errorf("proto message definition for %q was empty", v)
				}
				fmt.Fprintf(&additionalContext, "Protobuf message definition for %s:\n", v)
				fmt.Fprintf(&additionalContext, "```proto")
				fmt.Fprintf(&additionalContext, "%v", def)
				fmt.Fprintf(&additionalContext, "```")
				fmt.Fprintf(&additionalContext, "---\n")
			}
		}
		text = additionalContext.String() + strings.Join(tokens, " ")
		userParts = append(userParts, genai.Text(text))

		resp, err := chatSession.SendMessage(ctx, userParts...)
		if err != nil {
			return fmt.Errorf("generating content with gemini: %w", err)
		}

		// Print the usage metadata (includes token count i.e. cost)
		klog.Infof("UsageMetadata: %+v", resp.UsageMetadata)

		for _, candidate := range resp.Candidates {
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
					response, err := runFunctionCall(ctx, functionCall)
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
				chatSession.History = append(chatSession.History, &genai.Content{
					Role:  "model",
					Parts: functionResponses,
				})
			}

			break
		}
	}
}

func runFunctionCall(ctx context.Context, functionCall genai.FunctionCall) (map[string]any, error) {
	switch functionCall.Name {
	case "edit_file":
		result, err := runEditFile(ctx, functionCall.Args)
		if err != nil {
			result = make(map[string]any)
			result["result"] = "error"
			result["error"] = fmt.Sprintf("%v", err)
			return result, nil
		}
		return result, nil
	case "ast_edit":
		result, err := runASTEdit(ctx, functionCall.Args)
		if err != nil {
			result = make(map[string]any)
			result["result"] = "error"
			result["error"] = fmt.Sprintf("%v", err)
			return result, nil
		}
		return result, nil
	default:
		// TODO: Fatal or return an error?
		return nil, fmt.Errorf("unknown function %q", functionCall.Name)
	}
}

type EditFile struct {
	Find     string `json:"find"`
	Replace  string `json:"replace"`
	Filename string `json:"filename"`
}

func runEditFile(ctx context.Context, args map[string]any) (map[string]any, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	var editFile EditFile
	if err := json.Unmarshal(b, &editFile); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", &editFile, err)
	}

	result := make(map[string]any)

	klog.Infof("EditFile: %+v", editFile)

	fileContents, err := os.ReadFile(editFile.Filename)
	if err != nil {
		return nil, fmt.Errorf("reading file %q: %w", editFile.Filename, err)
	}

	if editFile.Find == "" {
		return nil, fmt.Errorf("the find argument is requiremnt")
	}

	ix := bytes.Index(fileContents, []byte(editFile.Find))
	if ix == -1 {
		return nil, fmt.Errorf("could not find the `find` string %q in the file %q", editFile.Find, editFile.Filename)
	}

	newContents := bytes.Replace(fileContents, []byte(editFile.Find), []byte(editFile.Replace), 1)
	if err := os.WriteFile(editFile.Filename, newContents, 0644); err != nil {
		return nil, fmt.Errorf("writing file %q: %w", editFile.Filename, err)
	}

	result["result"] = "success"
	return result, nil
}

type ASTEdit struct {
	Node     string `json:"node"`
	Action   string `json:"action"`
	Filename string `json:"filename"`
}

func runASTEdit(ctx context.Context, args map[string]any) (map[string]any, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	var astEdit ASTEdit
	if err := json.Unmarshal(b, &astEdit); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", &astEdit, err)
	}

	result := make(map[string]any)

	klog.Infof("ASTEdit: %+v", astEdit)

	fileContentsBytes, err := os.ReadFile(astEdit.Filename)
	if err != nil {
		return nil, fmt.Errorf("reading file %q: %w", astEdit.Filename, err)
	}

	fileContents := string(fileContentsBytes)

	if astEdit.Action == "" {
		return nil, fmt.Errorf("the action argument is requiremnt")
	}

	if astEdit.Node == "" {
		return nil, fmt.Errorf("the node argument is requiremnt")
	}

	switch astEdit.Action {
	case "ADD":
		fileContents = fileContents + "\n" + astEdit.Node

	case "REPLACE":
		klog.Fatalf("REPLACE node not yet implemented: %+v", astEdit)

	case "DELETE":
		klog.Fatalf("DELETE node not yet implemented: %+v", astEdit)

	default:
		return nil, fmt.Errorf("the action %q is not known; it should be ADD, REPLACE or DELETE", astEdit.Action)
	}
	if err := os.WriteFile(astEdit.Filename, []byte(fileContents), 0644); err != nil {
		return nil, fmt.Errorf("writing file %q: %w", astEdit.Filename, err)
	}

	result["result"] = "success"
	return result, nil
}
