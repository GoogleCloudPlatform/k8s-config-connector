// Copyright 2024 Google LLC
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

package typeupdater

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/kubectl-ai/gollm"
	"k8s.io/klog/v2"
)

func (u *FieldInserter) insertGoFieldGemini(llmClient gollm.Client, model string) error {
	klog.Infof("inserting the generated Go code for field %s", u.newField.proto.Name())
	ctx := context.Background()

	// Start new chat session
	systemPrompt := "" // TODO
	chat := llmClient.StartChat(systemPrompt, model)

	var userParts []any

	userParts = append(userParts, fmt.Sprintf(`
						I have some Go structs written in Go files under a directory.
						I will provide you the filename and content of each Go file under the directory.
						I will also provide the content of the new Go field.
						Could you find the Go struct which has comment "+kcc:proto=%s" with no following suffix,
						and insert the Go field into the found Go struct.
						In your response, only include what is asked for.
					`, u.newField.parent.FullName()))

	// provide the content of the new Go field
	userParts = append(userParts, fmt.Sprintf("new Go field:\n%s\n\n", u.newField.generatedContent))

	// provide content of the existing Go files
	files, err := listFiles(u.opts.APIDirectory)
	if err != nil {
		return fmt.Errorf("error listing files: %w", err)
	}
	for _, f := range files {
		content, err := readFile(f)
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}
		userParts = append(userParts, fmt.Sprintf("filename:\n%s\ncontent:\n%s\n\n", f, content))
	}

	// TODO: provide content of the proto message for reference on the field ordering.

	userParts = append(userParts, "What is the content of the modified Go file")

	// get response
	resp, err := chat.Send(ctx, userParts...)
	if err != nil {
		return fmt.Errorf("error receiving message: %w", err)
	}
	modifiedContent, err := extractModifiedContent(resp)
	if err != nil {
		return fmt.Errorf("error extracting modified content: %w", err)
	}

	resp, err = chat.Send(ctx, "What is the filename that was modified")
	if err != nil {
		return fmt.Errorf("error receiving message: %w", err)
	}
	modifiedFilename, err := extractModifiedFilename(resp)
	if err != nil {
		return fmt.Errorf("error extracting modified content: %w", err)
	}

	if err := writeFile(modifiedFilename, modifiedContent); err != nil {
		return fmt.Errorf("error writing file %s: %w", modifiedFilename, err)
	}

	return nil
}

func listFiles(directory string) ([]string, error) {
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func readFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func writeFile(filePath, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}

// extract the modified content of the file from the unstructured text response.
func extractModifiedContent(resp gollm.ChatResponse) (string, error) {
	candidates := resp.Candidates()
	if len(candidates) == 0 {
		return "", fmt.Errorf("no candidates found in response")
	}
	text, ok := candidates[0].Parts()[0].AsText()
	if !ok {
		return "", fmt.Errorf("unexpected response type %T", candidates[0].Parts()[0])
	}
	str := string(text)
	str = strings.TrimPrefix(str, "```go")
	str = strings.TrimSuffix(str, "```")
	str = strings.TrimLeft(str, "\n")  // trim all leading "\n"
	str = strings.TrimRight(str, "\n") // trim all trailing "\n"
	str = strings.TrimRight(str, " ")
	str += "\n" // add back the final "\n" at end of file
	return str, nil
}

// extract the name of the modified file from the unstructured text response.
func extractModifiedFilename(resp gollm.ChatResponse) (string, error) {
	candidates := resp.Candidates()
	if len(candidates) == 0 {
		return "", fmt.Errorf("no candidates found in response")
	}
	text, ok := candidates[0].Parts()[0].AsText()
	if !ok {
		return "", fmt.Errorf("unexpected response type %T", candidates[0].Parts()[0])
	}
	str := string(text)
	str = strings.TrimPrefix(str, "```go")
	str = strings.TrimSuffix(str, "```")
	str = strings.TrimRight(str, "\n")
	str = strings.TrimRight(str, " ")
	return str, nil
}
