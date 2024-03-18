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

package unmanageddetector

import "sigs.k8s.io/controller-runtime/pkg/event"

// Predicate is the predicate meant to be used by the
// unmanaged-detector controller. It sets up unmanaged-detector such that it
// handles Create events only.
type Predicate struct{}

func (Predicate) Create(_ event.CreateEvent) bool {
	return true
}

// There is no scenario where unmanaged-detector would need to handle Update
// events.
//
// Firstly, an Update event implies there had been a Create event already,
// which would have been handled by unmanaged-detector.
//
// Secondly, once unmanaged-detector successfully reconciles a resource once, it
// never has to reconcile the resource again. This is because if the resource
// is marked Unmanaged, then unmanaged-detector's job is done: the onus is now
// on the user to create a ConfigConnectorContext. On the other hand, if the
// resource is determined to be managed, then unmanaged-detector's job is also
// done: it is not possible for a resource to later become unmanaged since
// ConfigConnectorContext deletions are blocked if there are resources in its
// namespace.
//
// Even if the resource had been created before unmanaged-detector came online
// (e.g. if KCC is changed from cluster-mode to namespaced-mode), this is no
// problem as well since controllers reconcile all resources they are
// configured to watch when they first come up.
func (Predicate) Update(_ event.UpdateEvent) bool {
	return false
}

// KCC controllers in general do not handle Delete events since resources
// deleted directly on the API server should not be reconciled. Instead,
// user-requested deletions are handled via the updated DeletionTimestamp. That
// is, user-requested deletions are handled via Update events, and as explained
// above, there is no scenario where unmanaged-detector would need to handle
// Update events.
func (Predicate) Delete(_ event.DeleteEvent) bool {
	return false
}

func (Predicate) Generic(_ event.GenericEvent) bool {
	return false
}
