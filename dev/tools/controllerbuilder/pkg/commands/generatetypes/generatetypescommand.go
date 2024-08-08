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

package generatetypes

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/scaffold"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/spf13/cobra"
)

type GenerateCRDOptions struct {
	*options.GenerateOptions

	OutputAPIDirectory string
	KindNames          []string
}

func (o *GenerateCRDOptions) InitDefaults() {
}

func (o *GenerateCRDOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.OutputAPIDirectory, "output-api", o.OutputAPIDirectory, "base directory for writing APIs")
	// TODO: 1. only add API and mapper needed by this Kind 2. validate the kind should be in camel case
	cmd.Flags().StringSliceVarP(&o.KindNames, "kinds", "k", nil, "the GCP resource names under the GCP service. This will be used as the ConfigConnecter resource Kind names and should be in camel case")
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &GenerateCRDOptions{
		GenerateOptions: baseOptions,
	}

	opt.InitDefaults()

	cmd := &cobra.Command{
		Use:   "generate-types",
		Short: "generate KRM types for a proto service",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := RunGenerateCRD(ctx, opt); err != nil {
				return err
			}
			return nil
		},
	}

	opt.BindFlags(cmd)

	return cmd
}

func RunGenerateCRD(ctx context.Context, o *GenerateCRDOptions) error {
	if o.ServiceName == "" {
		return fmt.Errorf("`service` is required")
	}
	if o.GenerateOptions.ProtoSourcePath == "" {
		return fmt.Errorf("`proto-source-path` is required")
	}

	if o.KindNames != nil {
		errMsg := []string{}
		for _, k := range o.KindNames {
			if strings.ToLower(k) == k {
				errMsg = append(errMsg, "%q in `kinds` should be CamelCase.")
			}
		}
		if len(errMsg) > 0 {
			return fmt.Errorf(strings.Join(errMsg, "\n"))
		}
	}

	gv, err := schema.ParseGroupVersion(o.APIVersion)
	if err != nil {
		return fmt.Errorf("APIVersion %q is not valid: %w", o.APIVersion, err)
	}

	api, err := protoapi.LoadProto(o.GenerateOptions.ProtoSourcePath)
	if err != nil {
		return fmt.Errorf("loading proto: %w", err)
	}

	goPackage := strings.TrimSuffix(gv.Group, ".cnrm.cloud.google.com") + "/" + gv.Version

	pathForMessage := func(msg protoreflect.MessageDescriptor) (string, bool) {
		fullName := string(msg.FullName())
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
		if !strings.HasPrefix(fullName, o.ServiceName+".") {
			return "", false
		}

		return goPackage, true
	}
	typeGenerator := codegen.NewTypeGenerator(pathForMessage, o.OutputAPIDirectory)
	if err := typeGenerator.VisitProto(api); err != nil {
		return err
	}

	addCopyright := true
	if err := typeGenerator.WriteFiles(addCopyright); err != nil {
		return err
	}

	if o.KindNames != nil {
		if gv.Group == "" {
			return fmt.Errorf("--apiVersion must be specified with --kinds")
		}
		scaffolder := &scaffold.APIScaffolder{
			BaseDir:         o.OutputAPIDirectory,
			GoPackage:       goPackage,
			Group:           gv.Group,
			Version:         gv.Version,
			PackageProtoTag: o.ServiceName,
		}
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
		for _, kind := range o.KindNames {
			if !scaffolder.TypeFileNotExist(kind) {
				fmt.Printf("file %s already exists, skipping\n", scaffolder.GetTypeFile(kind))
				continue
			}
			if err := scaffolder.AddTypeFile(kind); err != nil {
				return fmt.Errorf("add type file %s: %w", scaffolder.GetTypeFile(kind), err)
			}
		}
	}
	return nil
}
