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
// krm.group: cloudsecuritycompliance.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: CloudSecurityComplianceFrameworkDeployment

package cloudsecuritycompliance

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudsecuritycompliance/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/cloudsecuritycompliance/apiv1"
	cloudsecuritycompliancepb "cloud.google.com/go/cloudsecuritycompliance/apiv1/cloudsecuritycompliancepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CloudSecurityComplianceFrameworkDeploymentGVK, NewFrameworkDeploymentModel)
}

func NewFrameworkDeploymentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelFrameworkDeployment{config: *config}, nil
}

var _ directbase.Model = &modelFrameworkDeployment{}

type modelFrameworkDeployment struct {
	config config.ControllerConfig
}

func (m *modelFrameworkDeployment) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudSecurityComplianceFrameworkDeployment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.CloudSecurityComplianceFrameworkDeploymentIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get cloudsecuritycompliance GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &FrameworkDeploymentAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelFrameworkDeployment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//cloudsecuritycompliance.googleapis.com/") {
		return nil, nil
	}
	id := &krm.CloudSecurityComplianceFrameworkDeploymentIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &FrameworkDeploymentAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelFrameworkDeployment) client(ctx context.Context) (*gcp.DeploymentClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDeploymentRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building FrameworkDeployment client: %w", err)
	}
	return gcpClient, err
}

type FrameworkDeploymentAdapter struct {
	id        *krm.CloudSecurityComplianceFrameworkDeploymentIdentity
	gcpClient *gcp.DeploymentClient
	desired   *krm.CloudSecurityComplianceFrameworkDeployment
	actual    *cloudsecuritycompliancepb.FrameworkDeployment
}

var _ directbase.Adapter = &FrameworkDeploymentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *FrameworkDeploymentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FrameworkDeployment", "name", a.id)

	req := &cloudsecuritycompliancepb.GetFrameworkDeploymentRequest{
		Name: a.id.String(),
	}
	frameworkdeploymentpb, err := a.gcpClient.GetFrameworkDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FrameworkDeployment %q: %w", a.id, err)
	}

	a.actual = frameworkdeploymentpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FrameworkDeploymentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FrameworkDeployment", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudSecurityComplianceFrameworkDeploymentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &cloudsecuritycompliancepb.CreateFrameworkDeploymentRequest{
		Parent:                a.id.ParentString(),
		FrameworkDeploymentId: a.id.FrameworkDeployment,
		FrameworkDeployment:   resource,
	}
	op, err := a.gcpClient.CreateFrameworkDeployment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FrameworkDeployment %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("FrameworkDeployment %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created FrameworkDeployment", "name", a.id)

	status := &krm.CloudSecurityComplianceFrameworkDeploymentStatus{}
	status.ObservedState = CloudSecurityComplianceFrameworkDeploymentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FrameworkDeploymentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FrameworkDeployment", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		CloudSecurityComplianceFrameworkDeploymentSpec_FromProto,
		CloudSecurityComplianceFrameworkDeploymentSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating CloudSecurityComplianceFrameworkDeployment is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.CloudSecurityComplianceFrameworkDeploymentStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = CloudSecurityComplianceFrameworkDeploymentObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *FrameworkDeploymentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudSecurityComplianceFrameworkDeployment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudSecurityComplianceFrameworkDeploymentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.CloudSecurityComplianceFrameworkDeploymentGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *FrameworkDeploymentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FrameworkDeployment", "name", a.id)

	req := &cloudsecuritycompliancepb.DeleteFrameworkDeploymentRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteFrameworkDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent FrameworkDeployment, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting FrameworkDeployment %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted FrameworkDeployment", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete FrameworkDeployment %s: %w", a.id, err)
	}
	return true, nil
}
