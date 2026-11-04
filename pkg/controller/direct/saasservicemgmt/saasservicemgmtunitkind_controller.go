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
// krm.group: saasservicemgmt.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: SaaSServiceMgmtUnitKind

package saasservicemgmt

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/saasservicemgmt/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1"
	saasservicemgmtpb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.SaaSServiceMgmtUnitKindGVK, NewUnitKindModel)
}

func NewUnitKindModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelUnitKind{config: *config}, nil
}

var _ directbase.Model = &modelUnitKind{}

type modelUnitKind struct {
	config config.ControllerConfig
}

func (m *modelUnitKind) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SaaSServiceMgmtUnitKind{}
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
	id := i.(*krm.SaaSServiceMgmtUnitKindIdentity)

	// Get saasservicemgmt GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &UnitKindAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelUnitKind) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//saasservicemgmt.googleapis.com/") {
		return nil, nil
	}
	id := &krm.SaaSServiceMgmtUnitKindIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &UnitKindAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelUnitKind) client(ctx context.Context) (*gcp.SaasDeploymentsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewSaasDeploymentsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building UnitKind client: %w", err)
	}
	return gcpClient, err
}

type UnitKindAdapter struct {
	id        *krm.SaaSServiceMgmtUnitKindIdentity
	gcpClient *gcp.SaasDeploymentsClient
	desired   *krm.SaaSServiceMgmtUnitKind
	actual    *saasservicemgmtpb.UnitKind
}

var _ directbase.Adapter = &UnitKindAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *UnitKindAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting UnitKind", "name", a.id)

	req := &saasservicemgmtpb.GetUnitKindRequest{
		Name: a.id.String(),
	}
	unitkindpb, err := a.gcpClient.GetUnitKind(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting UnitKind %q: %w", a.id, err)
	}

	a.actual = unitkindpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *UnitKindAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating UnitKind", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SaaSServiceMgmtUnitKindSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &saasservicemgmtpb.CreateUnitKindRequest{
		Parent:     a.id.ParentString(),
		UnitKindId: a.id.UnitKind,
		UnitKind:   resource,
	}
	created, err := a.gcpClient.CreateUnitKind(ctx, req)
	if err != nil {
		return fmt.Errorf("creating UnitKind %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created UnitKind", "name", a.id)

	status := &krm.SaaSServiceMgmtUnitKindStatus{}
	status.ObservedState = SaaSServiceMgmtUnitKindObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *UnitKindAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating UnitKind", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		SaaSServiceMgmtUnitKindSpec_FromProto,
		SaaSServiceMgmtUnitKindSpec_ToProto,
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
		req := &saasservicemgmtpb.UpdateUnitKindRequest{
			UpdateMask: updateMask,
			UnitKind:   resource,
		}
		updated, err = a.gcpClient.UpdateUnitKind(ctx, req)
		if err != nil {
			return fmt.Errorf("updating UnitKind %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated UnitKind", "name", a.id)
	}

	status := &krm.SaaSServiceMgmtUnitKindStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = SaaSServiceMgmtUnitKindObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *UnitKindAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SaaSServiceMgmtUnitKind{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SaaSServiceMgmtUnitKindSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.SaaSServiceMgmtUnitKindGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *UnitKindAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting UnitKind", "name", a.id)

	req := &saasservicemgmtpb.DeleteUnitKindRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteUnitKind(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent UnitKind, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting UnitKind %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted UnitKind", "name", a.id)
	return true, nil
}
