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
	"k8s.io/klog/v2"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type GenerateCRDOptions struct {
	*options.GenerateOptions

	OutputAPIDirectory string
	SkipScaffoldFiles  bool

	Resources ResourceList
}

type Resource struct {
	Kind      string
	ProtoName string
}

type ResourceList []Resource

var _ pflag.Value = &ResourceList{}

func (r *ResourceList) Type() string { return "resources" }

func (r *ResourceList) String() string {
	var sb strings.Builder
	for _, res := range *r {
		fmt.Fprintf(&sb, "%s:%s", res.Kind, res.ProtoName)
	}
	return sb.String()
}

func (r *ResourceList) Set(s string) error {
	tokens := strings.Split(s, ":")
	if len(tokens) != 2 || tokens[0] == "" || tokens[1] == "" {
		return fmt.Errorf("expected [KRMKind]:[ProtoResourceName], got %q", s)
	}
	*r = append(*r, Resource{
		Kind:      tokens[0],
		ProtoName: tokens[1],
	})
	return nil
}

func (o *GenerateCRDOptions) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return nil
	}
	o.OutputAPIDirectory = root + "/apis/"
	return nil
}

func (o *GenerateCRDOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.OutputAPIDirectory, "output-api", o.OutputAPIDirectory, "base directory for writing APIs")
	cmd.Flags().Var(&o.Resources, "resource", "the KRM Kind and the equivalent proto resource separated with a colon.  e.g. for resource google.storage.v1.Bucket, the flag should be `StorageBucket:Bucket`.  Can be specified multiple times.")
	cmd.Flags().BoolVar(&o.SkipScaffoldFiles, "skip-scaffold-files", false, "skip generating scaffold files (types, refs, and identity)")
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
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opt.ServiceName == "" {
				return fmt.Errorf("`service` is required")
			}
			if opt.GenerateOptions.ProtoSourcePath == "" {
				return fmt.Errorf("`proto-source-path` is required")
			}
			if len(opt.Resources) == 0 {
				return fmt.Errorf("`--resource` is required")
			}
			return nil
		},
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
	log := klog.FromContext(ctx)

	// o.ResourceProtoName = capitalizeFirstRune(o.ResourceProtoName)

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

	typeGenerator := codegen.NewTypeGenerator(goPackage, o.OutputAPIDirectory, api)

	for _, resource := range o.Resources {
		resourceProtoFullName := o.ServiceName + "." + resource.ProtoName
		log.Info("visting proto", "name", resourceProtoFullName)
		if err := typeGenerator.VisitProto(resourceProtoFullName); err != nil {
			return err
		}

		if o.SkipScaffoldFiles {
			log.Info("skipping scaffolding type, refs and identity files", "resource", resource.ProtoName)
		} else {
			kind := resource.Kind
			if !scaffolder.TypeFileNotExist(resource.ProtoName) {
				fmt.Printf("file %s already exists, skipping\n", scaffolder.PathToTypeFile(resource.ProtoName))
			} else {
				err := scaffolder.AddTypeFile(resource.ProtoName, kind)
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

	if err := typeGenerator.WriteVisitedMessages(); err != nil {
		return err
	}

	addCopyright := true
	if err := typeGenerator.WriteFiles(addCopyright); err != nil {
		return err
	}

	return nil
}

func capitalizeFirstRune(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
