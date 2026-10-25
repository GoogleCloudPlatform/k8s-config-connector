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
// krm.group: securityposture.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: SecurityposturePostureDeployment

package securityposture

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securityposture/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/securityposture/apiv1"
	securityposturepb "cloud.google.com/go/securityposture/apiv1/securityposturepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.SecurityposturePostureDeploymentGVK, NewPostureDeploymentModel)
}

func NewPostureDeploymentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelPostureDeployment{config: *config}, nil
}

var _ directbase.Model = &modelPostureDeployment{}

type modelPostureDeployment struct {
	config config.ControllerConfig
}

func (m *modelPostureDeployment) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SecurityposturePostureDeployment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.SecurityposturePostureDeploymentIdentity)

	// Get securityposture GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &PostureDeploymentAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelPostureDeployment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//securityposture.googleapis.com/") {
		return nil, nil
	}
	id := &krm.SecurityposturePostureDeploymentIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &PostureDeploymentAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelPostureDeployment) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building PostureDeployment client: %w", err)
	}
	return gcpClient, err
}

type PostureDeploymentAdapter struct {
	id        *krm.SecurityposturePostureDeploymentIdentity
	gcpClient *gcp.Client
	desired   *krm.SecurityposturePostureDeployment
	actual    *securityposturepb.PostureDeployment
}

var _ directbase.Adapter = &PostureDeploymentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *PostureDeploymentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting PostureDeployment", "name", a.id)

	req := &securityposturepb.GetPostureDeploymentRequest{
		Name: a.id.String(),
	}
	posturedeploymentpb, err := a.gcpClient.GetPostureDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting PostureDeployment %q: %w", a.id, err)
	}

	a.actual = posturedeploymentpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PostureDeploymentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating PostureDeployment", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SecurityposturePostureDeploymentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &securityposturepb.CreatePostureDeploymentRequest{
		Parent:              a.id.ParentString(),
		PostureDeploymentId: a.id.PostureDeployment,
		PostureDeployment:   resource,
	}
	op, err := a.gcpClient.CreatePostureDeployment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating PostureDeployment %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("PostureDeployment %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created PostureDeployment", "name", a.id)

	status := &krm.SecurityposturePostureDeploymentStatus{}
	status.ObservedState = SecurityposturePostureDeploymentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PostureDeploymentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating PostureDeployment", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		SecurityposturePostureDeploymentSpec_FromProto,
		SecurityposturePostureDeploymentSpec_ToProto,
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
		req := &securityposturepb.UpdatePostureDeploymentRequest{
			UpdateMask:        updateMask,
			PostureDeployment: resource,
		}
		op, err := a.gcpClient.UpdatePostureDeployment(ctx, req)
		if err != nil {
			return fmt.Errorf("updating PostureDeployment %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("PostureDeployment %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated PostureDeployment", "name", a.id)
	}

	status := &krm.SecurityposturePostureDeploymentStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = SecurityposturePostureDeploymentObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *PostureDeploymentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SecurityposturePostureDeployment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SecurityposturePostureDeploymentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.SecurityposturePostureDeploymentGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *PostureDeploymentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting PostureDeployment", "name", a.id)

	req := &securityposturepb.DeletePostureDeploymentRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeletePostureDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent PostureDeployment, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting PostureDeployment %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted PostureDeployment", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete PostureDeployment %s: %w", a.id, err)
	}
	return true, nil
}
