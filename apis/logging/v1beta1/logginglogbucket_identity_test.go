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
)

func TestLogBucketIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *LogBucketIdentity
		wantErr bool
	}{
		{
			name: "project log bucket",
			ref:  "projects/my-project/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				Project:  "my-project",
				Location: "global",
				Bucket:   "my-bucket",
			},
		},
		{
			name: "folder log bucket",
			ref:  "folders/my-folder/locations/us-central1/buckets/my-bucket",
			want: &LogBucketIdentity{
				Folder:   "my-folder",
				Location: "us-central1",
				Bucket:   "my-bucket",
			},
		},
		{
			name: "organization log bucket",
			ref:  "organizations/my-org/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				Organization: "my-org",
				Location:     "global",
				Bucket:       "my-bucket",
			},
		},
		{
			name: "billing account log bucket",
			ref:  "billingAccounts/my-billing/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				BillingAccount: "my-billing",
				Location:       "global",
				Bucket:         "my-bucket",
			},
		},
		{
			name: "access policy log bucket",
			ref:  "accessPolicies/my-policy/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				AccessPolicy: "my-policy",
				Location:     "global",
				Bucket:       "my-bucket",
			},
		},
		{
			name:    "invalid format",
			ref:     "projects/my-project/locations/global/buckets/my-bucket/extra",
			wantErr: true,
		},
		{
			name:    "another invalid format",
			ref:     "projects/my-project/buckets/my-bucket",
			wantErr: true,
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
				if i.Project != tt.want.Project {
					t.Errorf("FromExternal() got Project = %v, want %v", i.Project, tt.want.Project)
				}
				if i.Folder != tt.want.Folder {
					t.Errorf("FromExternal() got Folder = %v, want %v", i.Folder, tt.want.Folder)
				}
				if i.Organization != tt.want.Organization {
					t.Errorf("FromExternal() got Organization = %v, want %v", i.Organization, tt.want.Organization)
				}
				if i.BillingAccount != tt.want.BillingAccount {
					t.Errorf("FromExternal() got BillingAccount = %v, want %v", i.BillingAccount, tt.want.BillingAccount)
				}
				if i.AccessPolicy != tt.want.AccessPolicy {
					t.Errorf("FromExternal() got AccessPolicy = %v, want %v", i.AccessPolicy, tt.want.AccessPolicy)
				}
				if i.Location != tt.want.Location {
					t.Errorf("FromExternal() got Location = %v, want %v", i.Location, tt.want.Location)
				}
				if i.Bucket != tt.want.Bucket {
					t.Errorf("FromExternal() got Bucket = %v, want %v", i.Bucket, tt.want.Bucket)
				}
			}
		})
	}
}
