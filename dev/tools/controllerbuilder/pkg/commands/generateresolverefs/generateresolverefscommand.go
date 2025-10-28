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

package generateresolverefs

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type GenerateResolveRefsOptions struct {
	*options.GenerateOptions

	APIDirectory    string
	OutputDirectory string
}

func (o *GenerateResolveRefsOptions) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}
	o.APIDirectory = root + "/apis/"
	o.OutputDirectory = root + "/pkg/controller/direct/"
	return nil
}

func (o *GenerateResolveRefsOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.APIDirectory, "api-dir", o.APIDirectory, "base directory for reading APIs")
	cmd.Flags().StringVar(&o.OutputDirectory, "output-dir", o.OutputDirectory, "base directory for writing resolverefs")
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &GenerateResolveRefsOptions{
		GenerateOptions: baseOptions,
	}

	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "generate-resolverefs",
		Short: "generate resolverefs function for a KRM resource",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opt.loadAndApplyConfig(); err != nil {
				return err
			}
			return opt.validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := RunGenerateResolveRefs(ctx, opt); err != nil {
				return err
			}
			return nil
		},
	}

	opt.BindFlags(cmd)

	return cmd
}

func RunGenerateResolveRefs(ctx context.Context, o *GenerateResolveRefsOptions) error {
	gv, err := schema.ParseGroupVersion(o.APIVersion)
	if err != nil {
		return fmt.Errorf("APIVersion %q is not valid: %w", o.APIVersion, err)
	}

	service := strings.TrimSuffix(gv.Group, ".cnrm.cloud.google.com")
	version := gv.Version

	generator := codegen.NewResolveRefsGenerator(o.OutputDirectory)

	apiGoPackagePath := filepath.Join("github.com/GoogleCloudPlatform/k8s-config-connector/apis", service, version)
	apiDirectory := filepath.Join(o.APIDirectory, service, version)

	if err := generator.VisitGoCode(apiGoPackagePath, apiDirectory); err != nil {
		return fmt.Errorf("visiting go code: %w", err)
	}

	if err := generator.GenerateResolveRefs(service, version); err != nil {
		return fmt.Errorf("generating resolverefs: %w", err)
	}

	addCopyright := true
	writeEmptyFiles := true
	if err := generator.WriteFiles(addCopyright, writeEmptyFiles); err != nil {
		return err
	}

	return nil
}

func (o *GenerateResolveRefsOptions) loadAndApplyConfig() error {
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

	o.APIVersion = config.APIVersion
	return nil
}

func (o *GenerateResolveRefsOptions) validate() error {
	if o.APIVersion == "" {
		return fmt.Errorf("--api-version is required")
	}
	return nil
}
