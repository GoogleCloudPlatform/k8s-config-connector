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
	"os"
)

var _ UI = &TerminalUI{}

type TerminalUI struct {
	callback func(text string) error
}

func NewTerminalUI() *TerminalUI {
	return &TerminalUI{}
}

func (u *TerminalUI) Run() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(">>> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("reading from stdin: %w", err)
		}
		// fmt.Println(text)

		if err := u.callback(text); err != nil {
			return fmt.Errorf("error running callback: %w", err)
		}
	}
}

func (u *TerminalUI) SetCallback(callback func(text string) error) {
	u.callback = callback
}

func (u *TerminalUI) AddLLMOutput(output *LLMOutput) {
	fmt.Fprintf(os.Stdout, "%v\n", output.Text)
}
