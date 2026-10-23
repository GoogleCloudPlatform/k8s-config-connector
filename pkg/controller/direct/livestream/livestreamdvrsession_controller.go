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
// krm.group: livestream.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: LiveStreamDVRSession

package livestream

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/livestream/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/video/livestream/apiv1"
	livestreampb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.LiveStreamDVRSessionGVK, NewDvrSessionModel)
}

func NewDvrSessionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDvrSession{config: *config}, nil
}

var _ directbase.Model = &modelDvrSession{}

type modelDvrSession struct {
	config config.ControllerConfig
}

func (m *modelDvrSession) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.LiveStreamDVRSession{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.LiveStreamDVRSessionIdentity)

	// Get livestream GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DvrSessionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelDvrSession) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//livestream.googleapis.com/") {
		return nil, nil
	}
	id := &krm.LiveStreamDVRSessionIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DvrSessionAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelDvrSession) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DvrSession client: %w", err)
	}
	return gcpClient, err
}

type DvrSessionAdapter struct {
	id        *krm.LiveStreamDVRSessionIdentity
	gcpClient *gcp.Client
	desired   *krm.LiveStreamDVRSession
	actual    *livestreampb.DvrSession
}

var _ directbase.Adapter = &DvrSessionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DvrSessionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DvrSession", "name", a.id)

	req := &livestreampb.GetDvrSessionRequest{
		Name: a.id.String(),
	}
	dvrsessionpb, err := a.gcpClient.GetDvrSession(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DvrSession %q: %w", a.id, err)
	}

	a.actual = dvrsessionpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DvrSessionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DvrSession", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := LiveStreamDVRSessionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &livestreampb.CreateDvrSessionRequest{
		Parent:       a.id.ParentString(),
		DvrSessionId: a.id.DvrSession,
		DvrSession:   resource,
	}
	op, err := a.gcpClient.CreateDvrSession(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DvrSession %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DvrSession %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created DvrSession", "name", a.id)

	status := &krm.LiveStreamDVRSessionStatus{}
	status.ObservedState = LiveStreamDVRSessionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DvrSessionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DvrSession", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		LiveStreamDVRSessionSpec_FromProto,
		LiveStreamDVRSessionSpec_ToProto,
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
		req := &livestreampb.UpdateDvrSessionRequest{
			UpdateMask: updateMask,
			DvrSession: resource,
		}
		op, err := a.gcpClient.UpdateDvrSession(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DvrSession %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("DvrSession %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated DvrSession", "name", a.id)
	}

	status := &krm.LiveStreamDVRSessionStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = LiveStreamDVRSessionObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DvrSessionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.LiveStreamDVRSession{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LiveStreamDVRSessionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.LiveStreamDVRSessionGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DvrSessionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DvrSession", "name", a.id)

	req := &livestreampb.DeleteDvrSessionRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteDvrSession(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent DvrSession, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DvrSession %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DvrSession", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete DvrSession %s: %w", a.id, err)
	}
	return true, nil
}
