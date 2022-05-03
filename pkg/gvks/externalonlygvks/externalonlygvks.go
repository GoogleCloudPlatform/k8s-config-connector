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

// externalonlygvks contains variables and helpers for GroupVersionKinds that
// are not supported by KCC, but are commonly referenced by KCC resources.
package externalonlygvks

import "k8s.io/apimachinery/pkg/runtime/schema"

var (
	OrganizationGVK = schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Organization",
	}
	BillingAccountGVK = schema.GroupVersionKind{
		Group:   "billing.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "BillingAccount",
	}

	all = map[schema.GroupVersionKind]bool{
		OrganizationGVK:   true,
		BillingAccountGVK: true,
	}
)

// All returns GroupVersionKinds corresponding to GCP resources not supported
// by KCC, but are commonly referenced by KCC resources.
func All() []schema.GroupVersionKind {
	gvks := make([]schema.GroupVersionKind, 0)
	for gvk := range all {
		gvks = append(gvks, gvk)
	}
	return gvks
}

func IsExternalOnlyGVK(gvk schema.GroupVersionKind) bool {
	_, ok := all[gvk]
	return ok
}
