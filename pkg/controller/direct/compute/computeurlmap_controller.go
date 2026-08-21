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

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
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

type ComputeURLMapAdapter struct {
	urlMapsClient       *compute.UrlMapsClient
	regionUrlMapsClient *compute.RegionUrlMapsClient
	id                  *krm.ComputeURLMapIdentity
	desired             *pb.UrlMap
	actual              *pb.UrlMap
	reader              client.Reader
}

var _ directbase.Adapter = &ComputeURLMapAdapter{}

func (m *urlMapModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeURLMap{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always normalize references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	var urlMapsClient *compute.UrlMapsClient
	var regionUrlMapsClient *compute.RegionUrlMapsClient

	urlMapId := id.(*krm.ComputeURLMapIdentity)
	if urlMapId.IsGlobal() {
		urlMapsClient, err = gcpClient.newUrlMapsClient(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		regionUrlMapsClient, err = gcpClient.newRegionUrlMapsClient(ctx)
		if err != nil {
			return nil, err
		}
	}

	mapCtx := &direct.MapContext{}
	desired := ComputeURLMapSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set name in proto
	desired.Name = direct.LazyPtr(urlMapId.Urlmap)

	return &ComputeURLMapAdapter{
		urlMapsClient:       urlMapsClient,
		regionUrlMapsClient: regionUrlMapsClient,
		id:                  urlMapId,
		desired:             desired,
		reader:              reader,
	}, nil
}

func (m *urlMapModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *ComputeURLMapAdapter) Find(ctx context.Context) (bool, error) {
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

func (a *ComputeURLMapAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeURLMap", "name", a.id)

	var op *compute.Operation
	var err error

	if a.id.IsGlobal() {
		req := &pb.InsertUrlMapRequest{
			Project:        a.id.Project,
			UrlMapResource: a.desired,
		}
		op, err = a.urlMapsClient.Insert(ctx, req)
	} else {
		req := &pb.InsertRegionUrlMapRequest{
			Project:        a.id.Project,
			Region:         a.id.Region,
			UrlMapResource: a.desired,
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
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeURLMap %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *ComputeURLMapAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeURLMap", "name", a.id)

	// Handle output-only fields from GCP
	a.assignGCPDefaults(a.desired, a.actual)

	paths, report, err := common.CompareProtoMessageStructuredDiff(a.desired, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)

	updateOp.RecordUpdatingEvent()

	var op *compute.Operation

	if a.id.IsGlobal() {
		req := &pb.UpdateUrlMapRequest{
			Project:        a.id.Project,
			UrlMap:         a.id.Urlmap,
			UrlMapResource: a.desired,
		}
		op, err = a.urlMapsClient.Update(ctx, req)
	} else {
		req := &pb.UpdateRegionUrlMapRequest{
			Project:        a.id.Project,
			Region:         a.id.Region,
			UrlMap:         a.id.Urlmap,
			UrlMapResource: a.desired,
		}
		op, err = a.regionUrlMapsClient.Update(ctx, req)
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

	// Get the updated resource
	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeURLMap %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *ComputeURLMapAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeURLMap", "name", a.id)

	var op *compute.Operation
	var err error

	if a.id.IsGlobal() {
		req := &pb.DeleteUrlMapRequest{
			Project: a.id.Project,
			UrlMap:  a.id.Urlmap,
		}
		op, err = a.urlMapsClient.Delete(ctx, req)
	} else {
		req := &pb.DeleteRegionUrlMapRequest{
			Project: a.id.Project,
			Region:  a.id.Region,
			UrlMap:  a.id.Urlmap,
		}
		op, err = a.regionUrlMapsClient.Delete(ctx, req)
	}

	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeURLMap %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting ComputeURLMap %s delete failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeURLMap", "name", a.id)
	return true, nil
}

func (a *ComputeURLMapAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeURLMap{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeURLMapSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetGroupVersionKind(krm.ComputeURLMapGVK)
	return u, nil
}

func (a *ComputeURLMapAdapter) get(ctx context.Context) (*pb.UrlMap, error) {
	if a.id.IsGlobal() {
		req := &pb.GetUrlMapRequest{
			Project: a.id.Project,
			UrlMap:  a.id.Urlmap,
		}
		return a.urlMapsClient.Get(ctx, req)
	} else {
		req := &pb.GetRegionUrlMapRequest{
			Project: a.id.Project,
			Region:  a.id.Region,
			UrlMap:  a.id.Urlmap,
		}
		return a.regionUrlMapsClient.Get(ctx, req)
	}
}

func (a *ComputeURLMapAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.UrlMap) error {
	mapCtx := &direct.MapContext{}
	status := ComputeURLMapStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ComputeURLMapAdapter) assignGCPDefaults(desired, actual *pb.UrlMap) {
	if actual != nil {
		desired.Fingerprint = actual.Fingerprint
		desired.Id = actual.Id
		desired.SelfLink = actual.SelfLink
		desired.CreationTimestamp = actual.CreationTimestamp
		desired.Kind = actual.Kind
	}
}
