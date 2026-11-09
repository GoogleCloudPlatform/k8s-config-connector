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
// krm.group: aiplatform.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: AIPlatformCachedContent

package aiplatform

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/aiplatform/apiv1"
	aiplatformpb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AIPlatformCachedContentGVK, NewCachedContentModel)
}

func NewCachedContentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCachedContent{config: *config}, nil
}

var _ directbase.Model = &modelCachedContent{}

type modelCachedContent struct {
	config config.ControllerConfig
}

func (m *modelCachedContent) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AIPlatformCachedContent{}
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
	id := i.(*krm.AIPlatformCachedContentIdentity)

	// Get aiplatform GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &CachedContentAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelCachedContent) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//aiplatform.googleapis.com/") {
		return nil, nil
	}
	id := &krm.AIPlatformCachedContentIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &CachedContentAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelCachedContent) client(ctx context.Context, location string) (*gcp.GenAiCacheClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)))
	}
	gcpClient, err := gcp.NewGenAiCacheClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CachedContent client: %w", err)
	}
	return gcpClient, err
}

type CachedContentAdapter struct {
	id        *krm.AIPlatformCachedContentIdentity
	gcpClient *gcp.GenAiCacheClient
	desired   *krm.AIPlatformCachedContent
	actual    *aiplatformpb.CachedContent
}

var _ directbase.Adapter = &CachedContentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *CachedContentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CachedContent", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &aiplatformpb.GetCachedContentRequest{
		Name: a.id.String(),
	}
	cachedcontentpb, err := a.gcpClient.GetCachedContent(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CachedContent %q: %w", a.id, err)
	}

	a.actual = cachedcontentpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CachedContentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CachedContent", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AIPlatformCachedContentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &aiplatformpb.CreateCachedContentRequest{
		Parent:        a.id.ParentString(),
		CachedContent: resource,
	}
	created, err := a.gcpClient.CreateCachedContent(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CachedContent %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created CachedContent", "name", a.id)

	status := &krm.AIPlatformCachedContentStatus{}
	status.ObservedState = AIPlatformCachedContentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.AIPlatformCachedContentIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.CachedContent = parsedID.CachedContent
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CachedContentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CachedContent", "name", a.id)

	u := updateOp.GetUnstructured()
	normalize := func(ctx context.Context, pb *aiplatformpb.CachedContent) error {
		if u != nil && a.desired.Spec.TTL != nil && a.actual.GetExpireTime() != nil {
			obsGen, found, _ := unstructured.NestedInt64(u.Object, "status", "observedGeneration")
			if found && obsGen == u.GetGeneration() {
				// The spec has not changed since the last reconciliation.
				// Preserve actual expiration time to avoid re-triggering updates for input-only ttl.
				pb.Expiration = a.actual.GetExpiration()
			}
		}
		return nil
	}

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AIPlatformCachedContentSpec_FromProto,
		AIPlatformCachedContentSpec_ToProto,
		normalize,
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
		req := &aiplatformpb.UpdateCachedContentRequest{
			UpdateMask:    updateMask,
			CachedContent: resource,
		}
		updated, err = a.gcpClient.UpdateCachedContent(ctx, req)
		if err != nil {
			return fmt.Errorf("updating CachedContent %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated CachedContent", "name", a.id)
	}

	status := &krm.AIPlatformCachedContentStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = AIPlatformCachedContentObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *CachedContentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AIPlatformCachedContent{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AIPlatformCachedContentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.AIPlatformCachedContentGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *CachedContentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CachedContent", "name", a.id)

	req := &aiplatformpb.DeleteCachedContentRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteCachedContent(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent CachedContent, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting CachedContent %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted CachedContent", "name", a.id)
	return true, nil
}
