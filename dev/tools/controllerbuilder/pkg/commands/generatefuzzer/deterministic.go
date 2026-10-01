// Copyright 2026 Google LLC
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
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// deterministicOptions are the flags of generate-fuzzer --deterministic, which writes
// <kind>_fuzzer.generated.go from the KRM types and the mappers, instead of asking an LLM.
type deterministicOptions struct {
	enabled          bool
	service          string
	resources        options.ResourceList
	apiGoPackagePath string
	apiDirectory     string
	outputDirectory  string
}

// llmOnlyFlags are the flags that only the LLM mode uses.
var llmOnlyFlags = []string{"message", "llm-model", "max-attempts"}

// deterministicOnlyFlags are the flags that only --deterministic uses.
var deterministicOnlyFlags = []string{"service", "resource", "api-go-package-path", "api-dir", "output-dir"}

// InitDefaults sets the same defaults as generate-mapper.
func (o *deterministicOptions) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}
	o.apiGoPackagePath = "github.com/GoogleCloudPlatform/k8s-config-connector/apis/"
	o.apiDirectory = root + "/apis/"
	o.outputDirectory = root + "/pkg/controller/direct/"
	return nil
}

func (o *deterministicOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&o.enabled, "deterministic", o.enabled, "write <kind>_fuzzer.generated.go from the KRM types and mappers, without an LLM")
	cmd.Flags().StringVarP(&o.service, "service", "s", o.service, "[--deterministic] the proto package of the resources")
	cmd.Flags().Var(&o.resources, "resource", "[--deterministic] a KRM kind and its proto message, as Kind:Message; repeatable")
	cmd.Flags().StringVar(&o.apiGoPackagePath, "api-go-package-path", o.apiGoPackagePath, "[--deterministic] import path of the KRM API packages")
	cmd.Flags().StringVar(&o.apiDirectory, "api-dir", o.apiDirectory, "[--deterministic] base directory of the KRM API packages")
	cmd.Flags().StringVar(&o.outputDirectory, "output-dir", o.outputDirectory, "[--deterministic] base directory of the direct controllers, where the fuzzers are written")
}

// validateDeterministic checks the flags of generate-fuzzer --deterministic.
func (o *generateFuzzerOptions) validateDeterministic(cmd *cobra.Command) error {
	for _, name := range llmOnlyFlags {
		if cmd.Flags().Changed(name) {
			return fmt.Errorf("--%s is not used with --deterministic; name the resources with --resource Kind:Message", name)
		}
	}
	if o.deterministic.service == "" {
		return fmt.Errorf("--service is required with --deterministic")
	}
	if o.apiVersion == "" {
		return fmt.Errorf("--api-version flag is required")
	}
	if len(o.deterministic.resources) == 0 {
		return fmt.Errorf("at least one --resource Kind:Message is required with --deterministic")
	}
	return nil
}

// validateLLMMode rejects the flags of --deterministic when the command asks an LLM.
func validateLLMMode(cmd *cobra.Command) error {
	for _, name := range deterministicOnlyFlags {
		if cmd.Flags().Changed(name) {
			return fmt.Errorf("--%s requires --deterministic", name)
		}
	}
	return nil
}

// RunGenerateDeterministicFuzzers writes <kind>_fuzzer.generated.go for each --resource into
// the direct controller package of the API group, the directory generate-mapper writes
// mapper.generated.go to. A kind whose package already has a fuzzer that is not generated
// is skipped.
func RunGenerateDeterministicFuzzers(ctx context.Context, opt *generateFuzzerOptions) error {
	o := &opt.deterministic
	gv, err := schema.ParseGroupVersion(opt.apiVersion)
	if err != nil {
		return fmt.Errorf("--api-version %q is not valid: %w", opt.apiVersion, err)
	}
	if gv.Group == "" || gv.Version == "" {
		return fmt.Errorf("--api-version %q is not valid: want <group>/<version>", opt.apiVersion)
	}

	api, err := protoapi.LoadProto(opt.ProtoSourcePath, opt.ProtoOverlayPath)
	if err != nil {
		return fmt.Errorf("loading proto: %w", err)
	}

	goPackage := strings.TrimSuffix(gv.Group, ".cnrm.cloud.google.com")
	directDirectory := filepath.Join(append([]string{o.outputDirectory}, strings.Split(goPackage, ".")...)...)
	generator := codegen.NewFuzzerGenerator(codegen.FuzzerGeneratorOptions{
		Group:            gv.Group,
		Version:          gv.Version,
		ProtoService:     o.service,
		APIGoPackagePath: o.apiGoPackagePath,
		APIDirectory:     o.apiDirectory,
		DirectDirectory:  directDirectory,
	})

	for _, resource := range o.resources {
		fullName := resource.ProtoMessageFullName(o.service)
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
