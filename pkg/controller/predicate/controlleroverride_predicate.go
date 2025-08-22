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

package predicate

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/kccstate"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	k8scontrollertype "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	reconcilerAnnotation = "alpha.cnrm.cloud.google.com/reconciler"
	directReconciler     = "direct"
)

// ControllerOverridePredicate implements a predicate that filters events based on
// the ControllerOverrides field in ConfigConnectorContext.
type ControllerOverridePredicate struct {
	client client.Client
	gvk    schema.GroupVersionKind
}

// NewControllerOverridePredicate creates a new ControllerOverridePredicate.
func NewControllerOverridePredicate(c client.Client, gvk schema.GroupVersionKind) *ControllerOverridePredicate {
	return &ControllerOverridePredicate{
		client: c,
		gvk:    gvk,
	}
}

// Create implements Predicate.
func (p *ControllerOverridePredicate) Create(e event.CreateEvent) bool {
	return p.shouldReconcile(e.Object)
}

// Update implements Predicate.
func (p *ControllerOverridePredicate) Update(e event.UpdateEvent) bool {
	return p.shouldReconcile(e.ObjectNew)
}

// Delete implements Predicate.
func (p *ControllerOverridePredicate) Delete(e event.DeleteEvent) bool {
	// We don't want to filter delete events based on controller overrides.
	// Deletion should always be handled.
	return true
}

// Generic implements Predicate.
func (p *ControllerOverridePredicate) Generic(e event.GenericEvent) bool {
	return p.shouldReconcile(e.Object)
}

func (p *ControllerOverridePredicate) shouldReconcile(obj client.Object) bool {
	logger := log.FromContext(context.Background()) // Use background context for predicate

	// Check for the "alpha.cnrm.cloud.google.com/reconciler: direct" annotation
	annotations := obj.GetAnnotations()
	if val, ok := annotations[reconcilerAnnotation]; ok && val == directReconciler {
		return true
	}

	// Fetch the ConfigConnectorContext for the object's namespace.
	// The name of the ConfigConnectorContext is fixed to ConfigConnectorContextAllowedName.
	nn := types.NamespacedName{Namespace: obj.GetNamespace(), Name: operatorv1beta1.ConfigConnectorContextAllowedName}
	_, ccc, err := kccstate.FetchLiveKCCState(context.Background(), p.client, nn)
	if err != nil {
		logger.Error(err, "failed to fetch ConfigConnectorContext for predicate", "namespace", obj.GetNamespace())
		return false // If we can't fetch CCC, better not to reconcile.
	}

	// Check if an override exists for the current GVK.
	if overrideType, ok := ccc.Spec.Experiments.ControllerOverrides[p.gvk.GroupKind().String()]; ok {
		// If an override exists, this controller should only reconcile if the override
		// explicitly specifies the direct controller type.
		return overrideType == k8scontrollertype.ReconcilerTypeDirect
	}

	// No override specified for this GVK in CCC.
	// In this case, this predicate should allow reconciliation.
	// The actual default controller selection logic will be handled elsewhere.
	return true
}
