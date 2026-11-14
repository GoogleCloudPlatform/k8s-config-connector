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
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ComputeInterconnectAttachmentGroup

package compute

import (
	"context"
	"fmt"
	"sort"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ComputeInterconnectAttachmentGroupGVK, NewInterconnectAttachmentGroupModel)
}

func NewInterconnectAttachmentGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelInterconnectAttachmentGroup{config: *config}, nil
}

var _ directbase.Model = &modelInterconnectAttachmentGroup{}

type modelInterconnectAttachmentGroup struct {
	config config.ControllerConfig
}

func (m *modelInterconnectAttachmentGroup) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeInterconnectAttachmentGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ComputeInterconnectAttachmentGroupIdentity)

	// Get compute GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &InterconnectAttachmentGroupAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelInterconnectAttachmentGroup) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//compute.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ComputeInterconnectAttachmentGroupIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &InterconnectAttachmentGroupAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

func (m *modelInterconnectAttachmentGroup) client(ctx context.Context) (*gcp.InterconnectAttachmentGroupsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewInterconnectAttachmentGroupsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building InterconnectAttachmentGroups client: %w", err)
	}
	return gcpClient, err
}

type InterconnectAttachmentGroupAdapter struct {
	id        *krm.ComputeInterconnectAttachmentGroupIdentity
	gcpClient *gcp.InterconnectAttachmentGroupsClient
	desired   *krm.ComputeInterconnectAttachmentGroup
	actual    *computepb.InterconnectAttachmentGroup
}

var _ directbase.Adapter = &InterconnectAttachmentGroupAdapter{}

func (a *InterconnectAttachmentGroupAdapter) get(ctx context.Context) (*computepb.InterconnectAttachmentGroup, error) {
	req := &computepb.GetInterconnectAttachmentGroupRequest{
		Project:                     a.id.Project,
		InterconnectAttachmentGroup: a.id.Interconnectattachmentgroup,
	}
	return a.gcpClient.Get(ctx, req)
}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *InterconnectAttachmentGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeInterconnectAttachmentGroup", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeInterconnectAttachmentGroup %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *InterconnectAttachmentGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeInterconnectAttachmentGroup", "name", a.id)

	mapCtx := &direct.MapContext{}
	resource := ComputeInterconnectAttachmentGroupSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.Interconnectattachmentgroup)

	req := &computepb.InsertInterconnectAttachmentGroupRequest{
		Project:                             a.id.Project,
		InterconnectAttachmentGroupResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeInterconnectAttachmentGroup %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for ComputeInterconnectAttachmentGroup %s creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ComputeInterconnectAttachmentGroup", "name", a.id)

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeInterconnectAttachmentGroup %s after creation: %w", a.id, err)
	}

	return a.setStatus(ctx, createOp, created)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *InterconnectAttachmentGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating InterconnectAttachmentGroup", "name", a.id)

	mergedDesiredProto, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ComputeInterconnectAttachmentGroupSpec_v1alpha1_FromProto,
		ComputeInterconnectAttachmentGroupSpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.setStatus(ctx, updateOp, a.actual)
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)

	mergedDesiredProto.Name = direct.LazyPtr(a.id.Interconnectattachmentgroup)

	sort.Strings(updateMask.Paths)
	req := &computepb.PatchInterconnectAttachmentGroupRequest{
		Project:                             a.id.Project,
		InterconnectAttachmentGroup:         a.id.Interconnectattachmentgroup,
		InterconnectAttachmentGroupResource: mergedDesiredProto,
		UpdateMask:                          direct.PtrTo(strings.Join(updateMask.Paths, ",")),
	}
	op, err := a.gcpClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ComputeInterconnectAttachmentGroup %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for ComputeInterconnectAttachmentGroup %s update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated ComputeInterconnectAttachmentGroup", "name", a.id)

	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeInterconnectAttachmentGroup %s after update: %w", a.id, err)
	}

	return a.setStatus(ctx, updateOp, updated)
}

func (a *InterconnectAttachmentGroupAdapter) setStatus(ctx context.Context, op directbase.Operation, actual *computepb.InterconnectAttachmentGroup) error {
	status := &krm.ComputeInterconnectAttachmentGroupStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ComputeInterconnectAttachmentGroupObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *InterconnectAttachmentGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeInterconnectAttachmentGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeInterconnectAttachmentGroupSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Interconnectattachmentgroup)
	u.SetGroupVersionKind(krm.ComputeInterconnectAttachmentGroupGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *InterconnectAttachmentGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeInterconnectAttachmentGroup", "name", a.id)

	req := &computepb.DeleteInterconnectAttachmentGroupRequest{
		Project:                     a.id.Project,
		InterconnectAttachmentGroup: a.id.Interconnectattachmentgroup,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeInterconnectAttachmentGroup %s: %w", a.id.String(), err)
	}
	if err := op.Wait(ctx); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for ComputeInterconnectAttachmentGroup %s deletion: %w", a.id.String(), err)
	}
	return true, nil
}
