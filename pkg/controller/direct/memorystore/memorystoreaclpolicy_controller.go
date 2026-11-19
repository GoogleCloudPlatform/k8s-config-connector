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
// krm.group: memorystore.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: MemorystoreACLPolicy

package memorystore

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/memorystore/apiv1"
	memorystorepb "cloud.google.com/go/memorystore/apiv1/memorystorepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.MemorystoreACLPolicyGVK, NewAclPolicyModel)
}

func NewAclPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAclPolicy{config: *config}, nil
}

var _ directbase.Model = &modelAclPolicy{}

type modelAclPolicy struct {
	config config.ControllerConfig
}

func (m *modelAclPolicy) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.MemorystoreACLPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, err
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.MemorystoreACLPolicyIdentity)

	// Get memorystore GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AclPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelAclPolicy) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//memorystore.googleapis.com/") {
		return nil, nil
	}
	id := &krm.MemorystoreACLPolicyIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AclPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

func (m *modelAclPolicy) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AclPolicy client: %w", err)
	}
	return gcpClient, err
}

type AclPolicyAdapter struct {
	id        *krm.MemorystoreACLPolicyIdentity
	gcpClient *gcp.Client
	desired   *krm.MemorystoreACLPolicy
	actual    *memorystorepb.AclPolicy
}

var _ directbase.Adapter = &AclPolicyAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AclPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AclPolicy", "name", a.id)

	req := &memorystorepb.GetAclPolicyRequest{
		Name: a.id.String(),
	}
	aclpolicypb, err := a.gcpClient.GetAclPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AclPolicy %q: %w", a.id, err)
	}

	a.actual = aclpolicypb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AclPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AclPolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := MemorystoreACLPolicySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &memorystorepb.CreateAclPolicyRequest{
		Parent:      a.id.ParentString(),
		AclPolicyId: a.id.AclPolicy,
		AclPolicy:   resource,
	}
	created, err := a.gcpClient.CreateAclPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AclPolicy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created AclPolicy", "name", a.id)

	return a.setStatus(ctx, createOp, created)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AclPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AclPolicy", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		MemorystoreACLPolicySpec_FromProto,
		MemorystoreACLPolicySpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.setStatus(ctx, updateOp, a.actual)
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)
	resource.Name = a.id.String()
	req := &memorystorepb.UpdateAclPolicyRequest{
		UpdateMask: updateMask,
		AclPolicy:  resource,
	}
	op, err := a.gcpClient.UpdateAclPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AclPolicy %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("AclPolicy %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated AclPolicy", "name", a.id)

	if updated == nil {
		updated = a.actual
	}
	return a.setStatus(ctx, updateOp, updated)
}

func (a *AclPolicyAdapter) setStatus(ctx context.Context, op directbase.Operation, actual *memorystorepb.AclPolicy) error {
	status := &krm.MemorystoreACLPolicyStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = MemorystoreACLPolicyObservedState_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AclPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MemorystoreACLPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MemorystoreACLPolicySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.MemorystoreACLPolicyGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AclPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AclPolicy", "name", a.id)

	req := &memorystorepb.DeleteAclPolicyRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteAclPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent AclPolicy, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting AclPolicy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted AclPolicy", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete AclPolicy %s: %w", a.id, err)
	}
	return true, nil
}
