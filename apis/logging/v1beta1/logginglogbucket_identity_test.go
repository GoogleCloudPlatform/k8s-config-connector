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

func TestLoggingLogBucketIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *LogBucketIdentity
	}{
		{
			name: "valid project-scoped relative reference",
			ref:  "projects/my-project/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				Project:  "my-project",
				Location: "global",
				Bucket:   "my-bucket",
			},
		},
		{
			name: "valid colon-partitioned sovereign project relative reference",
			ref:  "projects/partition:sample-project/locations/u-region-1/buckets/my-bucket",
			want: &LogBucketIdentity{
				Project:  "partition:sample-project",
				Location: "u-region-1",
				Bucket:   "my-bucket",
			},
		},
		{
			name: "commercial full authority URI",
			ref:  "//logging.googleapis.com/projects/my-project/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				Project:  "my-project",
				Location: "global",
				Bucket:   "my-bucket",
			},
		},
		{
			name: "sovereign full authority URI",
			ref:  "//logging.custom.universe.goog/projects/partition:sample-project/locations/u-region-1/buckets/my-bucket",
			want: &LogBucketIdentity{
				Project:  "partition:sample-project",
				Location: "u-region-1",
				Bucket:   "my-bucket",
			},
		},
		{
			name: "sovereign HTTPS URL with v2 API version",
			ref:  "https://logging.custom.universe.goog/v2/projects/partition:sample-project/locations/u-region-1/buckets/my-bucket",
			want: &LogBucketIdentity{
				Project:  "partition:sample-project",
				Location: "u-region-1",
				Bucket:   "my-bucket",
			},
		},
		{
			name: "folder-scoped reference",
			ref:  "folders/12345/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				Folder:   "12345",
				Location: "global",
				Bucket:   "my-bucket",
			},
		},
		{
			name: "organization-scoped reference",
			ref:  "organizations/67890/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				Organization: "67890",
				Location:     "global",
				Bucket:       "my-bucket",
			},
		},
		{
			name: "billingAccount-scoped reference",
			ref:  "billingAccounts/123-456/locations/global/buckets/my-bucket",
			want: &LogBucketIdentity{
				BillingAccount: "123-456",
				Location:       "global",
				Bucket:         "my-bucket",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
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
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestLoggingLogBucketIdentity_Host(t *testing.T) {
	i := &LogBucketIdentity{}

	// Default domain
	if got := i.Host(); got != "logging.googleapis.com" {
		t.Errorf("expected logging.googleapis.com, got %s", got)
	}

	// Sovereign domain
	t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", "custom.universe.goog")
	if got := i.Host(); got != "logging.custom.universe.goog" {
		t.Errorf("expected logging.custom.universe.goog, got %s", got)
	}
}
