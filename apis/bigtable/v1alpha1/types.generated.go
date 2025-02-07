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


// +kcc:proto=google.bigtable.admin.v2.AuthorizedView
type AuthorizedView struct {
	// Identifier. The name of this AuthorizedView.
	//  Values are of the form
	//  `projects/{project}/instances/{instance}/tables/{table}/authorizedViews/{authorized_view}`
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.name
	Name *string `json:"name,omitempty"`

	// An AuthorizedView permitting access to an explicit subset of a Table.
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.subset_view
	SubsetView *AuthorizedView_SubsetView `json:"subsetView,omitempty"`

	// The etag for this AuthorizedView.
	//  If this is provided on update, it must match the server's etag. The server
	//  returns ABORTED error on a mismatched etag.
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.etag
	Etag *string `json:"etag,omitempty"`

	// Set to true to make the AuthorizedView protected against deletion.
	//  The parent Table and containing Instance cannot be deleted if an
	//  AuthorizedView has this bit set.
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.deletion_protection
	DeletionProtection *bool `json:"deletionProtection,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AuthorizedView.FamilySubsets
type AuthorizedView_FamilySubsets struct {
	// Individual exact column qualifiers to be included in the AuthorizedView.
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.FamilySubsets.qualifiers
	Qualifiers [][]byte `json:"qualifiers,omitempty"`

	// Prefixes for qualifiers to be included in the AuthorizedView. Every
	//  qualifier starting with one of these prefixes is included in the
	//  AuthorizedView. To provide access to all qualifiers, include the empty
	//  string as a prefix
	//  ("").
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.FamilySubsets.qualifier_prefixes
	QualifierPrefixes [][]byte `json:"qualifierPrefixes,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AuthorizedView.SubsetView
type AuthorizedView_SubsetView struct {
	// Row prefixes to be included in the AuthorizedView.
	//  To provide access to all rows, include the empty string as a prefix ("").
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.SubsetView.row_prefixes
	RowPrefixes [][]byte `json:"rowPrefixes,omitempty"`

	// TODO: unsupported map type with key string and value message

}
