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

package agentregistry

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/agentregistry/apiv1"
	agentregistrypb "cloud.google.com/go/agentregistry/apiv1/agentregistrypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/agentregistry/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AgentRegistryBindingGVK, NewAgentRegistryBindingModel)
}

func NewAgentRegistryBindingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAgentRegistryBinding{config: config}, nil
}

var _ directbase.Model = &modelAgentRegistryBinding{}

type modelAgentRegistryBinding struct {
	config *config.ControllerConfig
}

func (m *modelAgentRegistryBinding) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AgentRegistry REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelAgentRegistryBinding) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AgentRegistryBinding{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewAgentRegistryBindingIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Convert the KRM spec to API format
	mapCtx := &direct.MapContext{}
	desired := AgentRegistryBindingSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = id.String()

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &AgentRegistryBindingAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelAgentRegistryBinding) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.AgentRegistryBindingIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &AgentRegistryBindingAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type AgentRegistryBindingAdapter struct {
	id        *krm.AgentRegistryBindingIdentity
	gcpClient *gcp.Client
	desired   *agentregistrypb.Binding
	actual    *agentregistrypb.Binding
}

var _ directbase.Adapter = &AgentRegistryBindingAdapter{}

func (a *AgentRegistryBindingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting AgentRegistryBinding", "name", fqn)

	req := &agentregistrypb.GetBindingRequest{
		Name: fqn,
	}
	resource, err := a.gcpClient.GetBinding(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AgentRegistryBinding %q: %w", fqn, err)
	}

	a.actual = resource
	return true, nil
}

func (a *AgentRegistryBindingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	parent := a.id.Parent().String()
	fqn := a.id.String()
	log.V(2).Info("creating AgentRegistryBinding", "name", fqn)

	req := &agentregistrypb.CreateBindingRequest{
		Parent:    parent,
		BindingId: a.id.ID(),
		Binding:   a.desired,
	}
	op, err := a.gcpClient.CreateBinding(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AgentRegistryBinding %s: %w", a.id.ID(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AgentRegistryBinding %s creation: %w", a.id.ID(), err)
	}
	log.V(2).Info("successfully created AgentRegistryBinding", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *AgentRegistryBindingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating AgentRegistryBinding", "name", fqn)

	diffs, updateMask, err := compareBinding(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.Clone(a.desired).(*agentregistrypb.Binding)
		desired.Name = fqn

		req := &agentregistrypb.UpdateBindingRequest{
			Binding:    desired,
			UpdateMask: updateMask,
		}

		op, err := a.gcpClient.UpdateBinding(ctx, req)
		if err != nil {
			return fmt.Errorf("updating AgentRegistryBinding %s: %w", fqn, err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for AgentRegistryBinding %s update: %w", fqn, err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *AgentRegistryBindingAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *agentregistrypb.Binding) error {
	mapCtx := &direct.MapContext{}
	observedState := AgentRegistryBindingObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status := &krm.AgentRegistryBindingStatus{}
	status.ObservedState = observedState
	externalRef := latest.Name
	status.ExternalRef = &externalRef
	return op.UpdateStatus(ctx, status, nil)
}

func (a *AgentRegistryBindingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AgentRegistryBinding{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AgentRegistryBindingSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	parent := a.id.Parent()
	obj.Spec.ProjectRef = &refs.ProjectRef{External: parent.ProjectID}
	obj.Spec.Location = parent.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.AgentRegistryBindingGVK)
	return u, nil
}

func (a *AgentRegistryBindingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting AgentRegistryBinding", "name", fqn)

	req := &agentregistrypb.DeleteBindingRequest{Name: fqn}
	op, err := a.gcpClient.DeleteBinding(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent AgentRegistryBinding, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting AgentRegistryBinding %s: %w", fqn, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for AgentRegistryBinding %s deletion: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted AgentRegistryBinding", "name", fqn)
	return true, nil
}

func compareBinding(ctx context.Context, actual, desired *agentregistrypb.Binding) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, AgentRegistryBindingSpec_FromProto, AgentRegistryBindingSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
