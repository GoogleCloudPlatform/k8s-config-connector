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
// krm.group: chronicle.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ChronicleFindingsRefinement

package chronicle

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/chronicle/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/chronicle/apiv1"
	chroniclepb "cloud.google.com/go/chronicle/apiv1/chroniclepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ChronicleFindingsRefinementGVK, NewFindingsRefinementModel)
}

func NewFindingsRefinementModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelFindingsRefinement{config: *config}, nil
}

var _ directbase.Model = &modelFindingsRefinement{}

type modelFindingsRefinement struct {
	config config.ControllerConfig
}

func (m *modelFindingsRefinement) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ChronicleFindingsRefinement{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ChronicleFindingsRefinementIdentity)

	// Get chronicle GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &FindingsRefinementAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelFindingsRefinement) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//chronicle.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ChronicleFindingsRefinementIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &FindingsRefinementAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelFindingsRefinement) client(ctx context.Context) (*gcp.FindingsRefinementClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewFindingsRefinementRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building FindingsRefinement client: %w", err)
	}
	return gcpClient, err
}

type FindingsRefinementAdapter struct {
	id        *krm.ChronicleFindingsRefinementIdentity
	gcpClient *gcp.FindingsRefinementClient
	desired   *krm.ChronicleFindingsRefinement
	actual    *chroniclepb.FindingsRefinement
}

var _ directbase.Adapter = &FindingsRefinementAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *FindingsRefinementAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FindingsRefinement", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &chroniclepb.GetFindingsRefinementRequest{
		Name: a.id.String(),
	}
	findingsrefinementpb, err := a.gcpClient.GetFindingsRefinement(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FindingsRefinement %q: %w", a.id, err)
	}

	a.actual = findingsrefinementpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FindingsRefinementAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FindingsRefinement", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ChronicleFindingsRefinementSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &chroniclepb.CreateFindingsRefinementRequest{
		Parent:             a.id.ParentString(),
		FindingsRefinement: resource,
	}
	created, err := a.gcpClient.CreateFindingsRefinement(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FindingsRefinement %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created FindingsRefinement", "name", a.id)

	status := &krm.ChronicleFindingsRefinementStatus{}
	status.ObservedState = ChronicleFindingsRefinementObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.ChronicleFindingsRefinementIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.FindingsRefinement = parsedID.FindingsRefinement
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FindingsRefinementAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FindingsRefinement", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ChronicleFindingsRefinementSpec_FromProto,
		ChronicleFindingsRefinementSpec_ToProto,
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
		req := &chroniclepb.UpdateFindingsRefinementRequest{
			UpdateMask:         updateMask,
			FindingsRefinement: resource,
		}
		updated, err = a.gcpClient.UpdateFindingsRefinement(ctx, req)
		if err != nil {
			return fmt.Errorf("updating FindingsRefinement %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated FindingsRefinement", "name", a.id)
	}

	status := &krm.ChronicleFindingsRefinementStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ChronicleFindingsRefinementObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *FindingsRefinementAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ChronicleFindingsRefinement{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ChronicleFindingsRefinementSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ChronicleFindingsRefinementGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *FindingsRefinementAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deletion not supported for FindingsRefinement, removing finalizer without deleting GCP resource", "name", a.id)
	return true, nil
}
