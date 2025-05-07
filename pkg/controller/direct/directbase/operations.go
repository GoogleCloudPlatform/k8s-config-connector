// Copyright 2024 Google LLC
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

package directbase

import (
	"context"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// operationBase defines common functionality for multiple operation types.
type operationBase struct {
	client client.Client

	object *unstructured.Unstructured

	// HasSetReadyCondition tracks whether the controller explicitly set the ready condition
	HasSetReadyCondition bool

	// RequeueRequested tracks whether we need a re-reconciliation
	RequeueRequested bool
}

// Operation defines some functionality supported by all operation types.
type Operation interface {
	// Writes the status and ready condition to the object's status subresource.
	// We split out the readyCondition so that we will not write it from the reconcile loop if we wrote it here.
	UpdateStatus(ctx context.Context, typedStatus any, readyCondition *v1alpha1.Condition) error

	// RequestRequeue requests a requeue of the operation, by returning Requeue = true from the reconcile loop.
	RequestRequeue()
}

// GetUnstructured returns the object being reconciled, in unstructured format.
func (o *operationBase) GetUnstructured() *unstructured.Unstructured {
	return o.object
}

var _ Operation = &UpdateOperation{}

type UpdateOperation struct {
	operationBase

	lifecycleHandler lifecyclehandler.LifecycleHandler
	oldObject        *unstructured.Unstructured
}

func NewUpdateOperation(lifecycleHandler lifecyclehandler.LifecycleHandler, client client.Client, object *unstructured.Unstructured, oldObject *unstructured.Unstructured) *UpdateOperation {
	op := &UpdateOperation{}
	op.lifecycleHandler = lifecycleHandler
	op.client = client
	op.object = object
	op.oldObject = oldObject
	return op
}

func (o *UpdateOperation) HasOldObject() bool {
	return o.oldObject != nil
}

func (o *UpdateOperation) GetOldUnstructured() *unstructured.Unstructured {
	return o.oldObject
}

func (o *UpdateOperation) RecordUpdatingEvent() {
	r := o.lifecycleHandler.Recorder
	r.Event(o.object, corev1.EventTypeNormal, k8s.Updating, k8s.UpdatingMessage)
}

var _ Operation = &CreateOperation{}

type CreateOperation struct {
	operationBase
	lifecycleHandler lifecyclehandler.LifecycleHandler
}

func NewCreateOperation(lifecycleHandler lifecyclehandler.LifecycleHandler, client client.Client, object *unstructured.Unstructured) *CreateOperation {
	op := &CreateOperation{}
	op.lifecycleHandler = lifecycleHandler
	op.client = client
	op.object = object
	return op
}

func (o *CreateOperation) RecordUpdatingEvent() {
	r := o.lifecycleHandler.Recorder
	r.Event(o.object, corev1.EventTypeNormal, k8s.Updating, k8s.UpdatingMessage)
}

type DeleteOperation struct {
	object *unstructured.Unstructured
}

func NewDeleteOperation(client client.Client, object *unstructured.Unstructured) *DeleteOperation {
	return &DeleteOperation{
		object: object,
	}
}

func (o *DeleteOperation) GetUnstructured() *unstructured.Unstructured {
	return o.object
}

// UpdateStatus writes the status and ready condition to the object's status subresource.
// We split out the readyCondition so that we will not write it from the reconcile loop if we wrote it here.
func (o *operationBase) UpdateStatus(ctx context.Context, typedStatus any, readyCondition *v1alpha1.Condition) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(o.object.Object, "status")
	if old != nil {
		if status["conditions"] == nil {
			status["conditions"] = old["conditions"]
		}
		// status["observedGeneration"] = old["observedGeneration"]
		if status["externalRef"] == nil {
			status["externalRef"] = old["externalRef"]
		}
	}

	status["observedGeneration"] = o.object.GetGeneration()

	if readyCondition != nil {
		o.HasSetReadyCondition = true
		var statusWithConditions statusWithConditions
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(status, &statusWithConditions); err != nil {
			return fmt.Errorf("error converting status.conditions from structured: %w", err)
		}

		// Must be non-nil (for unclear reasons!)
		if statusWithConditions.Conditions == nil {
			statusWithConditions.Conditions = []v1alpha1.Condition{}
		}
		SetStatusCondition(&statusWithConditions.Conditions, *readyCondition)

		unstructuredStatusWithConditions, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&statusWithConditions)
		if err != nil {
			return fmt.Errorf("error converting status.conditions to unstructured: %w", err)
		}

		status["conditions"] = unstructuredStatusWithConditions["conditions"]
	}

	u := o.object
	u.Object["status"] = status

	if err := o.client.Status().Update(ctx, u); err != nil {
		return fmt.Errorf("updating object status: %w", err)
	}

	return nil
}

// RequestRequeue requests a requeue of the operation, by returning Requeue = true from the reconcile loop.
func (o *operationBase) RequestRequeue() {
	o.RequeueRequested = true
}

type statusWithConditions struct {
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
}

// SetStatusCondition is lifted from metav1.SetStatusCondition, but adapted to v1alpha1.Condition

// SetStatusCondition sets the corresponding condition in conditions to newCondition.
// conditions must be non-nil.
//  1. if the condition of the specified type already exists (all fields of the existing condition are updated to
//     newCondition, LastTransitionTime is set to now if the new status differs from the old status)
//  2. if a condition of the specified type does not exist (LastTransitionTime is set to now() if unset, and newCondition is appended)
func SetStatusCondition(conditions *[]v1alpha1.Condition, newCondition v1alpha1.Condition) {
	if conditions == nil {
		return
	}
	existingCondition := FindStatusCondition(*conditions, newCondition.Type)
	if existingCondition == nil {
		if newCondition.LastTransitionTime == "" {
			newCondition.LastTransitionTime = metav1.Now().Format(time.RFC3339)
		}
		*conditions = append(*conditions, newCondition)
		return
	}

	if existingCondition.Status != newCondition.Status {
		existingCondition.Status = newCondition.Status
		if newCondition.LastTransitionTime != "" {
			existingCondition.LastTransitionTime = newCondition.LastTransitionTime
		} else {
			existingCondition.LastTransitionTime = metav1.Now().Format(time.RFC3339)
		}
	}

	existingCondition.Reason = newCondition.Reason
	existingCondition.Message = newCondition.Message
	// TODO: Do we want ObservedGeneration in our conditions?
	// existingCondition.ObservedGeneration = newCondition.ObservedGeneration
}

// FindStatusCondition finds the conditionType in conditions.
func FindStatusCondition(conditions []v1alpha1.Condition, conditionType string) *v1alpha1.Condition {
	for i := range conditions {
		if conditions[i].Type == conditionType {
			return &conditions[i]
		}
	}

	return nil
}
