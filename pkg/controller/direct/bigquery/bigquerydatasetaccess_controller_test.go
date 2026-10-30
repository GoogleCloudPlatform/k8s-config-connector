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

package bigquery

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	api "google.golang.org/api/bigquery/v2"
	"google.golang.org/api/googleapi"
)

func TestCanonicalRole(t *testing.T) {
	tests := []struct {
		name string
		role string
		want string
	}{
		{
			name: "dataOwner translates to OWNER",
			role: "roles/bigquery.dataOwner",
			want: "OWNER",
		},
		{
			name: "dataEditor translates to WRITER",
			role: "roles/bigquery.dataEditor",
			want: "WRITER",
		},
		{
			name: "dataViewer translates to READER",
			role: "roles/bigquery.dataViewer",
			want: "READER",
		},
		{
			name: "primitive OWNER unchanged",
			role: "OWNER",
			want: "OWNER",
		},
		{
			name: "primitive WRITER unchanged",
			role: "WRITER",
			want: "WRITER",
		},
		{
			name: "primitive READER unchanged",
			role: "READER",
			want: "READER",
		},
		{
			name: "custom role unchanged",
			role: "roles/bigquery.admin",
			want: "roles/bigquery.admin",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			role := tt.role

			// Act
			got := canonicalRole(role)

			// Assert
			if got != tt.want {
				t.Errorf("canonicalRole(%q) = %q, want %q", role, got, tt.want)
			}
		})
	}
}

func TestMatchesAccessEntry(t *testing.T) {
	tests := []struct {
		name             string
		entry            *api.DatasetAccess
		desired          *api.DatasetAccess
		desiredIamMember string
		wantMatched      bool
		wantApiUpdated   bool
	}{
		{
			name:        "nil entry returns false",
			entry:       nil,
			desired:     &api.DatasetAccess{Role: "READER"},
			wantMatched: false,
		},
		{
			name:        "nil desired returns false",
			entry:       &api.DatasetAccess{Role: "READER"},
			desired:     nil,
			wantMatched: false,
		},
		// View branch
		{
			name: "matching view",
			entry: &api.DatasetAccess{
				View: &api.TableReference{
					ProjectId: "test-proj",
					DatasetId: "test-ds",
					TableId:   "test-table",
				},
			},
			desired: &api.DatasetAccess{
				View: &api.TableReference{
					ProjectId: "test-proj",
					DatasetId: "test-ds",
					TableId:   "test-table",
				},
			},
			wantMatched:    true,
			wantApiUpdated: false,
		},
		{
			name: "mismatched view table id",
			entry: &api.DatasetAccess{
				View: &api.TableReference{
					ProjectId: "test-proj",
					DatasetId: "test-ds",
					TableId:   "table-a",
				},
			},
			desired: &api.DatasetAccess{
				View: &api.TableReference{
					ProjectId: "test-proj",
					DatasetId: "test-ds",
					TableId:   "table-b",
				},
			},
			wantMatched: false,
		},
		{
			name: "mismatched view dataset id",
			entry: &api.DatasetAccess{
				View: &api.TableReference{
					ProjectId: "test-proj",
					DatasetId: "ds-a",
					TableId:   "table-a",
				},
			},
			desired: &api.DatasetAccess{
				View: &api.TableReference{
					ProjectId: "test-proj",
					DatasetId: "ds-b",
					TableId:   "table-a",
				},
			},
			wantMatched: false,
		},
		{
			name: "mismatched view project id",
			entry: &api.DatasetAccess{
				View: &api.TableReference{
					ProjectId: "proj-a",
					DatasetId: "test-ds",
					TableId:   "table-a",
				},
			},
			desired: &api.DatasetAccess{
				View: &api.TableReference{
					ProjectId: "proj-b",
					DatasetId: "test-ds",
					TableId:   "table-a",
				},
			},
			wantMatched: false,
		},
		{
			name:  "desired view with nil entry view",
			entry: &api.DatasetAccess{},
			desired: &api.DatasetAccess{
				View: &api.TableReference{
					ProjectId: "test-proj",
					DatasetId: "test-ds",
					TableId:   "test-table",
				},
			},
			wantMatched: false,
		},
		// Routine branch
		{
			name: "matching routine",
			entry: &api.DatasetAccess{
				Routine: &api.RoutineReference{
					ProjectId: "test-proj",
					DatasetId: "test-ds",
					RoutineId: "test-routine",
				},
			},
			desired: &api.DatasetAccess{
				Routine: &api.RoutineReference{
					ProjectId: "test-proj",
					DatasetId: "test-ds",
					RoutineId: "test-routine",
				},
			},
			wantMatched:    true,
			wantApiUpdated: false,
		},
		{
			name: "mismatched routine id",
			entry: &api.DatasetAccess{
				Routine: &api.RoutineReference{
					ProjectId: "test-proj",
					DatasetId: "test-ds",
					RoutineId: "routine-a",
				},
			},
			desired: &api.DatasetAccess{
				Routine: &api.RoutineReference{
					ProjectId: "test-proj",
					DatasetId: "test-ds",
					RoutineId: "routine-b",
				},
			},
			wantMatched: false,
		},
		{
			name: "mismatched routine dataset id",
			entry: &api.DatasetAccess{
				Routine: &api.RoutineReference{
					ProjectId: "test-proj",
					DatasetId: "ds-a",
					RoutineId: "routine-a",
				},
			},
			desired: &api.DatasetAccess{
				Routine: &api.RoutineReference{
					ProjectId: "test-proj",
					DatasetId: "ds-b",
					RoutineId: "routine-a",
				},
			},
			wantMatched: false,
		},
		{
			name: "mismatched routine project id",
			entry: &api.DatasetAccess{
				Routine: &api.RoutineReference{
					ProjectId: "proj-a",
					DatasetId: "test-ds",
					RoutineId: "routine-a",
				},
			},
			desired: &api.DatasetAccess{
				Routine: &api.RoutineReference{
					ProjectId: "proj-b",
					DatasetId: "test-ds",
					RoutineId: "routine-a",
				},
			},
			wantMatched: false,
		},
		{
			name:  "desired routine with nil entry routine",
			entry: &api.DatasetAccess{},
			desired: &api.DatasetAccess{
				Routine: &api.RoutineReference{
					ProjectId: "test-proj",
					DatasetId: "test-ds",
					RoutineId: "test-routine",
				},
			},
			wantMatched: false,
		},
		// Dataset branch
		{
			name: "matching dataset",
			entry: &api.DatasetAccess{
				Dataset: &api.DatasetAccessEntry{
					Dataset: &api.DatasetReference{
						ProjectId: "test-proj",
						DatasetId: "src-ds",
					},
					TargetTypes: []string{"VIEWS"},
				},
			},
			desired: &api.DatasetAccess{
				Dataset: &api.DatasetAccessEntry{
					Dataset: &api.DatasetReference{
						ProjectId: "test-proj",
						DatasetId: "src-ds",
					},
					TargetTypes: []string{"VIEWS"},
				},
			},
			wantMatched:    true,
			wantApiUpdated: false,
		},
		{
			name: "mismatched dataset target types",
			entry: &api.DatasetAccess{
				Dataset: &api.DatasetAccessEntry{
					Dataset: &api.DatasetReference{
						ProjectId: "test-proj",
						DatasetId: "src-ds",
					},
					TargetTypes: []string{"VIEWS"},
				},
			},
			desired: &api.DatasetAccess{
				Dataset: &api.DatasetAccessEntry{
					Dataset: &api.DatasetReference{
						ProjectId: "test-proj",
						DatasetId: "src-ds",
					},
					TargetTypes: []string{"ROUTINES"},
				},
			},
			wantMatched: false,
		},
		{
			name: "mismatched dataset id",
			entry: &api.DatasetAccess{
				Dataset: &api.DatasetAccessEntry{
					Dataset: &api.DatasetReference{
						ProjectId: "test-proj",
						DatasetId: "ds-1",
					},
					TargetTypes: []string{"VIEWS"},
				},
			},
			desired: &api.DatasetAccess{
				Dataset: &api.DatasetAccessEntry{
					Dataset: &api.DatasetReference{
						ProjectId: "test-proj",
						DatasetId: "ds-2",
					},
					TargetTypes: []string{"VIEWS"},
				},
			},
			wantMatched: false,
		},
		{
			name:  "desired dataset with nil entry dataset",
			entry: &api.DatasetAccess{},
			desired: &api.DatasetAccess{
				Dataset: &api.DatasetAccessEntry{
					Dataset: &api.DatasetReference{
						ProjectId: "test-proj",
						DatasetId: "src-ds",
					},
					TargetTypes: []string{"VIEWS"},
				},
			},
			wantMatched: false,
		},
		// Role mismatch
		{
			name: "mismatched roles",
			entry: &api.DatasetAccess{
				Role:        "READER",
				UserByEmail: "user@example.com",
			},
			desired: &api.DatasetAccess{
				Role:        "WRITER",
				UserByEmail: "user@example.com",
			},
			wantMatched: false,
		},
		// Role canonical matching
		{
			name: "matching canonical role predefined and primitive",
			entry: &api.DatasetAccess{
				Role:        "READER",
				UserByEmail: "user@example.com",
			},
			desired: &api.DatasetAccess{
				Role:        "roles/bigquery.dataViewer",
				UserByEmail: "user@example.com",
			},
			wantMatched:    true,
			wantApiUpdated: false,
		},
		// UserByEmail
		{
			name: "matching user by email case-insensitive",
			entry: &api.DatasetAccess{
				Role:        "READER",
				UserByEmail: "User@Example.Com",
			},
			desired: &api.DatasetAccess{
				Role:        "READER",
				UserByEmail: "user@example.com",
			},
			wantMatched:    true,
			wantApiUpdated: false,
		},
		{
			name: "mismatched user by email",
			entry: &api.DatasetAccess{
				Role:        "READER",
				UserByEmail: "other@example.com",
			},
			desired: &api.DatasetAccess{
				Role:        "READER",
				UserByEmail: "user@example.com",
			},
			wantMatched: false,
		},
		// GroupByEmail
		{
			name: "matching group by email case-insensitive",
			entry: &api.DatasetAccess{
				Role:         "READER",
				GroupByEmail: "Group@Example.Com",
			},
			desired: &api.DatasetAccess{
				Role:         "READER",
				GroupByEmail: "group@example.com",
			},
			wantMatched:    true,
			wantApiUpdated: false,
		},
		{
			name: "mismatched group by email",
			entry: &api.DatasetAccess{
				Role:         "READER",
				GroupByEmail: "other-group@example.com",
			},
			desired: &api.DatasetAccess{
				Role:         "READER",
				GroupByEmail: "group@example.com",
			},
			wantMatched: false,
		},
		// Domain
		{
			name: "matching domain case-insensitive",
			entry: &api.DatasetAccess{
				Role:   "READER",
				Domain: "Example.COM",
			},
			desired: &api.DatasetAccess{
				Role:   "READER",
				Domain: "example.com",
			},
			wantMatched:    true,
			wantApiUpdated: false,
		},
		{
			name: "mismatched domain",
			entry: &api.DatasetAccess{
				Role:   "READER",
				Domain: "other.com",
			},
			desired: &api.DatasetAccess{
				Role:   "READER",
				Domain: "example.com",
			},
			wantMatched: false,
		},
		// SpecialGroup
		{
			name: "matching special group",
			entry: &api.DatasetAccess{
				Role:         "READER",
				SpecialGroup: "projectReaders",
			},
			desired: &api.DatasetAccess{
				Role:         "READER",
				SpecialGroup: "projectReaders",
			},
			wantMatched:    true,
			wantApiUpdated: false,
		},
		{
			name: "mismatched special group",
			entry: &api.DatasetAccess{
				Role:         "READER",
				SpecialGroup: "projectWriters",
			},
			desired: &api.DatasetAccess{
				Role:         "READER",
				SpecialGroup: "projectReaders",
			},
			wantMatched: false,
		},
		// IamMember exact / direct match
		{
			name: "matching iamMember directly",
			entry: &api.DatasetAccess{
				Role:      "READER",
				IamMember: "allUsers",
			},
			desired: &api.DatasetAccess{
				Role: "READER",
			},
			desiredIamMember: "allUsers",
			wantMatched:      true,
			wantApiUpdated:   false,
		},
		{
			name: "matching deleted service account with tombstone",
			entry: &api.DatasetAccess{
				Role:      "READER",
				IamMember: "deleted:serviceAccount:my-sa@test.iam.gserviceaccount.com?uid=12345",
			},
			desired: &api.DatasetAccess{
				Role: "READER",
			},
			desiredIamMember: "deleted:serviceAccount:my-sa@test.iam.gserviceaccount.com?uid=12345",
			wantMatched:      true,
			wantApiUpdated:   false,
		},
		{
			name: "matching deleted service account entry with desired non-deleted service account",
			entry: &api.DatasetAccess{
				Role:      "READER",
				IamMember: "deleted:serviceAccount:my-sa@test.iam.gserviceaccount.com?uid=12345",
			},
			desired: &api.DatasetAccess{
				Role: "READER",
			},
			desiredIamMember: "serviceAccount:my-sa@test.iam.gserviceaccount.com",
			wantMatched:      true,
			wantApiUpdated:   false,
		},
		// IamMember mapped by API (apiUpdatedMember = true)
		{
			name: "iamMember user mapped to userByEmail",
			entry: &api.DatasetAccess{
				Role:        "READER",
				UserByEmail: "alice@example.com",
			},
			desired: &api.DatasetAccess{
				Role: "READER",
			},
			desiredIamMember: "user:alice@example.com",
			wantMatched:      true,
			wantApiUpdated:   true,
		},
		{
			name: "iamMember serviceAccount mapped to userByEmail",
			entry: &api.DatasetAccess{
				Role:        "READER",
				UserByEmail: "sa@test.iam.gserviceaccount.com",
			},
			desired: &api.DatasetAccess{
				Role: "READER",
			},
			desiredIamMember: "serviceAccount:sa@test.iam.gserviceaccount.com",
			wantMatched:      true,
			wantApiUpdated:   true,
		},
		{
			name: "iamMember deleted service account mapped to userByEmail",
			entry: &api.DatasetAccess{
				Role:        "READER",
				UserByEmail: "sa@test.iam.gserviceaccount.com",
			},
			desired: &api.DatasetAccess{
				Role: "READER",
			},
			desiredIamMember: "deleted:serviceAccount:sa@test.iam.gserviceaccount.com?uid=12345",
			wantMatched:      true,
			wantApiUpdated:   true,
		},
		{
			name: "iamMember group mapped to groupByEmail",
			entry: &api.DatasetAccess{
				Role:         "READER",
				GroupByEmail: "dev-team@example.com",
			},
			desired: &api.DatasetAccess{
				Role: "READER",
			},
			desiredIamMember: "group:dev-team@example.com",
			wantMatched:      true,
			wantApiUpdated:   true,
		},
		{
			name: "iamMember domain mapped to domain",
			entry: &api.DatasetAccess{
				Role:   "READER",
				Domain: "example.com",
			},
			desired: &api.DatasetAccess{
				Role: "READER",
			},
			desiredIamMember: "domain:example.com",
			wantMatched:      true,
			wantApiUpdated:   true,
		},
		{
			name: "iamMember specialGroup mapped to specialGroup",
			entry: &api.DatasetAccess{
				Role:         "READER",
				SpecialGroup: "projectOwners",
			},
			desired: &api.DatasetAccess{
				Role: "READER",
			},
			desiredIamMember: "projectOwners",
			wantMatched:      true,
			wantApiUpdated:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			entry := tt.entry
			desired := tt.desired
			desiredIamMember := tt.desiredIamMember

			// Act
			matched, apiUpdated := matchesAccessEntry(entry, desired, desiredIamMember)

			// Assert
			if matched != tt.wantMatched {
				t.Errorf("matchesAccessEntry() matched = %v, want %v", matched, tt.wantMatched)
			}
			if apiUpdated != tt.wantApiUpdated {
				t.Errorf("matchesAccessEntry() apiUpdated = %v, want %v", apiUpdated, tt.wantApiUpdated)
			}
		})
	}
}

func TestAdapterForURL(t *testing.T) {
	// Arrange
	ctx := context.Background()
	m := &datasetAccessModel{}
	url := "https://bigquery.googleapis.com/bigquery/v2/projects/p/datasets/d"

	// Act
	adapter, err := m.AdapterForURL(ctx, url)

	// Assert
	if adapter != nil {
		t.Errorf("AdapterForURL() returned adapter = %v, want nil", adapter)
	}
	if err != nil {
		t.Errorf("AdapterForURL() returned unexpected error = %v, want nil", err)
	}
}

func TestDescribeAccessDiff(t *testing.T) {
	tests := []struct {
		name             string
		actual           *api.DatasetAccess
		desired          *api.DatasetAccess
		desiredIamMember string
		wantSubstring    string
	}{
		{
			name: "role diff",
			actual: &api.DatasetAccess{
				Role:        "READER",
				UserByEmail: "alice@example.com",
			},
			desired: &api.DatasetAccess{
				Role:        "OWNER",
				UserByEmail: "alice@example.com",
			},
			wantSubstring: `role changed from "READER" to "OWNER"`,
		},
		{
			name: "grantee diff",
			actual: &api.DatasetAccess{
				Role:        "READER",
				UserByEmail: "alice@example.com",
			},
			desired: &api.DatasetAccess{
				Role:        "READER",
				UserByEmail: "bob@example.com",
			},
			wantSubstring: `userByEmail changed from "alice@example.com" to "bob@example.com"`,
		},
		{
			name: "view diff",
			actual: &api.DatasetAccess{
				View: &api.TableReference{
					ProjectId: "p",
					DatasetId: "d",
					TableId:   "table1",
				},
			},
			desired: &api.DatasetAccess{
				View: &api.TableReference{
					ProjectId: "p",
					DatasetId: "d",
					TableId:   "table2",
				},
			},
			wantSubstring: `view changed from p/d/table1 to p/d/table2`,
		},
		{
			name: "routine diff",
			actual: &api.DatasetAccess{
				Routine: &api.RoutineReference{
					ProjectId: "p",
					DatasetId: "d",
					RoutineId: "r1",
				},
			},
			desired: &api.DatasetAccess{
				Routine: &api.RoutineReference{
					ProjectId: "p",
					DatasetId: "d",
					RoutineId: "r2",
				},
			},
			wantSubstring: `routine changed from p/d/r1 to p/d/r2`,
		},
		{
			name: "dataset reference diff",
			actual: &api.DatasetAccess{
				Dataset: &api.DatasetAccessEntry{
					Dataset: &api.DatasetReference{
						ProjectId: "p",
						DatasetId: "ds1",
					},
					TargetTypes: []string{"VIEWS"},
				},
			},
			desired: &api.DatasetAccess{
				Dataset: &api.DatasetAccessEntry{
					Dataset: &api.DatasetReference{
						ProjectId: "p",
						DatasetId: "ds2",
					},
					TargetTypes: []string{"VIEWS"},
				},
			},
			wantSubstring: `dataset reference changed from p/ds1 to p/ds2`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			actual := tt.actual
			desired := tt.desired
			desiredIamMember := tt.desiredIamMember

			// Act
			diff := describeAccessDiff(actual, desired, desiredIamMember)

			// Assert
			if !strings.Contains(diff, tt.wantSubstring) {
				t.Errorf("describeAccessDiff() = %q, want substring %q", diff, tt.wantSubstring)
			}
		})
	}
}

func TestIsConflictOrPreconditionFailed(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "precondition failed (HTTP 412)",
			err: &googleapi.Error{
				Code:    http.StatusPreconditionFailed,
				Message: "Precondition Failed",
			},
			want: true,
		},
		{
			name: "conflict (HTTP 409)",
			err: &googleapi.Error{
				Code:    http.StatusConflict,
				Message: "Conflict",
			},
			want: true,
		},
		{
			name: "other http error (HTTP 400)",
			err: &googleapi.Error{
				Code:    http.StatusBadRequest,
				Message: "Bad Request",
			},
			want: false,
		},
		{
			name: "generic precondition message",
			err:  fmt.Errorf("condition failed: precondition check failed"),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			err := tt.err

			// Act
			got := isConflictOrPreconditionFailed(err)

			// Assert
			if got != tt.want {
				t.Errorf("isConflictOrPreconditionFailed(%v) = %v, want %v", err, got, tt.want)
			}
		})
	}
}
