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

package generateidentity

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type GenerateIdentityOptions struct {
	*options.GenerateOptions

	ServiceName        string
	OutputAPIDirectory string
	Resources          options.ResourceList
	Patterns           []string
	EmitTests          bool
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &GenerateIdentityOptions{
		GenerateOptions: baseOptions,
	}

	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "generate-identity",
		Short: "deterministically generate <kind>_identity.generated.go and <kind>_reference.generated.go for KRM resources",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opt.loadAndApplyConfig(); err != nil {
				return err
			}
			return opt.validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunGenerateIdentity(cmd.Context(), opt)
		},
	}

	opt.BindFlags(cmd)
	return cmd
}

func (o *GenerateIdentityOptions) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}
	o.OutputAPIDirectory = filepath.Join(root, "apis")
	o.EmitTests = true
	return nil
}

func (o *GenerateIdentityOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&o.ServiceName, "service", "s", o.ServiceName, "the GCP proto service name")
	cmd.Flags().StringVar(&o.OutputAPIDirectory, "output-api", o.OutputAPIDirectory, "base directory for writing APIs")
	cmd.Flags().Var(&o.Resources, "resource", "the KRM Kind and the equivalent proto resource separated with a colon (e.g. `StorageBucket:Bucket`). Can be specified multiple times.")
	cmd.Flags().StringSliceVar(&o.Patterns, "pattern", nil, "explicit URL pattern override per Kind in the format `Kind=pattern` or `Kind=//host/pattern`. Can be specified multiple times.")
	cmd.Flags().BoolVar(&o.EmitTests, "emit-tests", o.EmitTests, "emit <kind>_identity_generated_test.go calling identity.AssertConformance")
}

func (o *GenerateIdentityOptions) loadAndApplyConfig() error {
	if o.ConfigFilePath == "" {
		return nil
	}
	config, err := codegen.LoadConfig(o.ConfigFilePath)
	if err != nil {
		return fmt.Errorf("loading service config: %w", err)
	}
	if config == nil {
		return nil
	}

	o.ServiceName = config.Service
	o.APIVersion = config.APIVersion
	if o.ProtoOverlayPath == "" && config.ProtoOverlay != "" {
		if filepath.IsAbs(config.ProtoOverlay) {
			o.ProtoOverlayPath = config.ProtoOverlay
		} else {
			o.ProtoOverlayPath = filepath.Join(filepath.Dir(o.ConfigFilePath), config.ProtoOverlay)
		}
	}
	for _, res := range config.Resources {
		o.Resources = append(o.Resources, options.Resource{
			Kind:              res.Kind,
			ProtoName:         res.ProtoName,
			SkipScaffoldFiles: res.SkipScaffoldFiles,
		})
	}
	return nil
}

func (o *GenerateIdentityOptions) validate() error {
	if _, err := parsePatternFlags(o.Patterns); err != nil {
		return err
	}
	if o.ServiceName == "" {
		return fmt.Errorf("`--service` is required")
	}
	if o.APIVersion == "" {
		return fmt.Errorf("`--api-version` is required")
	}
	if o.GenerateOptions.ProtoSourcePath == "" {
		return fmt.Errorf("`--proto-source-path` is required")
	}
	if len(o.Resources) == 0 {
		return fmt.Errorf("`--resource` is required")
	}
	return nil
}

func RunGenerateIdentity(ctx context.Context, o *GenerateIdentityOptions) error {
	gv, err := schema.ParseGroupVersion(o.APIVersion)
	if err != nil {
		return fmt.Errorf("APIVersion %q is not valid: %w", o.APIVersion, err)
	}

	api, err := protoapi.LoadProto(o.GenerateOptions.ProtoSourcePath, o.GenerateOptions.ProtoOverlayPath)
	if err != nil {
		return fmt.Errorf("loading proto: %w", err)
	}

	patternOverrides, err := parsePatternFlags(o.Patterns)
	if err != nil {
		return err
	}

	gen := codegen.NewIdentityGenerator(codegen.IdentityGeneratorOptions{
		Group:            gv.Group,
		Version:          gv.Version,
		ProtoService:     o.ServiceName,
		APIDirectory:     o.OutputAPIDirectory,
		Proto:            api,
		EmitTests:        o.EmitTests,
		PatternOverrides: patternOverrides,
	})

	var queued []judgement.Entry
	for _, res := range o.Resources {
		if res.SkipScaffoldFiles {
			continue
		}
		msg, err := gen.ResolveMessage(res.ProtoName)
		if err != nil {
			return fmt.Errorf("resolving proto for %s: %w", res.Kind, err)
		}
		result, err := gen.Generate(res.Kind, msg)
		if err != nil {
			return fmt.Errorf("generating identity for %s: %w", res.Kind, err)
		}
		if result != nil && result.Plan != nil {
			queued = append(queued, result.Plan.Judgement...)
		}
	}

	if len(queued) > 0 {
		goPackage := strings.TrimSuffix(gv.Group, ".cnrm.cloud.google.com") + "/" + gv.Version
		if err := writeJudgementQueue(o.OutputAPIDirectory, goPackage, queued); err != nil {
			return fmt.Errorf("writing judgement queue: %w", err)
		}
	}

	return nil
}

func parsePatternFlags(raw []string) (map[string]string, error) {
	if len(raw) == 0 {
		return nil, nil
	}
	out := make(map[string]string, len(raw))
	for _, entry := range raw {
		entry = strings.TrimSpace(entry)
		if entry == "" {
			continue
		}
		kind, pat, ok := strings.Cut(entry, "=")
		kind = strings.TrimSpace(kind)
		pat = strings.TrimSpace(pat)
		if !ok || kind == "" || pat == "" {
			return nil, fmt.Errorf("invalid --pattern %q: expected Kind=pattern or Kind=//host/pattern", entry)
		}
		out[kind] = pat
	}
	return out, nil
}

func writeJudgementQueue(apiDir, goPackage string, entries []judgement.Entry) error {
	serviceDir := filepath.Dir(filepath.Join(apiDir, goPackage))
	path := filepath.Join(serviceDir, judgement.FileName)

	if err := os.MkdirAll(serviceDir, 0755); err != nil {
		return err
	}

	q, err := judgement.Read(path)
	if err != nil {
		return err
	}
	q.Merge(entries)
	return judgement.Write(path, q)
}
