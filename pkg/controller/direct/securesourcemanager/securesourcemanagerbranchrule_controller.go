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
// krm.group: securesourcemanager.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: SecureSourceManagerBranchRule

package securesourcemanager

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securesourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/securesourcemanager/apiv1"
	securesourcemanagerpb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.SecureSourceManagerBranchRuleGVK, NewBranchRuleModel)
}

func NewBranchRuleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBranchRule{config: *config}, nil
}

var _ directbase.Model = &modelBranchRule{}

type modelBranchRule struct {
	config config.ControllerConfig
}

func (m *modelBranchRule) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SecureSourceManagerBranchRule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.SecureSourceManagerBranchRuleIdentity)

	// Get securesourcemanager GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &BranchRuleAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelBranchRule) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//securesourcemanager.googleapis.com/") {
		return nil, nil
	}
	id := &krm.SecureSourceManagerBranchRuleIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &BranchRuleAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelBranchRule) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BranchRule client: %w", err)
	}
	return gcpClient, err
}

type BranchRuleAdapter struct {
	id        *krm.SecureSourceManagerBranchRuleIdentity
	gcpClient *gcp.Client
	desired   *krm.SecureSourceManagerBranchRule
	actual    *securesourcemanagerpb.BranchRule
}

var _ directbase.Adapter = &BranchRuleAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *BranchRuleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BranchRule", "name", a.id)

	req := &securesourcemanagerpb.GetBranchRuleRequest{
		Name: a.id.String(),
	}
	branchrulepb, err := a.gcpClient.GetBranchRule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BranchRule %q: %w", a.id, err)
	}

	a.actual = branchrulepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BranchRuleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BranchRule", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SecureSourceManagerBranchRuleSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &securesourcemanagerpb.CreateBranchRuleRequest{
		Parent:       a.id.ParentString(),
		BranchRuleId: a.id.BranchRule,
		BranchRule:   resource,
	}
	op, err := a.gcpClient.CreateBranchRule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BranchRule %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("BranchRule %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created BranchRule", "name", a.id)

	status := &krm.SecureSourceManagerBranchRuleStatus{}
	status.ObservedState = SecureSourceManagerBranchRuleObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BranchRuleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BranchRule", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		SecureSourceManagerBranchRuleSpec_v1alpha1_FromProto,
		SecureSourceManagerBranchRuleSpec_v1alpha1_ToProto,
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
		req := &securesourcemanagerpb.UpdateBranchRuleRequest{
			UpdateMask: updateMask,
			BranchRule: resource,
		}
		op, err := a.gcpClient.UpdateBranchRule(ctx, req)
		if err != nil {
			return fmt.Errorf("updating BranchRule %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("BranchRule %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated BranchRule", "name", a.id)
	}

	status := &krm.SecureSourceManagerBranchRuleStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = SecureSourceManagerBranchRuleObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BranchRuleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SecureSourceManagerBranchRule{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SecureSourceManagerBranchRuleSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.SecureSourceManagerBranchRuleGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *BranchRuleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BranchRule", "name", a.id)

	req := &securesourcemanagerpb.DeleteBranchRuleRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteBranchRule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent BranchRule, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BranchRule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BranchRule", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete BranchRule %s: %w", a.id, err)
	}
	return true, nil
}
