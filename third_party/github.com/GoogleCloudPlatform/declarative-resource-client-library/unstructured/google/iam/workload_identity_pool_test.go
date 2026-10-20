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

package iam

import (
	"testing"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
)

func TestWorkloadIdentityPoolUnstructuredMode(t *testing.T) {
	pool := &dclService.WorkloadIdentityPool{
		Name:        dcl.String("projects/p/locations/global/workloadIdentityPools/pool"),
		DisplayName: dcl.String("Pool"),
		Mode:        dcl.String("TRUST_DOMAIN"),
	}
	u := WorkloadIdentityPoolToUnstructured(pool)
	if mode, ok := u.Object["mode"].(string); !ok || mode != "TRUST_DOMAIN" {
		t.Fatalf("expected mode to be TRUST_DOMAIN, got %v", u.Object["mode"])
	}

	converted, err := UnstructuredToWorkloadIdentityPool(u)
	if err != nil {
		t.Fatalf("unexpected error converting to WorkloadIdentityPool: %v", err)
	}
	if converted.Mode == nil || *converted.Mode != "TRUST_DOMAIN" {
		t.Fatalf("expected converted Mode to be TRUST_DOMAIN, got %v", converted.Mode)
	}
}
