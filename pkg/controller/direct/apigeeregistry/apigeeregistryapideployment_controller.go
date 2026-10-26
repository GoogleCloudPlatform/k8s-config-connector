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
// krm.group: apigeeregistry.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ApigeeregistryAPIDeployment

package apigeeregistry

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigeeregistry/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/apigeeregistry/apiv1"
	apigeeregistrypb "cloud.google.com/go/apigeeregistry/apiv1/apigeeregistrypb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ApigeeregistryAPIDeploymentGVK, NewApiDeploymentModel)
}

func NewApiDeploymentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelApiDeployment{config: *config}, nil
}

var _ directbase.Model = &modelApiDeployment{}

type modelApiDeployment struct {
	config config.ControllerConfig
}

func (m *modelApiDeployment) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ApigeeregistryAPIDeployment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ApigeeregistryAPIDeploymentIdentity)

	// Get apigeeregistry GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ApiDeploymentAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelApiDeployment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//apigeeregistry.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ApigeeregistryAPIDeploymentIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ApiDeploymentAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelApiDeployment) client(ctx context.Context) (*gcp.RegistryClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRegistryClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ApiDeployment client: %w", err)
	}
	return gcpClient, err
}

type ApiDeploymentAdapter struct {
	id        *krm.ApigeeregistryAPIDeploymentIdentity
	gcpClient *gcp.RegistryClient
	desired   *krm.ApigeeregistryAPIDeployment
	actual    *apigeeregistrypb.ApiDeployment
}

var _ directbase.Adapter = &ApiDeploymentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ApiDeploymentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ApiDeployment", "name", a.id)

	req := &apigeeregistrypb.GetApiDeploymentRequest{
		Name: a.id.String(),
	}
	apideploymentpb, err := a.gcpClient.GetApiDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ApiDeployment %q: %w", a.id, err)
	}

	a.actual = apideploymentpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ApiDeploymentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ApiDeployment", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ApigeeregistryAPIDeploymentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &apigeeregistrypb.CreateApiDeploymentRequest{
		Parent:          a.id.ParentString(),
		ApiDeploymentId: a.id.Deployment,
		ApiDeployment:   resource,
	}
	created, err := a.gcpClient.CreateApiDeployment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ApiDeployment %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ApiDeployment", "name", a.id)

	status := &krm.ApigeeregistryAPIDeploymentStatus{}
	status.ObservedState = ApigeeregistryAPIDeploymentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ApiDeploymentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ApiDeployment", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ApigeeregistryAPIDeploymentSpec_FromProto,
		ApigeeregistryAPIDeploymentSpec_ToProto,
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
		req := &apigeeregistrypb.UpdateApiDeploymentRequest{
			UpdateMask:    updateMask,
			ApiDeployment: resource,
		}
		updated, err = a.gcpClient.UpdateApiDeployment(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ApiDeployment %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated ApiDeployment", "name", a.id)
	}

	status := &krm.ApigeeregistryAPIDeploymentStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ApigeeregistryAPIDeploymentObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ApiDeploymentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ApigeeregistryAPIDeployment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ApigeeregistryAPIDeploymentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ApigeeregistryAPIDeploymentGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ApiDeploymentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ApiDeployment", "name", a.id)

	req := &apigeeregistrypb.DeleteApiDeploymentRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteApiDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ApiDeployment, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ApiDeployment %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ApiDeployment", "name", a.id)
	return true, nil
}
