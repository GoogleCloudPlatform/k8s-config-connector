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

	"github.com/GoogleCloudPlatform/kubectl-ai/gollm"
)

func (u *FieldInserter) insertGoMessagesGemini(llmClient gollm.Client, model string) error {
	ctx := context.Background()

	// Start new chat session
	systemPrompt := "" // TODO
	session := llmClient.StartChat(systemPrompt, model)
	var userParts []any
	userParts = append(userParts, `
						I have some Go structs written in Go files under a directory.
						I will provide you the filename and content of each Go file under the directory.
						I will also provide the content of the Go structs I want to insert to the files.
						Could you insert the Go structs into one of the Go file?
						Insert the Go structs based on alphabetical order of the "+kcc:proto=" comment of the Go structs.
						You should only insert. Do not delete or edit the existing Go structs.
						In your response, only include what is asked for.
					`)

	// provide the content of the Go structs
	goStructs := ""
	for _, s := range u.dependentMessages {
		goStructs += string(s.generatedContent) + "\n"
	}

	userParts = append(userParts, fmt.Sprintf("the Go structs I want to insert:\n%s\n\n", goStructs))

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

	userParts = append(userParts, "What is the content of the modified Go file")
	// get response
	resp, err := session.Send(ctx, userParts...)
	if err != nil {
		return fmt.Errorf("error receiving message: %w", err)
	}
	modifiedContent, err := extractModifiedContent(resp)
	if err != nil {
		return fmt.Errorf("error extracting modified content: %w", err)
	}

	resp, err = session.Send(ctx, "What is the filename that was modified")
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
