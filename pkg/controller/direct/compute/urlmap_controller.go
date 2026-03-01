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
	"fmt"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
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
	id                   *krm.ComputeURLMapIdentity
	urlMapsClient        *gcp.UrlMapsClient
	regionalUrlMapClient *gcp.RegionUrlMapsClient
	desired              *krm.ComputeURLMap
	actual               *computepb.UrlMap
	reader               client.Reader
}

var _ directbase.Adapter = &urlMapAdapter{}

func (m *urlMapModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeURLMap{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewComputeURLMapIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	adapter := &urlMapAdapter{
		id:      id,
		desired: obj,
		reader:  reader,
	}

	// Get location
	location := id.Region
	if location == "" {
		location = "global"
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	if location == "global" {
		client, err := gcpClient.newUrlMapsClient(ctx)
		if err != nil {
			return nil, err
		}
		adapter.urlMapsClient = client
	} else {
		client, err := gcpClient.newRegionalUrlMapsClient(ctx)
		if err != nil {
			return nil, err
		}
		adapter.regionalUrlMapClient = client
	}
	return adapter, nil
}

func (m *urlMapModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

func (a *urlMapAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeURLMap", "name", a.id)

	urlMap, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeURLMap %q: %w", a.id, err)
	}
	a.actual = urlMap
	return true, nil
}

func (a *urlMapAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	return fmt.Errorf("Create not implemented for ComputeURLMap (Direct controller is in development, use Terraform reconciler)")
}

func (a *urlMapAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	return fmt.Errorf("Update not implemented for ComputeURLMap (Direct controller is in development, use Terraform reconciler)")
}

func (a *urlMapAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("urlMap %s not found", a.id)
	}

	mc := &direct.MapContext{}
	spec := ComputeURLMapSpec_v1beta1_FromProto(mc, a.actual)
	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting urlMap spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetGroupVersionKind(krm.ComputeURLMapGVK)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *urlMapAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	return false, fmt.Errorf("Delete not implemented for ComputeURLMap (Direct controller is in development, use Terraform reconciler)")
}

func (a *urlMapAdapter) get(ctx context.Context) (*computepb.UrlMap, error) {
	projectID := a.id.Project
	location := a.id.Region
	name := a.id.Name

	if location == "" {
		getReq := &computepb.GetUrlMapRequest{
			Project: projectID,
			UrlMap:  name,
		}
		return a.urlMapsClient.Get(ctx, getReq)
	} else {
		getReq := &computepb.GetRegionUrlMapRequest{
			Project: projectID,
			Region:  location,
			UrlMap:  name,
		}
		return a.regionalUrlMapClient.Get(ctx, getReq)
	}
}
