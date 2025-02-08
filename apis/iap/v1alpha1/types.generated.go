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

package v1alpha1


// +kcc:proto=google.cloud.iap.v1.TunnelDestGroup
type TunnelDestGroup struct {
	// Required. Immutable. Identifier for the TunnelDestGroup. Must be unique
	//  within the project and contain only lower case letters (a-z) and dashes
	//  (-).
	// +kcc:proto:field=google.cloud.iap.v1.TunnelDestGroup.name
	Name *string `json:"name,omitempty"`

	// Unordered list. List of CIDRs that this group applies to.
	// +kcc:proto:field=google.cloud.iap.v1.TunnelDestGroup.cidrs
	Cidrs []string `json:"cidrs,omitempty"`

	// Unordered list. List of FQDNs that this group applies to.
	// +kcc:proto:field=google.cloud.iap.v1.TunnelDestGroup.fqdns
	Fqdns []string `json:"fqdns,omitempty"`
}
