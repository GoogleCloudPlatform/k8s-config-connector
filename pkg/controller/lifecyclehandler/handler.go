// Copyright 2022 Google LLC
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

package lifecyclehandler

import (
	"context"
	"fmt"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/lease/leaser"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// The LifecycleHandler contains common methods to handle the lifecycle of the reconciliation
type LifecycleHandler struct {
	client.Client
	Recorder   record.EventRecorder
	fieldOwner string
}

func NewLifecycleHandler(c client.Client, r record.EventRecorder) LifecycleHandler {
	return NewLifecycleHandlerWithFieldOwner(c, r, k8s.ControllerManagedFieldManager)
}

func NewLifecycleHandlerWithFieldOwner(c client.Client, r record.EventRecorder, fieldOwner string) LifecycleHandler {
	return LifecycleHandler{
		Client:     c,
		Recorder:   r,
		fieldOwner: fieldOwner,
	}
}

func (r *LifecycleHandler) updateStatus(ctx context.Context, resource *k8s.Resource) error {
	u, err := resource.MarshalAsUnstructured()
	if err != nil {
		return err
	}
	if err := r.Client.Status().Update(ctx, u, client.FieldOwner(r.fieldOwner)); err != nil {
		if apierrors.IsConflict(err) {
			return fmt.Errorf("couldn't update the API server due to conflict. Re-enqueue the request for another reconciliation attempt: %w", err)
		}
		return fmt.Errorf("error with status update call to API server: %w", err)
	}
	// rejections by some validating webhooks won't be returned as an error; instead, they will be
	// objects of kind "Status" with a "Failure" status.
	if isFailureStatus(u) {
		return fmt.Errorf("error with status update call to API server: %v", u.Object["message"])
	}
	// sync the resource up with the updated metadata.
	if err := util.Marshal(u, resource); err != nil {
		return err
	}
	return resourceoverrides.Handler.PostUpdateStatusTransform(resource)
}

// WARNING: This function should NOT be exported and invoked directly outside the package.
// Controllers are supposed to call exported functions to handle lifecycle transitions.
func (r *LifecycleHandler) updateAPIServer(ctx context.Context, resource *k8s.Resource) error {
	// Preserve the intended status, as the client.Update call will ignore the given status
	// and return the stale existing status.
	status := deepcopy.MapStringInterface(resource.Status)
	// Get the current generation as the observed generation because the following client.Update
	// might increase the generation. We want the next reconciliation to handle the new generation.
	observedGeneration := resource.GetGeneration()
	u, err := resource.MarshalAsUnstructured()
	if err != nil {
		return err
	}
	removeSystemLabels(u)
	if err := r.Client.Update(ctx, u, client.FieldOwner(r.fieldOwner)); err != nil {
		if apierrors.IsConflict(err) {
			return fmt.Errorf("couldn't update the API server due to conflict. Re-enqueue the request for another reconciliation attempt: %w", err)
		}
		return fmt.Errorf("error with update call to API server: %w", err)
	}
	// rejections by validating webhooks won't be returned as an error; instead, they will be
	// objects of kind "Status" with a "Failure" status.
	if isFailureStatus(u) {
		return fmt.Errorf("error with update call to API server: %v", u.Object["message"])
	}
	// sync the resource up with the updated metadata
	if err := util.Marshal(u, resource); err != nil {
		return fmt.Errorf("error syncing updated resource metadata: %w", err)
	}
	if !u.GetDeletionTimestamp().IsZero() && len(u.GetFinalizers()) == 0 {
		// This resource is set for garbage collection and any status updates would be racey.
		// Status updates for successful deletions must be handled independently.
		return nil
	}
	resource.Status = status
	setObservedGeneration(resource, observedGeneration)
	return r.updateStatus(ctx, resource)
}

func isFailureStatus(u *unstructured.Unstructured) bool {
	return u.GetKind() == "Status" && u.Object["status"] == metav1.StatusFailure
}

// The system sets various labels on the resource that are not user facing and should not be saved in the API server
// this function removes any that may be present
func removeSystemLabels(u *unstructured.Unstructured) {
	labels := u.GetLabels()
	if labels == nil {
		return
	}
	keys := leaser.GetLabelKeys()
	keys = append(keys, label.CnrmManagedKey)
	for _, k := range keys {
		delete(labels, k)
	}
	// GetLabels(...) returns a new copy of the labels map so we must overwrite that value with our local value
	u.SetLabels(labels)
}

// CausedByUnreadyOrNonexistentResourceRefs checks to see if the input error
// is related to an unready or non-existent resource reference. Note that
// KeyInSecretNotFoundError is not included in this list.
func CausedByUnreadyOrNonexistentResourceRefs(err error) (refGVK schema.GroupVersionKind, refNN types.NamespacedName, ok bool) {
	if unwrappedErr, ok := k8s.AsReferenceNotReadyError(err); ok {
		return unwrappedErr.RefResourceGVK, unwrappedErr.RefResource, true
	}
	if unwrappedErr, ok := k8s.AsReferenceNotFoundError(err); ok {
		return unwrappedErr.RefResourceGVK, unwrappedErr.RefResource, true
	}
	if unwrappedErr, ok := k8s.AsTransitiveDependencyNotFoundError(err); ok {
		return unwrappedErr.ResourceGVK, unwrappedErr.Resource, true
	}
	if unwrappedErr, ok := k8s.AsTransitiveDependencyNotReadyError(err); ok {
		return unwrappedErr.ResourceGVK, unwrappedErr.Resource, true
	}
	if unwrappedErr, ok := k8s.AsSecretNotFoundError(err); ok {
		return schema.GroupVersionKind{Version: "v1", Kind: "Secret"}, unwrappedErr.Secret, true
	}
	return schema.GroupVersionKind{}, types.NamespacedName{}, false
}

func CausedByUnresolvableDeps(err error) (unwrappedErr error, ok bool) { //nolint:revive
	if unwrappedErr, ok := k8s.AsReferenceNotReadyError(err); ok {
		return unwrappedErr, true
	}
	if unwrappedErr, ok := k8s.AsReferenceNotFoundError(err); ok {
		return unwrappedErr, true
	}
	if unwrappedErr, ok := k8s.AsSecretNotFoundError(err); ok {
		return unwrappedErr, true
	}
	if unwrappedErr, ok := k8s.AsKeyInSecretNotFoundError(err); ok {
		return unwrappedErr, true
	}
	if unwrappedErr, ok := k8s.AsTransitiveDependencyNotFoundError(err); ok {
		return unwrappedErr, true
	}
	if unwrappedErr, ok := k8s.AsTransitiveDependencyNotReadyError(err); ok {
		return unwrappedErr, true
	}
	return nil, false
}

func reasonForUnresolvableDeps(err error) (string, error) {
	switch {
	case k8s.IsReferenceNotReadyError(err) || k8s.IsTransitiveDependencyNotReadyError(err):
		return k8s.DependencyNotReady, nil
	case k8s.IsReferenceNotFoundError(err) || k8s.IsSecretNotFoundError(err) || k8s.IsTransitiveDependencyNotFoundError(err):
		return k8s.DependencyNotFound, nil
	case k8s.IsKeyInSecretNotFoundError(err):
		return k8s.DependencyInvalid, nil
	default:
		return "", fmt.Errorf("unrecognized error caused by unresolvable dependencies: %w", err)
	}
}

func (r *LifecycleHandler) EnsureFinalizers(ctx context.Context, original, resource *k8s.Resource, finalizers ...string) error {
	if !k8s.EnsureFinalizers(resource, finalizers...) {
		uo, err := original.MarshalAsUnstructured()
		if err != nil {
			return err
		}
		originalCopy, err := k8s.NewResource(uo)
		if err != nil {
			return err
		}
		if !k8s.EnsureFinalizers(originalCopy, finalizers...) {
			if err := r.updateAPIServer(ctx, originalCopy); err != nil {
				return err
			}
			// Sync the resource up with the updated metadata except for the
			// defaulted / pre-processed annotations.
			originalCopy.ObjectMeta.Annotations = resource.ObjectMeta.Annotations
			resource.ObjectMeta = originalCopy.ObjectMeta
		}
	}
	return nil
}

func (r *LifecycleHandler) HandleUpToDate(ctx context.Context, resource *k8s.Resource) error {
	setCondition(resource, corev1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage)
	if err := r.updateAPIServer(ctx, resource); err != nil {
		return err
	}

	r.recordEvent(ctx, resource, corev1.EventTypeNormal, k8s.UpToDate, k8s.UpToDateMessage)
	return nil
}

func (r *LifecycleHandler) HandleUnresolvableDeps(ctx context.Context, resource *k8s.Resource, originErr error) error {
	reason, err := reasonForUnresolvableDeps(originErr)
	if err != nil {
		return r.HandleUpdateFailed(ctx, resource, err)
	}
	msg := originErr.Error()
	// Only update the API server if there's new information
	if !k8s.ReadyConditionMatches(resource, corev1.ConditionFalse, reason, msg) {
		setCondition(resource, corev1.ConditionFalse, reason, msg)
		setObservedGeneration(resource, resource.GetGeneration())
		if err := r.updateStatus(ctx, resource); err != nil {
			return err
		}
	}

	r.recordEvent(ctx, resource, corev1.EventTypeWarning, reason, msg)
	return nil
}

func (r *LifecycleHandler) HandleObtainLeaseFailed(ctx context.Context, resource *k8s.Resource, err error) error {
	msg := err.Error()
	// Only update the API server if there's new information
	if !k8s.ReadyConditionMatches(resource, corev1.ConditionFalse, k8s.ManagementConflict, msg) {
		setCondition(resource, corev1.ConditionFalse, k8s.ManagementConflict, msg)
		setObservedGeneration(resource, resource.GetGeneration())
		if err := r.updateStatus(ctx, resource); err != nil {
			return err
		}
	}

	r.recordEvent(ctx, resource, corev1.EventTypeWarning, k8s.ManagementConflict, msg)
	return err
}

func (r *LifecycleHandler) HandlePreActuationTransformFailed(ctx context.Context, resource *k8s.Resource, err error) error {
	msg := err.Error()
	// Only update the API server if there's new information
	if !k8s.ReadyConditionMatches(resource, corev1.ConditionFalse, k8s.PreActuationTransformFailed, msg) {
		setCondition(resource, corev1.ConditionFalse, k8s.PreActuationTransformFailed, msg)
		setObservedGeneration(resource, resource.GetGeneration())
		if err := r.updateStatus(ctx, resource); err != nil {
			return err
		}
	}

	r.recordEvent(ctx, resource, corev1.EventTypeWarning, k8s.PreActuationTransformFailed, msg)
	return err
}

func (r *LifecycleHandler) HandlePostActuationTransformFailed(ctx context.Context, resource *k8s.Resource, err error) error {
	msg := err.Error()
	// Only update the API server if there's new information
	if !k8s.ReadyConditionMatches(resource, corev1.ConditionFalse, k8s.PostActuationTransformFailed, msg) {
		setCondition(resource, corev1.ConditionFalse, k8s.PostActuationTransformFailed, msg)
		setObservedGeneration(resource, resource.GetGeneration())
		if err := r.updateStatus(ctx, resource); err != nil {
			return err
		}
	}

	r.recordEvent(ctx, resource, corev1.EventTypeWarning, k8s.PostActuationTransformFailed, msg)
	return err
}

func (r *LifecycleHandler) HandleUpdating(ctx context.Context, resource *k8s.Resource) error {
	setCondition(resource, corev1.ConditionFalse, k8s.Updating, k8s.UpdatingMessage)
	setObservedGeneration(resource, resource.GetGeneration())
	if err := r.updateStatus(ctx, resource); err != nil {
		return err
	}

	r.recordEvent(ctx, resource, corev1.EventTypeNormal, k8s.Updating, k8s.UpdatingMessage)
	return nil
}

func (r *LifecycleHandler) HandleUpdateFailed(ctx context.Context, resource *k8s.Resource, err error) error {
	structuredreporting.ReportError(ctx, err, resource)
	msg := fmt.Errorf("Update call failed: %w", err).Error()
	setCondition(resource, corev1.ConditionFalse, k8s.UpdateFailed, msg)
	setObservedGeneration(resource, resource.GetGeneration())
	if err := r.updateStatus(ctx, resource); err != nil {
		return err
	}

	r.recordEvent(ctx, resource, corev1.EventTypeWarning, k8s.UpdateFailed, msg)
	return fmt.Errorf("Update call failed: %w", err)
}

func (r *LifecycleHandler) HandleDeleting(ctx context.Context, resource *k8s.Resource) error {
	setCondition(resource, corev1.ConditionFalse, k8s.Deleting, k8s.DeletingMessage)
	setObservedGeneration(resource, resource.GetGeneration())
	if err := r.updateStatus(ctx, resource); err != nil {
		return err
	}

	r.recordEvent(ctx, resource, corev1.EventTypeNormal, k8s.Deleting, k8s.DeletingMessage)
	return nil
}

func (r *LifecycleHandler) HandleDeleted(ctx context.Context, resource *k8s.Resource) error {
	setCondition(resource, corev1.ConditionFalse, k8s.Deleted, k8s.DeletedMessage)
	setObservedGeneration(resource, resource.GetGeneration())
	// Do an explicit status update first to prevent a race between the status update and the API
	// server pruning the resource if there are no more finalizers present.
	if err := r.updateStatus(ctx, resource); err != nil {
		return fmt.Errorf("error updating status: %w", err)
	}

	r.recordEvent(ctx, resource, corev1.EventTypeNormal, k8s.Deleted, k8s.DeletedMessage)

	k8s.RemoveFinalizer(resource, k8s.ControllerFinalizerName)
	return r.updateAPIServer(ctx, resource)
}

func (r *LifecycleHandler) HandleDeleteFailed(ctx context.Context, resource *k8s.Resource, err error) error {
	msg := fmt.Sprintf(k8s.DeleteFailedMessageTmpl, err)
	setCondition(resource, corev1.ConditionFalse, k8s.DeleteFailed, msg)
	setObservedGeneration(resource, resource.GetGeneration())
	if err := r.updateStatus(ctx, resource); err != nil {
		return err
	}

	r.recordEvent(ctx, resource, corev1.EventTypeWarning, k8s.DeleteFailed, msg)
	return fmt.Errorf("Delete call failed: %w", err)
}

func (r *LifecycleHandler) HandleUnmanaged(ctx context.Context, resource *k8s.Resource) error {
	msg := fmt.Sprintf(k8s.UnmanagedMessageTmpl, resource.GetNamespace())
	setCondition(resource, corev1.ConditionFalse, k8s.Unmanaged, msg)
	setObservedGeneration(resource, resource.GetGeneration())
	if err := r.updateStatus(ctx, resource); err != nil {
		return err
	}

	r.recordEvent(ctx, resource, corev1.EventTypeWarning, k8s.Unmanaged, msg)
	return nil
}

func setCondition(resource *k8s.Resource, status corev1.ConditionStatus, reason, msg string) {
	if resource.Status == nil {
		resource.Status = make(map[string]interface{})
	}
	newReadyCondition := k8s.NewCustomReadyCondition(status, reason, msg)
	// We should only update the ready condition's last transition time if there was a transition
	// since its last state. The function sets it to time.Now(), so let's replace it if there was
	// no transition.
	if currentReadyCondition, found := k8s.GetReadyCondition(resource); found {
		if currentReadyCondition.Status == status {
			newReadyCondition.LastTransitionTime = currentReadyCondition.LastTransitionTime
		}
	}
	resource.Status["conditions"] = []k8sv1alpha1.Condition{newReadyCondition}
}

func setObservedGeneration(resource *k8s.Resource, observedGeneration int64) {
	if resource.Status == nil {
		resource.Status = make(map[string]interface{})
	}
	resource.Status["observedGeneration"] = observedGeneration
}

func (r *LifecycleHandler) recordEvent(ctx context.Context, resource *k8s.Resource, eventtype, reason, message string) {
	u, err := resource.MarshalAsUnstructured()
	if err != nil {
		log.FromContext(ctx).Error(err, "error recording event for resource", "resource", resource.GetName(), "namespace", resource.GetNamespace(), "reason", reason, "message", message, "event_type", eventtype)
		return
	}
	r.Recorder.Event(u, eventtype, reason, message)
}

func IsOrphaned(resource *k8s.Resource, parentReferenceConfigs []corekccv1alpha1.TypeConfig, kubeClient client.Client) (orphaned bool, parent *k8s.Resource, err error) {
	if len(parentReferenceConfigs) == 0 {
		return false, nil, nil
	}
	for _, refConfig := range parentReferenceConfigs {
		resourceRefRaw, ok := resource.Spec[refConfig.Key]
		if !ok {
			// This parent type isn't present. Check if another parent type is.
			continue
		}
		resourceRef := &corekccv1alpha1.ResourceReference{}
		if err := util.Marshal(resourceRefRaw, resourceRef); err != nil {
			return false, nil, fmt.Errorf("'spec.%v' is an unrecognized format", refConfig.Key)
		}
		if resourceRef.External != "" {
			return false, nil, nil
		}
		parent, err := k8s.GetReferencedResource(resourceRef, refConfig.GVK, resource.GetNamespace(), kubeClient)
		if err != nil {
			if k8s.IsReferenceNotFoundError(err) {
				return true, nil, nil
			}
			return false, nil, fmt.Errorf("error getting parent reference 'spec.%v': %w", refConfig.Key, err)
		}
		return false, parent, nil
	}
	return false, nil, fmt.Errorf("no parent reference found in resource")
}
