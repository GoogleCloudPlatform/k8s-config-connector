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

package networksecurity

import (
	"context"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityMirroringDeploymentGVK, NewMirroringDeploymentModel)
}

func NewMirroringDeploymentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &mirroringDeploymentModel{}, nil
}

var _ directbase.Model = &mirroringDeploymentModel{}

type mirroringDeploymentModel struct {
}

func (m *mirroringDeploymentModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	return &mirroringDeploymentAdapter{}, nil
}

func (m *mirroringDeploymentModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return &mirroringDeploymentAdapter{}, nil
}

type mirroringDeploymentAdapter struct {
}

var _ directbase.Adapter = &mirroringDeploymentAdapter{}

func (a *mirroringDeploymentAdapter) Find(ctx context.Context) (bool, error) {
	return false, nil
}

func (a *mirroringDeploymentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	return nil
}

func (a *mirroringDeploymentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	return nil
}

func (a *mirroringDeploymentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *mirroringDeploymentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	return false, nil
}
