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
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/scaffold"
	cctemplate "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template/controller"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type GenerateControllerOptions struct {
	*options.GenerateOptions

	Resource options.Resource
}

func (o *GenerateControllerOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().Var(&o.Resource, "resource", "the KRM Kind and the equivalent proto resource separated with a colon.  e.g. for resource google.storage.v1.Bucket, the flag should be `StorageBucket:Bucket`.")
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &GenerateControllerOptions{
		GenerateOptions: baseOptions,
	}

	cmd := &cobra.Command{
		Use:   "generate-controller",
		Short: "generate the direct controller",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opt.Resource == (options.Resource{}) {
				return fmt.Errorf("--resource flag is required")
			}
			if opt.Resource.Kind == "" {
				return fmt.Errorf("need to specify KRM Kind in --resource")
			}
			if opt.Resource.ProtoName == "" {
				return fmt.Errorf("need to specify proto resource name in --resource")
			}

			if baseOptions.APIVersion == "" {
				return fmt.Errorf("--api-version is required")
			}
			_, err := schema.ParseGroupVersion(baseOptions.APIVersion)
			if err != nil {
				return fmt.Errorf("unable to parse --api-version: %w", err)
			}

			if baseOptions.ServiceName == "" {
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

func RunController(ctx context.Context, o *GenerateControllerOptions) error {
	gv, _ := schema.ParseGroupVersion(o.GenerateOptions.APIVersion)
	gcpTokens := strings.Split(o.GenerateOptions.ServiceName, ".")
	version := gcpTokens[len(gcpTokens)-1]
	if version[0] != 'v' {
		return fmt.Errorf("--service does not contain GCP version")
	}
	serviceName := strings.TrimSuffix(gv.Group, ".cnrm.cloud.google.com")
	cArgs := &cctemplate.ControllerArgs{
		KCCService:    serviceName,
		KCCVersion:    gv.Version,
		Kind:          o.Resource.Kind,
		ProtoResource: o.Resource.ProtoName,
		ProtoVersion:  version,
	}
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}

	c := scaffold.NewControllerBuilder(root, serviceName, o.Resource.ProtoName)
	err = errors.Join(err, c.GenerateController(cArgs))
	err = errors.Join(err, c.RegisterController())
	return err
}
