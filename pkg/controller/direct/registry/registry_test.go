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

package registry

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestInPlaceUpdates(t *testing.T) {
	gvk := schema.GroupVersionKind{Group: "test.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "TestKind"}
	tests := []struct {
		name    string
		options []ModelOption
		want    []string
	}{
		{
			name: "no fields declared",
			want: nil,
		},
		{
			name:    "fields declared",
			options: []ModelOption{WithInPlaceUpdates("version", "remoteRepositoryConfig.description")},
			want:    []string{"version", "remoteRepositoryConfig.description"},
		},
		{
			name:    "fields from several options add up",
			options: []ModelOption{WithInPlaceUpdates("version"), WithIAM(), WithInPlaceUpdates("labels")},
			want:    []string{"version", "labels"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			RegisterModel(gvk, nil, tc.options...)
			t.Cleanup(func() { delete(singleton.registrations, gvk.GroupKind()) })

			got := InPlaceUpdates(gvk.GroupKind())

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("InPlaceUpdates(%v) mismatch (-want +got):\n%s", gvk.GroupKind(), diff)
			}
		})
	}
}

func TestInPlaceUpdatesForUnregisteredKind(t *testing.T) {
	gk := schema.GroupKind{Group: "test.cnrm.cloud.google.com", Kind: "UnregisteredKind"}

	got := InPlaceUpdates(gk)

	if got != nil {
		t.Errorf("InPlaceUpdates(%v) = %v, want nil", gk, got)
	}
}
