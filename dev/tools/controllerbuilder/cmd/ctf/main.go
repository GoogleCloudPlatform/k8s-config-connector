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
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot/repocontext"
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
	klog.InitFlags(nil)
	flag.Parse()

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

	contextFiles := make(map[string]*repocontext.FileInfo)

	{
		// main.go
		s := `
	package main
	
	import (
		"context"
		"fmt"
		"os"
	)
	
	func main() {
	  fmt.Fprintf(os.Stdout, "Hello world\n")
	}
	`

		contextFiles["main.go"] = &repocontext.FileInfo{
			Content: s,
			Path:    "main.go",
		}

		p := filepath.Join(tmpDir, "main.go")
		if err := os.WriteFile(p, []byte(s), 0644); err != nil {
			return fmt.Errorf("writing file %q: %w", p, err)
		}
	}

	{
		// go.mod
		s := `
module mymodule

go 1.21
	`

		contextFiles["go.mod"] = &repocontext.FileInfo{
			Content: s,
			Path:    "go.mod",
		}

		p := filepath.Join(tmpDir, "go.mod")
		if err := os.WriteFile(p, []byte(s), 0644); err != nil {
			return fmt.Errorf("writing file %q: %w", p, err)
		}
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

	chat, err := codebot.NewChat(ctx, llmClient, u, &codebot.Options{BaseDir: tmpDir, ContextFiles: contextFiles})
	if err != nil {
		return err
	}
	defer chat.Close()

	msg := `
I am trying to write a simple go program, and when I run go build I get the following errors:\n
{{stdout}}\n
{{stderr}}\n

Can you fix the problems?
`

	msg = strings.ReplaceAll(msg, "{{stdout}}", buildResults.Stdout)
	msg = strings.ReplaceAll(msg, "{{stderr}}", buildResults.Stderr)
	var userParts []string
	userParts = append(userParts, msg)
	if err := chat.SendMessage(ctx, userParts...); err != nil {
		return err
	}

	p := filepath.Join(tmpDir, "main.go")
	newMain, err := os.ReadFile(p)
	if err != nil {
		return fmt.Errorf("reading file %q: %w", p, err)
	}
	fmt.Fprintf(os.Stdout, "updated main.go is:\n%s", string(newMain))

	buildResults2, err := runGoBuild(ctx, tmpDir)
	if err != nil {
		return err
	}
	log.Info("new build results", "results", buildResults2)

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
