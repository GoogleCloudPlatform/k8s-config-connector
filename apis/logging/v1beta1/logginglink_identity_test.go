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

func TestLoggingLinkIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *LoggingLinkIdentity
	}{
		{
			name: "valid project-scoped link",
			ref:  "projects/my-project/locations/global/buckets/my-bucket/links/my-link",
			want: &LoggingLinkIdentity{
				Project:  "my-project",
				Location: "global",
				Bucket:   "my-bucket",
				Link:     "my-link",
			},
		},
		{
			name: "valid folder-scoped link",
			ref:  "folders/123456/locations/global/buckets/my-bucket/links/my-link",
			want: &LoggingLinkIdentity{
				Folder:   "123456",
				Location: "global",
				Bucket:   "my-bucket",
				Link:     "my-link",
			},
		},
		{
			name: "valid organization-scoped link",
			ref:  "organizations/789012/locations/global/buckets/my-bucket/links/my-link",
			want: &LoggingLinkIdentity{
				Organization: "789012",
				Location:     "global",
				Bucket:       "my-bucket",
				Link:         "my-link",
			},
		},
		{
			name: "valid billingAccount-scoped link",
			ref:  "billingAccounts/999999/locations/global/buckets/my-bucket/links/my-link",
			want: &LoggingLinkIdentity{
				BillingAccount: "999999",
				Location:       "global",
				Bucket:         "my-bucket",
				Link:           "my-link",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "project-scoped full url",
			ref:  "https://logging.googleapis.com/projects/my-project/locations/global/buckets/my-bucket/links/my-link",
			want: &LoggingLinkIdentity{
				Project:  "my-project",
				Location: "global",
				Bucket:   "my-bucket",
				Link:     "my-link",
			},
		},
		{
			name: "folder-scoped full url",
			ref:  "https://logging.googleapis.com/folders/123456/locations/global/buckets/my-bucket/links/my-link",
			want: &LoggingLinkIdentity{
				Folder:   "123456",
				Location: "global",
				Bucket:   "my-bucket",
				Link:     "my-link",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &LoggingLinkIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				// Also verify the Round-Trip: Formatting back to string should equal the original canonical/relative ref
				var wantStr string
				if i.Project != "" {
					wantStr = "projects/" + i.Project + "/locations/" + i.Location + "/buckets/" + i.Bucket + "/links/" + i.Link
				} else if i.Folder != "" {
					wantStr = "folders/" + i.Folder + "/locations/" + i.Location + "/buckets/" + i.Bucket + "/links/" + i.Link
				} else if i.Organization != "" {
					wantStr = "organizations/" + i.Organization + "/locations/" + i.Location + "/buckets/" + i.Bucket + "/links/" + i.Link
				} else if i.BillingAccount != "" {
					wantStr = "billingAccounts/" + i.BillingAccount + "/locations/" + i.Location + "/buckets/" + i.Bucket + "/links/" + i.Link
				}
				if gotStr := i.String(); gotStr != wantStr {
					t.Errorf("String() got = %q, want = %q", gotStr, wantStr)
				}
			}
		})
	}
}
