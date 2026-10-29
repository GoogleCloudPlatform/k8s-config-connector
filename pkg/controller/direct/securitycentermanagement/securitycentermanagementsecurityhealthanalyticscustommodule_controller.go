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
// krm.group: securitycentermanagement.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: SecurityCenterManagementSecurityHealthAnalyticsCustomModule

package securitycentermanagement

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycentermanagement/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/securitycentermanagement/apiv1"
	securitycentermanagementpb "cloud.google.com/go/securitycentermanagement/apiv1/securitycentermanagementpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.SecurityCenterManagementSecurityHealthAnalyticsCustomModuleGVK, NewSecurityHealthAnalyticsCustomModuleModel)
}

func NewSecurityHealthAnalyticsCustomModuleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelSecurityHealthAnalyticsCustomModule{config: *config}, nil
}

var _ directbase.Model = &modelSecurityHealthAnalyticsCustomModule{}

type modelSecurityHealthAnalyticsCustomModule struct {
	config config.ControllerConfig
}

func (m *modelSecurityHealthAnalyticsCustomModule) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SecurityCenterManagementSecurityHealthAnalyticsCustomModule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity)

	// Get securitycentermanagement GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SecurityHealthAnalyticsCustomModuleAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelSecurityHealthAnalyticsCustomModule) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//securitycentermanagement.googleapis.com/") {
		return nil, nil
	}
	id := &krm.SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SecurityHealthAnalyticsCustomModuleAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelSecurityHealthAnalyticsCustomModule) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SecurityHealthAnalyticsCustomModule client: %w", err)
	}
	return gcpClient, err
}

type SecurityHealthAnalyticsCustomModuleAdapter struct {
	id        *krm.SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity
	gcpClient *gcp.Client
	desired   *krm.SecurityCenterManagementSecurityHealthAnalyticsCustomModule
	actual    *securitycentermanagementpb.SecurityHealthAnalyticsCustomModule
}

var _ directbase.Adapter = &SecurityHealthAnalyticsCustomModuleAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *SecurityHealthAnalyticsCustomModuleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting SecurityHealthAnalyticsCustomModule", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &securitycentermanagementpb.GetSecurityHealthAnalyticsCustomModuleRequest{
		Name: a.id.String(),
	}
	securityhealthanalyticscustommodulepb, err := a.gcpClient.GetSecurityHealthAnalyticsCustomModule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SecurityHealthAnalyticsCustomModule %q: %w", a.id, err)
	}

	a.actual = securityhealthanalyticscustommodulepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *SecurityHealthAnalyticsCustomModuleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating SecurityHealthAnalyticsCustomModule", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SecurityCenterManagementSecurityHealthAnalyticsCustomModuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &securitycentermanagementpb.CreateSecurityHealthAnalyticsCustomModuleRequest{
		Parent:                              a.id.ParentString(),
		SecurityHealthAnalyticsCustomModule: resource,
	}
	created, err := a.gcpClient.CreateSecurityHealthAnalyticsCustomModule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating SecurityHealthAnalyticsCustomModule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created SecurityHealthAnalyticsCustomModule", "name", a.id)

	status := &krm.SecurityCenterManagementSecurityHealthAnalyticsCustomModuleStatus{}
	status.ObservedState = SecurityCenterManagementSecurityHealthAnalyticsCustomModuleObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.SecurityHealthAnalyticsCustomModule = parsedID.SecurityHealthAnalyticsCustomModule
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *SecurityHealthAnalyticsCustomModuleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating SecurityHealthAnalyticsCustomModule", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		SecurityCenterManagementSecurityHealthAnalyticsCustomModuleSpec_FromProto,
		SecurityCenterManagementSecurityHealthAnalyticsCustomModuleSpec_ToProto,
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
		req := &securitycentermanagementpb.UpdateSecurityHealthAnalyticsCustomModuleRequest{
			UpdateMask:                          updateMask,
			SecurityHealthAnalyticsCustomModule: resource,
		}
		updated, err = a.gcpClient.UpdateSecurityHealthAnalyticsCustomModule(ctx, req)
		if err != nil {
			return fmt.Errorf("updating SecurityHealthAnalyticsCustomModule %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated SecurityHealthAnalyticsCustomModule", "name", a.id)
	}

	status := &krm.SecurityCenterManagementSecurityHealthAnalyticsCustomModuleStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = SecurityCenterManagementSecurityHealthAnalyticsCustomModuleObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *SecurityHealthAnalyticsCustomModuleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SecurityCenterManagementSecurityHealthAnalyticsCustomModule{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SecurityCenterManagementSecurityHealthAnalyticsCustomModuleSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.SecurityCenterManagementSecurityHealthAnalyticsCustomModuleGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *SecurityHealthAnalyticsCustomModuleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting SecurityHealthAnalyticsCustomModule", "name", a.id)

	req := &securitycentermanagementpb.DeleteSecurityHealthAnalyticsCustomModuleRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteSecurityHealthAnalyticsCustomModule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent SecurityHealthAnalyticsCustomModule, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting SecurityHealthAnalyticsCustomModule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted SecurityHealthAnalyticsCustomModule", "name", a.id)
	return true, nil
}
