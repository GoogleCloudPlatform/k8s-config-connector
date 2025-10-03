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
	"strings"

	"k8s.io/klog/v2"
)

type TerminalUI struct {
	callback func(text string) error
}

func NewTerminalUI() UI {
	return &TerminalUI{}
}

func (u *TerminalUI) Run() error {
	reader := bufio.NewReader(os.Stdin)
	var text strings.Builder
	lastLine := "\n"
	for {
		if text.Len() == 0 {
			fmt.Printf(">>> ")
		} else {
			fmt.Printf("... ")
		}
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return fmt.Errorf("reading from stdin: %w", err)
			}
		}
		text.WriteString(line)

		// Only process if we have two empty lines in a row
		// This allows users to enter multiline prompts
		// (at the expense of requiring them to hit enter three times!)
		if err == io.EOF || (line == "\n" && lastLine == "\n") {
			if text.String() == "" {
				if err == io.EOF {
					return nil
				} else {
					fmt.Printf("I am but an LLM, I need instruction\n")
					text.Reset()
					continue
				}
			}
			klog.V(2).Infof("sending text: %s", text.String())
			if err := u.callback(text.String()); err != nil {
				return fmt.Errorf("error running callback: %w", err)
			}
			text.Reset()
		}
		lastLine = line
	}
}

func (u *TerminalUI) SetCallback(callback func(text string) error) {
	u.callback = callback
}

func (u *TerminalUI) AddLLMOutput(output *LLMOutput) {
	fmt.Fprintf(os.Stdout, "%v\n", output.Text)
}
