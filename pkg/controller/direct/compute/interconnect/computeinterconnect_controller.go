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
// See the License for the applicable language governing permissions and
// limitations under the License.

package interconnect

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	gcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/compute"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ComputeInterconnectGVK, NewInterconnectModel)
}

func NewInterconnectModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &interconnectModel{config: config}, nil
}

var _ directbase.Model = &interconnectModel{}

type interconnectModel struct {
	config *config.ControllerConfig
}

func (m *interconnectModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeInterconnect{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Resolve identity using the new IdentityV2 interface
	identityObj, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := identityObj.(*krm.ComputeInterconnectIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", identityObj)
	}

	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, fmt.Errorf("building rest client options: %w", err)
	}
	interconnectsClient, err := compute.NewInterconnectsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute Interconnects REST client: %w", err)
	}

	return &InterconnectAdapter{
		gcpClient: interconnectsClient,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *interconnectModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type InterconnectAdapter struct {
	gcpClient *compute.InterconnectsClient
	id        *krm.ComputeInterconnectIdentity
	desired   *krm.ComputeInterconnect
	actual    *computepb.Interconnect
}

var _ directbase.Adapter = &InterconnectAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *InterconnectAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeInterconnect", "name", a.id)

	req := &computepb.GetInterconnectRequest{
		Project:      a.id.Project,
		Interconnect: a.id.Interconnect,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeInterconnect %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *InterconnectAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeInterconnect", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := gcp.ComputeInterconnectSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.Interconnect)

	req := &computepb.InsertInterconnectRequest{
		Project:              a.id.Project,
		InterconnectResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeInterconnect %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeInterconnect %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute ComputeInterconnect in gcp", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeInterconnect %s: %w", a.id, err)
	}

	status := &krm.ComputeInterconnectStatus{}
	status.ObservedState = gcp.ComputeInterconnectObservedState_v1alpha1_FromProto(mapCtx, created)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *InterconnectAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeInterconnect", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := gcp.ComputeInterconnectSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.Interconnect)
	// An up-to-date fingerprint must be provided in order to patch
	resource.LabelFingerprint = a.actual.LabelFingerprint

	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	var updated *computepb.Interconnect
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
		for path := range paths {
			report.AddField(path, nil, nil)
		}
		structuredreporting.ReportDiff(ctx, report)

		req := &computepb.PatchInterconnectRequest{
			Project:              a.id.Project,
			Interconnect:         a.id.Interconnect,
			InterconnectResource: resource,
		}
		op, err := a.gcpClient.Patch(ctx, req)
		if err != nil {
			return fmt.Errorf("updating compute ComputeInterconnect %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated compute ComputeInterconnect", "name", a.id.String())

		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("compute ComputeInterconnect %s waiting for update: %w", a.id.String(), err)
		}

		// Get the updated resource
		updated, err = a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeInterconnect %s: %w", a.id, err)
		}
	}

	status := &krm.ComputeInterconnectStatus{}
	status.ObservedState = gcp.ComputeInterconnectObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *InterconnectAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeInterconnect{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(gcp.ComputeInterconnectSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.Parent.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ComputeInterconnectGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *InterconnectAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeInterconnect", "name", a.id)

	req := &computepb.DeleteInterconnectRequest{
		Project:      a.id.Project,
		Interconnect: a.id.Interconnect,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting compute ComputeInterconnect %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted compute ComputeInterconnect", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of compute ComputeInterconnect %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *InterconnectAdapter) get(ctx context.Context) (*computepb.Interconnect, error) {
	getReq := &computepb.GetInterconnectRequest{
		Project:      a.id.Project,
		Interconnect: a.id.Interconnect,
	}
	resource, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return nil, fmt.Errorf("getting ComputeInterconnect %s: %w", a.id, err)
	}
	return resource, nil
}
