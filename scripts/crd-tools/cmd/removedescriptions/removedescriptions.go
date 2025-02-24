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

package removedescriptions

import (
	"context"
	"errors"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/crd-tools/pkg/objectvisitor"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func AddCommand(parent *cobra.Command) {
	var opt Options
	cmd := &cobra.Command{
		Use:   "remove-descriptions",
		Short: "Remove descriptions in objects, for easier schema-only comparison.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(cmd.Context(), opt)
		},
	}
	cmd.Flags().StringVar(&opt.Dir, "dir", "", "Directory to process")
	parent.AddCommand(cmd)
}

type Options struct {
	Dir string
}

func (o *Options) Validate() error {
	return nil
}

func (o *Options) InitDefaults() {
	o.Dir = ""
}

type visitor struct {
	errors []error
}

func (v *visitor) VisitObject(crd *unstructured.Unstructured) error {
	v.visitMap(crd.Object)
	return errors.Join(v.errors...)
}

func (v *visitor) visitMap(m map[string]any) map[string]any {
	delete(m, "description")
	for k, val := range m {
		m[k] = v.visitAny(val)
	}
	return m
}

func (v *visitor) visitSlice(s []any) []any {
	for i, val := range s {
		s[i] = v.visitAny(val)
	}
	return s
}

func (v *visitor) visitAny(val any) any {
	switch val := val.(type) {
	case map[string]any:
		return v.visitMap(val)
	case []any:
		return v.visitSlice(val)
	case string:
		return val
	case bool, int32, int64, float32, float64:
		return val
	case nil:
		return val
	default:
		v.errors = append(v.errors, fmt.Errorf("unexpected type %T", val))
		return val
	}
}

func Run(ctx context.Context, options Options) error {
	if err := options.Validate(); err != nil {
		return err
	}

	visitor := &visitor{}
	if options.Dir != "" {
		return objectvisitor.VisitObjectsInDirectory(ctx, options.Dir, visitor)
	} else {
		return objectvisitor.VisitObjectsFromStdin(ctx, visitor)
	}
}
