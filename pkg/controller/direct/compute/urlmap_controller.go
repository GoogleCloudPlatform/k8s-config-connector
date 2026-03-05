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

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ComputeURLMapIdentity)

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	adapter := &urlMapAdapter{
		id:      id,
		desired: obj,
		reader:  reader,
	}

	if id.Region == "" {
		c, err := gcpClient.newUrlMapsClient(ctx)
		if err != nil {
			return nil, err
		}
		adapter.urlMapsClient = c
	} else {
		c, err := gcpClient.newRegionalUrlMapsClient(ctx)
		if err != nil {
			return nil, err
		}
		adapter.regionalUrlMapClient = c
	}

	return adapter, nil
}

func (m *urlMapModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *urlMapAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeURLMap", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeURLMap %q: %w", a.id, err)
	}
	a.actual = actual
	return true, nil
}

func (a *urlMapAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeURLMap", "name", a.id)

	if err := resolveURLMapRefs(ctx, a.reader, a.desired); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desiredpb := ComputeURLMapSpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredpb.Name = &a.id.Name

	var op *gcp.Operation
	var err error

	if a.id.Region == "" {
		req := &computepb.InsertUrlMapRequest{
			Project:        a.id.Project,
			UrlMapResource: desiredpb,
		}
		op, err = a.urlMapsClient.Insert(ctx, req)
	} else {
		req := &computepb.InsertRegionUrlMapRequest{
			Project:        a.id.Project,
			Region:         a.id.Region,
			UrlMapResource: desiredpb,
		}
		op, err = a.regionalUrlMapClient.Insert(ctx, req)
	}

	if err != nil {
		return fmt.Errorf("creating ComputeURLMap %s: %w", a.id, err)
	}

	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for ComputeURLMap %s create: %w", a.id, err)
	}

	actual, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting created ComputeURLMap %s: %w", a.id, err)
	}

	status := &krm.ComputeURLMapStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.ObservedState = &krm.ComputeURLMapObservedState{
		CreationTimestamp: actual.CreationTimestamp,
		SelfLink:          actual.SelfLink,
		ID:                actual.Id,
		Fingerprint:       actual.Fingerprint,
	}
	if actual.Id != nil {
		mapID := int64(*actual.Id)
		status.MapID = &mapID
	}

	return setStatus(u, status)
}

func (a *urlMapAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeURLMap", "name", a.id)

	if err := resolveURLMapRefs(ctx, a.reader, a.desired); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desiredpb := ComputeURLMapSpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// name is immutable
	desiredpb.Name = &a.id.Name
	// fingerprint is needed for optimistic locking
	desiredpb.Fingerprint = a.actual.Fingerprint

	var op *gcp.Operation
	var err error

	if a.id.Region == "" {
		req := &computepb.UpdateUrlMapRequest{
			Project:        a.id.Project,
			UrlMap:         a.id.Name,
			UrlMapResource: desiredpb,
		}
		op, err = a.urlMapsClient.Update(ctx, req)
	} else {
		req := &computepb.UpdateRegionUrlMapRequest{
			Project:        a.id.Project,
			Region:         a.id.Region,
			UrlMap:         a.id.Name,
			UrlMapResource: desiredpb,
		}
		op, err = a.regionalUrlMapClient.Update(ctx, req)
	}

	if err != nil {
		return fmt.Errorf("updating ComputeURLMap %s: %w", a.id, err)
	}

	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for ComputeURLMap %s update: %w", a.id, err)
	}

	actual, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting updated ComputeURLMap %s: %w", a.id, err)
	}

	status := &krm.ComputeURLMapStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.ObservedState = &krm.ComputeURLMapObservedState{
		CreationTimestamp: actual.CreationTimestamp,
		SelfLink:          actual.SelfLink,
		ID:                actual.Id,
		Fingerprint:       actual.Fingerprint,
	}
	if actual.Id != nil {
		mapID := int64(*actual.Id)
		status.MapID = &mapID
	}

	return setStatus(u, status)
}

func (a *urlMapAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("ComputeURLMap %q not found", a.id)
	}

	mc := &direct.MapContext{}
	spec := ComputeURLMapSpec_v1beta1_FromProto(mc, a.actual)
	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting ComputeURLMap spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.id.Name)
	u.SetGroupVersionKind(krm.ComputeURLMapGVK)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *urlMapAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeURLMap", "name", a.id)

	var op *gcp.Operation
	var err error

	if a.id.Region == "" {
		req := &computepb.DeleteUrlMapRequest{
			Project: a.id.Project,
			UrlMap:  a.id.Name,
		}
		op, err = a.urlMapsClient.Delete(ctx, req)
	} else {
		req := &computepb.DeleteRegionUrlMapRequest{
			Project: a.id.Project,
			Region:  a.id.Region,
			UrlMap:  a.id.Name,
		}
		op, err = a.regionalUrlMapClient.Delete(ctx, req)
	}

	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting ComputeURLMap %s: %w", a.id, err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for ComputeURLMap %s delete: %w", a.id, err)
	}

	return true, nil
}

func (a *urlMapAdapter) get(ctx context.Context) (*computepb.UrlMap, error) {
	if a.id.Region == "" {
		getReq := &computepb.GetUrlMapRequest{
			Project: a.id.Project,
			UrlMap:  a.id.Name,
		}
		return a.urlMapsClient.Get(ctx, getReq)
	} else {
		getReq := &computepb.GetRegionUrlMapRequest{
			Project: a.id.Project,
			Region:  a.id.Region,
			UrlMap:  a.id.Name,
		}
		return a.regionalUrlMapClient.Get(ctx, getReq)
	}
}
