// Copyright 2025 Google LLC
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
// krm.group: cloudtalentsolution.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: CloudTalentSolutionTenant

package cloudtalentsolution

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudtalentsolution/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/talent/apiv4"
	talentpb "cloud.google.com/go/talent/apiv4/talentpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CloudTalentSolutionTenantGVK, NewTenantModel)
}

func NewTenantModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTenant{config: *config}, nil
}

var _ directbase.Model = &modelTenant{}

type modelTenant struct {
	config config.ControllerConfig
}

func (m *modelTenant) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudTalentSolutionTenant{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.CloudTalentSolutionTenantIdentity)

	// Get cloudtalentsolution GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &TenantAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelTenant) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//jobs.googleapis.com/") {
		return nil, nil
	}
	id := &krm.CloudTalentSolutionTenantIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &TenantAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelTenant) client(ctx context.Context) (*gcp.TenantClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewTenantRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Tenant client: %w", err)
	}
	return gcpClient, err
}

type TenantAdapter struct {
	id        *krm.CloudTalentSolutionTenantIdentity
	gcpClient *gcp.TenantClient
	desired   *krm.CloudTalentSolutionTenant
	actual    *talentpb.Tenant
}

var _ directbase.Adapter = &TenantAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *TenantAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Tenant", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &talentpb.GetTenantRequest{
		Name: a.id.String(),
	}
	tenantpb, err := a.gcpClient.GetTenant(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Tenant %q: %w", a.id, err)
	}

	a.actual = tenantpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TenantAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Tenant", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudTalentSolutionTenantSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &talentpb.CreateTenantRequest{
		Parent: a.id.ParentString(),
		Tenant: resource,
	}
	created, err := a.gcpClient.CreateTenant(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Tenant %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Tenant", "name", a.id)

	status := &krm.CloudTalentSolutionTenantStatus{}
	parsedID := &krm.CloudTalentSolutionTenantIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.Tenant = parsedID.Tenant
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TenantAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Tenant", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		CloudTalentSolutionTenantSpec_FromProto,
		CloudTalentSolutionTenantSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &talentpb.UpdateTenantRequest{
			UpdateMask: updateMask,
			Tenant:     resource,
		}
		_, err = a.gcpClient.UpdateTenant(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Tenant %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Tenant", "name", a.id)
	}

	status := &krm.CloudTalentSolutionTenantStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *TenantAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudTalentSolutionTenant{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudTalentSolutionTenantSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.CloudTalentSolutionTenantGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *TenantAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Tenant", "name", a.id)

	req := &talentpb.DeleteTenantRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteTenant(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Tenant, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Tenant %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Tenant", "name", a.id)
	return true, nil
}
