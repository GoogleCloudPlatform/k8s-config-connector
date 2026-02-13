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

package compute

import (
	"context"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeURLMapGVK, NewURLMapModel)
}

func NewURLMapModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &urlMapModel{config: config}, nil
}

type urlMapModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &urlMapModel{}

type urlMapAdapter struct {
	id      *krm.ComputeURLMapIdentity
	desired *krm.ComputeURLMap
	reader  client.Reader
}

var _ directbase.Adapter = &urlMapAdapter{}

func (m *urlMapModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	// TODO: Implement
	return nil, nil
}

func (m *urlMapModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Implement
	return nil, nil
}
