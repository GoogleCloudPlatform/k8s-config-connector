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
	"bytes"
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot/ui"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/llm"
	"k8s.io/klog/v2"
)

func main() {
	ctx := context.Background()
	err := run(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	scenario := ""
	flag.StringVar(&scenario, "scenario", scenario, "scenario to run")

	klog.InitFlags(nil)
	flag.Parse()

	if scenario == "" {
		return fmt.Errorf("must specify --scenario")
	}

	log := klog.FromContext(ctx)

	tmpDir, err := os.MkdirTemp("", "codebot-ctf")
	if err != nil {
		return fmt.Errorf("creating temp dir: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			klog.Error(err, "error cleaning up temp dir", "path", tmpDir)
		}
	}()

	scenarioDir, err := filepath.Abs(scenario)
	if err != nil {
		return fmt.Errorf("getting absolute path for %q: %w", scenario, err)
	}

	srcDir := filepath.Join(scenarioDir, "src")

	contextFiles := make(map[string]*codebot.FileInfo)

	// Walk the files in srcDir, copy them to tmpDir, and add them to contextFiles
	if err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		relativePath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return fmt.Errorf("getting relative path for %q: %w", path, err)
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading file %q: %w", path, err)
		}

		contextFiles[path] = &codebot.FileInfo{
			Content: string(b),
			Path:    relativePath,
		}

		tmpPath := filepath.Join(tmpDir, relativePath)
		if err := os.MkdirAll(filepath.Dir(tmpPath), 0755); err != nil {
			return fmt.Errorf("creating directory %q: %w", filepath.Dir(tmpPath), err)
		}
		if err := os.WriteFile(tmpPath, b, 0644); err != nil {
			return fmt.Errorf("writing file %q: %w", tmpPath, err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("walking scenario dir %q: %w", scenarioDir, err)
	}
	buildResults, err := runGoBuild(ctx, tmpDir)
	if err != nil {
		return err
	}
	log.Info("go build results", "results", buildResults)

	if buildResults.ExitCode == 0 {
		return fmt.Errorf("expected build error from scenario, but got no error")
	}

	llmClient, err := llm.BuildVertexAIClient(ctx)
	if err != nil {
		return fmt.Errorf("initializing LLM: %w", err)
	}
	defer llmClient.Close()

	u := ui.NewTerminalUI()

	toolbox := codebot.NewToolbox(codebot.GetAllTools())

	chat, err := codebot.NewChat(ctx, llmClient, tmpDir, contextFiles, toolbox, u)
	if err != nil {
		return err
	}
	defer chat.Close()

	msg := `
I am trying to write a simple go program, and when I run go build I get the following errors:\n
{{stdout}}\n
{{stderr}}\n

Can you fix the problems? Make the minimal code changes, do not make any changes not needed to fix the problem with go build.

Use function calling to fix the problems; do not ask me follow-on questions.
`

	msg = strings.ReplaceAll(msg, "{{stdout}}", buildResults.Stdout)
	msg = strings.ReplaceAll(msg, "{{stderr}}", buildResults.Stderr)
	var userParts []string
	userParts = append(userParts, msg)
	if err := chat.SendMessage(ctx, userParts...); err != nil {
		return err
	}

	{
		newMain, err := os.ReadFile(p)
		if err != nil {
			return fmt.Errorf("reading file %q: %w", p, err)
	}
	log.Info("new build results", "results", buildResults2)

	if _, err := runGoFormat(ctx, tmpDir); err != nil {
		return err
	}

	{
		p := filepath.Join(tmpDir, "main.go")
		newMain, err := os.ReadFile(p)
		if err != nil {
			return fmt.Errorf("reading file %q: %w", p, err)
		}
		fmt.Fprintf(os.Stdout, "final main.go is:\n%s", string(newMain))
	}

	if buildResults2.ExitCode == 0 {
		fmt.Printf("SUCCESS\n")
	} else {
		fmt.Printf("FAIL\n")
	}

	return nil
}

type GoBuildResults struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

func runGoBuild(ctx context.Context, dir string) (*GoBuildResults, error) {
	results := &GoBuildResults{}

	buildCmd := exec.CommandContext(ctx, "go", "build", ".")
	buildCmd.Dir = dir
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	buildCmd.Stdout = &stdout
	buildCmd.Stderr = &stderr

	err := buildCmd.Run()
	if err != nil {
		switch err := err.(type) {
		case *exec.ExitError:
			results.ExitCode = err.ExitCode()
		default:
			return nil, fmt.Errorf("unexpected error running %q: %w", strings.Join(buildCmd.Args, " "), err)
		}
	}

	results.Stdout = stdout.String()
	results.Stderr = stderr.String()

	return results, nil
}

type GoFormatResults struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

func runGoFormat(ctx context.Context, dir string) (*GoFormatResults, error) {
	results := &GoFormatResults{}

	buildCmd := exec.CommandContext(ctx, "gofmt", "-w", ".")
	buildCmd.Dir = dir
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	buildCmd.Stdout = &stdout
	buildCmd.Stderr = &stderr

	err := buildCmd.Run()
	if err != nil {
		switch err := err.(type) {
		case *exec.ExitError:
			results.ExitCode = err.ExitCode()
		default:
			return nil, fmt.Errorf("unexpected error running %q: %w", strings.Join(buildCmd.Args, " "), err)
		}
	}

	results.Stdout = stdout.String()
	results.Stderr = stderr.String()

	klog.Infof("gofmt results: %+v", results)

	return results, nil
}
