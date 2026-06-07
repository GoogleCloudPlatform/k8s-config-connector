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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/google/go-cmp/cmp"
)

func TestLoggingLogSinkIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *LoggingLogSinkIdentity
	}{
		{
			name: "valid project-scoped sink",
			ref:  "projects/my-project/sinks/my-sink",
			want: &LoggingLogSinkIdentity{
				Project: "my-project",
				Sink:    "my-sink",
			},
		},
		{
			name: "valid folder-scoped sink",
			ref:  "folders/123456/sinks/my-sink",
			want: &LoggingLogSinkIdentity{
				Folder: "123456",
				Sink:   "my-sink",
			},
		},
		{
			name: "valid organization-scoped sink",
			ref:  "organizations/789012/sinks/my-sink",
			want: &LoggingLogSinkIdentity{
				Organization: "789012",
				Sink:         "my-sink",
			},
		},
		{
			name: "valid billingAccount-scoped sink",
			ref:  "billingAccounts/999999/sinks/my-sink",
			want: &LoggingLogSinkIdentity{
				BillingAccount: "999999",
				Sink:           "my-sink",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "project-scoped full url",
			ref:  "https://logging.googleapis.com/projects/my-project/sinks/my-sink",
			want: &LoggingLogSinkIdentity{
				Project: "my-project",
				Sink:    "my-sink",
			},
		},
		{
			name: "folder-scoped full url",
			ref:  "https://logging.googleapis.com/folders/123456/sinks/my-sink",
			want: &LoggingLogSinkIdentity{
				Folder: "123456",
				Sink:   "my-sink",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &LoggingLogSinkIdentity{}
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
					wantStr = "projects/" + i.Project + "/sinks/" + i.Sink
				} else if i.Folder != "" {
					wantStr = "folders/" + i.Folder + "/sinks/" + i.Sink
				} else if i.Organization != "" {
					wantStr = "organizations/" + i.Organization + "/sinks/" + i.Sink
				} else if i.BillingAccount != "" {
					wantStr = "billingAccounts/" + i.BillingAccount + "/sinks/" + i.Sink
				}
				if gotStr := i.String(); gotStr != wantStr {
					t.Errorf("String() got = %q, want = %q", gotStr, wantStr)
				}
			}
		})
	}
}

func TestLoggingLogSinkIdentity_GetIdentityValidation(t *testing.T) {
	tests := []struct {
		name    string
		obj     *LoggingLogSink
		wantErr bool
	}{
		{
			name: "only projectRef",
			obj: &LoggingLogSink{
				Spec: LoggingLogSinkSpec{
					ProjectRef: &refs.ProjectRef{External: "projects/my-project"},
				},
			},
			wantErr: false,
		},
		{
			name: "projectRef and folderRef",
			obj: &LoggingLogSink{
				Spec: LoggingLogSinkSpec{
					ProjectRef: &refs.ProjectRef{External: "projects/my-project"},
					FolderRef:  &refs.FolderRef{External: "folders/123456"},
				},
			},
			wantErr: true,
		},
		{
			name: "projectRef, folderRef and organizationRef",
			obj: &LoggingLogSink{
				Spec: LoggingLogSinkSpec{
					ProjectRef:      &refs.ProjectRef{External: "projects/my-project"},
					FolderRef:       &refs.FolderRef{External: "folders/123456"},
					OrganizationRef: &refs.OrganizationRef{External: "organizations/789012"},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Give it a metadata name so refs.GetResourceID has something
			tt.obj.SetName("my-sink")
			_, err := getIdentityFromLoggingLogSinkSpec(nil, nil, tt.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIdentityFromLoggingLogSinkSpec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
