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

package generatetypes

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/scaffold"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type GenerateTypesV2Options struct {
	*options.GenerateOptions

	OutputAPIDirectory string

	ServiceMetadata *codegen.ServiceMetadata
}

func (o *GenerateTypesV2Options) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}
	o.OutputAPIDirectory = root + "/apis/"
	return nil
}

func (o *GenerateTypesV2Options) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.OutputAPIDirectory, "output-api", o.OutputAPIDirectory, "base directory for writing APIs")
}

func (o *GenerateTypesV2Options) validate() error {
	if o.MetadataFile == "" {
		return fmt.Errorf("metadata file is required")
	}
	if o.ProtoSourcePath == "" {
		return fmt.Errorf("proto source path is required")
	}
	if o.ServiceName != "" {
		return fmt.Errorf("service name is not supported for v2, use metadata file instead")
	}
	if o.APIVersion != "" {
		return fmt.Errorf("api version is not supported for v2, use metadata file instead")
	}
	return nil
}

func BuildV2Command(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &GenerateTypesV2Options{
		GenerateOptions: baseOptions,
	}

	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "generate-types-v2",
		Short: "generate KRM types for a proto service (v2)",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opt.loadMetadata(); err != nil {
				return err
			}
			return opt.validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateTypes(cmd.Context(), opt); err != nil {
				return err
			}
			return nil
		},
	}

	opt.BindFlags(cmd)
	return cmd
}

func generateTypes(_ context.Context, o *GenerateTypesV2Options) error {
	gv, err := schema.ParseGroupVersion(o.ServiceMetadata.APIVersion)
	if err != nil {
		return fmt.Errorf("APIVersion %q is not valid: %w", o.ServiceMetadata.APIVersion, err)
	}

	api, err := protoapi.LoadProto(o.GenerateOptions.ProtoSourcePath)
	if err != nil {
		return fmt.Errorf("loading proto: %w", err)
	}

	goPackage := strings.TrimSuffix(gv.Group, ".cnrm.cloud.google.com") + "/" + gv.Version

	scaffolder := &scaffold.APIScaffolder{
		BaseDir:         o.OutputAPIDirectory,
		GoPackage:       goPackage,
		Group:           gv.Group,
		Version:         gv.Version,
		PackageProtoTag: o.ServiceMetadata.Service,
	}

	// Generate doc.go and groupversion_info.go if they don't exist
	if scaffolder.DocFileNotExist() {
		if err := scaffolder.AddDocFile(); err != nil {
			return fmt.Errorf("add doc.go file: %w", err)
		}
	}
	if scaffolder.GroupVersionFileNotExist() {
		if err := scaffolder.AddGroupVersionFile(); err != nil {
			return fmt.Errorf("add groupversion_info.go file: %w", err)
		}
	}

	// Create a new type generator and generate Go types
	typeGenerator := codegen.NewTypeGeneratorV2(
		goPackage,
		api,
		o.OutputAPIDirectory,
		o.ServiceMetadata,
	)

	if err := typeGenerator.VisitProto(); err != nil {
		return err
	}
	if err := typeGenerator.WriteVisitedMessages(); err != nil {
		return err
	}
	if err := typeGenerator.WriteOutputMessages(); err != nil {
		return err
	}

	// Generate template files
	fmt.Printf("scaffolding for service %+v\n", o.ServiceMetadata)
	for _, resource := range o.ServiceMetadata.Resources {
		if !resource.SkipScaffoldFiles {
			fmt.Printf("scaffolding for resource %+v\n", resource)
			kind := resource.Kind
			if !scaffolder.TypeFileNotExist(resource.ProtoName) {
				fmt.Printf("file %s already exists, skipping\n", scaffolder.PathToTypeFile(resource.ProtoName))
			} else {
				err := scaffolder.AddTypeFileV2(resource.ProtoName, kind)
				if err != nil {
					return fmt.Errorf("add type file %s: %w", scaffolder.PathToTypeFile(resource.ProtoName), err)
				}
			}
			if scaffolder.RefsFileExist(kind, resource.ProtoName) {
				fmt.Printf("file %s already exists, skipping\n", scaffolder.PathToRefsFile(kind, resource.ProtoName))
			} else {
				err := scaffolder.AddRefsFile(kind, resource.ProtoName)
				if err != nil {
					return fmt.Errorf("add refs file %s: %w", scaffolder.PathToRefsFile(kind, resource.ProtoName), err)
				}
			}
			if scaffolder.IdentityFileExist(kind, resource.ProtoName) {
				fmt.Printf("file %s already exists, skipping\n", scaffolder.PathToIdentityFile(kind, resource.ProtoName))
			} else {
				err := scaffolder.AddIdentityFile(kind, resource.ProtoName)
				if err != nil {
					return fmt.Errorf("add identity file %s: %w", scaffolder.PathToIdentityFile(kind, resource.ProtoName), err)
				}
			}
		}
	}

	addCopyright := true
	if err := typeGenerator.WriteFiles(addCopyright); err != nil {
		return err
	}

	return nil
}

func (o *GenerateTypesV2Options) loadMetadata() error {
	metadata, err := codegen.LoadServiceMetadata(o.MetadataFile)
	if err != nil {
		return fmt.Errorf("loading metadata: %w", err)
	}
	o.ServiceMetadata = metadata
	return nil
}
