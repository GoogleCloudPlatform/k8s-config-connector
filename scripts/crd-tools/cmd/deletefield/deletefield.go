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

package deletefield

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/crd-tools/pkg/objectvisitor"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func AddCommand(parent *cobra.Command) {
	var opt Options
	cmd := &cobra.Command{
		Use:   "delete-field",
		Short: "Delete field in objects",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(cmd.Context(), opt, args)
		},
	}
	cmd.Flags().StringVar(&opt.Dir, "dir", "", "Directory to process")
	parent.AddCommand(cmd)
}

type Options struct {
	Dir string
}

func (o *Options) Validate() error {
	if o.Dir == "" {
		return fmt.Errorf("--dir is required")
	}
	return nil
}

func (o *Options) InitDefaults() {
	o.Dir = ""
}

type visitor struct {
	args []string
}

func (v *visitor) VisitObject(obj *unstructured.Unstructured) error {
	for _, arg := range v.args {
		// We will "likely" want to enhance this, but starting with what we know
		if arg == "status" {
			unstructured.RemoveNestedField(obj.Object, "status")
		} else {
			return fmt.Errorf("command %q not implemented", arg)
		}
	}
	return nil
}

func Run(ctx context.Context, options Options, args []string) error {
	if err := options.Validate(); err != nil {
		return err
	}

	visitor := &visitor{args: args}
	return objectvisitor.VisitObjectsInDirectory(ctx, options.Dir, visitor)
}
