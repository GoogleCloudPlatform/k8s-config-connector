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
	"os"
	"strings"
	"unicode"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/scaffold"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/spf13/cobra"
)

type GenerateCRDOptions struct {
	*options.GenerateOptions

	OutputAPIDirectory string
	ResourceKindName   string
	ResourceProtoName  string
}

func (o *GenerateCRDOptions) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return nil
	}
	o.ProtoSourcePath = root + "/dev/tools/proto-to-mapper/build/googleapis.pb"
	o.OutputAPIDirectory = root + "/apis/"
	return nil
}

func (o *GenerateCRDOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.OutputAPIDirectory, "output-api", o.OutputAPIDirectory, "base directory for writing APIs")
	cmd.Flags().StringVarP(&o.ResourceProtoName, "proto-resource", "p", "", "the GCP resource proto name. It should match the name in the proto apis. i.e. For resource google.storage.v1.bucket, the `--proto-resource` should be `bucket`. If `--kind` is not given, the `--proto-resource` value will also be used as the kind name with a capital letter `Storage`.")
	cmd.Flags().StringVarP(&o.ResourceKindName, "kind", "k", "", "the KCC resource Kind. requires `--proto-resource`.")
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &GenerateCRDOptions{
		GenerateOptions: baseOptions,
	}

	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

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
	if o.ResourceProtoName == "" {
		return fmt.Errorf("`--proto-resource` is required")
	}
	if o.ResourceKindName == "" {
		return fmt.Errorf("`--kind` is required")
	}
	o.ResourceProtoName = capitalizeFirstRune(o.ResourceProtoName)

	gv, err := schema.ParseGroupVersion(o.APIVersion)
	if err != nil {
		return fmt.Errorf("APIVersion %q is not valid: %w", o.APIVersion, err)
	}

	api, err := protoapi.LoadProto(o.GenerateOptions.ProtoSourcePath)
	if err != nil {
		return fmt.Errorf("loading proto: %w", err)
	}

	goPackage := strings.TrimSuffix(gv.Group, ".cnrm.cloud.google.com") + "/" + gv.Version

	if gv.Group == "" {
		return fmt.Errorf("--api-version must be specified with --kind")
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

	resources, err := parseResourcesToGenerate(o.ResourceKindName, o.ResourceProtoName)
	if err != nil {
		return err
	}

	// generate types file
	var protoFQNs []string
	for _, r := range resources {
		protoFQNs = append(protoFQNs, o.ServiceName+"."+r.protoName)
	}
	typeGenerator := codegen.NewTypeGenerator(goPackage, o.OutputAPIDirectory, protoFQNs)
	if err := typeGenerator.VisitProto(api); err != nil {
		return err
	}
	addCopyright := true
	if err := typeGenerator.WriteFiles(addCopyright); err != nil {
		return err
	}

	// add type template files
	for _, r := range resources {
		kind := r.kind
		if !scaffolder.TypeFileNotExist(kind) {
			fmt.Printf("file %s already exists, skipping\n", scaffolder.GetTypeFile(kind))
		} else {
			err := scaffolder.AddTypeFile(kind, r.protoName)
			if err != nil {
				return fmt.Errorf("add type file %s: %w", scaffolder.GetTypeFile(kind), err)
			}
		}
	}
	return nil
}

// resource represents a resource to be generated
type resource struct {
	kind      string
	protoName string
}

func parseResourcesToGenerate(kindNames, protoNames string) ([]resource, error) {
	kinds := strings.Split(kindNames, ",")
	protos := strings.Split(protoNames, ",")
	if len(kinds) != len(protos) {
		return nil, fmt.Errorf("unexpected number of Kind and ProtoName provided, got %d Kinds but %d Proto names", len(kinds), len(protos))
	}
	resources := make([]resource, len(kinds))
	for i := range kinds {
		resources[i] = resource{
			kind:      strings.TrimSpace(kinds[i]),
			protoName: strings.TrimSpace(protos[i]),
		}
	}
	return resources, nil
}

func capitalizeFirstRune(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
