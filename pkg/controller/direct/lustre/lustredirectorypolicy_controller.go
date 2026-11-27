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
// krm.group: lustre.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: LustreDirectoryPolicy

package lustre

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/lustre/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/lustre/apiv1"
	lustrepb "cloud.google.com/go/lustre/apiv1/lustrepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.LustreDirectoryPolicyGVK, NewDirectoryPolicyModel)
}

func NewDirectoryPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDirectoryPolicy{config: *config}, nil
}

var _ directbase.Model = &modelDirectoryPolicy{}

type modelDirectoryPolicy struct {
	config config.ControllerConfig
}

func (m *modelDirectoryPolicy) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.LustreDirectoryPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.LustreDirectoryPolicyIdentity)

	// Get lustre GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DirectoryPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelDirectoryPolicy) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//lustre.googleapis.com/") {
		return nil, nil
	}
	id := &krm.LustreDirectoryPolicyIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DirectoryPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelDirectoryPolicy) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryPolicy client: %w", err)
	}
	return gcpClient, err
}

type DirectoryPolicyAdapter struct {
	id        *krm.LustreDirectoryPolicyIdentity
	gcpClient *gcp.Client
	desired   *krm.LustreDirectoryPolicy
	actual    *lustrepb.DirectoryPolicy
}

var _ directbase.Adapter = &DirectoryPolicyAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DirectoryPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DirectoryPolicy", "name", a.id)

	req := &lustrepb.GetDirectoryPolicyRequest{
		Name: a.id.String(),
	}
	directorypolicypb, err := a.gcpClient.GetDirectoryPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DirectoryPolicy %q: %w", a.id, err)
	}

	a.actual = directorypolicypb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DirectoryPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DirectoryPolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := LustreDirectoryPolicySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &lustrepb.CreateDirectoryPolicyRequest{
		Parent:            a.id.ParentString(),
		DirectoryPolicyId: a.id.DirectoryPolicy,
		DirectoryPolicy:   resource,
	}
	op, err := a.gcpClient.CreateDirectoryPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DirectoryPolicy %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DirectoryPolicy %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created DirectoryPolicy", "name", a.id)

	status := &krm.LustreDirectoryPolicyStatus{}
	status.ObservedState = LustreDirectoryPolicyObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DirectoryPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DirectoryPolicy", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		LustreDirectoryPolicySpec_FromProto,
		LustreDirectoryPolicySpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating LustreDirectoryPolicy is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.LustreDirectoryPolicyStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = LustreDirectoryPolicyObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DirectoryPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.LustreDirectoryPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LustreDirectoryPolicySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.LustreDirectoryPolicyGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DirectoryPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DirectoryPolicy", "name", a.id)

	req := &lustrepb.DeleteDirectoryPolicyRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteDirectoryPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent DirectoryPolicy, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DirectoryPolicy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DirectoryPolicy", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete DirectoryPolicy %s: %w", a.id, err)
	}
	return true, nil
}
