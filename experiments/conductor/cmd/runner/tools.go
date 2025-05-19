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

package runner

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

func BuildToolsCommand() *cobra.Command {
	var opts ToolOptions

	cmd := &cobra.Command{
		Use:   "tools",
		Short: "Run a single tool.",
	}

	cmd.PersistentFlags().BoolVar(&opts.Force, "force", opts.Force, "Force operation even if files already exist.")
	cmd.PersistentFlags().BoolVar(&opts.Verbose, "verbose", opts.Verbose, "Enable verbose output logging")
	cmd.PersistentFlags().StringVar(&opts.RepoRoot, "repo-root", opts.RepoRoot, "Directory in which to do the work.")

	cmd.AddCommand(BuildMockgcpGenerateGcloudTestCommand(&opts))

	return cmd
}

func BuildMockgcpGenerateGcloudTestCommand(ToolOptions *ToolOptions) *cobra.Command {
	var opts MockgcpGenerateGcloudTestOptions
	opts.ToolOptions = ToolOptions
	cmd := &cobra.Command{
		Use:   "mockgcp-generate-gcloud-test",
		Short: "Create a mockgcp test that uses gcloud.",
		RunE: func(cmd *cobra.Command, args []string) error {
			out := cmd.OutOrStdout()

			results, err := RunMockgcpGenerateGcloudTest(cmd.Context(), opts)
			if err != nil {
				return err
			}
			fmt.Fprintf(out, "%d affectedPaths:\n", len(results.AffectedPaths))
			for _, affectedPath := range results.AffectedPaths {
				fmt.Fprintf(out, "  %s\n", affectedPath)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&opts.GcloudCommand, "gcloud-command", opts.GcloudCommand, "The gcloud command to generate")
	cmd.Flags().StringVar(&opts.Resource, "resource", opts.Resource, "The resource name we are generating")
	cmd.Flags().StringVar(&opts.Group, "group", opts.Group, "The resource group we are generating")

	return cmd
}

// ToolOptions are the common options for a single tool.
type ToolOptions struct {
	// ScratchDir is the directory to use for temporary files (that we preserve for diagnostics), including log files.
	// We default this directory if not specified.
	ScratchDir string

	// RepoRoot is the root of the KCC source code repo.
	RepoRoot string

	// Force indicates if we should re-execute, even if files already exists.
	Force bool

	// Verbose indicates if we should show verbose output
	Verbose bool
}

func (o *ToolOptions) DefaultAndValidate() error {
	if o.ScratchDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("getting working directory: %w", err)
		}
		o.ScratchDir = filepath.Join(cwd, "conductor", "logs", time.Now().Format(time.RFC3339))
	}
	if err := os.MkdirAll(o.ScratchDir, 0755); err != nil {
		return fmt.Errorf("creating scratch directory %q: %w", o.ScratchDir, err)
	}
	if o.RepoRoot == "" {
		return fmt.Errorf("RepoRoot is required")
	}
	return nil
}

// MockgcpGenerateGcloudTestOptions is the options for the MockgcpGenerateGcloudTest tool.
type MockgcpGenerateGcloudTestOptions struct {
	*ToolOptions

	// GcloudCommand is the command we are going to generate a test script for.
	GcloudCommand string

	// Group is the name of the proto service.
	Group string

	// Resource is the name of the proto resource.
	Resource string
}

func (o *MockgcpGenerateGcloudTestOptions) DefaultAndValidate() error {
	if err := o.ToolOptions.DefaultAndValidate(); err != nil {
		return err
	}
	if o.GcloudCommand == "" {
		return fmt.Errorf("gcloudCommand is required")
	}
	if o.Resource == "" {
		return fmt.Errorf("resource is required")
	}
	if o.Group == "" {
		return fmt.Errorf("group is required")
	}
	return nil
}

// MockgcpGenerateGcloudTestResults holds the results from executing the MockgcpGenerateGcloudTest tool.
type MockgcpGenerateGcloudTestResults struct {
	AffectedPaths []string

	// TODO: This feels like a layering violation
	ExecResults *ExecResults
}

// RunMockgcpGenerateGcloudTest executes the MockgcpGenerateGcloudTest tool.
func RunMockgcpGenerateGcloudTest(ctx context.Context, opts MockgcpGenerateGcloudTestOptions) (*MockgcpGenerateGcloudTestResults, error) {
	results := &MockgcpGenerateGcloudTestResults{}

	if err := opts.DefaultAndValidate(); err != nil {
		return nil, err
	}

	// Check to see if the script file already exists
	scriptFileRelativePath := filepath.Join("mockgcp", fmt.Sprintf("mock%s", opts.Group), "testdata", opts.Resource, "crud", "script.yaml")
	scriptFilePath := filepath.Join(opts.RepoRoot, scriptFileRelativePath)

	// Check if file exists and handle force flag
	if _, err := os.Stat(scriptFilePath); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Printf("ERROR checking file %s: %v", scriptFilePath, err)
			return nil, err
		}
	} else if opts.Force {
		if err := os.Remove(scriptFilePath); err != nil {
			log.Printf("ERROR deleting existing file %s: %v", scriptFilePath, err)
			return nil, err
		}
	} else {
		log.Printf("SKIPPING %s already exists", scriptFilePath)
		return nil, nil
	}

	// Delete then write the prompt file.
	// promptPath := filepath.Join(opts.loggingDir, branch.Name, "create-script-prompt.txt")
	promptPath := filepath.Join(opts.ScratchDir, "create-script-prompt.txt")

	promptOptions := MockgcpGenerateGcloudTestPrompt{
		GcloudCommand: opts.GcloudCommand,
		Group:         opts.Group,
		Resource:      opts.Resource,
	}
	b, err := promptOptions.Generate()
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(promptPath, b, 0644); err != nil {
		return nil, fmt.Errorf("writing file %q: %w", promptPath, err)
	}

	// Run the LLM to generate the file.
	cfg := CommandConfig{
		Name:         "Generate script",
		Cmd:          "codebot",
		Args:         []string{"--ui-type=prompt", "--prompt=" + promptPath},
		WorkDir:      filepath.Join(opts.RepoRoot, "mockgcp"),
		RetryBackoff: GenerativeCommandRetryBackoff,
	}

	runOpts := &RunnerOptions{
		verbose: opts.Verbose,
	}
	execResults, err := executeCommand(runOpts, cfg)
	results.ExecResults = &execResults
	results.AffectedPaths = append(results.AffectedPaths, scriptFileRelativePath)
	return results, err
}

// createScriptYaml bridges the all-in-one conductor engine into the tool execution approach.
func createScriptYaml(ctx context.Context, runnerOptions *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	var opts MockgcpGenerateGcloudTestOptions

	opts.RepoRoot = runnerOptions.branchRepoDir
	opts.ScratchDir = filepath.Join(runnerOptions.loggingDir, branch.Name)
	opts.Force = runnerOptions.force

	opts.Resource = branch.Resource
	opts.Group = branch.Group
	opts.GcloudCommand = branch.Command

	results, err := RunMockgcpGenerateGcloudTest(ctx, opts)
	if err != nil {
		return nil, nil, err
	}
	return results.AffectedPaths, results.ExecResults, nil
}
