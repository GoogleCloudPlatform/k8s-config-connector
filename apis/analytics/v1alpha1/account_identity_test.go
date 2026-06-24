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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestAccountIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name        string
		external    string
		wantAccount string
		wantErr     bool
	}{
		{
			name:        "valid external",
			external:    "accounts/12345",
			wantAccount: "12345",
			wantErr:     false,
		},
		{
			name:     "invalid external format",
			external: "invalid/format",
			wantErr:  true,
		},
		{
			name:     "wrong service prefix",
			external: "notaccounts/12345",
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
			id := &AccountIdentity{}
			err := id.FromExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && id.Account != tt.wantAccount {
				t.Errorf("FromExternal() gotAccount = %v, want %v", id.Account, tt.wantAccount)
			}
		})
	}
}

func TestAnalyticsAccount_GetIdentity(t *testing.T) {
	ctx := context.Background()
	fakeClient := fake.NewClientBuilder().Build()

	tests := []struct {
		name    string
		obj     *AnalyticsAccount
		wantID  string
		wantErr bool
	}{
		{
			name: "spec.resourceID is set",
			obj: &AnalyticsAccount{
				Spec: AnalyticsAccountSpec{
					ResourceID: direct.PtrTo("12345"),
				},
			},
			wantID:  "12345",
			wantErr: false,
		},
		{
			name: "spec.resourceID is empty, status.externalRef is set",
			obj: &AnalyticsAccount{
				Status: AnalyticsAccountStatus{
					ExternalRef: direct.PtrTo("accounts/67890"),
				},
			},
			wantID:  "67890",
			wantErr: false,
		},
		{
			name: "spec.resourceID and status.externalRef are both set and consistent",
			obj: &AnalyticsAccount{
				Spec: AnalyticsAccountSpec{
					ResourceID: direct.PtrTo("12345"),
				},
				Status: AnalyticsAccountStatus{
					ExternalRef: direct.PtrTo("accounts/12345"),
				},
			},
			wantID:  "12345",
			wantErr: false,
		},
		{
			name: "spec.resourceID and status.externalRef are inconsistent",
			obj: &AnalyticsAccount{
				Spec: AnalyticsAccountSpec{
					ResourceID: direct.PtrTo("12345"),
				},
				Status: AnalyticsAccountStatus{
					ExternalRef: direct.PtrTo("accounts/67890"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIdentity, err := tt.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				got := gotIdentity.(*AccountIdentity)
				if got.ID() != tt.wantID {
					t.Errorf("GetIdentity() ID = %v, want %v", got.ID(), tt.wantID)
				}
				if got.String() != "accounts/"+tt.wantID {
					t.Errorf("GetIdentity() String = %v, want %v", got.String(), "accounts/"+tt.wantID)
				}
			}
		})
	}
}
