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
// krm.group: designcenter.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DesignCenterSpace

package designcenter

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/designcenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/designcenter/apiv1"
	designcenterpb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DesignCenterSpaceGVK, NewSpaceModel)
}

func NewSpaceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelSpace{config: *config}, nil
}

var _ directbase.Model = &modelSpace{}

type modelSpace struct {
	config config.ControllerConfig
}

func (m *modelSpace) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DesignCenterSpace{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DesignCenterSpaceIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get designcenter GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SpaceAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelSpace) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//designcenter.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DesignCenterSpaceIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SpaceAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelSpace) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Space client: %w", err)
	}
	return gcpClient, err
}

type SpaceAdapter struct {
	id        *krm.DesignCenterSpaceIdentity
	gcpClient *gcp.Client
	desired   *krm.DesignCenterSpace
	actual    *designcenterpb.Space
}

var _ directbase.Adapter = &SpaceAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *SpaceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Space", "name", a.id)

	req := &designcenterpb.GetSpaceRequest{
		Name: a.id.String(),
	}
	spacepb, err := a.gcpClient.GetSpace(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Space %q: %w", a.id, err)
	}

	a.actual = spacepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *SpaceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Space", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DesignCenterSpaceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &designcenterpb.CreateSpaceRequest{
		Parent:  a.id.ParentString(),
		SpaceId: a.id.Space,
		Space:   resource,
	}
	created, err := a.gcpClient.CreateSpace(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Space %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Space", "name", a.id)

	status := &krm.DesignCenterSpaceStatus{}
	status.ObservedState = DesignCenterSpaceObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *SpaceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Space", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DesignCenterSpaceSpec_FromProto,
		DesignCenterSpaceSpec_ToProto,
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
		req := &designcenterpb.UpdateSpaceRequest{
			UpdateMask: updateMask,
			Space:      resource,
		}
		updated, err = a.gcpClient.UpdateSpace(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Space %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Space", "name", a.id)
	}

	status := &krm.DesignCenterSpaceStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DesignCenterSpaceObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *SpaceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DesignCenterSpace{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DesignCenterSpaceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DesignCenterSpaceGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *SpaceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Space", "name", a.id)

	req := &designcenterpb.DeleteSpaceRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteSpace(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Space, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Space %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Space", "name", a.id)
	return true, nil
}
