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

package datalineage

import (
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func TestAttributesMapping(t *testing.T) {
	tests := []struct {
		name     string
		krmAttrs map[string]apiextensionsv1.JSON
	}{
		{
			name:     "nil map",
			krmAttrs: nil,
		},
		{
			name: "empty raw value ignored",
			krmAttrs: map[string]apiextensionsv1.JSON{
				"empty": {Raw: nil},
				"blank": {Raw: []byte{}},
			},
		},
		{
			name: "valid values roundtrip",
			krmAttrs: map[string]apiextensionsv1.JSON{
				"strKey":  {Raw: []byte(`"hello"`)},
				"numKey":  {Raw: []byte(`123`)},
				"boolKey": {Raw: []byte(`true`)},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			protoAttrs := Attributes_ToProto(mapCtx, tc.krmAttrs)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error converting to proto: %v", err)
			}

			if tc.name == "empty raw value ignored" {
				if len(protoAttrs) != 0 {
					t.Fatalf("expected empty proto map, got %v", protoAttrs)
				}
				return
			}

			if tc.krmAttrs == nil {
				if protoAttrs != nil {
					t.Fatalf("expected nil proto map, got %v", protoAttrs)
				}
				return
			}

			convertedBack := Attributes_FromProto(mapCtx, protoAttrs)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error converting from proto: %v", err)
			}

			if len(convertedBack) != len(tc.krmAttrs) {
				t.Fatalf("mismatched length: expected %d, got %d", len(tc.krmAttrs), len(convertedBack))
			}
			for k, v := range tc.krmAttrs {
				got, ok := convertedBack[k]
				if !ok {
					t.Fatalf("missing key %s in converted back", k)
				}
				if !reflect.DeepEqual(got.Raw, v.Raw) {
					t.Errorf("key %s: expected %s, got %s", k, string(v.Raw), string(got.Raw))
				}
			}
		})
	}
}
