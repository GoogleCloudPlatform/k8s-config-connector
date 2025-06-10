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

package storagecontrol

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/storage/control/apiv2"
	pb "cloud.google.com/go/storage/control/apiv2/controlpb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// anywhereCacheURLPrefix is the expected prefix for AnywhereCache URLs.
	anywhereCacheURLPrefix = "//storage.googleapis.com/"

	// AnywhereCache States (lowercase representation from GCP API).
	anywhereCacheStateCreating = "creating"
	anywhereCacheStateRunning  = "running"
	anywhereCacheStatePaused   = "paused"
	anywhereCacheStateDisabled = "disabled"
	anywhereCacheStateInvalid  = "invalid" // For cases where the state is unexpected
)

func init() {
	registry.RegisterModel(krm.StorageAnywhereCacheGVK, NewAnywhereCacheModel)
}

func NewAnywhereCacheModel(_ context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAnywhereCache{config: *config}, nil
}

var _ directbase.Model = &modelAnywhereCache{}

type modelAnywhereCache struct {
	config config.ControllerConfig
}

func (m *modelAnywhereCache) client(ctx context.Context) (*gcp.StorageControlClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewStorageControlClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AnywhereCache client: %w", err)
	}
	return gcpClient, err
}

func (m *modelAnywhereCache) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.StorageAnywhereCache{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewAnywhereCacheIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get storage GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AnywhereCacheAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

// AdapterForURL creates an Adapter for a given GCP resource URL.
func (m *modelAnywhereCache) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Expected URL format: //storage.googleapis.com/projects/_/buckets/{bucket_id}/anywhereCaches/{anywhere_cache_id}

	if !strings.HasPrefix(url, anywhereCacheURLPrefix) {
		return nil, nil // Not an AnywhereCache URL
	}

	tokens := strings.Split(strings.TrimPrefix(url, anywhereCacheURLPrefix), "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "buckets" && tokens[4] == "anywhereCaches" {
		gcpClient, err := m.client(ctx)
		if err != nil {
			return nil, err
		}

		parent, id, err := krm.ParseAnywhereCacheExternal(strings.TrimPrefix(url, anywhereCacheURLPrefix))
		if err != nil {
			return nil, fmt.Errorf("error parsing AnywhereCache external ID from URL %q: %w", url, err)
		}

		return &AnywhereCacheAdapter{
			id:        krm.GetAnywhereCacheIdentity(parent, id),
			gcpClient: gcpClient,
		}, nil
	}

	return nil, nil // Not an AnywhereCache URL or invalid format
}

type AnywhereCacheAdapter struct {
	id        *krm.AnywhereCacheIdentity
	gcpClient AnywhereCacheAPI
	desired   *krm.StorageAnywhereCache
	actual    *pb.AnywhereCache
}

var _ directbase.Adapter = &AnywhereCacheAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AnywhereCacheAdapter) Find(ctx context.Context) (bool, error) {
	/*
			The resourceID for an AnywhereCache resource is
			backend-generated. We first check if resourceID field
			exists in Adapter.

			1. If a resourceID exists, we check for existence of
			   cache with that ID. If found, we proceed with update
			   path; otherwise, we create a new one.

			2. If no resourceID is provided, we create a new cache.

		    During creation, we'll populate the externalRef in the
			status. This externalRef will be used in subsequent
			reconciliations to extract the resourceID, leading to
			follow update path next time.
	*/
	if !(a.id.HasKnownId()) {
		return false, nil
	}
	req := &pb.GetAnywhereCacheRequest{Name: a.id.String()}
	anywherecachepb, err := a.gcpClient.GetAnywhereCache(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AnywhereCache %q: %w", a.id, err)
	}

	a.actual = anywherecachepb
	return true, nil
}

// getReadyCondition is a helper to create a KRM condition for resource status.
func getReadyCondition(status v1.ConditionStatus, reason string, message string) *v1alpha1.Condition {
	return &v1alpha1.Condition{
		Type:    v1alpha1.ReadyConditionType,
		Status:  status,
		Reason:  reason,
		Message: message,
	}
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AnywhereCacheAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	desiredPb := StorageAnywhereCacheSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateAnywhereCacheRequest{
		Parent:        a.id.Parent().String(),
		AnywhereCache: desiredPb,
	}

	op, err := a.gcpClient.CreateAnywhereCache(ctx, req)
	if err != nil {
		return fmt.Errorf("Creating AnywhereCache: %w", err)
	}

	metadata, err := op.Metadata()
	if err != nil {
		return fmt.Errorf("Failed to GET anywhere cache metadata: %w", err)
	}

	a.id.SetResourceID(metadata.GetAnywhereCacheId())
	resource, err := a.gcpClient.GetAnywhereCache(ctx, &pb.GetAnywhereCacheRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("Failed to get anywhere cache metadata: %w", err)
	}

	status := &krm.StorageAnywhereCacheStatus{}
	status.ObservedState = StorageAnywhereCacheObservedState_FromProto(mapCtx, resource)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())

	if resource.GetState() == anywhereCacheStateCreating {
		return createOp.UpdateStatus(ctx, status, getReadyCondition(v1.ConditionFalse, k8s.Creating, k8s.CreatingMessage))
	}

	return createOp.UpdateStatus(ctx, status, getReadyCondition(v1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage))
}

func (a *AnywhereCacheAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	return a.UpdateCache(ctx, updateOp)
}

/*
We primarily receive two types of updates: state changes and metadata changes. To maintain simplicity,
these are kept separate, with state changes prioritized. This means that if an update involves both
the cache state and metadata changes, the cache state is updated first, followed by the metadata fields.
*/
// Update updates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *AnywhereCacheAdapter) UpdateCache(ctx context.Context, updateOp DirectBaseUpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("attempting to update AnywhereCache", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	// Get the current status, initially using `a.actual` from the Find() call.
	status, err := a.GetUpdateStatus(ctx, false)
	if err != nil {
		return fmt.Errorf("error getting current update status: %w", err)
	}

	/*
		We cannot perform any operation if the cache is in the "creating" state.
		This state means the resource is still being provisioned in GCP.
		We delay our updates until it reaches a "running" state, as performing
		updates during "creating" would likely result in an error.
	*/
	currentState := a.GetCurrentState()
	if currentState == anywhereCacheStateCreating {
		log.V(2).Info("Delaying update as AnywhereCache is in 'creating' state.", "name", a.id.String())
		return updateOp.UpdateStatus(ctx, status, getReadyCondition(v1.ConditionFalse, k8s.Creating, k8s.CreatingMessage))
	}

	// Compute the metadata changes (Note: This doesn't include state changes, which are handled separately below).
	desiredPb := StorageAnywhereCacheSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return fmt.Errorf("error mapping desired spec to proto for update: %w", mapCtx.Err())
	}
	desiredPb.Name = a.id.String() // Ensure the name is set for comparison and request.

	paths := make(sets.Set[string])
	paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return fmt.Errorf("error comparing desired and actual protobuf messages: %w", err)
	}

	/*
		Our cache can be in one of three states: "running", "disabled", or "paused".
		If a user requests a state change, we prioritize and perform it immediately,
		as such changes are generally possible regardless of any ongoing metadata updates.
	*/
	if a.IsStateChangeRequested() {
		err = a.UpdateState(ctx)
		if err != nil {
			return fmt.Errorf("error updating AnywhereCache state: %w", err)
		}
		// We need a fresh status, as updating the state would have changed some state attributes.
		status, err = a.GetUpdateStatus(ctx, true) // Fetch from GCP to get the latest state.
		if err != nil {
			return fmt.Errorf("error getting update status after state change: %w", err)
		}

		if len(paths) != 0 || a.actual.GetPendingUpdate() {
			// There are still metadata updates or a previous update is pending.
			// These updates will be handled in the next reconcile cycle.
			log.V(2).Info("State change completed, but metadata updates pending or in progress. Requeuing.", "name", a.id.String())

			// Do not wait for next reconcile, instead use RequestReque()
			updateOp.RequestRequeue()
			return updateOp.UpdateStatus(ctx, status, getReadyCondition(v1.ConditionFalse, k8s.Updating, k8s.UpdatingMessage))
		}
		// No other updates pending.
		log.V(2).Info("State change completed, no further updates pending.", "name", a.id.String())
		return updateOp.UpdateStatus(ctx, status, getReadyCondition(v1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage))
	}

	// With no state changes requested, we proceed to check for metadata changes.
	if len(paths) == 0 {
		// Nothing remains to update for metadata.
		log.V(2).Info("No metadata fields need update.", "name", a.id.String())
		return updateOp.UpdateStatus(ctx, status, getReadyCondition(v1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage))
	} else if a.actual.GetPendingUpdate() {
		// A previous metadata update is currently in progress. Attempting another update would result in failure.
		// Therefore, this update must be delayed until the previous request is fulfilled.
		log.V(2).Info("Delaying metadata update, as a previous update is already in progress.", "name", a.id.String())
		return updateOp.UpdateStatus(ctx, status, getReadyCondition(v1.ConditionFalse, k8s.Updating, k8s.UpdatingMessage))
	} else if currentState == anywhereCacheStatePaused || currentState == anywhereCacheStateDisabled {
		// This update cannot be delayed by requeueing, as it requires the user to modify the cache state.
		// Therefore, an error is thrown to inform the user about the unmodifiable state.
		return fmt.Errorf("AnywhereCache cannot be updated in '%s' or '%s' state. Please change the cache state to '%s' to apply metadata updates",
			anywhereCacheStatePaused, anywhereCacheStateDisabled, anywhereCacheStateRunning)
	}

	// At this point, cache is in "running" state, with no previous updates pending, and metadata changes are required.
	log.Info(fmt.Sprintf("Metadata fields needing update: %v", sets.List(paths)))
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}
	req := &pb.UpdateAnywhereCacheRequest{
		UpdateMask:    updateMask,
		AnywhereCache: desiredPb,
	}

	_, err = a.gcpClient.UpdateAnywhereCache(ctx, req)
	if err != nil {
		return fmt.Errorf("error while updating AnywhereCache metadata %s: %w", a.id, err)
	}

	// We need a fresh status, as the metadata would have changed post-update,
	// and the update is an LRO, so its completion needs to be observed.
	status, err = a.GetUpdateStatus(ctx, true) // Fetch from GCP to get the latest state.
	if err != nil {
		return fmt.Errorf("error getting update status after metadata update: %w", err)
	}

	if status.ObservedState.PendingUpdate != nil && *status.ObservedState.PendingUpdate {
		// Metadata updates are performed via LROs, hence the status is 'Updating' instead of 'UpToDate' immediately.
		return updateOp.UpdateStatus(ctx, status, getReadyCondition(v1.ConditionFalse, k8s.Updating, k8s.UpdatingMessage))
	}

	// dead code, as updates are not instantaneous.
	return updateOp.UpdateStatus(ctx, status, getReadyCondition(v1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage))
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AnywhereCacheAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.StorageAnywhereCache{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(StorageAnywhereCacheSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.BucketRef = &refs.StorageBucketRef{External: a.id.Parent().String()}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ID())
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.StorageAnywhereCacheGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service by disabling it when the corresponding Config Connector resource is deleted.
func (a *AnywhereCacheAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("attempting to disable (delete) AnywhereCache", "name", a.id.String())

	// If the cache is already disabled
	if a.GetCurrentState() == anywhereCacheStateDisabled {
		log.V(2).Info("AnywhereCache is already disabled", "name", a.id.String())
		return true, nil
	}

	req := &pb.DisableAnywhereCacheRequest{Name: a.id.String()}
	_, err := a.gcpClient.DisableAnywhereCache(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted/disabled).
			log.V(2).Info("AnywhereCache not found in GCP, assuming it was already deleted/disabled", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("disabling (deleting) AnywhereCache %s: %w", a.id, err)
	}
	log.V(2).Info("successfully disabled AnywhereCache", "name", a.id.String())
	return true, nil
}

// GetCurrentState returns the current state of the AnywhereCache from its actual GCP representation.
func (a *AnywhereCacheAdapter) GetCurrentState() string {
	return strings.ToLower(a.actual.GetState())
}

// GetDesiredState returns the desired state of the AnywhereCache from the KRM spec.
// Defaults to "running" if not explicitly specified in the spec.
func (a *AnywhereCacheAdapter) GetDesiredState() string {
	desiredState := a.desired.Spec.DesiredState
	if desiredState == nil {
		return anywhereCacheStateRunning // Default behavior if not specified
	}
	return strings.ToLower(*desiredState)
}

// IsStateChangeRequested checks if a state change is requested by comparing current and desired states.
func (a *AnywhereCacheAdapter) IsStateChangeRequested() bool {
	return a.GetCurrentState() != a.GetDesiredState()
}

// UpdateState changes the cache state to the desired state if different from current.
func (a *AnywhereCacheAdapter) UpdateState(ctx context.Context) error {
	currentState := a.GetCurrentState()
	desiredState := a.GetDesiredState()

	if currentState == desiredState {
		return nil // No state change needed.
	}

	if currentState == anywhereCacheStateCreating || desiredState == anywhereCacheStateCreating {
		return fmt.Errorf("cannot change state from or to '%s' state", anywhereCacheStateCreating)
	}

	var err error
	switch desiredState {
	case anywhereCacheStateRunning:
		// Possible actual states - "disabled", "paused"
		// To transition to 'running', if it's 'disabled' or 'paused', a resume is required.
		// The GCP API allows resuming from both paused and disabled states to 'running'.
		_, err = a.gcpClient.ResumeAnywhereCache(ctx, &pb.ResumeAnywhereCacheRequest{Name: a.id.String()})
	case anywhereCacheStatePaused:
		// Possible actual states - "running", "disabled"
		// To transition to 'paused', if it's 'disabled', it must first be resumed (to 'running') then paused.
		if currentState == anywhereCacheStateDisabled {
			klog.FromContext(ctx).V(2).Info(fmt.Sprintf("Resuming AnywhereCache %s from '%s' state before pausing.", a.id.String(), currentState))
			_, err = a.gcpClient.ResumeAnywhereCache(ctx, &pb.ResumeAnywhereCacheRequest{Name: a.id.String()})
			if err != nil {
				return fmt.Errorf("error resuming AnywhereCache from '%s' before pausing: %w", currentState, err)
			}
		}
		// After resuming (or if already running), pause the cache.
		_, err = a.gcpClient.PauseAnywhereCache(ctx, &pb.PauseAnywhereCacheRequest{Name: a.id.String()})
	case anywhereCacheStateDisabled:
		// Possible actual states - "running", "paused"
		// Direct transition to 'disabled' from 'running' or 'paused'.
		_, err = a.gcpClient.DisableAnywhereCache(ctx, &pb.DisableAnywhereCacheRequest{Name: a.id.String()})
	default:
		err = fmt.Errorf("invalid desired state specified: %q", desiredState)
	}

	if err != nil {
		return fmt.Errorf("failed to update AnywhereCache state to %q: %w", desiredState, err)
	}
	return nil
}

// GetUpdateStatus returns the update status from GCP metadata or KRM metadata based on the provided argument.
// If `fetchGcp` is true, it makes a fresh GCP API call to get the latest resource state.
func (a *AnywhereCacheAdapter) GetUpdateStatus(ctx context.Context, fetchGcp bool) (*krm.StorageAnywhereCacheStatus, error) {
	mapCtx := &direct.MapContext{}
	resource := a.actual // Start with the state from the last Find() call.
	var err error

	if fetchGcp {
		// Fetch the latest state from GCP if explicitly requested.
		resource, err = a.gcpClient.GetAnywhereCache(ctx, &pb.GetAnywhereCacheRequest{Name: a.id.String()})
		if err != nil {
			return nil, fmt.Errorf("error fetching AnywhereCache metadata from GCP: %w", err)
		}
	}

	status := &krm.StorageAnywhereCacheStatus{}
	status.ObservedState = StorageAnywhereCacheObservedState_FromProto(mapCtx, resource)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return status, nil
}
