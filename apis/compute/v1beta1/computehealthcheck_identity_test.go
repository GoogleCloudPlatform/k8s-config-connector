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

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestComputeHealthCheckIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeHealthCheckIdentity
	}{
		{
			name: "valid regional reference",
			ref:  "projects/my-project/regions/us-central1/healthChecks/my-healthcheck",
			want: &ComputeHealthCheckIdentity{
				Project:     "my-project",
				Region:      "us-central1",
				HealthCheck: "my-healthcheck",
			},
		},
		{
			name: "valid global reference",
			ref:  "projects/my-project/global/healthChecks/my-healthcheck",
			want: &ComputeHealthCheckIdentity{
				Project:     "my-project",
				Region:      "global",
				HealthCheck: "my-healthcheck",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full regional url",
			ref:  "https://compute.googleapis.com/projects/my-project/regions/us-central1/healthChecks/my-healthcheck",
			want: &ComputeHealthCheckIdentity{
				Project:     "my-project",
				Region:      "us-central1",
				HealthCheck: "my-healthcheck",
			},
		},
		{
			name: "full global url",
			ref:  "https://compute.googleapis.com/projects/my-project/global/healthChecks/my-healthcheck",
			want: &ComputeHealthCheckIdentity{
				Project:     "my-project",
				Region:      "global",
				HealthCheck: "my-healthcheck",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeHealthCheckIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project {
					t.Errorf("Project = %v, want %v", i.Project, tt.want.Project)
				}
				if i.Region != tt.want.Region {
					t.Errorf("Region = %v, want %v", i.Region, tt.want.Region)
				}
				if i.HealthCheck != tt.want.HealthCheck {
					t.Errorf("HealthCheck = %v, want %v", i.HealthCheck, tt.want.HealthCheck)
				}
			}
		})
	}
}

func TestComputeHealthCheck_GetIdentity(t *testing.T) {
	ctx := context.Background()
	var reader client.Reader = nil // not needed since project-id annotation is used

	// Test 1: Successful global ComputeHealthCheck GetIdentity
	hcGlobal := &ComputeHealthCheck{}
	hcGlobal.SetName("my-global-hc")
	hcGlobal.SetAnnotations(map[string]string{
		"cnrm.cloud.google.com/project-id": "my-project",
	})
	hcGlobal.Spec.Location = "global"

	id, err := hcGlobal.GetIdentity(ctx, reader)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantStr := "projects/my-project/global/healthChecks/my-global-hc"
	if id.String() != wantStr {
		t.Errorf("GetIdentity() String = %q, want %q", id.String(), wantStr)
	}

	// Test 2: Successful regional ComputeHealthCheck GetIdentity
	hcRegional := &ComputeHealthCheck{}
	hcRegional.SetName("my-regional-hc")
	hcRegional.SetAnnotations(map[string]string{
		"cnrm.cloud.google.com/project-id": "my-project",
	})
	hcRegional.Spec.Location = "us-central1"

	idReg, err := hcRegional.GetIdentity(ctx, reader)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantStrReg := "projects/my-project/regions/us-central1/healthChecks/my-regional-hc"
	if idReg.String() != wantStrReg {
		t.Errorf("GetIdentity() String = %q, want %q", idReg.String(), wantStrReg)
	}

	// Test 3: GetIdentity with custom resourceID
	customID := "custom-hc-id"
	hcCustom := &ComputeHealthCheck{}
	hcCustom.SetName("my-hc")
	hcCustom.SetAnnotations(map[string]string{
		"cnrm.cloud.google.com/project-id": "my-project",
	})
	hcCustom.Spec.Location = "global"
	hcCustom.Spec.ResourceID = &customID

	idCustom, err := hcCustom.GetIdentity(ctx, reader)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantStrCustom := "projects/my-project/global/healthChecks/custom-hc-id"
	if idCustom.String() != wantStrCustom {
		t.Errorf("GetIdentity() String = %q, want %q", idCustom.String(), wantStrCustom)
	}

	// Test 4: GetIdentity failure cases
	// A: Missing location
	hcNoLoc := &ComputeHealthCheck{}
	hcNoLoc.SetName("my-hc")
	hcNoLoc.SetAnnotations(map[string]string{
		"cnrm.cloud.google.com/project-id": "my-project",
	})
	_, err = hcNoLoc.GetIdentity(ctx, reader)
	if err == nil {
		t.Error("expected error for missing location, got nil")
	}

	// B: Missing project ID
	hcNoProj := &ComputeHealthCheck{}
	hcNoProj.SetName("my-hc")
	hcNoProj.Spec.Location = "global"
	_, err = hcNoProj.GetIdentity(ctx, reader)
	if err == nil {
		t.Error("expected error for missing project, got nil")
	}

	// Test 5: SelfLink mismatch error
	hcMismatch := &ComputeHealthCheck{}
	hcMismatch.SetName("my-hc")
	hcMismatch.SetAnnotations(map[string]string{
		"cnrm.cloud.google.com/project-id": "my-project",
	})
	hcMismatch.Spec.Location = "global"
	mismatchedSelfLink := "projects/my-project/global/healthChecks/different-hc"
	hcMismatch.Status.SelfLink = &mismatchedSelfLink
	_, err = hcMismatch.GetIdentity(ctx, reader)
	if err == nil {
		t.Error("expected error for SelfLink mismatch, got nil")
	}
}
