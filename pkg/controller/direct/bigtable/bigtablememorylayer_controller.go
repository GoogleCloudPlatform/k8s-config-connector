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

package bigtable

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	admin "cloud.google.com/go/bigtable/admin/apiv2"
	bigtablepb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

// memoryConfigPath is the only field of MemoryLayer that is settable. Sending
// an empty memory_config enables the memory layer; clearing it disables the
// memory layer.
const memoryConfigPath = "memory_config"

func init() {
	registry.RegisterModel(krm.BigtableMemoryLayerGVK, NewMemoryLayerModel)
}

func NewMemoryLayerModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelMemoryLayer{config: *config}, nil
}

var _ directbase.Model = &modelMemoryLayer{}

type modelMemoryLayer struct {
	config config.ControllerConfig
}

func (m *modelMemoryLayer) client(ctx context.Context) (*admin.BigtableInstanceAdminClient, error) {
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, fmt.Errorf("building BigtableMemoryLayer client options: %w", err)
	}
	gcpClient, err := admin.NewBigtableInstanceAdminClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigtableMemoryLayer client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelMemoryLayer) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigtableMemoryLayer{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewMemoryLayerIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &MemoryLayerAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelMemoryLayer) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type MemoryLayerAdapter struct {
	id        *krm.MemoryLayerIdentity
	gcpClient *admin.BigtableInstanceAdminClient
	desired   *krm.BigtableMemoryLayer
	actual    *bigtablepb.MemoryLayer
}

var _ directbase.Adapter = &MemoryLayerAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
//
// A memory layer is a singleton that the Bigtable API reports for every
// cluster, whether or not it is turned on, so "found" here means "enabled":
// a memory layer with no memory_config is disabled, and is reported as not
// found so that Create turns it on.
func (a *MemoryLayerAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigtableMemoryLayer", "name", a.id)

	ml, err := a.gcpClient.GetMemoryLayer(ctx, &bigtablepb.GetMemoryLayerRequest{Name: a.id.String()})
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigtableMemoryLayer %q: %w", a.id, err)
	}
	// Keep the response even when the memory layer is disabled: it carries the
	// etag that Create needs to send back.
	a.actual = ml

	if ml.GetMemoryConfig() == nil {
		log.V(2).Info("BigtableMemoryLayer is disabled", "name", a.id)
		return false, nil
	}
	return true, nil
}

// Create enables the memory layer on the parent cluster.
func (a *MemoryLayerAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigtableMemoryLayer", "name", a.id)

	// An empty memory_config enables the memory layer.
	if err := a.setMemoryConfig(ctx, &bigtablepb.MemoryLayer_MemoryConfig{}); err != nil {
		return fmt.Errorf("creating BigtableMemoryLayer %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created BigtableMemoryLayer", "name", a.id)

	return a.setStatus(ctx, func(status *krm.BigtableMemoryLayerStatus) error {
		return createOp.UpdateStatus(ctx, status, nil)
	})
}

// Update refreshes the observed state.
//
// spec.clusterRef is the only spec field, and it is immutable: a change is
// rejected by NewMemoryLayerIdentity before we get here. Find only reports the
// memory layer as found once it is enabled, so there is never a settable field
// out of sync at this point and no GCP call is needed.
func (a *MemoryLayerAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("no field needs update", "name", a.id)

	return a.setStatus(ctx, func(status *krm.BigtableMemoryLayerStatus) error {
		return updateOp.UpdateStatus(ctx, status, nil)
	})
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *MemoryLayerAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigtableMemoryLayer{}
	obj.Spec.ClusterRef = &krm.ClusterRef{External: a.id.ParentString()}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Parent().ID())
	u.SetGroupVersionKind(krm.BigtableMemoryLayerGVK)
	u.Object = uObj
	return u, nil
}

// Delete disables the memory layer on the parent cluster.
//
// A memory layer cannot be removed, only turned off, so this clears
// memory_config rather than deleting a resource.
func (a *MemoryLayerAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigtableMemoryLayer", "name", a.id)

	if err := a.setMemoryConfig(ctx, nil); err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent BigtableMemoryLayer, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BigtableMemoryLayer %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BigtableMemoryLayer", "name", a.id)
	return true, nil
}

// setMemoryConfig sets (enables) or clears (disables) the memory layer's
// memory_config and waits for the resulting long-running operation.
func (a *MemoryLayerAdapter) setMemoryConfig(ctx context.Context, memoryConfig *bigtablepb.MemoryLayer_MemoryConfig) error {
	ml := &bigtablepb.MemoryLayer{
		Name:         a.id.String(),
		MemoryConfig: memoryConfig,
	}
	// Send the etag we last read, so the server rejects the write if the memory
	// layer changed underneath us.
	if a.actual != nil {
		ml.Etag = a.actual.GetEtag()
	}

	op, err := a.gcpClient.UpdateMemoryLayer(ctx, &bigtablepb.UpdateMemoryLayerRequest{
		MemoryLayer: ml,
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{memoryConfigPath}},
	})
	if err != nil {
		return err
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for update of memory layer: %w", err)
	}
	a.actual = updated
	return nil
}

func (a *MemoryLayerAdapter) setStatus(ctx context.Context, updateFn func(status *krm.BigtableMemoryLayerStatus) error) error {
	status := &krm.BigtableMemoryLayerStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())

	mapCtx := &direct.MapContext{}
	status.ObservedState = BigtableMemoryLayerObservedState_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateFn(status)
}
