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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
)

func TestAppEngineDomainMappingIdentity_SpecProjectFallback(t *testing.T) {
	ctx := context.Background()
	obj := &AppEngineDomainMapping{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-mapping",
		},
		Spec: AppEngineDomainMappingSpec{
			Project: ptr.To("project-from-spec"),
		},
	}

	id, err := getIdentityFromAppEngineDomainMappingSpec(ctx, nil, obj)
	if err != nil {
		t.Fatalf("unexpected error resolving identity: %v", err)
	}
	if id.App != "project-from-spec" {
		t.Errorf("got app %q, want %q", id.App, "project-from-spec")
	}
	if id.DomainMapping != "test-mapping" {
		t.Errorf("got domainMapping %q, want %q", id.DomainMapping, "test-mapping")
	}
}
