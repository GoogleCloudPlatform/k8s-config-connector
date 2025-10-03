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

package orgpolicy

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/orgpolicy/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/orgpolicy/apiv2"

	orgpolicypb "cloud.google.com/go/orgpolicy/apiv2/orgpolicypb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.OrgPolicyCustomConstraintGVK, NewCustomConstraintModel)
}

func NewCustomConstraintModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCustomConstraint{config: *config}, nil
}

var _ directbase.Model = &modelCustomConstraint{}

type modelCustomConstraint struct {
	config config.ControllerConfig
}

func (m *modelCustomConstraint) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CustomConstraint client: %w", err)
	}
	return gcpClient, err
}

func (m *modelCustomConstraint) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.OrgPolicyCustomConstraint{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewCustomConstraintIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get orgpolicy GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &CustomConstraintAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelCustomConstraint) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type CustomConstraintAdapter struct {
	id        *krm.CustomConstraintIdentity
	gcpClient *gcp.Client
	desired   *krm.OrgPolicyCustomConstraint
	actual    *orgpolicypb.CustomConstraint
}

var _ directbase.Adapter = &CustomConstraintAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *CustomConstraintAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CustomConstraint", "name", a.id)

	req := &orgpolicypb.GetCustomConstraintRequest{Name: a.id.String()}
	customconstraintpb, err := a.gcpClient.GetCustomConstraint(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CustomConstraint %q: %w", a.id, err)
	}

	a.actual = customconstraintpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CustomConstraintAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CustomConstraint", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := OrgPolicyCustomConstraintSpec_ToProto(mapCtx, &desired.Spec)
	resource.Name = a.id.String()
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &orgpolicypb.CreateCustomConstraintRequest{
		Parent:           a.id.Parent().String(),
		CustomConstraint: resource,
	}
	created, err := a.gcpClient.CreateCustomConstraint(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CustomConstraint %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created CustomConstraint", "name", a.id)

	status := &krm.OrgPolicyCustomConstraintStatus{}
	status.ObservedState = OrgPolicyCustomConstraintObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CustomConstraintAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CustomConstraint", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := OrgPolicyCustomConstraintSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	desiredPb.Name = a.id.String()
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &orgpolicypb.UpdateCustomConstraintRequest{
		CustomConstraint: desiredPb,
	}
	updated, err := a.gcpClient.UpdateCustomConstraint(ctx, req)
	if err != nil {
		return fmt.Errorf("updating CustomConstraint %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated CustomConstraint", "name", a.id)

	status := &krm.OrgPolicyCustomConstraintStatus{}
	status.ObservedState = OrgPolicyCustomConstraintObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *CustomConstraintAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.OrgPolicyCustomConstraint{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(OrgPolicyCustomConstraintSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.OrganizationRef = &refs.OrganizationRef{External: a.id.Parent().OrganizationID}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.OrgPolicyCustomConstraintGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *CustomConstraintAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CustomConstraint", "name", a.id)

	req := &orgpolicypb.DeleteCustomConstraintRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteCustomConstraint(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent CustomConstraint, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting CustomConstraint %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted CustomConstraint", "name", a.id)

	return true, nil
}
