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

package v1beta1

// +kcc:proto=google.bigtable.admin.v2.AppProfile.DataBoostIsolationReadOnly
type AppProfile_DataBoostIsolationReadOnly struct {
	// The Compute Billing Owner for this Data Boost App Profile.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.DataBoostIsolationReadOnly.compute_billing_owner
	ComputeBillingOwner *string `json:"computeBillingOwner,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AppProfile.MultiClusterRoutingUseAny
type AppProfile_MultiClusterRoutingUseAny struct {
	// The set of clusters to route to. The order is ignored; clusters will be
	//  tried in order of distance. If left empty, all clusters are eligible.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.MultiClusterRoutingUseAny.cluster_ids
	ClusterIds []string `json:"clusterIds,omitempty"`

	// Row affinity sticky routing based on the row key of the request.
	//  Requests that span multiple rows are routed non-deterministically.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.MultiClusterRoutingUseAny.row_affinity
	RowAffinity *AppProfile_MultiClusterRoutingUseAny_RowAffinity `json:"rowAffinity,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AppProfile.MultiClusterRoutingUseAny.RowAffinity
type AppProfile_MultiClusterRoutingUseAny_RowAffinity struct {
}

// +kcc:proto=google.bigtable.admin.v2.AppProfile.StandardIsolation
type AppProfile_StandardIsolation struct {
	// The priority of requests sent using this app profile.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.StandardIsolation.priority
	Priority *string `json:"priority,omitempty"`
}
