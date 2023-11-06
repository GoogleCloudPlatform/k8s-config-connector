// Copyright 2023 Google LLC
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

package manifestengine

import (
	"context"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

type Engine struct {
	restMapper    meta.RESTMapper
	dynamicClient dynamic.Interface
}

func NewEngine(restMapper meta.RESTMapper, dynamicClient dynamic.Interface) *Engine {
	return &Engine{
		restMapper:    restMapper,
		dynamicClient: dynamicClient,
	}
}

func (e *Engine) BuildObjects(ctx context.Context, fileName string, definition string, subject *unstructured.Unstructured) ([]*unstructured.Unstructured, error) {
	objects, err := manifest.ParseObjects(ctx, definition)
	if err != nil {
		return nil, err
	}
	var ret []*unstructured.Unstructured
	for _, obj := range objects.GetItems() {
		ret = append(ret, obj.UnstructuredObject())
	}
	return ret, nil
}
