// Copyright 2023 Google LLC
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

package main

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/fielddesc"
)

func TestIsResourceField(ot *testing.T) {
	testCases := []struct {
		name     string
		field    fielddesc.FieldDescription
		expected bool
	}{
		{
			// This field has the Ref suffix, and the expected children.
			name: "reference-field",
			field: fielddesc.FieldDescription{
				ShortName: "ThisIsARef",
				Children: []fielddesc.FieldDescription{
					{
						ShortName: "external",
					},
					{
						ShortName: "namespace",
					},
					{
						ShortName: "name",
					},
				},
			},
			expected: true,
		},
		{
			// This field does not have the Ref name suffix.
			name: "not-reference-field",
			field: fielddesc.FieldDescription{
				ShortName: "NotRefField",
				Children: []fielddesc.FieldDescription{
					{
						ShortName: "other",
					},
					{
						ShortName: "children",
					},
				},
			},
			expected: false,
		},
		{
			// This field has the Ref name suffix, but does not match the expected children fields.
			// refer to the v1alpha1.SecretKeyReference struct at pkg/apis/core/v1alpha1/krm_types.go
			name: "secret-key-ref",
			field: fielddesc.FieldDescription{
				ShortName: "SecretKeyRef",
				Children: []fielddesc.FieldDescription{
					{
						ShortName: "name",
					},
					{
						ShortName: "key",
					},
				},
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		ot.Run(tc.name, func(t *testing.T) {
			actual := isResourceReference(tc.field)
			if tc.expected && !actual {
				t.Errorf("expected field to be resource ref: %+v", tc.field)
			} else if !tc.expected && actual {
				t.Errorf("expected field to not be resource ref: %+v", tc.field)
			}
		})
	}
}
