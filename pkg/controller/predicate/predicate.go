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

package predicate

import (
	"context"
	"reflect"
	"strconv"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// This predicate will react to changes only when there is something relevant to
// send to the underlying API.
type UnderlyingResourceOutOfSyncPredicate struct {
	predicate.Funcs
}

// Update implements default UpdateEvent filter for validating changes that require
// updates to the underlying API.
func (UnderlyingResourceOutOfSyncPredicate) Update(e event.UpdateEvent) bool {
	ctx := context.Background()
	log := klog.FromContext(ctx)

	// Kubernetes deletions manifest as API updates with the deletion
	// timestamp set when it previously was not.
	if e.ObjectOld.GetDeletionTimestamp().IsZero() &&
		!e.ObjectNew.GetDeletionTimestamp().IsZero() {
		return true
	}

	// The deletion defender finalizer being removed signifies the controller
	// may now proceed with finalizing deletion on GCP.
	if hasDeletionDefenderFinalizerBeenRemoved(e.ObjectOld, e.ObjectNew) {
		return true
	}

	// Container annotation changes should trigger reconciliations as some
	// resources rely on container annotation changes for cross-container
	// migrations on GCP (e.g. moving Projects/Folders to different parent
	// Folders/Organizations)
	if !areContainerAnnotationsEqual(e.ObjectOld.GetAnnotations(), e.ObjectNew.GetAnnotations()) {
		return true
	}

	// Labels updates should be propagated to the underlying API
	if !reflect.DeepEqual(e.ObjectOld.GetLabels(), e.ObjectNew.GetLabels()) {
		return true
	}

	// Recognize an internal annotation which can be used to force reconciles
	if e.ObjectOld.GetAnnotations()[k8s.InternalForceReconcileAnnotation] != e.ObjectNew.GetAnnotations()[k8s.InternalForceReconcileAnnotation] {
		return true
	}

	// Changes to the reconcile interval annotation should trigger a reconcile
	if oldValue, newValue := e.ObjectOld.GetAnnotations()[k8s.ReconcileIntervalInSecondsAnnotation], e.ObjectNew.GetAnnotations()[k8s.ReconcileIntervalInSecondsAnnotation]; oldValue != newValue {
		newValueInt, err := strconv.ParseInt(newValue, 10, 32)
		if err == nil && newValueInt == 0 {
			// A special case: if the new value is 0, we do not want to trigger a reconcile.
			// This is because setting the annotation to 0 disables drift correction,
			// and we do not want to immediately reconcile after the user sets this annotation:
			// they presumably want to stop reconciliations.
			log.V(2).Info("Reconcile interval annotation changed to 0, will not trigger reconcile")
		} else {
			return true
		}
	}

	// The object's generation will increment when the spec is updated, so a different
	// generation implies potential work to be done on the underlying API.
	if e.ObjectNew.GetGeneration() != e.ObjectOld.GetGeneration() {
		return true
	}

	return false
}

// Delete always returns false, as resources deleted directly from the
// API server should not be reconciled. We process user-requested deletions
// via the updated DeletionTimestamp.
func (UnderlyingResourceOutOfSyncPredicate) Delete(_ event.DeleteEvent) bool {
	return false
}

func hasDeletionDefenderFinalizerBeenRemoved(ObjectOld, ObjectNew metav1.Object) bool {
	return slice.StringSliceContains(ObjectOld.GetFinalizers(), k8s.DeletionDefenderFinalizerName) &&
		!slice.StringSliceContains(ObjectNew.GetFinalizers(), k8s.DeletionDefenderFinalizerName)
}

func areContainerAnnotationsEqual(a, b map[string]string) bool {
	for _, x := range k8s.ContainerAnnotations {
		if a[x] != b[x] {
			return false
		}
	}
	return true
}
