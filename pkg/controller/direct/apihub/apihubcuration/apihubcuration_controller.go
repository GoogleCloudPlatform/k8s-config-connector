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

package apihubcuration

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func init() {
	registry.RegisterModel(krm.APIHubCurationGVK, newAPIHubCurationModel)
}

func newAPIHubCurationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

type model struct {
	config *config.ControllerConfig
}

func (m *model) client(ctx context.Context) error {
	return nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krm.APIHubCuration{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(op.Object.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, op.Reader)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		id:    id.(*krm.APIHubCurationIdentity),
		model: m,
		obj:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id    *krm.APIHubCurationIdentity
	model *model
	obj   *krm.APIHubCuration
}

// Delete removes the GCP object.
func (a *Adapter) Delete(ctx context.Context, op *directbase.DeleteOperation) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

// Find returns whether the corresponding GCP object was found.
func (a *Adapter) Find(ctx context.Context) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

// Create creates a new GCP object.
func (a *Adapter) Create(ctx context.Context, op *directbase.CreateOperation) error {
	return fmt.Errorf("not implemented")
}

// Update updates an existing GCP object.
func (a *Adapter) Update(ctx context.Context, op *directbase.UpdateOperation) error {
	return fmt.Errorf("not implemented")
}

// Export returns the state of the GCP object in a form that can be applied to Kubernetes.
func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, fmt.Errorf("not implemented")
}
