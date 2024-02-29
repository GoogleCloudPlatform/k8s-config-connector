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

package registration

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// This predicate will react only to Create requests from CRDs that KCC manages.
type ManagedByKCCPredicate struct {
	predicate.Funcs
}

// Create returns true if the given resource has the KCC management label.
func (ManagedByKCCPredicate) Create(e event.CreateEvent) bool {
	return isManagedByKCC(e.Object)
}

// Update returns true if the given resource has the KCC management label.
// When CRD is changed, the controller should reload its jsonSchema from the
// newly updated CRD.
func (ManagedByKCCPredicate) Update(e event.UpdateEvent) bool {
	return isManagedByKCC(e.ObjectNew)
}

// Delete always returns false, as currently there is no support for removing controllers
// on CRD deletion.
func (ManagedByKCCPredicate) Delete(_ event.DeleteEvent) bool {
	return false
}

func isManagedByKCC(o metav1.Object) bool {
	if o.GetName() == "securesourcemanagerinstances.securesourcemanager.cnrm.cloud.google.com" {
		return true
	}
	val, ok := o.GetLabels()[crdgeneration.ManagedByKCCLabel]
	return ok && val == "true"
}
