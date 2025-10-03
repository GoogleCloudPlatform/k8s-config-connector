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

// AnnotationKeyAlphaReconciler allows customers to opt-in to using the direct reconciler.
const AnnotationKeyAlphaReconciler = "alpha.cnrm.cloud.google.com/reconciler"

// AnnotationUnmanaged allows customer to specify the fields that are not managed by the controller.
const AnnotationUnmanaged = "cnrm.cloud.google.com/unmanaged"

// OptInToDirectReconciliation allows users to opt in to direct reconciliation
// by specifying an AnnotationKeyAlphaReconciler annotation.
type OptInToDirectReconciliation struct {
}
