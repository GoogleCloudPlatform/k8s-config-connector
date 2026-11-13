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
// krm.group: dialogflowcx.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DialogflowCXAgent

package dialogflowcx

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/dialogflow/cx/apiv3"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflowcx/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DialogflowCXAgentGVK, NewAgentModel)
}

func NewAgentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAgent{config: *config}, nil
}

var _ directbase.Model = &modelAgent{}

type modelAgent struct {
	config config.ControllerConfig
}

func (m *modelAgent) client(ctx context.Context, location string) (*gcp.AgentsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-dialogflow.googleapis.com", location)))
	}
	gcpClient, err := gcp.NewAgentsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Agents REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelAgent) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DialogflowCXAgent{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DialogflowCXAgentIdentity)

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &AgentAdapter{
		id:         id,
		gcpClient:  gcpClient,
		desiredKRM: obj,
	}, nil
}

func (m *modelAgent) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//dialogflow.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DialogflowCXAgentIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &AgentAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type AgentAdapter struct {
	id         *krm.DialogflowCXAgentIdentity
	gcpClient  *gcp.AgentsClient
	desiredKRM *krm.DialogflowCXAgent
	actual     *cxpb.Agent
}

var _ directbase.Adapter = &AgentAdapter{}

func (a *AgentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DialogflowCXAgent", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &cxpb.GetAgentRequest{
		Name: a.id.String(),
	}
	agentpb, err := a.gcpClient.GetAgent(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DialogflowCXAgent %q: %w", a.id, err)
	}

	a.actual = agentpb
	return true, nil
}

func (a *AgentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DialogflowCXAgent", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desiredKRM.DeepCopy()
	resource := DialogflowCXAgentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &cxpb.CreateAgentRequest{
		Parent: a.id.ParentString(),
		Agent:  resource,
	}
	created, err := a.gcpClient.CreateAgent(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DialogflowCXAgent %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created DialogflowCXAgent", "name", created.GetName())

	parsedID := &krm.DialogflowCXAgentIdentity{}
	if err := parsedID.FromExternal(created.GetName()); err != nil {
		return err
	}
	a.id.Agent = parsedID.Agent

	status := &krm.DialogflowCXAgentStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.Name = direct.LazyPtr(a.id.Agent)
	if created.GetStartFlow() != "" {
		status.StartFlow = direct.LazyPtr(created.GetStartFlow())
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *AgentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DialogflowCXAgent", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desiredKRM.Spec,
		a.actual,
		DialogflowCXAgentSpec_FromProto,
		DialogflowCXAgentSpec_ToProto,
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
		req := &cxpb.UpdateAgentRequest{
			UpdateMask: updateMask,
			Agent:      resource,
		}
		updated, err = a.gcpClient.UpdateAgent(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DialogflowCXAgent %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated DialogflowCXAgent", "name", a.id)
	}

	status := &krm.DialogflowCXAgentStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.Name = direct.LazyPtr(a.id.Agent)
	if updated.GetStartFlow() != "" {
		status.StartFlow = direct.LazyPtr(updated.GetStartFlow())
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *AgentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DialogflowCXAgent{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DialogflowCXAgentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Agent)
	u.SetGroupVersionKind(krm.DialogflowCXAgentGVK)

	return u, nil
}

func (a *AgentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DialogflowCXAgent", "name", a.id)

	if !a.id.HasIdentitySpecified() {
		return true, nil
	}

	req := &cxpb.DeleteAgentRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteAgent(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent DialogflowCXAgent, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DialogflowCXAgent %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DialogflowCXAgent", "name", a.id)
	return true, nil
}
