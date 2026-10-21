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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	admin "cloud.google.com/go/bigtable/admin/apiv2"
	bigtablepb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	corev1 "k8s.io/api/core/v1"
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
// a disabled memory layer is reported as not found so that Create turns it on.
//
// The service reports a MemoryLayer with state DISABLED for a cluster that has
// never turned one on, so state is what we key off. memory_config is not a
// reliable signal: storage_size_gib is output only, so a freshly enabled memory
// layer that has not been provisioned yet serializes the same way an empty
// memory_config would.
//
// "Enabled" is not the same as "serving": the state is ENABLING or RESIZING
// while the memory layer is still coming up, so a Ready condition on this object
// means the enable has been accepted, not that reads are being served from
// memory. Callers that need the latter should watch
// status.observedState.state for READY.
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

	switch ml.GetState() {
	case bigtablepb.MemoryLayer_DISABLED:
		log.V(2).Info("BigtableMemoryLayer is disabled", "name", a.id)
		return false, nil
	case bigtablepb.MemoryLayer_STATE_NOT_KNOWN:
		// Acting on this either way risks enabling a memory layer that is
		// already on, or reporting one as ready when it is off; requeue instead.
		return false, fmt.Errorf("BigtableMemoryLayer %s is in an unknown state", a.id)
	default:
		// READY, ENABLING, RESIZING, or a state added after this was written.
		return true, nil
	}
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

	return a.setStatus(ctx, createOp)
}

// Update refreshes the observed state.
//
// spec.clusterRef is the only spec field, and it is immutable: a change is
// rejected by NewMemoryLayerIdentity before we get here. Find only reports the
// memory layer as found once it is enabled, so there is never a settable field
// out of sync at this point and no GCP call is needed.
//
// The memory layer can still be coming up, though: see setStatus.
func (a *MemoryLayerAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("no field needs update", "name", a.id)

	return a.setStatus(ctx, updateOp)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *MemoryLayerAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	obj := &krm.BigtableMemoryLayer{}
	obj.Spec.ClusterRef = &krm.ClusterRef{External: a.id.ParentString()}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	// Set the name and GVK after assigning u.Object; assigning it replaces the
	// whole map and would otherwise discard them.
	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Parent().ID())
	u.SetGroupVersionKind(krm.BigtableMemoryLayerGVK)
	return u, nil
}

// Delete disables the memory layer on the parent cluster.
//
// A memory layer cannot be removed, only turned off, so this clears
// memory_config rather than deleting a resource.
//
// There is no already-deleted case to handle here: directbase only calls Delete
// when Find reported the memory layer as enabled, and Find already turns a
// missing parent cluster into "not found".
func (a *MemoryLayerAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigtableMemoryLayer", "name", a.id)

	if err := a.setMemoryConfig(ctx, nil); err != nil {
		return false, fmt.Errorf("deleting BigtableMemoryLayer %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BigtableMemoryLayer", "name", a.id)
	return true, nil
}

// setMemoryConfig sets (enables) or clears (disables) the memory layer's
// memory_config and waits for the resulting long-running operation.
func (a *MemoryLayerAdapter) setMemoryConfig(ctx context.Context, memoryConfig *bigtablepb.MemoryLayer_MemoryConfig) error {
	// The service requires the etag we last read, so that it can reject the
	// write if the memory layer changed underneath us. Find always populates
	// a.actual before Create or Delete runs, including for a disabled memory
	// layer, so a missing etag here is a programming error rather than
	// something to paper over by sending an empty one.
	if a.actual == nil || a.actual.GetEtag() == "" {
		return fmt.Errorf("no etag available for memory layer %s; Find() must run first", a.id)
	}

	ml := &bigtablepb.MemoryLayer{
		Name:         a.id.String(),
		MemoryConfig: memoryConfig,
		Etag:         a.actual.GetEtag(),
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

// setStatus writes observedState, and reports the object as not ready while the
// memory layer is still coming up.
//
// Enabling and disabling a memory layer is a long-running operation that can
// take tens of minutes. setMemoryConfig waits for it, but that wait is lost if
// the controller restarts, and an out-of-band resize can start one we never
// issued. In both cases a later reconcile finds the memory layer in ENABLING or
// RESIZING with no work of its own to do, and would otherwise let directbase
// mark the object UpToDate while reads are not yet served from memory. Requeue
// instead, so the object goes Ready when the state does.
func (a *MemoryLayerAdapter) setStatus(ctx context.Context, op directbase.Operation) error {
	status := &krm.BigtableMemoryLayerStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())

	mapCtx := &direct.MapContext{}
	status.ObservedState = BigtableMemoryLayerObservedState_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if state := a.actual.GetState(); state != bigtablepb.MemoryLayer_READY {
		klog.FromContext(ctx).V(2).Info("BigtableMemoryLayer is not ready yet", "name", a.id, "state", state)
		op.RequestRequeue()
		return op.UpdateStatus(ctx, status, readyCondition(corev1.ConditionFalse, k8s.Updating, k8s.UpdatingMessage))
	}
	return op.UpdateStatus(ctx, status, nil)
}

// readyCondition is a helper to create a KRM condition for resource status.
func readyCondition(status corev1.ConditionStatus, reason, message string) *v1alpha1.Condition {
	return &v1alpha1.Condition{
		Type:    v1alpha1.ReadyConditionType,
		Status:  status,
		Reason:  reason,
		Message: message,
	}
}
