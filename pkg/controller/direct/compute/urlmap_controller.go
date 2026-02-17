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
	id                  *krm.ComputeURLMapIdentity
	urlMapsClient       *gcp.UrlMapsClient
	regionUrlMapsClient *gcp.RegionUrlMapsClient
	desired             *krm.ComputeURLMap
	actual              *computepb.UrlMap
	reader              client.Reader
}

var _ directbase.Adapter = &urlMapAdapter{}

func (m *urlMapModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeURLMap{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ComputeURLMapIdentity)

	urlMapAdapter := &urlMapAdapter{
		id:      id,
		desired: obj,
		reader:  reader,
	}

	// Get GCP client
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	if id.Region == "" {
		urlMapsClient, err := gcpClient.newGlobalUrlMapsClient(ctx)
		if err != nil {
			return nil, err
		}
		urlMapAdapter.urlMapsClient = urlMapsClient
	} else {
		regionUrlMapsClient, err := gcpClient.newRegionUrlMapsClient(ctx)
		if err != nil {
			return nil, err
		}
		urlMapAdapter.regionUrlMapsClient = regionUrlMapsClient
	}
	return urlMapAdapter, nil
}

func (m *urlMapModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *urlMapAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeURLMap", "name", a.id)

	var err error
	urlMap := &computepb.UrlMap{}
	urlMap, err = a.get(ctx)
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
	u := createOp.GetUnstructured()
	var err error

	err = resolveURLMapRefs(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeURLMap", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	urlMap := ComputeURLMapSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	urlMap.Name = direct.LazyPtr(a.id.UrlMap)

	op := &gcp.Operation{}
	if a.id.Region == "" {
		req := &computepb.InsertUrlMapRequest{
			UrlMapResource: urlMap,
			Project:        a.id.Project,
		}
		op, err = a.urlMapsClient.Insert(ctx, req)
	} else {
		req := &computepb.InsertRegionUrlMapRequest{
			UrlMapResource: urlMap,
			Region:         a.id.Region,
			Project:        a.id.Project,
		}
		op, err = a.regionUrlMapsClient.Insert(ctx, req)
	}
	if err != nil {
		return fmt.Errorf("creating ComputeURLMap %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeURLMap %s create failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully created ComputeURLMap", "name", a.id)

	// Get the created resource
	created := &computepb.UrlMap{}
	created, err = a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeURLMap %q: %w", a.id, err)
	}

	status := &krm.ComputeURLMapStatus{
		CreationTimestamp: created.CreationTimestamp,
		Fingerprint:       created.Fingerprint,
		ID:                created.Id,
		SelfLink:          created.SelfLink,
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return setStatus(u, status)
}

func (a *urlMapAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()
	var err error

	err = resolveURLMapRefs(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeURLMap", "name", a.id.UrlMap)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	urlMap := ComputeURLMapSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	urlMap.Name = direct.LazyPtr(a.id.UrlMap)

	op := &gcp.Operation{}
	if a.id.Region == "" {
		req := &computepb.UpdateUrlMapRequest{
			UrlMap:         a.id.UrlMap,
			UrlMapResource: urlMap,
			Project:        a.id.Project,
		}
		op, err = a.urlMapsClient.Update(ctx, req)
	} else {
		req := &computepb.UpdateRegionUrlMapRequest{
			UrlMap:         a.id.UrlMap,
			UrlMapResource: urlMap,
			Region:         a.id.Region,
			Project:        a.id.Project,
		}
		op, err = a.regionUrlMapsClient.Update(ctx, req)
	}
	if err != nil {
		return fmt.Errorf("updating ComputeURLMap %s: %w", a.id, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting ComputeURLMap %s update failed: %w", a.id, err)
	}
	log.V(2).Info("successfully updated ComputeURLMap", "name", a.id)

	// Get the updated resource
	updated := &computepb.UrlMap{}
	updated, err = a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeURLMap %q: %w", a.id.UrlMap, err)
	}

	status := &krm.ComputeURLMapStatus{
		CreationTimestamp: updated.CreationTimestamp,
		Fingerprint:       updated.Fingerprint,
		ID:                updated.Id,
		SelfLink:          updated.SelfLink,
	}
	return setStatus(u, status)
}

func (a *urlMapAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("urlmap %q not found", a.id)
	}

	mc := &direct.MapContext{}
	spec := ComputeURLMapSpec_v1beta1_FromProto(mc, a.actual)
	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting urlmap spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.id.UrlMap)
	u.SetGroupVersionKind(krm.ComputeURLMapGVK)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

// Delete implements the Adapter interface.
func (a *urlMapAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeURLMap", "name", a.id.UrlMap)

	var err error
	op := &gcp.Operation{}
	if a.id.Region == "" {
		req := &computepb.DeleteUrlMapRequest{
			UrlMap:  a.id.UrlMap,
			Project: a.id.Project,
		}
		op, err = a.urlMapsClient.Delete(ctx, req)
	} else {
		req := &computepb.DeleteRegionUrlMapRequest{
			UrlMap:  a.id.UrlMap,
			Region:  a.id.Region,
			Project: a.id.Project,
		}
		op, err = a.regionUrlMapsClient.Delete(ctx, req)
	}
	if err != nil {
		return false, fmt.Errorf("deleting ComputeURLMap %s: %w", a.id.UrlMap, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting ComputeURLMap %s delete failed: %w", a.id.UrlMap, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeURLMap", "name", a.id.UrlMap)
	return true, nil
}

func (a *urlMapAdapter) get(ctx context.Context) (*computepb.UrlMap, error) {
	if a.id.Region == "" {
		getReq := &computepb.GetUrlMapRequest{
			UrlMap:  a.id.UrlMap,
			Project: a.id.Project,
		}
		return a.urlMapsClient.Get(ctx, getReq)
	} else {
		getReq := &computepb.GetRegionUrlMapRequest{
			UrlMap:  a.id.UrlMap,
			Region:  a.id.Region,
			Project: a.id.Project,
		}
		return a.regionUrlMapsClient.Get(ctx, getReq)
	}
}
