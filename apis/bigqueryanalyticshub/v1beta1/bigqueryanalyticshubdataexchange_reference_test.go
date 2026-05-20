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

	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestBigQueryAnalyticsHubDataExchangeRef_GetGVK(t *testing.T) {
	tests := []struct {
		name string
		want schema.GroupVersionKind
	}{
		{
			name: "valid GVK",
			want: schema.GroupVersionKind{
				Group:   "bigqueryanalyticshub.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "BigQueryAnalyticsHubDataExchange",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &BigQueryAnalyticsHubDataExchangeRef{}
			if got := r.GetGVK(); got != tt.want {
				t.Errorf("GetGVK() = %v, want %v", got, tt.want)
			}
		})
	}
}
