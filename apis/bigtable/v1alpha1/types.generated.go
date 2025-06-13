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

// +generated:types
// krm.group: bigtable.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.bigtable.admin.v2
// resource: BigtableLogicalView:LogicalView

package v1alpha1


// +kcc:proto=google.bigtable.admin.v2.LogicalView
type LogicalView struct {
	// Identifier. The unique name of the logical view.
	//  Format:
	//  `projects/{project}/instances/{instance}/logicalViews/{logical_view}`
	// +kcc:proto:field=google.bigtable.admin.v2.LogicalView.name
	Name *string `json:"name,omitempty"`

	// Required. The logical view's select query.
	// +kcc:proto:field=google.bigtable.admin.v2.LogicalView.query
	Query *string `json:"query,omitempty"`

	// Optional. The etag for this logical view.
	//  This may be sent on update requests to ensure that the client has an
	//  up-to-date value before proceeding. The server returns an ABORTED error on
	//  a mismatched etag.
	// +kcc:proto:field=google.bigtable.admin.v2.LogicalView.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Set to true to make the LogicalView protected against deletion.
	// +kcc:proto:field=google.bigtable.admin.v2.LogicalView.deletion_protection
	DeletionProtection *bool `json:"deletionProtection,omitempty"`
}
