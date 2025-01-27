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
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	userScheme  = colorScheme{foreground: tcell.ColorBlue}
	robotScheme = colorScheme{foreground: tcell.ColorWhite}
	errorScheme = colorScheme{foreground: tcell.ColorRed}
)

func NewTViewUI(prompt string) UI {
	flex := tview.NewFlex()
	flex.SetBorder(true).SetTitle("codebot")
	flex.SetDirection(tview.FlexRow)

	app := tview.NewApplication()
	app.SetRoot(flex, true)
	app.EnableMouse(true)
	app.EnablePaste(true)

	interactive := prompt == ""
	ui := &TViewUI{
		flex:        flex,
		app:         app,
		interactive: interactive,
		prompt:      prompt,
	}

	if interactive {
		inputField := tview.NewInputField().SetLabel("Enter text:")
		inputField.SetDoneFunc(ui.onInputFieldDone)
		ui.inputField = inputField

		// inputField.SetText("Can you write hello world in go?")

		ui.flex.AddItem(inputField, 1, 0, true)
		app.SetFocus(inputField)
	}

	return ui
}

type TViewUI struct {
	flex        *tview.Flex
	app         *tview.Application
	inputField  *tview.InputField
	callback    func(text string) error
	prompt      string
	interactive bool
}

type colorScheme struct {
	background tcell.Color
	foreground tcell.Color
}

func (u *TViewUI) onInputFieldDone(key tcell.Key) {
	inputField := u.inputField
	if key == tcell.KeyEnter {
		text := inputField.GetText()
		u.flex.RemoveItem(inputField)
		u.addMessage(text, userScheme)
		go func() {
			if err := u.callback(text); err != nil {
				//klog.Errorf("error running callback: %v", err)
				u.app.QueueUpdateDraw(func() {
					u.addMessage(fmt.Sprintf("Error: %v", err), errorScheme)
					u.flex.AddItem(u.inputField, 1, 0, false)
				})
			}
		}()
	}
}

func measureText(s string) int {
	height := 0
	for _, line := range strings.Split(s, "\n") {
		height++
		if len(line) > 80 {
			height++
		}
	}
	return height
}

func (u *TViewUI) addMessage(msg string, colors colorScheme) {
	height := measureText(msg)
	// [<foreground>:<background>:<attribute flags>:<url>]
	foreground := colors.foreground.CSS()
	background := colors.background.CSS()
	text := fmt.Sprintf("[%s:%s]%s", foreground, background, msg)
	u.flex.AddItem(tview.NewTextView().SetText(text).SetDynamicColors(true), height, 0, false)
}

func (u *TViewUI) Run(ctx context.Context) error {
	go func() {
		if !u.interactive {
			text := u.prompt
			u.app.QueueUpdateDraw(func() {
				u.addMessage(text, userScheme)
			})
			if err := u.callback(text); err != nil {
				u.app.QueueUpdateDraw(func() {
					u.addMessage(fmt.Sprintf("Error: %v", err), errorScheme)
				})
			}
		}
	}()
	return u.app.Run()
}

func (u *TViewUI) SetCallback(callback func(text string) error) {
	u.callback = callback
}

func (u *TViewUI) AddLLMOutput(output *LLMOutput) {
	u.app.QueueUpdateDraw(func() {
		u.addMessage(output.Text, robotScheme)
		if u.inputField != nil {
			u.flex.AddItem(u.inputField, 1, 0, false)
		}
	})
}
