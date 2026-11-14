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

package v1beta1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLogBucketIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name       string
		ref        string
		wantErr    bool
		want       *LogBucketIdentity
		wantParent string
	}{
		{
			name: "valid project-scoped bucket",
			ref:  "projects/my-project/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				Project:  "my-project",
				Location: "global",
				Bucket:   "my-bucket",
			},
			wantParent: "projects/my-project/locations/global",
		},
		{
			name: "valid folder-scoped bucket",
			ref:  "folders/123456/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				Folder:   "123456",
				Location: "global",
				Bucket:   "my-bucket",
			},
			wantParent: "folders/123456/locations/global",
		},
		{
			name: "valid organization-scoped bucket",
			ref:  "organizations/789012/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				Organization: "789012",
				Location:     "global",
				Bucket:       "my-bucket",
			},
			wantParent: "organizations/789012/locations/global",
		},
		{
			name: "valid billingAccount-scoped bucket",
			ref:  "billingAccounts/999999/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				BillingAccount: "999999",
				Location:       "global",
				Bucket:         "my-bucket",
			},
			wantParent: "billingAccounts/999999/locations/global",
		},
		{
			name: "valid accessPolicy-scoped bucket",
			ref:  "accessPolicies/111222/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				AccessPolicy: "111222",
				Location:     "global",
				Bucket:       "my-bucket",
			},
			wantParent: "accessPolicies/111222/locations/global",
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "project-scoped commercial https url",
			ref:  "https://logging.googleapis.com/projects/my-project/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				Project:  "my-project",
				Location: "global",
				Bucket:   "my-bucket",
			},
			wantParent: "projects/my-project/locations/global",
		},
		{
			name: "project-scoped commercial authority prefix (//logging.googleapis.com/)",
			ref:  "//logging.googleapis.com/projects/my-project/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				Project:  "my-project",
				Location: "global",
				Bucket:   "my-bucket",
			},
			wantParent: "projects/my-project/locations/global",
		},
		{
			name: "project-scoped sovereign authority prefix (//logging.custom.universe.goog/)",
			ref:  "//logging.custom.universe.goog/projects/my-project/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				Project:  "my-project",
				Location: "global",
				Bucket:   "my-bucket",
			},
			wantParent: "projects/my-project/locations/global",
		},
		{
			name: "folder-scoped sovereign https url (https://logging.custom.universe.goog/)",
			ref:  "https://logging.custom.universe.goog/folders/123456/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				Folder:   "123456",
				Location: "global",
				Bucket:   "my-bucket",
			},
			wantParent: "folders/123456/locations/global",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &LogBucketIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				// Verify String() outputs the relative resource name without commercial authority
				var wantStr string
				if i.Project != "" {
					wantStr = "projects/" + i.Project + "/locations/" + i.Location + "/buckets/" + i.Bucket
				} else if i.Folder != "" {
					wantStr = "folders/" + i.Folder + "/locations/" + i.Location + "/buckets/" + i.Bucket
				} else if i.Organization != "" {
					wantStr = "organizations/" + i.Organization + "/locations/" + i.Location + "/buckets/" + i.Bucket
				} else if i.BillingAccount != "" {
					wantStr = "billingAccounts/" + i.BillingAccount + "/locations/" + i.Location + "/buckets/" + i.Bucket
				} else if i.AccessPolicy != "" {
					wantStr = "accessPolicies/" + i.AccessPolicy + "/locations/" + i.Location + "/buckets/" + i.Bucket
				}
				if gotStr := i.String(); gotStr != wantStr {
					t.Errorf("String() got = %q, want = %q", gotStr, wantStr)
				}
				if gotParent := i.ParentString(); gotParent != tt.wantParent {
					t.Errorf("ParentString() got = %q, want = %q", gotParent, tt.wantParent)
				}
				if gotID := i.ID(); gotID != tt.want.Bucket {
					t.Errorf("ID() got = %q, want = %q", gotID, tt.want.Bucket)
				}
			}
		})
	}
}
