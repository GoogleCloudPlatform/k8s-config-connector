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
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestCloudIdentityDeviceIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name                  string
		ref                   string
		wantErr               bool
		want                  *CloudIdentityDeviceIdentity
		wantIdentitySpecified bool
	}{
		{
			name: "valid reference",
			ref:  "devices/my-device",
			want: &CloudIdentityDeviceIdentity{
				Device: "my-device",
			},
			wantIdentitySpecified: true,
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format/extra",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://cloudidentity.googleapis.com/devices/my-device",
			want: &CloudIdentityDeviceIdentity{
				Device: "my-device",
			},
			wantIdentitySpecified: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CloudIdentityDeviceIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.HasIdentitySpecified(); got != tt.wantIdentitySpecified {
					t.Errorf("HasIdentitySpecified() = %v, want %v", got, tt.wantIdentitySpecified)
				}
			}
		})
	}
}

func TestCloudIdentityDevice_GetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name                  string
		obj                   *CloudIdentityDevice
		wantErr               bool
		want                  *CloudIdentityDeviceIdentity
		wantIdentitySpecified bool
	}{
		{
			name: "GetIdentity with specified resourceID",
			obj: &CloudIdentityDevice{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-device",
					Namespace: "my-namespace",
				},
				Spec: CloudIdentityDeviceSpec{
					ResourceID: common.LazyPtr("C02Z12345"),
				},
			},
			want: &CloudIdentityDeviceIdentity{
				Device: "C02Z12345",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with empty resourceID (not yet created/assigned)",
			obj: &CloudIdentityDevice{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-device",
					Namespace: "my-namespace",
				},
				Spec: CloudIdentityDeviceSpec{},
			},
			want: &CloudIdentityDeviceIdentity{
				Device: "",
			},
			wantIdentitySpecified: false,
		},
		{
			name: "GetIdentity with empty resourceID adopting status.externalRef",
			obj: &CloudIdentityDevice{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-device",
					Namespace: "my-namespace",
				},
				Spec: CloudIdentityDeviceSpec{},
				Status: CloudIdentityDeviceStatus{
					ExternalRef: common.LazyPtr("devices/C02Z54321"),
				},
			},
			want: &CloudIdentityDeviceIdentity{
				Device: "C02Z54321",
			},
			wantIdentitySpecified: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
				if gotIdentitySpecified := got.(*CloudIdentityDeviceIdentity).HasIdentitySpecified(); gotIdentitySpecified != tt.wantIdentitySpecified {
					t.Errorf("HasIdentitySpecified() = %v, want %v", gotIdentitySpecified, tt.wantIdentitySpecified)
				}
			}
		})
	}
}
