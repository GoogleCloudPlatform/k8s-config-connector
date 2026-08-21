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

package reflowdescriptions

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/crd-tools/pkg/objectvisitor"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func AddCommand(parent *cobra.Command) {
	var opt Options
	cmd := &cobra.Command{
		Use:   "reflow-descriptions",
		Short: "Reflow descriptions in objects to match the format previously generated.",
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
	if o.Dir == "" {
		return fmt.Errorf("--dir is required")
	}
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
	for k, val := range m {
		if k == "description" {
			s, ok := val.(string)
			if ok {
				val = v.visitDescription(s)
			}
		}
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

func (v *visitor) visitDescription(s string) string {
	// Reflow the text to be more similar to what we had before;
	// remove line breaks and trim excess white space around those line breaks.

	// However, if the string has a double line break, we assume that
	// it is nicely formatted, and try to preserve that.
	if strings.Contains(s, "\n\n") {
		return s
	}

	var out strings.Builder
	for {
		ix := strings.Index(s, "\n")
		if ix == -1 {
			out.WriteString(s)
			break
		}
		if ix == 0 {
			s = s[1:]
			continue
		}
		if ix == len(s)-1 {
			s = s[:ix]
			continue
		}
		head := s[:ix]
		tail := s[ix+1:]

		head = strings.TrimSpace(head)
		s = strings.TrimSpace(tail)
		out.WriteString(head)
		out.WriteString(" ")
	}
	return out.String()
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
	return objectvisitor.VisitObjectsInDirectory(ctx, options.Dir, visitor)
}
