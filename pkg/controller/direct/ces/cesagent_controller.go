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
// krm.group: ces.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: CesAgent

package ces

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ces/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/ces/apiv1"
	cespb "cloud.google.com/go/ces/apiv1/cespb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CesAgentGVK, NewAgentModel)
}

func NewAgentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAgent{config: *config}, nil
}

var _ directbase.Model = &modelAgent{}

type modelAgent struct {
	config config.ControllerConfig
}

func (m *modelAgent) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CesAgent{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.CesAgentIdentity)

	// Get ces GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AgentAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelAgent) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//ces.googleapis.com/") {
		return nil, nil
	}
	id := &krm.CesAgentIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AgentAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelAgent) client(ctx context.Context) (*gcp.AgentClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAgentRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Agent client: %w", err)
	}
	return gcpClient, err
}

type AgentAdapter struct {
	id        *krm.CesAgentIdentity
	gcpClient *gcp.AgentClient
	desired   *krm.CesAgent
	actual    *cespb.Agent
}

var _ directbase.Adapter = &AgentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AgentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Agent", "name", a.id)

	req := &cespb.GetAgentRequest{
		Name: a.id.String(),
	}
	agentpb, err := a.gcpClient.GetAgent(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Agent %q: %w", a.id, err)
	}

	a.actual = agentpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AgentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Agent", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CesAgentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &cespb.CreateAgentRequest{
		Parent:  a.id.ParentString(),
		AgentId: a.id.Agent,
		Agent:   resource,
	}
	created, err := a.gcpClient.CreateAgent(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Agent %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Agent", "name", a.id)

	status := &krm.CesAgentStatus{}
	status.ObservedState = CesAgentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AgentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Agent", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		CesAgentSpec_FromProto,
		CesAgentSpec_ToProto,
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
		req := &cespb.UpdateAgentRequest{
			UpdateMask: updateMask,
			Agent:      resource,
		}
		updated, err = a.gcpClient.UpdateAgent(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Agent %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Agent", "name", a.id)
	}

	status := &krm.CesAgentStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = CesAgentObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AgentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CesAgent{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CesAgentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.CesAgentGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AgentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Agent", "name", a.id)

	req := &cespb.DeleteAgentRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteAgent(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Agent, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Agent %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Agent", "name", a.id)
	return true, nil
}
