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

package ui

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var _ UI = &BashUI{}

type BashUI struct {
	callback func(text string) error
}

func NewBashUI() UI {
	return &BashUI{}
}

func (u *BashUI) Run() error {
	scanner := bufio.NewScanner(os.Stdin)
	totalText := []string{}
	// Read and process lines from stdin
	lastLine := ""
	text := ""
	for scanner.Scan() {
		line := scanner.Text()
		if line == "\n" && lastLine == "\n" {
			totalText = append(totalText, text)
			text = ""
			lastLine = ""
		} else if line != "\n" {
			text += line
		} else {
			lastLine = "\n"
		}
	}
	totalText = append(totalText, text)

	for _, text := range totalText {
		if text == "" {
			continue
		}
		if err := u.callback(text); err != nil {
			return fmt.Errorf("error running callback: %w", err)
		}
	}
	// Check for errors during scanning
	if err := scanner.Err(); err != io.EOF {
		return err
	}
	return nil
}

func (u *BashUI) SetCallback(callback func(text string) error) {
	u.callback = callback
}

func (u *BashUI) AddLLMOutput(output *LLMOutput) {
	fmt.Printf("%v\n", output.Text)
}
