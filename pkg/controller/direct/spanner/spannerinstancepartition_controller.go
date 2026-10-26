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

// +generated:controller
// krm.group: spanner.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: SpannerInstancePartition

package spanner

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/spanner/admin/instance/apiv1"
	instancepb "cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.SpannerInstancePartitionGVK, NewInstancePartitionModel)
}

func NewInstancePartitionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelInstancePartition{config: *config}, nil
}

var _ directbase.Model = &modelInstancePartition{}

type modelInstancePartition struct {
	config config.ControllerConfig
}

func (m *modelInstancePartition) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SpannerInstancePartition{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.SpannerInstancePartitionIdentity)

	// Get spanner GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &InstancePartitionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelInstancePartition) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//spanner.googleapis.com/") {
		return nil, nil
	}
	id := &krm.SpannerInstancePartitionIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &InstancePartitionAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelInstancePartition) client(ctx context.Context) (*gcp.InstanceAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewInstanceAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building InstancePartition client: %w", err)
	}
	return gcpClient, err
}

type InstancePartitionAdapter struct {
	id        *krm.SpannerInstancePartitionIdentity
	gcpClient *gcp.InstanceAdminClient
	desired   *krm.SpannerInstancePartition
	actual    *instancepb.InstancePartition
}

var _ directbase.Adapter = &InstancePartitionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *InstancePartitionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting InstancePartition", "name", a.id)

	req := &instancepb.GetInstancePartitionRequest{
		Name: a.id.String(),
	}
	instancepartitionpb, err := a.gcpClient.GetInstancePartition(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting InstancePartition %q: %w", a.id, err)
	}

	a.actual = instancepartitionpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *InstancePartitionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating InstancePartition", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SpannerInstancePartitionSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &instancepb.CreateInstancePartitionRequest{
		Parent:              a.id.ParentString(),
		InstancePartitionId: a.id.InstancePartition,
		InstancePartition:   resource,
	}
	op, err := a.gcpClient.CreateInstancePartition(ctx, req)
	if err != nil {
		return fmt.Errorf("creating InstancePartition %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("InstancePartition %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created InstancePartition", "name", a.id)

	status := &krm.SpannerInstancePartitionStatus{}
	status.ObservedState = SpannerInstancePartitionObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *InstancePartitionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating InstancePartition", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		SpannerInstancePartitionSpec_v1alpha1_FromProto,
		SpannerInstancePartitionSpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &instancepb.UpdateInstancePartitionRequest{
			FieldMask:         updateMask,
			InstancePartition: resource,
		}
		op, err := a.gcpClient.UpdateInstancePartition(ctx, req)
		if err != nil {
			return fmt.Errorf("updating InstancePartition %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("InstancePartition %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated InstancePartition", "name", a.id)
	}

	status := &krm.SpannerInstancePartitionStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = SpannerInstancePartitionObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *InstancePartitionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SpannerInstancePartition{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SpannerInstancePartitionSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.SpannerInstancePartitionGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *InstancePartitionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting InstancePartition", "name", a.id)

	req := &instancepb.DeleteInstancePartitionRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteInstancePartition(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent InstancePartition, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting InstancePartition %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted InstancePartition", "name", a.id)
	return true, nil
}
