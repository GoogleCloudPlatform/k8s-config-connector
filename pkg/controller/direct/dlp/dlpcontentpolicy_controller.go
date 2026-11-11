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
// krm.group: dlp.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DLPContentPolicy

package dlp

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dlp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/dlp/apiv2"
	dlppb "cloud.google.com/go/dlp/apiv2/dlppb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DLPContentPolicyGVK, NewContentPolicyModel)
}

func NewContentPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelContentPolicy{config: *config}, nil
}

var _ directbase.Model = &modelContentPolicy{}

type modelContentPolicy struct {
	config config.ControllerConfig
}

func (m *modelContentPolicy) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DLPContentPolicy{}
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
	id := i.(*krm.DLPContentPolicyIdentity)

	// Get dlp GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ContentPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelContentPolicy) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//dlp.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DLPContentPolicyIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ContentPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelContentPolicy) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ContentPolicy client: %w", err)
	}
	return gcpClient, err
}

type ContentPolicyAdapter struct {
	id        *krm.DLPContentPolicyIdentity
	gcpClient *gcp.Client
	desired   *krm.DLPContentPolicy
	actual    *dlppb.ContentPolicy
}

var _ directbase.Adapter = &ContentPolicyAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ContentPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ContentPolicy", "name", a.id)

	req := &dlppb.GetContentPolicyRequest{
		Name: a.id.String(),
	}
	contentpolicypb, err := a.gcpClient.GetContentPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ContentPolicy %q: %w", a.id, err)
	}

	a.actual = contentpolicypb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ContentPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ContentPolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DLPContentPolicySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &dlppb.CreateContentPolicyRequest{
		Parent:          a.id.ParentString(),
		ContentPolicyId: a.id.ContentPolicy,
		ContentPolicy:   resource,
	}
	created, err := a.gcpClient.CreateContentPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ContentPolicy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ContentPolicy", "name", a.id)

	status := &krm.DLPContentPolicyStatus{}
	status.ObservedState = DLPContentPolicyObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ContentPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ContentPolicy", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DLPContentPolicySpec_FromProto,
		DLPContentPolicySpec_ToProto,
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
		req := &dlppb.UpdateContentPolicyRequest{
			Name:          a.id.String(),
			UpdateMask:    updateMask,
			ContentPolicy: resource,
		}
		updated, err = a.gcpClient.UpdateContentPolicy(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ContentPolicy %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated ContentPolicy", "name", a.id)
	}

	status := &krm.DLPContentPolicyStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DLPContentPolicyObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ContentPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DLPContentPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DLPContentPolicySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DLPContentPolicyGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ContentPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ContentPolicy", "name", a.id)

	req := &dlppb.DeleteContentPolicyRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteContentPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ContentPolicy, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ContentPolicy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ContentPolicy", "name", a.id)
	return true, nil
}
