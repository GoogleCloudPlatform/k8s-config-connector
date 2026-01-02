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

package generatemapper

import (
	"context"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/annotations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/spf13/cobra"
)

type GenerateMapperOptions struct {
	*options.GenerateOptions

	// IncludeProtoPackages lists proto packages to generate mappers for.
	IncludeProtoPackages []string

	// IncludeProtoMessages lists additional fully-qualified proto messages to generate mappers for.
	IncludeProtoMessages []string

	// SourceDirs will only generate mappers for go types that appear in the specified source directories.
	SourceDirs []string

	APIGoPackagePath      string
	APIDirectory          string
	OutputMapperDirectory string

	Multiversion bool
}

func (o *GenerateMapperOptions) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}
	o.APIGoPackagePath = "github.com/GoogleCloudPlatform/k8s-config-connector/apis/"
	o.APIDirectory = root + "/apis/"
	o.OutputMapperDirectory = root + "/pkg/controller/direct/"
	return nil
}

func (o *GenerateMapperOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringSliceVarP(&o.IncludeProtoPackages, "proto-package", "p", o.IncludeProtoPackages, "the proto packages to generate mappers for")
	cmd.Flags().StringSliceVarP(&o.IncludeProtoPackages, "service", "s", o.IncludeProtoPackages, "the proto services to generate mappers for")
	cmd.Flags().MarkDeprecated("service", "use --proto-package instead of --service to specify packages to generate mappers for")
	cmd.Flags().StringSliceVarP(&o.IncludeProtoMessages, "proto-message", "m", o.IncludeProtoMessages, "an additional fully-qualified proto message to generate (used to add messages from additional packages)")

	cmd.Flags().StringSliceVar(&o.SourceDirs, "source-dir", o.SourceDirs, "the go source directories to generate mappers for")

	cmd.Flags().StringVar(&o.APIGoPackagePath, "api-go-package-path", o.APIGoPackagePath, "package path")
	cmd.Flags().StringVar(&o.APIDirectory, "api-dir", o.APIDirectory, "base directory for reading APIs")
	cmd.Flags().StringVar(&o.OutputMapperDirectory, "output-dir", o.OutputMapperDirectory, "base directory for writing mappers")
	cmd.Flags().BoolVar(&o.Multiversion, "multiversion", o.Multiversion, "generate mappers with version specifiers, to support mixed versions")
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &GenerateMapperOptions{
		GenerateOptions: baseOptions,
	}

	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "generate-mapper",
		Short: "generate mapper functions for a proto service",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opt.loadAndApplyConfig(); err != nil {
				return err
			}
			return opt.validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := RunGenerateMapper(ctx, opt); err != nil {
				return err
			}
			return nil
		},
	}

	opt.BindFlags(cmd)

	return cmd
}

func RunGenerateMapper(ctx context.Context, o *GenerateMapperOptions) error {
	gv, err := schema.ParseGroupVersion(o.APIVersion)
	if err != nil {
		return fmt.Errorf("APIVersion %q is not valid: %w", o.APIVersion, err)
	}

	api, err := protoapi.LoadProto(o.GenerateOptions.ProtoSourcePath)
	if err != nil {
		return fmt.Errorf("loading proto: %w", err)
	}

	goPackage := strings.TrimSuffix(gv.Group, ".cnrm.cloud.google.com")

	pathForMessage := func(msg protoreflect.MessageDescriptor) (string, bool) {
		fullName := string(msg.FullName())

		matched := false

		if slices.Contains(o.IncludeProtoMessages, fullName) {
			matched = true
		}

		if !matched {
			if strings.HasSuffix(fullName, "Request") {
				return "", false
			}
			if strings.HasSuffix(fullName, "Response") {
				return "", false
			}
			if strings.HasSuffix(fullName, "OperationMetadata") {
				return "", false
			}
			if strings.HasSuffix(fullName, "Metadata") {
				return "", false
			}
			for _, protoPackage := range o.IncludeProtoPackages {
				if strings.HasPrefix(fullName, protoPackage+".") {
					matched = true
				}
			}
		}
		if !matched {
			return "", false
		}

		return goPackage, true
	}

	generatedFileAnnotation := &annotations.FileAnnotation{
		Key: "+generated:mapper",
		Attributes: map[string][]string{
			"proto.service": o.IncludeProtoPackages,
			"krm.group":     {gv.Group},
			"krm.version":   {gv.Version},
		},
	}

	mapperGenerator := codegen.NewMapperGenerator(pathForMessage, o.OutputMapperDirectory, generatedFileAnnotation, o.Multiversion)

	// Ensure that our first proto package is always imported with the "pb" alias.
	firstService, err := api.GetFileDescriptorByPackage(o.IncludeProtoPackages[0])
	if err != nil {
		return err
	}
	mapperGenerator.AddGoImportAlias(codegen.GoPackageForProto(firstService[0]), "pb")

	filterGoPackages := func(pkg *gocode.Package) bool {
		if len(o.SourceDirs) == 0 {
			return true
		}

		for _, dir := range o.SourceDirs {
			if strings.HasPrefix(pkg.SourceDir, dir) {
				return true
			}
		}
		return false
	}

	if err := mapperGenerator.VisitGoCode(o.APIGoPackagePath, o.APIDirectory, filterGoPackages); err != nil {
		return err
	}

	if err := mapperGenerator.VisitProto(api); err != nil {
		return err
	}

	goImports := map[string]string{
		"krm": "github.com/GoogleCloudPlatform/k8s-config-connector/apis/" + strings.TrimSuffix(gv.Group, ".cnrm.cloud.google.com") + "/" + gv.Version,
	}
	if err := mapperGenerator.GenerateMappers(goImports); err != nil {
		return err
	}

	addCopyright := true
	writeEmptyFiles := true
	if err := mapperGenerator.WriteFiles(addCopyright, writeEmptyFiles); err != nil {
		return err
	}

	return nil

}

func (o *GenerateMapperOptions) loadAndApplyConfig() error {
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

	if !config.GenerateMapper {
		return fmt.Errorf("mapper generation is disabled for this service in config file %s", o.ConfigFilePath)
	}

	o.IncludeProtoPackages = []string{config.Service}
	o.APIVersion = config.APIVersion
	return nil
}

func (o *GenerateMapperOptions) validate() error {
	if len(o.IncludeProtoPackages) == 0 {
		return fmt.Errorf("ServiceName is required")
	}
	if o.GenerateOptions.ProtoSourcePath == "" {
		return fmt.Errorf("ProtoSourcePath is required")
	}
	if o.APIGoPackagePath == "" {
		return fmt.Errorf("GoPackagePath is required")
	}
	if o.OutputMapperDirectory == "" {
		return fmt.Errorf("OutputMapperDirectory is required")
	}
	if o.APIVersion == "" {
		return fmt.Errorf("APIVersion is required")
	}
	return nil
}
