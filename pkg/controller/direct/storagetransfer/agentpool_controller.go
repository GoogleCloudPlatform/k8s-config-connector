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

package storagetransfer

import (
	"context"
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storagetransfer/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/storagetransfer/apiv1"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	storagetransferpb "cloud.google.com/go/storagetransfer/apiv1/storagetransferpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.StorageTransferAgentPoolGVK, NewAgentPoolModel)
}

func NewAgentPoolModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAgentPool{config: *config}, nil
}

var _ directbase.Model = &modelAgentPool{}

type modelAgentPool struct {
	config config.ControllerConfig
}

func (m *modelAgentPool) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AgentPool client: %w", err)
	}
	return gcpClient, err
}

func (m *modelAgentPool) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.StorageTransferAgentPool{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewAgentPoolIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get storagetransfer GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AgentPoolAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelAgentPool) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type AgentPoolAdapter struct {
	id        *krm.AgentPoolIdentity
	gcpClient *gcp.Client
	desired   *krm.StorageTransferAgentPool
	actual    *storagetransferpb.AgentPool
}

var _ directbase.Adapter = &AgentPoolAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AgentPoolAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AgentPool", "name", a.id)

	req := &storagetransferpb.GetAgentPoolRequest{Name: a.id.String()}
	agentpoolpb, err := a.gcpClient.GetAgentPool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AgentPool %q: %w", a.id, err)
	}

	a.actual = agentpoolpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AgentPoolAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AgentPool", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := StorageTransferAgentPoolSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &storagetransferpb.CreateAgentPoolRequest{
		ProjectId: a.id.Parent().ProjectID,
		AgentPool: resource,
	}
	created, err := a.gcpClient.CreateAgentPool(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AgentPool %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created AgentPool", "name", a.id)

	status := &krm.StorageTransferAgentPoolStatus{}
	status.ObservedState = StorageTransferAgentPoolObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AgentPoolAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AgentPool", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := StorageTransferAgentPoolSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var err error
	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.StorageTransferAgentPoolStatus{}
		status.ObservedState = StorageTransferAgentPoolObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
        }

	desiredPb.Name = a.id.String()
	req := &storagetransferpb.UpdateAgentPoolRequest{
		AgentPool:  desiredPb,
		UpdateMask: updateMask,
	}
	updated, err := a.gcpClient.UpdateAgentPool(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AgentPool %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated AgentPool", "name", a.id)

	status := &krm.StorageTransferAgentPoolStatus{}
	status.ObservedState = StorageTransferAgentPoolObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AgentPoolAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.StorageTransferAgentPool{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(StorageTransferAgentPoolSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.StorageTransferAgentPoolGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AgentPoolAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AgentPool", "name", a.id)

	req := &storagetransferpb.DeleteAgentPoolRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteAgentPool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent AgentPool, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting AgentPool %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted AgentPool", "name", a.id)

	return true, nil
}
