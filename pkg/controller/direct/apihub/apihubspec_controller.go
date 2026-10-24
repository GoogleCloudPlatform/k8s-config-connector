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
// krm.group: apihub.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ApihubSpec

package apihub

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/apihub/apiv1"
	apihubpb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ApihubSpecGVK, NewSpecModel)
}

func NewSpecModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelSpec{config: *config}, nil
}

var _ directbase.Model = &modelSpec{}

type modelSpec struct {
	config config.ControllerConfig
}

func (m *modelSpec) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ApihubSpec{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ApihubSpecIdentity)

	// Get apihub GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SpecAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelSpec) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//apihub.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ApihubSpecIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SpecAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelSpec) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Spec client: %w", err)
	}
	return gcpClient, err
}

type SpecAdapter struct {
	id        *krm.ApihubSpecIdentity
	gcpClient *gcp.Client
	desired   *krm.ApihubSpec
	actual    *apihubpb.Spec
}

var _ directbase.Adapter = &SpecAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *SpecAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Spec", "name", a.id)

	req := &apihubpb.GetSpecRequest{
		Name: a.id.String(),
	}
	specpb, err := a.gcpClient.GetSpec(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Spec %q: %w", a.id, err)
	}

	a.actual = specpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *SpecAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Spec", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ApihubSpecSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &apihubpb.CreateSpecRequest{
		Parent: a.id.ParentString(),
		SpecId: a.id.Spec,
		Spec:   resource,
	}
	created, err := a.gcpClient.CreateSpec(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Spec %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Spec", "name", a.id)

	status := &krm.ApihubSpecStatus{}
	status.ObservedState = ApihubSpecObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *SpecAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Spec", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ApihubSpecSpec_FromProto,
		ApihubSpecSpec_ToProto,
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
		req := &apihubpb.UpdateSpecRequest{
			UpdateMask: updateMask,
			Spec:       resource,
		}
		updated, err = a.gcpClient.UpdateSpec(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Spec %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Spec", "name", a.id)
	}

	status := &krm.ApihubSpecStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ApihubSpecObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *SpecAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ApihubSpec{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ApihubSpecSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ApihubSpecGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *SpecAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Spec", "name", a.id)

	req := &apihubpb.DeleteSpecRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteSpec(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Spec, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Spec %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Spec", "name", a.id)
	return true, nil
}
