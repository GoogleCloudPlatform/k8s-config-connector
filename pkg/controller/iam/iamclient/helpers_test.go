// Copyright 2022 Google LLC
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

package iamclient

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
)

func TestResolveMemberIdentity(t *testing.T) {
	ctx := context.TODO()

	tests := []struct {
		name       string
		member     v1beta1.Member
		memberFrom *v1beta1.MemberSource
		wantID     string
		wantErr    bool
	}{
		{
			name:   "service account string",
			member: "serviceAccount:test-sa@project.iam.gserviceaccount.com",
			wantID: "serviceAccount:test-sa@project.iam.gserviceaccount.com",
		},
		{
			name:   "user string",
			member: "user:test-user@example.com",
			wantID: "user:test-user@example.com",
		},
		{
			name:   "service account with project number",
			member: "serviceAccount:123456789012@cloudbuild.gserviceaccount.com",
			wantID: "serviceAccount:123456789012@cloudbuild.gserviceaccount.com",
		},
		{
			name:    "both empty",
			member:  "",
			wantErr: true,
		},
		{
			name:       "both set",
			member:     "user:foo",
			memberFrom: &v1beta1.MemberSource{},
			wantErr:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id, err := ResolveMemberIdentity(ctx, tc.member, tc.memberFrom, "default", nil)
			if tc.wantErr {
				if err == nil {
					t.Errorf("ResolveMemberIdentity() error = nil, wantErr %v", tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("ResolveMemberIdentity() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if id != tc.wantID {
				t.Errorf("ResolveMemberIdentity() = %v, want %v", id, tc.wantID)
			}
		})
	}
}
