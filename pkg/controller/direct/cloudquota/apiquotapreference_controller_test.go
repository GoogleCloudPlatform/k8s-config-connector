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

package cloudquota

import (
	"testing"

	pb "cloud.google.com/go/cloudquotas/apiv1beta/cloudquotaspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestQuotaConfigEqual(t *testing.T) {
	tests := []struct {
		name    string
		desired *pb.QuotaConfig
		actual  *pb.QuotaConfig
		want    bool
	}{
		{
			name:    "both nil",
			desired: nil,
			actual:  nil,
			want:    true,
		},
		{
			name: "equal when actual has output-only fields populated",
			desired: &pb.QuotaConfig{
				PreferredValue: 100,
				Annotations:    map[string]string{"env": "prod"},
			},
			actual: &pb.QuotaConfig{
				PreferredValue: 100,
				Annotations:    map[string]string{"env": "prod"},
				GrantedValue:   wrapperspb.Int64(100),
				StateDetail:    "Quota granted",
				TraceId:        "trace-123",
				RequestOrigin:  pb.QuotaConfig_CLOUD_CONSOLE,
			},
			want: true,
		},
		{
			name: "not equal when preferred_value differs",
			desired: &pb.QuotaConfig{
				PreferredValue: 200,
			},
			actual: &pb.QuotaConfig{
				PreferredValue: 100,
				GrantedValue:   wrapperspb.Int64(100),
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			if got := quotaConfigEqual(mapCtx, tc.desired, tc.actual); got != tc.want {
				t.Errorf("quotaConfigEqual() = %v, want %v", got, tc.want)
			}
		})
	}
}
