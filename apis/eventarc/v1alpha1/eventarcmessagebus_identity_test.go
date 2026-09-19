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

package v1alpha1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEventarcMessageBusIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name         string
		ref          string
		wantIdentity *EventarcMessageBusIdentity
		wantErr      bool
	}{
		{
			name: "valid full URL",
			ref:  "//eventarc.googleapis.com/projects/my-project/locations/us-central1/messageBuses/my-bus",
			wantIdentity: &EventarcMessageBusIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				MessageBus: "my-bus",
			},
		},
		{
			name: "valid relative path",
			ref:  "projects/my-project/locations/us-central1/messageBuses/my-bus",
			wantIdentity: &EventarcMessageBusIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				MessageBus: "my-bus",
			},
		},
		{
			name:    "invalid format - missing location",
			ref:     "projects/my-project/messageBuses/my-bus",
			wantErr: true,
		},
		{
			name:    "invalid format - extra segments",
			ref:     "projects/my-project/locations/us-central1/messageBuses/my-bus/extra",
			wantErr: true,
		},
		{
			name:    "empty messageBus",
			ref:     "projects/my-project/locations/us-central1/messageBuses/",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &EventarcMessageBusIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.wantIdentity, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestEventarcMessageBusIdentity_Methods(t *testing.T) {
	i := &EventarcMessageBusIdentity{
		Project:    "my-project",
		Location:   "us-central1",
		MessageBus: "my-bus",
	}

	want := "projects/my-project/locations/us-central1/messageBuses/my-bus"
	if got := i.String(); got != want {
		t.Errorf("String() got = %v, want = %v", got, want)
	}

	wantParent := "projects/my-project/locations/us-central1"
	if got := i.ParentString(); got != wantParent {
		t.Errorf("ParentString() got = %v, want = %v", got, wantParent)
	}
}
