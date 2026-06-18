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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/google/go-cmp/cmp"
)

func TestPrivateCACertificateTemplateIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *PrivateCACertificateTemplateIdentity
		wantErr bool
	}{
		{
			name: "valid certificateTemplate reference",
			ref:  "projects/test-project/locations/us-central1/certificateTemplates/test-template",
			want: &PrivateCACertificateTemplateIdentity{
				Project:             "test-project",
				Location:            "us-central1",
				CertificateTemplate: "test-template",
			},
			wantErr: false,
		},
		{
			name:    "invalid format - extra segment",
			ref:     "projects/test-project/locations/us-central1/certificateTemplates/test-template/extra",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid format - missing location",
			ref:     "projects/test-project/certificateTemplates/test-template",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &PrivateCACertificateTemplateIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("PrivateCACertificateTemplateIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("PrivateCACertificateTemplateIdentity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestPrivateCACertificateTemplateRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		wantErr  bool
	}{
		{
			name:     "valid external",
			external: "projects/my-project/locations/global/certificateTemplates/my-template",
			wantErr:  false,
		},
		{
			name:     "invalid external",
			external: "projects/my-project/certificateTemplates/my-template",
			wantErr:  true,
		},
		{
			name:     "empty external",
			external: "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PrivateCACertificateTemplateRef{}
			if err := r.ValidateExternal(tt.external); (err != nil) != tt.wantErr {
				t.Errorf("PrivateCACertificateTemplateRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPrivateCACertificateTemplateImplementsIdentity(t *testing.T) {
	var _ identity.Identity = &PrivateCACertificateTemplateIdentity{}
	var _ identity.IdentityV2 = &PrivateCACertificateTemplateIdentity{}
}
