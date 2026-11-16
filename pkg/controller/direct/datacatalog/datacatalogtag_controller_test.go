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

package datacatalog

import (
	"testing"

	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
)

func TestTagFieldsEqual(t *testing.T) {
	tests := []struct {
		name    string
		desired map[string]*pb.TagField
		actual  map[string]*pb.TagField
		want    bool
	}{
		{
			name:    "both nil",
			desired: nil,
			actual:  nil,
			want:    true,
		},
		{
			name: "equal when actual has output-only display_name and order",
			desired: map[string]*pb.TagField{
				"source": {
					Kind: &pb.TagField_StringValue{StringValue: "crm"},
				},
			},
			actual: map[string]*pb.TagField{
				"source": {
					DisplayName: "Source System",
					Order:       1,
					Kind:        &pb.TagField_StringValue{StringValue: "crm"},
				},
			},
			want: true,
		},
		{
			name: "not equal when value differs",
			desired: map[string]*pb.TagField{
				"source": {
					Kind: &pb.TagField_StringValue{StringValue: "erp"},
				},
			},
			actual: map[string]*pb.TagField{
				"source": {
					DisplayName: "Source System",
					Order:       1,
					Kind:        &pb.TagField_StringValue{StringValue: "crm"},
				},
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tagFieldsEqual(tc.desired, tc.actual); got != tc.want {
				t.Errorf("tagFieldsEqual() = %v, want %v", got, tc.want)
			}
		})
	}
}
