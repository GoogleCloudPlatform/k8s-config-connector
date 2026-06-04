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
)

func TestBillingAccountIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *BillingAccountIdentity
	}{
		{
			name: "valid reference",
			ref:  "billingAccounts/012345-567890-ABCDEF",
			want: &BillingAccountIdentity{
				BillingAccount: "012345-567890-ABCDEF",
			},
		},
		{
			name: "raw ID format",
			ref:  "012345-567890-ABCDEF",
			want: &BillingAccountIdentity{
				BillingAccount: "012345-567890-ABCDEF",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://cloudbilling.googleapis.com/billingAccounts/012345-567890-ABCDEF",
			want: &BillingAccountIdentity{
				BillingAccount: "012345-567890-ABCDEF",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BillingAccountIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.BillingAccount != tt.want.BillingAccount {
					t.Errorf("BillingAccount = %v, want %v", i.BillingAccount, tt.want.BillingAccount)
				}
			}
		})
	}
}
