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

package e2e

import "k8s.io/klog/v2"

// fieldReplacement is a visitor that replaces the field at FindPath in a object,
// setting the value to Replace.
type fieldReplacement struct {
	FindPath string
	Replace  any
}

// DoReplacement is the entrypoint for the fieldVisitor
func (f *fieldReplacement) DoReplacement(m map[string]any) {
	f.visit(m, "")
}

func (f *fieldReplacement) visit(v any, fieldPath string) any {
	switch v := v.(type) {
	case map[string]any:
		return f.visitMap(v, fieldPath)
	case []any:
		return f.visitSlice(v, fieldPath)
	case string, int, int32, int64, bool, float32, float64:
		return f.visitLeaf(v, fieldPath)
	default:
		klog.Fatalf("unsupported type %T", v)
		return nil
	}
}

func (f *fieldReplacement) visitSlice(s []any, fieldPath string) []any {
	for i, v := range s {
		s[i] = f.visit(v, fieldPath+"[]")
	}
	return s
}

func (f *fieldReplacement) visitMap(m map[string]any, fieldPath string) map[string]any {
	for k, v := range m {
		m[k] = f.visit(v, fieldPath+"."+k)
	}
	return m
}

func (f *fieldReplacement) visitLeaf(v any, fieldPath string) any {
	if f.FindPath == fieldPath {
		return f.Replace
	}
	return v
}
