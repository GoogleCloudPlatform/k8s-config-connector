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

package mustache

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

func Interpret(ctx context.Context, raw []byte, activation *Activation) ([]*unstructured.Unstructured, error) {
	objects, err := manifest.ParseObjects(ctx, string(raw))
	if err != nil {
		return nil, fmt.Errorf("parsing yaml: %w", err)
	}

	var out []*unstructured.Unstructured

	for _, obj := range objects.Items {
		u := obj.UnstructuredObject()
		walker := &YAMLWalker{
			activation: activation,
			obj:        u,
		}
		if err := walker.Walk(); err != nil {
			return nil, err
		}

		out = append(out, u)
	}

	return out, nil
}

type YAMLWalker struct {
	errors     []error
	obj        *unstructured.Unstructured
	activation *Activation
}

func (v *YAMLWalker) visitString(s string) (string, error) {
	ctx := context.TODO()

	if !strings.Contains(s, "${") {
		return s, nil
	}

	parser := Parser{}
	parser.Init(s)

	expressionList, err := parser.ParseExpressionList()
	if err != nil {
		return "", fmt.Errorf("error parsing expression list: %w", err)
	}

	result, err := expressionList.Eval(ctx, v.activation)
	if err != nil {
		return "", fmt.Errorf("error evaluating expression %q: %w", s, err)
	}
	return result, nil
}

func (w *YAMLWalker) Walk() error {
	w.walkMap(w.obj.Object)
	if len(w.errors) != 0 {
		return errors.Join(w.errors...)
	}
	return nil
}

func (w *YAMLWalker) walkMap(m map[string]interface{}) {
	for k, v := range m {
		switch v := v.(type) {
		// Most primitive values cannot contain an expression
		case float64:
		case int64:
		case bool:

		case string:
			v, err := w.visitString(v)
			if err != nil {
				w.errors = append(w.errors, err)
				return
			}
			m[k] = v

		case map[string]interface{}:
			w.walkMap(v)

		case []interface{}:
			w.walkSlice(v)

		default:
			klog.Fatalf("unhandled type %T", v)
		}
	}
}

func (w *YAMLWalker) walkSlice(m []interface{}) {
	for i, v := range m {
		switch v := v.(type) {
		// Most primitive values cannot contain an expression
		case float64:
		case int64:
		case bool:

		case string:
			v, err := w.visitString(v)
			if err != nil {
				w.errors = append(w.errors, err)
				return
			}
			m[i] = v

		case map[string]interface{}:
			w.walkMap(v)

		case []interface{}:
			w.walkSlice(v)

		default:
			klog.Fatalf("unhandled type %T", v)
		}
	}
}
