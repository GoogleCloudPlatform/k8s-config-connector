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
	registry.RegisterModel(krm.OrgPolicyPolicyGVK, NewPolicyModel)
}

func NewPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelPolicy{config: *config}, nil
}

var _ directbase.Model = &modelPolicy{}

type modelPolicy struct {
	config config.ControllerConfig
}

func (m *modelPolicy) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Policy client: %w", err)
	}
	return gcpClient, err
}

func (m *modelPolicy) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.OrgPolicyPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewPolicyIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get orgpolicy GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &PolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelPolicy) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type PolicyAdapter struct {
	id        *krm.PolicyIdentity
	gcpClient *gcp.Client
	desired   *krm.OrgPolicyPolicy
	actual    *orgpolicypb.Policy
}

var _ directbase.Adapter = &PolicyAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *PolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Policy", "name", a.id)

	req := &orgpolicypb.GetPolicyRequest{Name: a.id.String()}
	policypb, err := a.gcpClient.GetPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Policy %q: %w", a.id, err)
	}

	a.actual = policypb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Policy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := OrgPolicyPolicySpec_ToProto(mapCtx, &desired.Spec)
	resource.Name = a.id.String()
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &orgpolicypb.CreatePolicyRequest{
		Parent: a.id.Parent().String(),
		Policy: resource,
	}
	created, err := a.gcpClient.CreatePolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Policy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Policy", "name", a.id)

	status := &krm.OrgPolicyPolicyStatus{}
	status.ObservedState = OrgPolicyPolicyObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Policy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := OrgPolicyPolicySpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	desiredPb.Name = a.id.String()
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &orgpolicypb.UpdatePolicyRequest{
		Policy: desiredPb,
	}
	updated, err := a.gcpClient.UpdatePolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Policy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Policy", "name", a.id)

	status := &krm.OrgPolicyPolicyStatus{}
	status.ObservedState = OrgPolicyPolicyObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *PolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.OrgPolicyPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(OrgPolicyPolicySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	parentRef, _, err := krm.ParsePolicyExternal(a.actual.Name)
	if parentRef.ProjectID != "" {
		obj.Spec.ProjectRef = &refs.ProjectRef{External: parentRef.String()}
	} else if parentRef.FolderID != "" {
		obj.Spec.FolderRef = &refs.FolderRef{External: parentRef.String()}
	} else if parentRef.OrganizationID != "" {
		obj.Spec.OrganizationRef = &refs.OrganizationRef{External: parentRef.String()}
	} else {
		return nil, fmt.Errorf("unknown parent type in name %q", a.actual.Name)
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.OrgPolicyPolicyGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *PolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Policy", "name", a.id)

	req := &orgpolicypb.DeletePolicyRequest{Name: a.id.String()}
	err := a.gcpClient.DeletePolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Policy, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Policy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Policy", "name", a.id)

	return true, nil
}
