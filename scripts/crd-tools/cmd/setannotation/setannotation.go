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

package setannotation

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/crd-tools/pkg/objectvisitor"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func AddCommand(parent *cobra.Command) {
	var opt Options
	cmd := &cobra.Command{
		Use:   "set-annotation",
		Short: "Set annotation in objects",
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
		fields := strings.SplitN(arg, "=", 2)
		if len(fields) != 2 {
			return fmt.Errorf("invalid argument %q", arg)
		}
		key, value := fields[0], fields[1]
		annotations := obj.GetAnnotations()
		if annotations == nil {
			annotations = make(map[string]string)
		}
		annotations[key] = value
		obj.SetAnnotations(annotations)
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
