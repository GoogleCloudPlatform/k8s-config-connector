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
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	codebotui "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot/ui"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/llm"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/toolbot"
	"k8s.io/klog/v2"
)

func main() {
	ctx := context.Background()
	codebot := &CodeBot{
		ctx: ctx,
	}
	if err := codebot.run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type Options struct {
	llm.Options

	// ProtoDir is the base directory for the checkout of the proto API definitions
	ProtoDir string
	// BaseDir is the base directory for the project code
	BaseDir string
	// Prompt is the prompt to be passed in non-interactive mode
	Prompt string

	UIType string
}

type CodeBot struct {
	ctx           context.Context
	protoEnhancer *toolbot.EnhanceWithProtoDefinition
	chatSession   *codebot.Chat
}

func (cb *CodeBot) run(ctx context.Context) error {
	var o Options
	o.InitDefaults()

	klog.InitFlags(nil)

	flag.StringVar(&o.ProtoDir, "proto-dir", o.ProtoDir, "base directory for checkout of proto API definitions")
	flag.StringVar(&o.BaseDir, "base-dir", o.BaseDir, "base directory for the project code")
	flag.StringVar(&o.Prompt, "prompt", o.Prompt, "prompt to be passed in non-interactive mode")
	flag.StringVar(&o.UIType, "ui-type", o.UIType, "available value is terminal, tview, prompt or bash.")
	o.AddFlags(flag.CommandLine)
	flag.Parse()

	if o.ProtoDir == "" {
		klog.Warningf("proto-dir not set; protobuf assistance will be disabled")
	}
	var protoEnhancer *toolbot.EnhanceWithProtoDefinition
	if o.ProtoDir != "" {
		enhancer, err := toolbot.NewEnhanceWithProtoDefinition(o.ProtoDir)
		if err != nil {
			return fmt.Errorf("loading proto definitions: %w", err)
		}
		protoEnhancer = enhancer
	}
	cb.protoEnhancer = protoEnhancer

	if o.BaseDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("getting current working directory: %w", err)
		}
		cwdAbs, err := filepath.Abs(cwd)
		if err != nil {
			return fmt.Errorf("getting absolute path for current working directory %q: %w", cwd, err)
		}
		o.BaseDir = cwdAbs
	}

	files := flag.Args()

	contextFiles := make(map[string]*codebot.FileInfo)
	for _, file := range files {
		p := filepath.Join(o.BaseDir, file)
		b, err := os.ReadFile(p)
		if err != nil {
			return fmt.Errorf("reading file %q (in %q): %w", file, o.BaseDir, err)
		}
		contextFiles[file] = &codebot.FileInfo{
			Path:    file,
			Content: string(b),
		}
	}

	llmClient, err := o.NewLLMClient(ctx)
	if err != nil {
		return fmt.Errorf("initializing LLM: %w", err)
	}

	defer llmClient.Close()

	toolbox := codebot.NewToolbox(codebot.GetAllTools())

	var ui codebotui.UI
	switch o.UIType {
	case "terminal":
		ui = codebotui.NewTerminalUI()
	case "tview":
		ui = codebotui.NewTViewUI()
	case "bash":
		ui = codebotui.NewBashUI()
	case "prompt":
		ui = codebotui.NewNoInteractTerminal(o.Prompt)
	default:
		ui = codebotui.NewTerminalUI()
	}

	ui.SetCallback(cb.sendToLlm)

	session, err := codebot.NewChat(ctx, llmClient, o.Model, o.BaseDir, contextFiles, toolbox, ui)
	if err != nil {
		return err
	}
	cb.chatSession = session
	defer session.Close()

	if err := ui.Run(); err != nil {
		return fmt.Errorf("running tview: %w", err)
	}

	return nil
}

func (cb *CodeBot) sendToLlm(text string) error {
	var userParts []any

	var additionalContext strings.Builder

	tokens := strings.Split(text, " ")
	for i, token := range tokens {
		if cb.protoEnhancer != nil {
			if strings.HasPrefix(token, "@proto.service:") {
				tokens[i] = ""
				v := strings.TrimPrefix(token, "@proto.service:")
				dataPoint := &toolbot.DataPoint{}
				dataPoint.SetInput("proto.service", v)
				if err := cb.protoEnhancer.EnhanceDataPoint(cb.ctx, dataPoint); err != nil {
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
				if err := cb.protoEnhancer.EnhanceDataPoint(cb.ctx, dataPoint); err != nil {
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
	}
	text = additionalContext.String() + strings.Join(tokens, " ")
	userParts = append(userParts, text)

	if err := cb.chatSession.SendMessage(cb.ctx, userParts...); err != nil {
		return fmt.Errorf("generating content with LLM: %w", err)
	}

	return nil
}
