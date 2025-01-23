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
	"context"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"k8s.io/klog/v2"
)

func NewTViewUI(interactive bool) UI {
	flex := tview.NewFlex()
	flex.SetBorder(true).SetTitle("Hello, world!")
	flex.SetDirection(tview.FlexRow)

	app := tview.NewApplication()
	app.SetRoot(flex, true)
	app.EnableMouse(true)
	app.EnablePaste(true)

	ui := &TViewUI{flex: flex, app: app}

	if interactive {
		inputField := tview.NewInputField().SetLabel("Enter text:")
		inputField.SetDoneFunc(ui.onInputFieldDone)
		ui.inputField = inputField

		inputField.SetText("Can you write hello world in go?")

		ui.flex.AddItem(inputField, 1, 0, true)
		app.SetFocus(inputField)
	}

	return ui
}

type TViewUI struct {
	flex       *tview.Flex
	app        *tview.Application
	inputField *tview.InputField
	callback   func(text string) error
}

func (u *TViewUI) onInputFieldDone(key tcell.Key) {
	inputField := u.inputField
	if key == tcell.KeyEnter {
		text := inputField.GetText()
		u.flex.RemoveItem(inputField)
		u.flex.AddItem(tview.NewTextView().SetText(text), 1, 0, false)
		go func() {
			if err := u.callback(text); err != nil {
				klog.Errorf("error running callback: %v", err)
				u.app.QueueUpdateDraw(func() {
					u.flex.AddItem(tview.NewTextView().SetText(fmt.Sprintf("Error: %v", err)), 1, 0, false)
					u.flex.AddItem(u.inputField, 1, 0, false)
				})
			}
		}()
	}
}

func (u *TViewUI) Run(ctx context.Context) error {
	return u.app.Run()
}

func (u *TViewUI) SetCallback(callback func(text string) error) {
	u.callback = callback
}

func (u *TViewUI) AddLLMOutput(output *LLMOutput) {
	u.app.QueueUpdateDraw(func() {
		u.flex.AddItem(tview.NewTextView().SetText(output.Text), 1, 0, false)
		if u.inputField != nil {
			u.flex.AddItem(u.inputField, 1, 0, false)
		}
	})
}
