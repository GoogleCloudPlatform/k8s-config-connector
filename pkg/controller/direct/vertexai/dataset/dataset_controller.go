// Copyright 2026 Google LLC
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

package dataset

import (
	"context"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func init() {
	registry.RegisterModel(krm.VertexAIDatasetGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{}, nil
}

type model struct{}

// AdapterForObject implements the Model interface.
func (m *model) AdapterForObject(ctx context.Context, reader *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	return &Adapter{}, nil
}

// AdapterForURL implements the Model interface.
func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return &Adapter{}, nil
}

// Adapter implements the Adapter interface.
type Adapter struct{}

func (a *Adapter) Find(ctx context.Context) (bool, error)                                 { return false, nil }
func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error { return nil }
func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error { return nil }
func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error)         { return nil, nil }
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	return false, nil
}
