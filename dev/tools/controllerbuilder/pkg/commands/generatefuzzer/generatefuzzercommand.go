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

package generatefuzzer

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type generateFuzzerOptions struct {
	*options.GenerateOptions
	message     string
	apiVersion  string
	Kind        string
	maxAttempts int
	llmModel    string
}

func (o *generateFuzzerOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.message, "message", o.message, "Proto message to generate fuzzer for")
	cmd.Flags().StringVar(&o.apiVersion, "api-version", o.apiVersion, "API version to generate fuzzer for")
	cmd.Flags().StringVar(&o.Kind, "kind", o.Kind, "Kind to generate fuzzer for")
	cmd.Flags().StringVar(&o.llmModel, "llm-model", o.llmModel, "LLM model to use for fuzzer generation")
	cmd.Flags().IntVar(&o.maxAttempts, "max-attempts", 5, "Maximum number of attempts to generate a valid fuzzer")
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &generateFuzzerOptions{
		GenerateOptions: baseOptions,
	}
	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "generate-fuzzer",
		Short: "Generate fuzzer tests for proto messages",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opt.message == "" {
				// TODO: extract messages from generate.sh or api directory
				return fmt.Errorf("--message flag is required")
			}
			if opt.apiVersion == "" {
				return fmt.Errorf("--api-version flag is required")
			}
			if opt.Kind == "" {
				return fmt.Errorf("--kind flag is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := RunGenerateFuzzer(ctx, opt); err != nil {
				return err
			}
			return nil
		},
	}

	opt.BindFlags(cmd)
	return cmd
}

func RunGenerateFuzzer(ctx context.Context, opt *generateFuzzerOptions) error {
	root, err := options.RepoRoot()
	if err != nil {
		return fmt.Errorf("failed to get repo root: %v", err)
	}

	gv, err := schema.ParseGroupVersion(opt.apiVersion)
	if err != nil {
		return fmt.Errorf("failed to parse api version: %v", err)
	}
	goPackage := strings.TrimSuffix(gv.Group, ".cnrm.cloud.google.com")              // e.g. bigquerydatatransfer.cnrm.cloud.google.com -> bigquerydatatransfer
	resource := strings.ToLower(opt.message[strings.LastIndex(opt.message, ".")+1:]) // e.g. google.cloud.bigquery.datatransfer.v1.TransferConfig -> transferconfig

	attempts := 0
	for attempts < opt.maxAttempts {
		attempts++

		fmt.Printf("Generating fuzzer for %s.%s (attempt %d/%d)...\n", goPackage, resource, attempts, opt.maxAttempts)

		// 1. Create output directory if it doesn't exist
		outputPath := filepath.Join(root, "pkg/controller/direct", goPackage)
		if err := os.MkdirAll(outputPath, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}

		// 2. Generate fuzzer using the prompt command
		cmd := exec.CommandContext(ctx, "go", "run", "main.go", "prompt")
		cmd.Dir = filepath.Join(root, "dev/tools/controllerbuilder")
		cmd.Env = append(os.Environ(), fmt.Sprintf("REPO_ROOT=%s", root))
		if opt.llmModel != "" {
			cmd.Env = append(cmd.Env, fmt.Sprintf("LLM_MODEL=%s", opt.llmModel))
		}

		input := fmt.Sprintf("// +tool:fuzz-gen\n// proto.message: %s\n// api.group: %s\n// crd.kind: %s\n", opt.message, gv.Group, opt.Kind)
		cmd.Stdin = strings.NewReader(input)

		outputFile := filepath.Join(root, "pkg/controller/direct", goPackage, fmt.Sprintf("%s_fuzzer.go", resource))
		cmd.Stdout, err = os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("failed to create output file: %v", err)
		}

		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Fuzzer generation failed: %v\n%s\n", err, stderr.String())
			continue // keep trying
		}

		// 3. Format the generated fuzzer
		formatCmd := exec.CommandContext(ctx, "gofmt", "-w", outputFile)
		var formatStderr bytes.Buffer
		formatCmd.Stderr = &formatStderr

		if err := formatCmd.Run(); err != nil {
			fmt.Printf("gofmt failed: %v\nstderr: %s", err, formatStderr.String())
			continue // gofmt failed, keep trying
		}

		// 4. Verify the generated fuzzer
		fmt.Printf("Verifying fuzzer for %s.%s (attempt %d/%d)...\n", goPackage, resource, attempts, opt.maxAttempts)

		testCmd := exec.CommandContext(ctx, "go", "test", "-v", filepath.Join(root, "pkg/fuzztesting/fuzztests"), "-fuzz=FuzzAllMappers", "-fuzztime", "60s")
		var testStderr bytes.Buffer
		testCmd.Stderr = &testStderr
		if err := testCmd.Run(); err != nil {
			fmt.Printf("Fuzzer verification failed: %v\nstderr: %s\n", err, testStderr.String())
			continue // verification failed, keep trying
		}

		fmt.Printf("Fuzzer generated and verified for %s.%s\n", goPackage, resource)
		break
	}

	return nil
}
