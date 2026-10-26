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
// krm.group: dialogflow.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DialogflowTransitionRouteGroup

package dialogflow

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/dialogflow/cx/apiv3"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DialogflowTransitionRouteGroupGVK, NewTransitionRouteGroupModel)
}

func NewTransitionRouteGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTransitionRouteGroup{config: *config}, nil
}

var _ directbase.Model = &modelTransitionRouteGroup{}

type modelTransitionRouteGroup struct {
	config config.ControllerConfig
}

func (m *modelTransitionRouteGroup) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DialogflowTransitionRouteGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DialogflowTransitionRouteGroupIdentity)

	// Get dialogflow GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &TransitionRouteGroupAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelTransitionRouteGroup) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//dialogflow.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DialogflowTransitionRouteGroupIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &TransitionRouteGroupAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelTransitionRouteGroup) client(ctx context.Context) (*gcp.TransitionRouteGroupsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewTransitionRouteGroupsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building TransitionRouteGroup client: %w", err)
	}
	return gcpClient, err
}

type TransitionRouteGroupAdapter struct {
	id        *krm.DialogflowTransitionRouteGroupIdentity
	gcpClient *gcp.TransitionRouteGroupsClient
	desired   *krm.DialogflowTransitionRouteGroup
	actual    *cxpb.TransitionRouteGroup
}

var _ directbase.Adapter = &TransitionRouteGroupAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *TransitionRouteGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting TransitionRouteGroup", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &cxpb.GetTransitionRouteGroupRequest{
		Name: a.id.String(),
	}
	transitionroutegrouppb, err := a.gcpClient.GetTransitionRouteGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting TransitionRouteGroup %q: %w", a.id, err)
	}

	a.actual = transitionroutegrouppb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TransitionRouteGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating TransitionRouteGroup", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DialogflowTransitionRouteGroupSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &cxpb.CreateTransitionRouteGroupRequest{
		Parent:               a.id.ParentString(),
		TransitionRouteGroup: resource,
	}
	created, err := a.gcpClient.CreateTransitionRouteGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating TransitionRouteGroup %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created TransitionRouteGroup", "name", a.id)

	status := &krm.DialogflowTransitionRouteGroupStatus{}
	status.ObservedState = DialogflowTransitionRouteGroupObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.DialogflowTransitionRouteGroupIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.TransitionRouteGroup = parsedID.TransitionRouteGroup
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TransitionRouteGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating TransitionRouteGroup", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DialogflowTransitionRouteGroupSpec_FromProto,
		DialogflowTransitionRouteGroupSpec_ToProto,
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
		req := &cxpb.UpdateTransitionRouteGroupRequest{
			UpdateMask:           updateMask,
			TransitionRouteGroup: resource,
		}
		updated, err = a.gcpClient.UpdateTransitionRouteGroup(ctx, req)
		if err != nil {
			return fmt.Errorf("updating TransitionRouteGroup %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated TransitionRouteGroup", "name", a.id)
	}

	status := &krm.DialogflowTransitionRouteGroupStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DialogflowTransitionRouteGroupObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *TransitionRouteGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DialogflowTransitionRouteGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DialogflowTransitionRouteGroupSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DialogflowTransitionRouteGroupGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *TransitionRouteGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting TransitionRouteGroup", "name", a.id)

	req := &cxpb.DeleteTransitionRouteGroupRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteTransitionRouteGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent TransitionRouteGroup, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting TransitionRouteGroup %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted TransitionRouteGroup", "name", a.id)
	return true, nil
}
