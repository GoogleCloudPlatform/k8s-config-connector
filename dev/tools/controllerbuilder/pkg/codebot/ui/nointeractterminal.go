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
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"k8s.io/klog/v2"
)

type NoInteractTerminal struct {
	promptFile string
	callback   func(text string) error
}

func NewNoInteractTerminal(promptFile string) UI {
	return &NoInteractTerminal{
		promptFile: promptFile,
	}
}

func (u *NoInteractTerminal) Run() error {
	file, err := os.Open(u.promptFile) // For read access.
	if err != nil {
		return fmt.Errorf("opening prompt file: %w", err)
	}
	defer file.Close()

	promptb, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("reading prompt file: %w", err)
	}
	prompt := string(promptb)
	// for {
	klog.V(2).Infof("sending text: %s", prompt)
	if err := u.callback(prompt); err != nil {
		return fmt.Errorf("error running callback: %w", err)
	}
	klog.V(2).Infof("returned from callback: %s", prompt)
	time.Sleep(8 * time.Second) // Trying to keep usage below 10 QPM
	// }
	return nil
}

func (u *NoInteractTerminal) SetCallback(callback func(text string) error) {
	u.callback = callback
}

func (u *NoInteractTerminal) AddLLMOutput(output *LLMOutput) {
	fmt.Fprintf(os.Stdout, "%v\n", output.Text)
}
