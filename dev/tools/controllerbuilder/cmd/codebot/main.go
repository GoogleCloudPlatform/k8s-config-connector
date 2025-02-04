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

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot/repocontext"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot/rules"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot/ui"
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

type Options struct {
	// ProtoDir is the base directory for the checkout of the proto API definitions
	ProtoDir string
	// BaseDir is the base directory for the project code
	BaseDir string

	// MockGcp base dir
	MockGcpDir string

	// APIsDir for where we have the direct APIs only
	APIsDir string
}

func run(ctx context.Context) error {
	var o Options
	files := []string{}

	klog.InitFlags(nil)

	flag.StringVar(&o.ProtoDir, "proto-dir", o.ProtoDir, "base directory for checkout of proto API definitions")
	flag.StringVar(&o.BaseDir, "base-dir", o.BaseDir, "base directory for the project code")
	flag.StringVar(&o.MockGcpDir, "mockgcp-dir", o.MockGcpDir, "MockGcp base dir")
	flag.StringVar(&o.APIsDir, "apis-dir", o.APIsDir, "APIsDir for where we have the direct APIs only")

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
		//files = append(files, cwdAbs) // todo acpana doing this will make it stutter as /a/b => /a/b/a/b
	}

	files = append(files, flag.Args()...)

	contextFiles := make(map[string]*repocontext.FileInfo)
	for _, file := range files {
		p := filepath.Join(o.BaseDir, file)
		b, err := os.ReadFile(p)
		if err != nil {
			return fmt.Errorf("reading file %q (in %q): %w", file, o.BaseDir, err)
		}
		contextFiles[file] = &repocontext.FileInfo{
			Path:    file,
			Content: string(b),
		}
	}

	llmClient, err := llm.BuildVertexAIClient(ctx)
	if err != nil {
		return fmt.Errorf("initializing LLM: %w", err)
	}

	defer llmClient.Close()

	var chatSession *codebot.Chat

	// ui := ui.NewTViewUI()
	ui := ui.NewTerminalUI()

	ui.SetCallback(func(text string) error {
		var userParts []string

		var additionalContext strings.Builder

		tokens := strings.Split(text, " ")
		for i, token := range tokens {
			if protoEnhancer != nil {
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
		}
		text = additionalContext.String() + strings.Join(tokens, " ")
		userParts = append(userParts, text)

		if err := chatSession.SendMessage(ctx, userParts...); err != nil {
			return fmt.Errorf("generating content with gemini: %w", err)
		}

		return nil
	})

	codebotOpts := &codebot.Options{BaseDir: o.BaseDir, ContextFiles: contextFiles}
	codebotOpts.Rules = []rules.Rule{}
	codebotOpts.MockGCPFiles = map[string]string{}
	codebotOpts.APIsFiles = map[string]string{}

	// if present, gather mocks
	if o.MockGcpDir != "" {
		mocks, err := repocontext.GatherMockContents(o.MockGcpDir, 10)
		if err != nil {
			return err
		}
		codebotOpts.MockGCPFiles["mocks"] = strings.Join(mocks, "\n ==separator== \n")
	}

	// if present, gather apis
	if o.APIsDir != "" {		
		apis, err := repocontext.GatherAPIContents(o.APIsDir, 10)
		if err != nil {
			return err
		}
		codebotOpts.APIsFiles["apis"] = strings.Join(apis, "\n ==separator== \n")
		codebotOpts.Rules = append(codebotOpts.Rules, rules.CRDShortNames)
	}
	codebotOpts.Rules = append(codebotOpts.Rules, rules.NoNakedReturns)

	session, err := codebot.NewChat(ctx, llmClient, ui, codebotOpts)
	if err != nil {
		return err
	}
	chatSession = session
	defer chatSession.Close()

	if err := ui.Run(); err != nil {
		return fmt.Errorf("running tview: %w", err)
	}

	return nil
}
