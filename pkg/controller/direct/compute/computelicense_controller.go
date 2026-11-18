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
// krm.kind: ComputeLicense

package compute

import (
	"context"
	"fmt"
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
	registry.RegisterModel(krm.ComputeLicenseGVK, NewLicenseModel)
}

func NewLicenseModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLicense{config: *config}, nil
}

var _ directbase.Model = &modelLicense{}

type modelLicense struct {
	config config.ControllerConfig
}

func (m *modelLicense) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeLicense{}
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
	id := i.(*krm.ComputeLicenseIdentity)

	// Get compute GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &LicenseAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelLicense) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//compute.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ComputeLicenseIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &LicenseAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

func (m *modelLicense) client(ctx context.Context) (*gcp.LicensesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewLicensesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building License client: %w", err)
	}
	return gcpClient, err
}

type LicenseAdapter struct {
	id        *krm.ComputeLicenseIdentity
	gcpClient *gcp.LicensesClient
	desired   *krm.ComputeLicense
	actual    *computepb.License
}

var _ directbase.Adapter = &LicenseAdapter{}

func (a *LicenseAdapter) get(ctx context.Context) (*computepb.License, error) {
	req := &computepb.GetLicenseRequest{
		Project: a.id.Project,
		License: a.id.License,
	}
	return a.gcpClient.Get(ctx, req)
}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *LicenseAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeLicense", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeLicense %q: %w", a.id, err)
	}
	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *LicenseAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeLicense", "name", a.id)

	mapCtx := &direct.MapContext{}
	resource := ComputeLicenseSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.License)

	req := &computepb.InsertLicenseRequest{
		Project:         a.id.Project,
		LicenseResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeLicense %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for ComputeLicense %s creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ComputeLicense", "name", a.id)

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeLicense %s after creation: %w", a.id, err)
	}

	return a.setStatus(ctx, createOp, created)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *LicenseAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeLicense", "name", a.id)

	normalize := func(ctx context.Context, pb *computepb.License) error {
		pb.ResourceRequirements = nil
		return nil
	}
	mergedDesiredProto, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ComputeLicenseSpec_v1alpha1_FromProto,
		ComputeLicenseSpec_v1alpha1_ToProto,
		normalize,
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

	resource := mergedDesiredProto
	resource.Name = direct.LazyPtr(a.id.License)

	req := &computepb.UpdateLicenseRequest{
		Project:         a.id.Project,
		License:         a.id.License,
		LicenseResource: resource,
	}
	if len(updateMask.Paths) > 0 {
		req.UpdateMask = direct.LazyPtr(strings.Join(updateMask.Paths, ","))
	}
	op, err := a.gcpClient.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ComputeLicense %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for ComputeLicense %s update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated ComputeLicense", "name", a.id)

	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeLicense %s after update: %w", a.id, err)
	}

	return a.setStatus(ctx, updateOp, updated)
}

func (a *LicenseAdapter) setStatus(ctx context.Context, op directbase.Operation, actual *computepb.License) error {
	status := &krm.ComputeLicenseStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ComputeLicenseObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *LicenseAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeLicense{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeLicenseSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.License)
	u.SetGroupVersionKind(krm.ComputeLicenseGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *LicenseAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeLicense", "name", a.id)

	req := &computepb.DeleteLicenseRequest{
		Project: a.id.Project,
		License: a.id.License,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeLicense %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for ComputeLicense %s deletion: %w", a.id, err)
	}
	return true, nil
}
