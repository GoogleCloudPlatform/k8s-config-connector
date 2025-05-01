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
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type generateTestOptions struct {
	*options.GenerateOptions
	//message     string
	//apiVersion  string
	maxAttempts int
	llmModel    string
	kind        string
}

func (o *generateTestOptions) BindFlags(cmd *cobra.Command) {
	//cmd.Flags().StringVar(&o.message, "message", o.message, "Proto message to generate fuzzer for")
	//cmd.Flags().StringVar(&o.apiVersion, "api-version", o.apiVersion, "API version to generate fuzzer for")
	cmd.Flags().StringVar(&o.llmModel, "llm-model", o.llmModel, "LLM model to use for test generation")
	cmd.Flags().IntVar(&o.maxAttempts, "max-attempts", 5, "Maximum number of attempts to generate a test with maximal coverage")
	cmd.Flags().StringVar(&o.kind, "kind", "", "the KRM Kind for test generation")
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &generateTestOptions{
		GenerateOptions: baseOptions,
	}
	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "generate-test",
		Short: "Generate test data for maximal coverage",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			//if opt.message == "" {
			//	// TODO: extract messages from generate.sh or api directory
			//	return fmt.Errorf("--message flag is required")
			//}
			if baseOptions.APIVersion == "" {
				return fmt.Errorf("--api-version flag is required")
			}
			if opt.kind == "" {
				return fmt.Errorf("--kind flag is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := RunGenerateTest(ctx, opt); err != nil {
				return err
			}
			return nil
		},
	}

	opt.BindFlags(cmd)
	return cmd
}

func RunGenerateTest(ctx context.Context, opt *generateTestOptions) error {
	root, err := options.RepoRoot()
	if err != nil {
		return fmt.Errorf("failed to get repo root: %v", err)
	}

	gv, err := schema.ParseGroupVersion(opt.APIVersion)
	if err != nil {
		return fmt.Errorf("failed to parse api version: %v", err)
	}
	groupPrefix := strings.TrimSuffix(gv.Group, ".cnrm.cloud.google.com") // e.g. bigquerydatatransfer.cnrm.cloud.google.com -> bigquerydatatransfer
	//resource := strings.ToLower(opt.message[strings.LastIndex(opt.message, ".")+1:]) // e.g. google.cloud.bigquery.datatransfer.v1.TransferConfig -> transferconfig

	attempts := 0
	for attempts < opt.maxAttempts {
		attempts++

		fmt.Printf("Generating test for %s (attempt %d/%d)...\n", opt.kind, attempts, opt.maxAttempts)

		// 1. Create output directory if it doesn't exist
		outputDir := filepath.Join(
			root, "pkg", "test", "resourcefixture", "testdata", "basic",
			groupPrefix, gv.Version, strings.ToLower(opt.kind),
			fmt.Sprintf("%s-full", strings.ToLower(opt.kind)),
		)
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}

		// 2. Generate create.yaml in test using the prompt command
		cmd := exec.CommandContext(ctx, "go", "run", "main.go", "prompt")
		cmd.Dir = filepath.Join(root, "dev/tools/controllerbuilder")
		cmd.Env = append(os.Environ(), fmt.Sprintf("REPO_ROOT=%s", root))
		if opt.llmModel != "" {
			cmd.Env = append(cmd.Env, fmt.Sprintf("LLM_MODEL=%s", opt.llmModel))
		}

		currentYear := time.Now().Year()
		input := fmt.Sprintf(`Generate a %s/create.yaml file for testing a Kubernetes controller.

First, read the CRD file at config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_<pluralized-kind>.%s.yaml to understand the schema.
Replace <pluralized-kind> with the pluralized version of the kind: %s in the filename.

The file should follow these requirements:
- Add an Apache 2.0 license header with Copyright %d Google LLC
- Use apiVersion: %s
- Use kind: %s
- Include metadata.name: %s-${uniqueId}
- If the CRD has a "spec.projectRef" field, use projectRef.external: ${projectId}
- If the CRD has a "spec.location" field, use location: us-central1
- If the CRD has any field name end with "Ref", it's a reference field. Use its subfield ".name" and set value to be <kind>-${uniqueId}. Replace <kind> with the kind of the reference field. The kind is indicated in the description of the reference field's subfield, "external" field.
- Follow the schema defined in the CRD file
- Use as many fields under spec as possible to try to reach full coverage
- Do not use any field not defined in the CRD

Use ReadFile to read the CRD file.
Use CreateFile to write the YAML content to the %s/create.yaml file.
Respond only with the YAML content, no explanations.`,
			outputDir, gv.Group, opt.kind, currentYear, gv.String(), opt.kind, strings.ToLower(opt.kind), outputDir)

		cmd.Stdin = strings.NewReader(input)

		outputFile := filepath.Join(outputDir, "create.yaml")
		cmd.Stdout, err = os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("failed to create output file %s: %w", outputFile, err)
		}

		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Test generation failed: %v\n%s\n", err, stderr.String())
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
		//fmt.Printf("Verifying fuzzer for %s.%s (attempt %d/%d)...\n", goPackage, resource, attempts, opt.maxAttempts)

		testCmd := exec.CommandContext(ctx, "go", "test", "-v", filepath.Join(root, "pkg/fuzztesting/fuzztests"), "-fuzz=FuzzAllMappers", "-fuzztime", "60s")
		var testStderr bytes.Buffer
		testCmd.Stderr = &testStderr
		if err := testCmd.Run(); err != nil {
			fmt.Printf("Fuzzer verification failed: %v\nstderr: %s\n", err, testStderr.String())
			continue // verification failed, keep trying
		}

		//fmt.Printf("Fuzzer generated and verified for %s.%s\n", goPackage, resource)
		break
	}

	return nil
}
