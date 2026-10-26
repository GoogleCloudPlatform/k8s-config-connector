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
// krm.kind: ApigeeregistryAPISpec

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
	registry.RegisterModel(krm.ApigeeregistryAPISpecGVK, NewApiSpecModel)
}

func NewApiSpecModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelApiSpec{config: *config}, nil
}

var _ directbase.Model = &modelApiSpec{}

type modelApiSpec struct {
	config config.ControllerConfig
}

func (m *modelApiSpec) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ApigeeregistryAPISpec{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ApigeeregistryAPISpecIdentity)

	// Get apigeeregistry GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ApiSpecAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelApiSpec) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//apigeeregistry.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ApigeeregistryAPISpecIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ApiSpecAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelApiSpec) client(ctx context.Context) (*gcp.RegistryClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRegistryClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ApiSpec client: %w", err)
	}
	return gcpClient, err
}

type ApiSpecAdapter struct {
	id        *krm.ApigeeregistryAPISpecIdentity
	gcpClient *gcp.RegistryClient
	desired   *krm.ApigeeregistryAPISpec
	actual    *apigeeregistrypb.ApiSpec
}

var _ directbase.Adapter = &ApiSpecAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ApiSpecAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ApiSpec", "name", a.id)

	req := &apigeeregistrypb.GetApiSpecRequest{
		Name: a.id.String(),
	}
	apispecpb, err := a.gcpClient.GetApiSpec(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ApiSpec %q: %w", a.id, err)
	}

	a.actual = apispecpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ApiSpecAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ApiSpec", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ApigeeregistryAPISpecSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &apigeeregistrypb.CreateApiSpecRequest{
		Parent:    a.id.ParentString(),
		ApiSpecId: a.id.Spec,
		ApiSpec:   resource,
	}
	created, err := a.gcpClient.CreateApiSpec(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ApiSpec %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ApiSpec", "name", a.id)

	status := &krm.ApigeeregistryAPISpecStatus{}
	status.ObservedState = ApigeeregistryAPISpecObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ApiSpecAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ApiSpec", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ApigeeregistryAPISpecSpec_FromProto,
		ApigeeregistryAPISpecSpec_ToProto,
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
		req := &apigeeregistrypb.UpdateApiSpecRequest{
			UpdateMask: updateMask,
			ApiSpec:    resource,
		}
		updated, err = a.gcpClient.UpdateApiSpec(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ApiSpec %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated ApiSpec", "name", a.id)
	}

	status := &krm.ApigeeregistryAPISpecStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ApigeeregistryAPISpecObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ApiSpecAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ApigeeregistryAPISpec{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ApigeeregistryAPISpecSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ApigeeregistryAPISpecGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ApiSpecAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ApiSpec", "name", a.id)

	req := &apigeeregistrypb.DeleteApiSpecRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteApiSpec(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ApiSpec, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ApiSpec %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ApiSpec", "name", a.id)
	return true, nil
}
