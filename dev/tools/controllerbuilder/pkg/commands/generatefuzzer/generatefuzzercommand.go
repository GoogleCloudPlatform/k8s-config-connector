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

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type generateFuzzerOptions struct {
	*options.GenerateOptions
	message     string
	apiVersion  string
	maxAttempts int
	llmModel    string

	prepopulateResource bool
	deterministic       bool
	service             string
	resources           options.ResourceList
	apiGoPackagePath    string
	apiDirectory        string
	outputDirectory     string
}

func (o *generateFuzzerOptions) isDeterministic() bool {
	return o.prepopulateResource || o.deterministic
}

func (o *generateFuzzerOptions) InitDefaults() error {
	if err := o.GenerateOptions.InitDefaults(); err != nil {
		return err
	}
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}
	o.apiGoPackagePath = "github.com/GoogleCloudPlatform/k8s-config-connector/apis/"
	o.apiDirectory = root + "/apis/"
	o.outputDirectory = root + "/pkg/controller/direct/"
	return nil
}

func (o *generateFuzzerOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.message, "message", o.message, "Proto message to generate fuzzer for")
	cmd.Flags().StringVar(&o.apiVersion, "api-version", o.apiVersion, "API version to generate fuzzer for")
	cmd.Flags().StringVar(&o.llmModel, "llm-model", o.llmModel, "LLM model to use for fuzzer generation")
	cmd.Flags().IntVar(&o.maxAttempts, "max-attempts", 5, "Maximum number of attempts to generate a valid fuzzer")
	cmd.Flags().BoolVar(&o.prepopulateResource, "prepopulate-resource", o.prepopulateResource, "write <kind>_fuzzer.generated.go deterministically from KRM types and mappers")
	cmd.Flags().BoolVar(&o.deterministic, "deterministic", o.deterministic, "alias for --prepopulate-resource")
	cmd.Flags().StringVarP(&o.service, "service", "s", o.service, "[--prepopulate-resource] the proto package of the resources")
	cmd.Flags().Var(&o.resources, "resource", "[--prepopulate-resource] a KRM kind and its proto message, as Kind:Message; repeatable")
	cmd.Flags().StringVar(&o.apiGoPackagePath, "api-go-package-path", o.apiGoPackagePath, "[--prepopulate-resource] import path of the KRM API packages")
	cmd.Flags().StringVar(&o.apiDirectory, "api-dir", o.apiDirectory, "[--prepopulate-resource] base directory of the KRM API packages")
	cmd.Flags().StringVar(&o.outputDirectory, "output-dir", o.outputDirectory, "[--prepopulate-resource] base directory of the direct controllers")
}

func (o *generateFuzzerOptions) Validate(cmd *cobra.Command) error {
	if o.isDeterministic() {
		for _, name := range []string{"message", "llm-model", "max-attempts"} {
			if cmd.Flags().Changed(name) {
				return fmt.Errorf("--%s is not used with --prepopulate-resource; name the resources with --resource Kind:Message", name)
			}
		}
		if o.service == "" {
			return fmt.Errorf("--service is required with --prepopulate-resource")
		}
		if o.apiVersion == "" {
			return fmt.Errorf("--api-version flag is required")
		}
		if len(o.resources) == 0 {
			return fmt.Errorf("at least one --resource Kind:Message is required with --prepopulate-resource")
		}
		return nil
	}
	for _, name := range []string{"service", "resource", "api-go-package-path", "api-dir", "output-dir"} {
		if cmd.Flags().Changed(name) {
			return fmt.Errorf("--%s requires --prepopulate-resource", name)
		}
	}
	if o.message == "" {
		return fmt.Errorf("--message flag is required")
	}
	if o.apiVersion == "" {
		return fmt.Errorf("--api-version flag is required")
	}
	return nil
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
			return opt.Validate(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if opt.isDeterministic() {
				return RunGenerateDeterministicFuzzers(ctx, opt)
			}
			return RunGenerateFuzzer(ctx, opt)
		},
	}

	opt.BindFlags(cmd)
	return cmd
}

func RunGenerateDeterministicFuzzers(ctx context.Context, opt *generateFuzzerOptions) error {
	gv, err := schema.ParseGroupVersion(opt.apiVersion)
	if err != nil || gv.Group == "" || gv.Version == "" {
		return fmt.Errorf("--api-version %q is not valid: want <group>/<version>", opt.apiVersion)
	}

	api, err := protoapi.LoadProto(opt.ProtoSourcePath, opt.ProtoOverlayPath)
	if err != nil {
		return fmt.Errorf("loading proto: %w", err)
	}

	goPackage := strings.TrimSuffix(gv.Group, ".cnrm.cloud.google.com")
	directDirectory := filepath.Join(append([]string{opt.outputDirectory}, strings.Split(goPackage, ".")...)...)
	generator := codegen.NewFuzzerGenerator(codegen.FuzzerGeneratorOptions{
		Group:            gv.Group,
		Version:          gv.Version,
		ProtoService:     opt.service,
		APIGoPackagePath: opt.apiGoPackagePath,
		APIDirectory:     opt.apiDirectory,
		DirectDirectory:  directDirectory,
	})

	for _, resource := range opt.resources {
		fullName := resource.ProtoMessageFullName(opt.service)
		desc, err := api.Files().FindDescriptorByName(protoreflect.FullName(fullName))
		if err != nil {
			return fmt.Errorf("finding proto message %q of kind %s: %w", fullName, resource.Kind, err)
		}
		msg, ok := desc.(protoreflect.MessageDescriptor)
		if !ok {
			return fmt.Errorf("%q of kind %s is not a proto message", fullName, resource.Kind)
		}
		result, err := generator.Generate(resource.Kind, msg)
		if err != nil {
			return fmt.Errorf("generating the fuzzer of kind %s: %w", resource.Kind, err)
		}
		if result.Skipped {
			fmt.Printf("Skipped the fuzzer of %s: %s\n", resource.Kind, result.SkipReason)
			continue
		}
		fmt.Printf("Wrote %s\n", result.Path)
	}
	return nil
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
	goPackage := strings.TrimSuffix(gv.Group, ".cnrm.cloud.google.com")
	resource := strings.ToLower(opt.message[strings.LastIndex(opt.message, ".")+1:])

	attempts := 0
	for attempts < opt.maxAttempts {
		attempts++

		fmt.Printf("Generating fuzzer for %s.%s (attempt %d/%d)...\n", goPackage, resource, attempts, opt.maxAttempts)

		outputPath := filepath.Join(root, "pkg/controller/direct", goPackage)
		if err := os.MkdirAll(outputPath, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}

		cmd := exec.CommandContext(ctx, "go", "run", "main.go", "prompt")
		cmd.Dir = filepath.Join(root, "dev/tools/controllerbuilder")
		cmd.Env = append(os.Environ(), fmt.Sprintf("REPO_ROOT=%s", root))
		if opt.llmModel != "" {
			cmd.Env = append(cmd.Env, fmt.Sprintf("LLM_MODEL=%s", opt.llmModel))
		}

		input := fmt.Sprintf("// +tool:fuzz-gen\n// proto.message: %s\n// api.group: %s\n", opt.message, gv.Group)
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
			continue
		}

		formatCmd := exec.CommandContext(ctx, "gofmt", "-w", outputFile)
		var formatStderr bytes.Buffer
		formatCmd.Stderr = &formatStderr

		if err := formatCmd.Run(); err != nil {
			fmt.Printf("gofmt failed: %v\nstderr: %s", err, formatStderr.String())
			continue
		}

		fmt.Printf("Verifying fuzzer for %s.%s (attempt %d/%d)...\n", goPackage, resource, attempts, opt.maxAttempts)

		testCmd := exec.CommandContext(ctx, "go", "test", "-v", filepath.Join(root, "pkg/fuzztesting/fuzztests"), "-fuzz=FuzzAllMappers", "-fuzztime", "60s")
		var testStderr bytes.Buffer
		testCmd.Stderr = &testStderr
		if err := testCmd.Run(); err != nil {
			fmt.Printf("Fuzzer verification failed: %v\nstderr: %s\n", err, testStderr.String())
			continue
		}

		fmt.Printf("Fuzzer generated and verified for %s.%s\n", goPackage, resource)
		break
	}

	return nil
}
