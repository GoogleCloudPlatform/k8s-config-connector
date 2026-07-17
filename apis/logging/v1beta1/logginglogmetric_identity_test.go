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

func TestLoggingLogMetricIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *LoggingLogMetricIdentity
	}{
		{
			name: "valid project-scoped metric",
			ref:  "projects/my-project/metrics/my-metric",
			want: &LoggingLogMetricIdentity{
				Project: "my-project",
				Metric:  "my-metric",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "project-scoped full url",
			ref:  "https://logging.googleapis.com/projects/my-project/metrics/my-metric",
			want: &LoggingLogMetricIdentity{
				Project: "my-project",
				Metric:  "my-metric",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &LoggingLogMetricIdentity{}
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
				wantStr := "projects/" + i.Project + "/metrics/" + i.Metric
				if gotStr := i.String(); gotStr != wantStr {
					t.Errorf("String() got = %q, want = %q", gotStr, wantStr)
				}
			}
		})
	}
}
