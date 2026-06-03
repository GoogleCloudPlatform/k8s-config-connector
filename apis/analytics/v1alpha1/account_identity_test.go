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

	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func ptrTo(s string) *string {
	return &s
}

func TestParseAccountExternal(t *testing.T) {
	tests := []struct {
		name           string
		external       string
		wantResourceID string
		wantErr        bool
	}{
		{
			name:           "valid external",
			external:       "accounts/12345",
			wantResourceID: "12345",
			wantErr:        false,
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
			gotResourceID, err := ParseAccountExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAccountExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && gotResourceID != tt.wantResourceID {
				t.Errorf("ParseAccountExternal() gotResourceID = %v, want %v", gotResourceID, tt.wantResourceID)
			}
		})
	}
}

func TestNewAccountIdentity(t *testing.T) {
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
					ResourceID: ptrTo("12345"),
				},
			},
			wantID:  "12345",
			wantErr: false,
		},
		{
			name: "spec.resourceID is empty, status.externalRef is set",
			obj: &AnalyticsAccount{
				Status: AnalyticsAccountStatus{
					ExternalRef: ptrTo("accounts/67890"),
				},
			},
			wantID:  "67890",
			wantErr: false,
		},
		{
			name: "spec.resourceID and status.externalRef are both set and consistent",
			obj: &AnalyticsAccount{
				Spec: AnalyticsAccountSpec{
					ResourceID: ptrTo("12345"),
				},
				Status: AnalyticsAccountStatus{
					ExternalRef: ptrTo("accounts/12345"),
				},
			},
			wantID:  "12345",
			wantErr: false,
		},
		{
			name: "spec.resourceID and status.externalRef are inconsistent",
			obj: &AnalyticsAccount{
				Spec: AnalyticsAccountSpec{
					ResourceID: ptrTo("12345"),
				},
				Status: AnalyticsAccountStatus{
					ExternalRef: ptrTo("accounts/67890"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAccountIdentity(ctx, fakeClient, tt.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAccountIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.ID() != tt.wantID {
					t.Errorf("NewAccountIdentity() ID = %v, want %v", got.ID(), tt.wantID)
				}
				if got.String() != "accounts/"+tt.wantID {
					t.Errorf("NewAccountIdentity() String = %v, want` %v", got.String(), "accounts/"+tt.wantID)
				}
			}
		})
	}
}
