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

package compute

import (
	"context"
	"fmt"
	"reflect"

	"k8s.io/klog/v2"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var ComputeURLMapGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputeURLMap",
}

func init() {
	registry.RegisterModel(ComputeURLMapGVK, NewComputeURLMapModel)
}

func NewComputeURLMapModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeURLMapModel{config: config}, nil
}

type computeURLMapModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &computeURLMapModel{}

type computeURLMapAdapter struct {
	id            *krm.ComputeURLMapIdentity
	urlMapsClient *gcp.RegionUrlMapsClient
	globalClient  *gcp.UrlMapsClient
	desired       *krm.ComputeURLMap
	actual        *computepb.UrlMap
	reader        client.Reader
}

var _ directbase.Adapter = &computeURLMapAdapter{}

func (m *computeURLMapModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ComputeURLMap{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ComputeURLMapIdentity)

	adapter := &computeURLMapAdapter{
		id:      id,
		desired: obj,
		reader:  reader,
	}

	// Get region
	region := id.ParentID.Location

	// Get GCP client
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	if region == "global" {
		globalClient, err := gcpClient.urlMapsClient(ctx)
		if err != nil {
			return nil, err
		}
		adapter.globalClient = globalClient
	} else {
		urlMapsClient, err := gcpClient.regionUrlMapsClient(ctx)
		if err != nil {
			return nil, err
		}
		adapter.urlMapsClient = urlMapsClient
	}
	return adapter, nil
}

func (m *computeURLMapModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

func (a *computeURLMapAdapter) Find(ctx context.Context) (bool, error) {
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

func (a *computeURLMapAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()
	var err error

	err = resolveComputeURLMapRefs(ctx, a.reader, a.desired)
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
	urlMap.Name = direct.LazyPtr(a.id.ResourceID)

	op := &gcp.Operation{}
	if a.id.ParentID.Location == "global" {
		req := &computepb.InsertUrlMapRequest{
			UrlMapResource: urlMap,
			Project:        a.id.ParentID.ProjectID,
		}
		op, err = a.globalClient.Insert(ctx, req)
	} else {
		req := &computepb.InsertRegionUrlMapRequest{
			UrlMapResource: urlMap,
			Region:         a.id.ParentID.Location,
			Project:        a.id.ParentID.ProjectID,
		}
		op, err = a.urlMapsClient.Insert(ctx, req)
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
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeURLMap %q: %w", a.id, err)
	}

	status := &krm.ComputeURLMapStatus{
		CreationTimestamp: created.CreationTimestamp,
		Fingerprint:       created.Fingerprint,
		MapId:             int64PtrFromUint64Ptr(created.Id),
		SelfLink:          created.SelfLink,
	}
	return setStatus(u, status)
}

func int64PtrFromUint64Ptr(v *uint64) *int64 {
	if v == nil {
		return nil
	}
	val := int64(*v)
	return &val
}

func (a *computeURLMapAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()
	var err error

	if a.id.ResourceID == "" {
		return fmt.Errorf("resourceID is empty")
	}

	err = resolveComputeURLMapRefs(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeURLMap", "name", a.id.ResourceID)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	urlMap := ComputeURLMapSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	urlMap.Name = direct.LazyPtr(a.id.ResourceID)
	// Fingerprint is required for update
	urlMap.Fingerprint = a.actual.Fingerprint

	// Compare desired spec with actual spec to avoid diffs on output-only fields
	actualSpec := ComputeURLMapSpec_v1beta1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Ignore Location as it is not in the proto
	actualSpec.Location = desired.Spec.Location

	// Ignore ResourceID if it is not set in desired
	if desired.Spec.ResourceID == nil {
		actualSpec.ResourceID = nil
	}

	if !reflect.DeepEqual(desired.Spec, *actualSpec) {
		op := &gcp.Operation{}
		if a.id.ParentID.Location == "global" {
			req := &computepb.UpdateUrlMapRequest{
				UrlMap:         a.id.ResourceID,
				UrlMapResource: urlMap,
				Project:        a.id.ParentID.ProjectID,
			}
			op, err = a.globalClient.Update(ctx, req)
		} else {
			req := &computepb.UpdateRegionUrlMapRequest{
				UrlMap:         a.id.ResourceID,
				UrlMapResource: urlMap,
				Project:        a.id.ParentID.ProjectID,
				Region:         a.id.ParentID.Location,
			}
			op, err = a.urlMapsClient.Update(ctx, req)
		}
		if err != nil {
			return fmt.Errorf("updating ComputeURLMap %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting ComputeURLMap %s update failed: %w", a.id, err)
			}
		}
		log.V(2).Info("successfully updated ComputeURLMap", "name", a.id)
	}

	// Get the updated resource
	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeURLMap %q: %w", a.id.ResourceID, err)
	}

	status := &krm.ComputeURLMapStatus{
		CreationTimestamp: updated.CreationTimestamp,
		Fingerprint:       updated.Fingerprint,
		MapId:             int64PtrFromUint64Ptr(updated.Id),
		SelfLink:          updated.SelfLink,
	}
	return setStatus(u, status)
}

func (a *computeURLMapAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("urlMap %q not found", a.id)
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
	u.SetName(a.id.ResourceID)
	u.SetGroupVersionKind(ComputeURLMapGVK)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

// Delete implements the Adapter interface.
func (a *computeURLMapAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	if a.id.ResourceID == "" {
		return false, fmt.Errorf("resourceID is empty")
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeURLMap", "name", a.id.ResourceID)

	var err error
	op := &gcp.Operation{}
	if a.id.ParentID.Location == "global" {
		req := &computepb.DeleteUrlMapRequest{
			UrlMap:  a.id.ResourceID,
			Project: a.id.ParentID.ProjectID,
		}
		op, err = a.globalClient.Delete(ctx, req)
	} else {
		req := &computepb.DeleteRegionUrlMapRequest{
			UrlMap:  a.id.ResourceID,
			Region:  a.id.ParentID.Location,
			Project: a.id.ParentID.ProjectID,
		}
		op, err = a.urlMapsClient.Delete(ctx, req)
	}
	if err != nil {
		return false, fmt.Errorf("deleting ComputeURLMap %s: %w", a.id.ResourceID, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting ComputeURLMap %s delete failed: %w", a.id.ResourceID, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeURLMap", "name", a.id.ResourceID)
	return true, nil
}

func (a *computeURLMapAdapter) get(ctx context.Context) (*computepb.UrlMap, error) {
	if a.id.ParentID.Location == "global" {
		getReq := &computepb.GetUrlMapRequest{
			UrlMap:  a.id.ResourceID,
			Project: a.id.ParentID.ProjectID,
		}
		return a.globalClient.Get(ctx, getReq)
	} else {
		getReq := &computepb.GetRegionUrlMapRequest{
			UrlMap:  a.id.ResourceID,
			Region:  a.id.ParentID.Location,
			Project: a.id.ParentID.ProjectID,
		}
		return a.urlMapsClient.Get(ctx, getReq)
	}
}
