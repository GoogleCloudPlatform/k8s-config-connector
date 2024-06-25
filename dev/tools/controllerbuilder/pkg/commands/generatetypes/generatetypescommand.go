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
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/spf13/cobra"
)

type GenerateCRDOptions struct {
	*options.GenerateOptions

	OutputAPIDirectory string
}

func (o *GenerateCRDOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.OutputAPIDirectory, "output-api", o.OutputAPIDirectory, "base directory for writing APIs")
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &GenerateCRDOptions{
		GenerateOptions: baseOptions,
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
		return fmt.Errorf("ServiceName is required")
	}
	if o.GenerateOptions.ProtoSourcePath == "" {
		return fmt.Errorf("ProtoSourcePath is required")
	}

	gv, err := schema.ParseGroupVersion(o.APIVersion)
	if err != nil {
		return fmt.Errorf("APIVersion %q is not valid: %w", o.APIVersion, err)
	}

	api, err := protoapi.LoadProto(o.GenerateOptions.ProtoSourcePath)
	if err != nil {
		return fmt.Errorf("loading proto: %w", err)
	}

	pathForMessage := func(msg protoreflect.MessageDescriptor) (string, bool) {
		fullName := string(msg.FullName())
		if strings.HasSuffix(fullName, "Request") {
			return "", false
		}
		if strings.HasSuffix(fullName, "Response") {
			return "", false
		}

		if !strings.HasPrefix(fullName, o.ServiceName) {
			return "", false
		}

		protoPackagePath := string(msg.ParentFile().Package())
		protoPackagePath = strings.TrimPrefix(protoPackagePath, "google.")
		protoPackagePath = strings.TrimPrefix(protoPackagePath, "cloud.")
		protoPackagePath = strings.TrimSuffix(protoPackagePath, ".v1")
		protoPackagePath = strings.TrimSuffix(protoPackagePath, ".v1beta1")
		goPackage := "apis/" + strings.Join(strings.Split(protoPackagePath, "."), "/") + "/" + gv.Version

		return goPackage, true
	}
	typeGenerator := codegen.NewTypeGenerator(pathForMessage)
	if err := typeGenerator.VisitProto(api); err != nil {
		return err
	}

	if err := typeGenerator.WriteFiles(o.OutputAPIDirectory); err != nil {
		return err
	}

	return nil

}
