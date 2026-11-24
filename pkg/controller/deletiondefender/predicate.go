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

package deletiondefender

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

// Predicate prevents live-resource resyncs from delaying deletion finalization.
// The watch provides all metadata needed to filter before an uncached API read.
type Predicate struct{}

func (Predicate) Create(e event.CreateEvent) bool {
	// Initial lists must resume deletions that started before this process.
	return needsDeletionDefenseFinalization(e.Object)
}

func (Predicate) Update(e event.UpdateEvent) bool {
	return needsDeletionDefenseFinalization(e.ObjectNew)
}

func (Predicate) Delete(_ event.DeleteEvent) bool {
	return false
}

func (Predicate) Generic(e event.GenericEvent) bool {
	return needsDeletionDefenseFinalization(e.Object)
}

func needsDeletionDefenseFinalization(obj client.Object) bool {
	return obj != nil &&
		!obj.GetDeletionTimestamp().IsZero() &&
		k8s.HasFinalizer(obj, k8s.DeletionDefenderFinalizerName)
}
