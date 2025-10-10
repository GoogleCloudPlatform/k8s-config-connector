// Copyright 2024 Google LLC
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

package generatedirectreconciler

import (
	"context"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/generatecontroller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/generatemapper"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/generatetypes"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type GenerateBasicReconcilerOptions struct {
	*options.GenerateOptions
	Resource options.Resource

	ServiceName           string
	APIGoPackagePath      string
	APIDirectory          string
	OutputMapperDirectory string
}

func (o *GenerateBasicReconcilerOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&o.ServiceName, "service", "s", o.ServiceName, "the GCP service name")
	cmd.Flags().Var(&o.Resource, "resource", "the KRM Kind and the equivalent proto resource separated with a colon. e.g. for resource google.storage.v1.Bucket, the flag should be `StorageBucket:Bucket`")
	cmd.Flags().StringVar(&o.APIGoPackagePath, "api-go-package-path", o.APIGoPackagePath, "package path")
	cmd.Flags().StringVar(&o.APIDirectory, "api-dir", o.APIDirectory, "base directory for reading APIs")
	cmd.Flags().StringVar(&o.OutputMapperDirectory, "output-dir", o.OutputMapperDirectory, "base directory for writing mappers")
}

func (o *GenerateBasicReconcilerOptions) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}
	o.APIGoPackagePath = "github.com/GoogleCloudPlatform/k8s-config-connector/apis/"
	o.APIDirectory = root + "/apis/"
	//	o.OutputAPIDirectory = root + "/apis/"
	o.OutputMapperDirectory = root + "/pkg/controller/direct/"
	return nil
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &GenerateBasicReconcilerOptions{
		GenerateOptions: baseOptions,
	}

	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "generate-direct-reconciler",
		Short: "[ALPHA] generate a basic direct reconciler that is up and run",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opt.Resource == (options.Resource{}) {
				return fmt.Errorf("--resource flag is required")
			}

			if baseOptions.APIVersion == "" {
				return fmt.Errorf("--api-version is required")
			}
			_, err := schema.ParseGroupVersion(baseOptions.APIVersion)
			if err != nil {
				return fmt.Errorf("unable to parse --api-version: %w", err)
			}

			if opt.ServiceName == "" {
				return fmt.Errorf("--service is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := RunGenerateBasicReconciler(ctx, opt); err != nil {
				return err
			}
			return nil
		},
	}
	opt.BindFlags(cmd)

	return cmd
}

func RunGenerateBasicReconciler(ctx context.Context, o *GenerateBasicReconcilerOptions) error {
	crdOps := &generatetypes.GenerateCRDOptions{
		GenerateOptions:    o.GenerateOptions,
		OutputAPIDirectory: o.APIDirectory,
		Resources:          options.ResourceList{o.Resource},
	}
	if err := generatetypes.RunGenerateCRD(ctx, crdOps); err != nil {
		return fmt.Errorf("generate types: %w", err)
	}
	mapperOps := &generatemapper.GenerateMapperOptions{
		GenerateOptions:       o.GenerateOptions,
		APIGoPackagePath:      o.APIGoPackagePath,
		APIDirectory:          o.APIDirectory,
		OutputMapperDirectory: o.OutputMapperDirectory,
	}
	if err := generatemapper.RunGenerateMapper(ctx, mapperOps); err != nil {
		return fmt.Errorf("generate mapper: %w", err)
	}
	controllerOps := &generatecontroller.GenerateControllerOptions{
		GenerateOptions: o.GenerateOptions,
		Resource:        o.Resource,
	}
	if err := generatecontroller.RunController(ctx, controllerOps); err != nil {
		return fmt.Errorf("generate controller: %w", err)
	}
	return nil
}
