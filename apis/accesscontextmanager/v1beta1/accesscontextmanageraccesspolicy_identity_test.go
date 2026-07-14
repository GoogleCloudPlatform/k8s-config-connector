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
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAccessPolicyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		want     *AccessPolicyIdentityV2
		wantStr  string
		wantErr  bool
	}{
		{
			name:     "canonical",
			external: "accessPolicies/12345",
			want: &AccessPolicyIdentityV2{
				AccessPolicy: "12345",
			},
			wantStr: "accessPolicies/12345",
		},
		{
			name:     "full URL",
			external: "//accesscontextmanager.googleapis.com/accessPolicies/12345",
			want: &AccessPolicyIdentityV2{
				AccessPolicy: "12345",
			},
			wantStr: "accessPolicies/12345",
		},
		{
			name:     "invalid format",
			external: "projects/my-project/locations/global/accessPolicies/12345",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AccessPolicyIdentityV2{}
			err := i.FromExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if i.AccessPolicy != tt.want.AccessPolicy {
				t.Errorf("FromExternal() got = %v, want %v", i, tt.want)
			}
			if i.String() != tt.wantStr {
				t.Errorf("String() got = %v, want %v", i.String(), tt.wantStr)
			}
		})
	}
}

func TestAccessPolicyIdentityV2(t *testing.T) {
	var _ identity.IdentityV2 = &AccessPolicyIdentityV2{}
	var _ identity.Resource = &AccessContextManagerAccessPolicy{}
}

func TestGetIdentityFromAccessContextManagerAccessPolicySpec(t *testing.T) {
	obj := &AccessContextManagerAccessPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-policy",
		},
		Spec: AccessContextManagerAccessPolicySpec{
			ResourceID: commonPtr("12345"),
			Title:      commonPtr("My Policy"),
		},
	}

	id, err := obj.GetIdentity(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetIdentity failed: %v", err)
	}

	if id.String() != "accessPolicies/12345" {
		t.Errorf("Expected accessPolicies/12345, got %s", id.String())
	}
}

func commonPtr[T any](val T) *T {
	return &val
}
