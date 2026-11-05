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

package generatecontroller

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/scaffold"
	cctemplate "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template/controller"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
)

type GenerateControllerOptions struct {
	*options.GenerateOptions

	ServiceName         string
	Resources           options.ResourceList
	PrepopulateResource bool
	PlaceholderOnly     bool
	RegionalEndpoint    string
	Overwrite           bool
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &GenerateControllerOptions{
		GenerateOptions: baseOptions,
	}

	cmd := &cobra.Command{
		Use:   "generate-controller",
		Short: "generate the direct controller",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(opt.Resources) == 0 {
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
			if err := RunController(ctx, opt); err != nil {
				return err
			}
			return nil
		},
	}
	opt.BindFlags(cmd)

	return cmd
}

func (o *GenerateControllerOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().Var(&o.Resources, "resource", "the KRM Kind and the equivalent proto resource separated with a colon. e.g. for resource google.storage.v1.Bucket, the flag should be `StorageBucket:Bucket`. Can be specified multiple times.")
	cmd.Flags().StringVarP(&o.ServiceName, "service", "s", "", "the GCP service name")
	cmd.Flags().BoolVar(&o.PrepopulateResource, "prepopulate-resource", false, "if true, run proto- and KRM-driven deterministic passes to generate a complete direct controller (leaving existing handwritten controllers unmodified unless --overwrite is passed)")
	cmd.Flags().BoolVar(&o.PlaceholderOnly, "placeholder-only", false, "if true, only generate the Pass 1 base CRUD placeholder without proto-driven deterministic specialization passes")
	cmd.Flags().StringVar(&o.RegionalEndpoint, "regional-endpoint", "", "optional regional endpoint pattern ('prefix' for {loc}-{service} or 'rep' for {service}.{loc}.rep.googleapis.com)")
	cmd.Flags().BoolVar(&o.Overwrite, "overwrite", false, "if true, overwrite an existing controller file")
}

func RunController(ctx context.Context, o *GenerateControllerOptions) error {
	gv, err := schema.ParseGroupVersion(o.GenerateOptions.APIVersion)
	if err != nil {
		return fmt.Errorf("unable to parse --api-version: %w", err)
	}
	gcpTokens := strings.Split(o.ServiceName, ",")
	firstService := strings.TrimSpace(gcpTokens[0])
	svcParts := strings.Split(firstService, ".")
	version := svcParts[len(svcParts)-1]
	if len(version) == 0 || version[0] != 'v' {
		return fmt.Errorf("--service does not contain GCP version")
	}
	serviceName := strings.TrimSuffix(gv.Group, ".cnrm.cloud.google.com")
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}

	goPackage := serviceName + "/" + gv.Version
	scaffolder := &scaffold.APIScaffolder{
		BaseDir:         root + "/apis/",
		GoPackage:       goPackage,
		Group:           gv.Group,
		Version:         gv.Version,
		PackageProtoTag: firstService,
	}

	var loadedProto *protoapi.Proto
	if api, err := protoapi.LoadProto(o.GenerateOptions.ProtoSourcePath, o.GenerateOptions.ProtoOverlayPath); err != nil {
		klog.Warningf("could not load proto %q, scaffolding will guess the resource name shape: %v",
			o.GenerateOptions.ProtoSourcePath, err)
	} else {
		loadedProto = api
		scaffolder.Proto = api
	}

	var allErrs error
	for _, resource := range o.Resources {
		if err := runControllerForResource(o, scaffolder, loadedProto, root, serviceName, gv.Version, version, resource); err != nil {
			allErrs = errors.Join(allErrs, err)
		}
	}
	return allErrs
}

func runControllerForResource(
	o *GenerateControllerOptions,
	scaffolder *scaffold.APIScaffolder,
	loadedProto *protoapi.Proto,
	root, serviceName, krmVersion, gcpVersion string,
	resource options.Resource,
) error {
	cArgs := cctemplate.NewBaseControllerArgs(
		serviceName,
		krmVersion,
		resource.Kind,
		resource.ProtoMessageName(),
		gcpVersion,
	)

	if !o.PrepopulateResource && !o.PlaceholderOnly {
		if scaffolder.RefsFileExist(resource) {
			klog.V(1).Infof("file %s already exists, skipping\n", scaffolder.PathToRefsFile(resource))
		} else if err := scaffolder.AddRefsFile(resource); err != nil {
			return fmt.Errorf("add refs file %s: %w", scaffolder.PathToRefsFile(resource), err)
		}
		if scaffolder.IdentityFileExist(resource) {
			klog.V(1).Infof("file %s already exists, skipping\n", scaffolder.PathToIdentityFile(resource))
		} else if err := scaffolder.AddIdentityFile(resource); err != nil {
			return fmt.Errorf("add identity file %s: %w", scaffolder.PathToIdentityFile(resource), err)
		}
	}

	prepopulate := o.PrepopulateResource && !o.PlaceholderOnly
	if prepopulate {
		krmInfo := scaffold.InspectKRMPackageForController(
			root,
			serviceName,
			krmVersion,
			resource.Kind,
			resource.ProtoMessageName(),
		)
		resAPI := resolveResourceAPI(loadedProto, o.ServiceName, resource.ProtoName)
		scaffold.ApplyDeterministicPasses(cArgs, resAPI, krmInfo, o.RegionalEndpoint)
	}

	c := scaffold.NewControllerBuilder(root, serviceName, resource.Kind).
		WithOverwrite(o.Overwrite).
		WithPrepopulateResource(prepopulate)
	var err error
	err = errors.Join(err, c.GenerateController(cArgs))
	err = errors.Join(err, c.RegisterController())
	return err
}

func resolveResourceAPI(p *protoapi.Proto, servicePkgTag, protoName string) *protoapi.ResourceAPI {
	if p == nil || protoName == "" {
		return nil
	}
	var candidates []string
	if strings.Contains(protoName, ".") {
		candidates = append(candidates, protoName)
	} else {
		for _, rawPkg := range strings.Split(servicePkgTag, ",") {
			pkg := strings.TrimSpace(rawPkg)
			if pkg != "" {
				candidates = append(candidates, pkg+"."+protoName)
			}
		}
	}
	for _, fullName := range candidates {
		if api, err := p.ResourceAPIByName(protoreflect.FullName(fullName)); err == nil {
			return api
		}
	}
	return nil
}
