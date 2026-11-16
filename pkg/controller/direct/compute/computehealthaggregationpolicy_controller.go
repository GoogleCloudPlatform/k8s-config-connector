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
// krm.kind: ComputeHealthAggregationPolicy

package compute

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ComputeHealthAggregationPolicyGVK, NewHealthAggregationPolicyModel)
}

func NewHealthAggregationPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelHealthAggregationPolicy{config: *config}, nil
}

var _ directbase.Model = &modelHealthAggregationPolicy{}

type modelHealthAggregationPolicy struct {
	config config.ControllerConfig
}

func (m *modelHealthAggregationPolicy) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeHealthAggregationPolicy{}
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
	id := i.(*krm.ComputeHealthAggregationPolicyIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &HealthAggregationPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelHealthAggregationPolicy) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//compute.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ComputeHealthAggregationPolicyIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &HealthAggregationPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

func (m *modelHealthAggregationPolicy) client(ctx context.Context) (*gcp.RegionHealthAggregationPoliciesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRegionHealthAggregationPoliciesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building HealthAggregationPolicy client: %w", err)
	}
	return gcpClient, err
}

type HealthAggregationPolicyAdapter struct {
	id        *krm.ComputeHealthAggregationPolicyIdentity
	gcpClient *gcp.RegionHealthAggregationPoliciesClient
	desired   *krm.ComputeHealthAggregationPolicy
	actual    *computepb.HealthAggregationPolicy
}

var _ directbase.Adapter = &HealthAggregationPolicyAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *HealthAggregationPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting HealthAggregationPolicy", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting HealthAggregationPolicy %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *HealthAggregationPolicyAdapter) get(ctx context.Context) (*computepb.HealthAggregationPolicy, error) {
	req := &computepb.GetRegionHealthAggregationPolicyRequest{
		Project:                 a.id.Project,
		Region:                  a.id.Region,
		HealthAggregationPolicy: a.id.Healthaggregationpolicy,
	}
	return a.gcpClient.Get(ctx, req)
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *HealthAggregationPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating HealthAggregationPolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ComputeHealthAggregationPolicySpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.Healthaggregationpolicy)

	req := &computepb.InsertRegionHealthAggregationPolicyRequest{
		Project:                         a.id.Project,
		Region:                          a.id.Region,
		HealthAggregationPolicyResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating HealthAggregationPolicy %s: %w", a.id, err)
	}
	if !op.Done() {
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for creation of HealthAggregationPolicy %s: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully created HealthAggregationPolicy", "name", a.id)

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting HealthAggregationPolicy %s: %w", a.id, err)
	}

	return a.setStatus(ctx, createOp, created)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *HealthAggregationPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating HealthAggregationPolicy", "name", a.id)

	resource, report, _, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ComputeHealthAggregationPolicySpec_v1alpha1_FromProto,
		ComputeHealthAggregationPolicySpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return err
	}

	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.setStatus(ctx, updateOp, a.actual)
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)

	// An up-to-date fingerprint must be provided in order to patch
	resource.Fingerprint = a.actual.Fingerprint
	resource.Name = direct.LazyPtr(a.id.Healthaggregationpolicy)

	req := &computepb.PatchRegionHealthAggregationPolicyRequest{
		Project:                         a.id.Project,
		Region:                          a.id.Region,
		HealthAggregationPolicy:         a.id.Healthaggregationpolicy,
		HealthAggregationPolicyResource: resource,
	}
	op, err := a.gcpClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("updating HealthAggregationPolicy %s: %w", a.id, err)
	}
	if !op.Done() {
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for update of HealthAggregationPolicy %s: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully updated HealthAggregationPolicy", "name", a.id)

	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting HealthAggregationPolicy %s: %w", a.id, err)
	}

	return a.setStatus(ctx, updateOp, updated)
}

func (a *HealthAggregationPolicyAdapter) setStatus(ctx context.Context, op directbase.Operation, actual *computepb.HealthAggregationPolicy) error {
	status := &krm.ComputeHealthAggregationPolicyStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ComputeHealthAggregationPolicyObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *HealthAggregationPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeHealthAggregationPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeHealthAggregationPolicySpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Region
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Healthaggregationpolicy)
	u.SetGroupVersionKind(krm.ComputeHealthAggregationPolicyGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *HealthAggregationPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting HealthAggregationPolicy", "name", a.id)

	req := &computepb.DeleteRegionHealthAggregationPolicyRequest{
		Project:                 a.id.Project,
		Region:                  a.id.Region,
		HealthAggregationPolicy: a.id.Healthaggregationpolicy,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting HealthAggregationPolicy %s: %w", a.id.String(), err)
	}

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of HealthAggregationPolicy %s: %w", a.id.String(), err)
		}
	}
	log.V(2).Info("successfully deleted HealthAggregationPolicy", "name", a.id)

	return true, nil
}
