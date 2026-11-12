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
// krm.group: dialogflow.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DialogflowCompanionAgent

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

	gcp "cloud.google.com/go/dialogflow/apiv2"
	dialogflowpb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DialogflowCompanionAgentGVK, NewCompanionAgentModel)
}

func NewCompanionAgentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCompanionAgent{config: *config}, nil
}

var _ directbase.Model = &modelCompanionAgent{}

type modelCompanionAgent struct {
	config config.ControllerConfig
}

func (m *modelCompanionAgent) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DialogflowCompanionAgent{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DialogflowCompanionAgentIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get dialogflow GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &CompanionAgentAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelCompanionAgent) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//dialogflow.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DialogflowCompanionAgentIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &CompanionAgentAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelCompanionAgent) client(ctx context.Context, location string) (*gcp.CompanionAgentsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-dialogflow.googleapis.com:443", location)))
	}
	gcpClient, err := gcp.NewCompanionAgentsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CompanionAgent client: %w", err)
	}
	return gcpClient, err
}

type CompanionAgentAdapter struct {
	id        *krm.DialogflowCompanionAgentIdentity
	gcpClient *gcp.CompanionAgentsClient
	desired   *krm.DialogflowCompanionAgent
	actual    *dialogflowpb.CompanionAgent
}

var _ directbase.Adapter = &CompanionAgentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *CompanionAgentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CompanionAgent", "name", a.id)

	req := &dialogflowpb.GetCompanionAgentRequest{
		Name: a.id.String(),
	}
	companionagentpb, err := a.gcpClient.GetCompanionAgent(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CompanionAgent %q: %w", a.id, err)
	}

	a.actual = companionagentpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CompanionAgentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CompanionAgent", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DialogflowCompanionAgentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &dialogflowpb.CreateCompanionAgentRequest{
		Parent:           a.id.ParentString(),
		CompanionAgentId: a.id.CompanionAgent,
		CompanionAgent:   resource,
	}
	created, err := a.gcpClient.CreateCompanionAgent(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CompanionAgent %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created CompanionAgent", "name", a.id)

	status := &krm.DialogflowCompanionAgentStatus{}
	status.ObservedState = DialogflowCompanionAgentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CompanionAgentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CompanionAgent", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DialogflowCompanionAgentSpec_FromProto,
		DialogflowCompanionAgentSpec_ToProto,
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
		req := &dialogflowpb.UpdateCompanionAgentRequest{
			UpdateMask:     updateMask,
			CompanionAgent: resource,
		}
		updated, err = a.gcpClient.UpdateCompanionAgent(ctx, req)
		if err != nil {
			return fmt.Errorf("updating CompanionAgent %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated CompanionAgent", "name", a.id)
	}

	status := &krm.DialogflowCompanionAgentStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DialogflowCompanionAgentObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *CompanionAgentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DialogflowCompanionAgent{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DialogflowCompanionAgentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DialogflowCompanionAgentGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *CompanionAgentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CompanionAgent", "name", a.id)

	req := &dialogflowpb.DeleteCompanionAgentRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteCompanionAgent(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent CompanionAgent, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting CompanionAgent %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted CompanionAgent", "name", a.id)
	return true, nil
}
