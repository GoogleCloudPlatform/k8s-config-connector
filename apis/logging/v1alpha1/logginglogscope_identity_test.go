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

	billingv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestLoggingLogScopeIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name       string
		ref        string
		wantErr    bool
		want       *LoggingLogScopeIdentity
		wantString string
		wantParent string
	}{
		{
			name:       "project scope",
			ref:        "projects/test-project/locations/global/logScopes/my-scope",
			want:       &LoggingLogScopeIdentity{Project: "test-project", Location: "global", LogScope: "my-scope"},
			wantString: "projects/test-project/locations/global/logScopes/my-scope",
			wantParent: "projects/test-project/locations/global",
		},
		{
			name:       "folder scope",
			ref:        "folders/123456/locations/global/logScopes/my-scope",
			want:       &LoggingLogScopeIdentity{Folder: "123456", Location: "global", LogScope: "my-scope"},
			wantString: "folders/123456/locations/global/logScopes/my-scope",
			wantParent: "folders/123456/locations/global",
		},
		{
			name:       "organization scope",
			ref:        "organizations/789012/locations/global/logScopes/my-scope",
			want:       &LoggingLogScopeIdentity{Organization: "789012", Location: "global", LogScope: "my-scope"},
			wantString: "organizations/789012/locations/global/logScopes/my-scope",
			wantParent: "organizations/789012/locations/global",
		},
		{
			name:       "billing account scope",
			ref:        "billingAccounts/012345-678901-ABCDEF/locations/global/logScopes/my-scope",
			want:       &LoggingLogScopeIdentity{BillingAccount: "012345-678901-ABCDEF", Location: "global", LogScope: "my-scope"},
			wantString: "billingAccounts/012345-678901-ABCDEF/locations/global/logScopes/my-scope",
			wantParent: "billingAccounts/012345-678901-ABCDEF/locations/global",
		},
		{
			name:    "invalid format",
			ref:     "invalid/format",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &LoggingLogScopeIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.String(); got != tt.wantString {
					t.Errorf("String() = %q, want %q", got, tt.wantString)
				}
				if got := i.ParentString(); got != tt.wantParent {
					t.Errorf("ParentString() = %q, want %q", got, tt.wantParent)
				}
			}
		})
	}
}

func TestLoggingLogScopeIdentity_GetIdentityFromSpec(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	_ = billingv1alpha1.AddToScheme(scheme)
	client := fake.NewClientBuilder().WithScheme(scheme).Build()

	loc := "global"
	resourceID := "test-scope"

	t.Run("projectRef specified", func(t *testing.T) {
		obj := &LoggingLogScope{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-scope",
				Namespace: "default",
			},
			Spec: LoggingLogScopeSpec{
				ProjectRef: &refsv1beta1.ProjectRef{External: "spec-project"},
				Location:   &loc,
				ResourceID: &resourceID,
			},
		}
		id, err := getIdentityFromLoggingLogScopeSpec(context.Background(), client, obj)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id.Project != "spec-project" || id.Location != "global" || id.LogScope != "test-scope" {
			t.Errorf("unexpected identity: %+v", id)
		}
	})

	t.Run("folderRef specified", func(t *testing.T) {
		obj := &LoggingLogScope{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-scope",
				Namespace: "default",
			},
			Spec: LoggingLogScopeSpec{
				FolderRef:  &refsv1beta1.FolderRef{External: "folders/123456"},
				Location:   &loc,
				ResourceID: &resourceID,
			},
		}
		id, err := getIdentityFromLoggingLogScopeSpec(context.Background(), client, obj)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id.Folder != "123456" || id.Location != "global" || id.LogScope != "test-scope" {
			t.Errorf("unexpected identity: %+v", id)
		}
	})

	t.Run("organizationRef specified", func(t *testing.T) {
		obj := &LoggingLogScope{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-scope",
				Namespace: "default",
			},
			Spec: LoggingLogScopeSpec{
				OrganizationRef: &refsv1beta1.OrganizationRef{External: "organizations/789012"},
				Location:        &loc,
				ResourceID:      &resourceID,
			},
		}
		id, err := getIdentityFromLoggingLogScopeSpec(context.Background(), client, obj)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id.Organization != "789012" || id.Location != "global" || id.LogScope != "test-scope" {
			t.Errorf("unexpected identity: %+v", id)
		}
	})

	t.Run("billingAccountRef specified", func(t *testing.T) {
		obj := &LoggingLogScope{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-scope",
				Namespace: "default",
			},
			Spec: LoggingLogScopeSpec{
				BillingAccountRef: &billingv1alpha1.BillingAccountRef{External: "billingAccounts/012345-678901-ABCDEF"},
				Location:          &loc,
				ResourceID:        &resourceID,
			},
		}
		id, err := getIdentityFromLoggingLogScopeSpec(context.Background(), client, obj)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id.BillingAccount != "012345-678901-ABCDEF" || id.Location != "global" || id.LogScope != "test-scope" {
			t.Errorf("unexpected identity: %+v", id)
		}
	})

	t.Run("multiple parents error", func(t *testing.T) {
		obj := &LoggingLogScope{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-scope",
				Namespace: "default",
			},
			Spec: LoggingLogScopeSpec{
				ProjectRef: &refsv1beta1.ProjectRef{External: "spec-project"},
				FolderRef:  &refsv1beta1.FolderRef{External: "123456"},
				Location:   &loc,
				ResourceID: &resourceID,
			},
		}
		_, err := getIdentityFromLoggingLogScopeSpec(context.Background(), client, obj)
		if err == nil {
			t.Fatal("expected error with multiple parents specified, got nil")
		}
	})
}
