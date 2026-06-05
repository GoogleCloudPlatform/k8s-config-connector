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

package refs

import (
	"testing"

	"k8s.io/apimachinery/pkg/types"
)

func TestDataprocMetastoreServiceRef(t *testing.T) {
	ref := &DataprocMetastoreServiceRef{
		Name:      "my-service",
		Namespace: "my-namespace",
	}

	if ref.GetGVK() != DataprocMetastoreServiceGVK {
		t.Errorf("GetGVK() = %v, want %v", ref.GetGVK(), DataprocMetastoreServiceGVK)
	}

	expectedNamespacedName := types.NamespacedName{
		Name:      "my-service",
		Namespace: "my-namespace",
	}
	if ref.GetNamespacedName() != expectedNamespacedName {
		t.Errorf("GetNamespacedName() = %v, want %v", ref.GetNamespacedName(), expectedNamespacedName)
	}

	externalRef := "projects/my-project/locations/us-central1/services/my-service"
	ref.SetExternal(externalRef)
	if ref.GetExternal() != externalRef {
		t.Errorf("GetExternal() = %v, want %v", ref.GetExternal(), externalRef)
	}

	if err := ref.ValidateExternal(externalRef); err != nil {
		t.Errorf("ValidateExternal() unexpected error: %v", err)
	}

	if err := ref.ValidateExternal("invalid-ref"); err == nil {
		t.Errorf("ValidateExternal() expected error on invalid ref, got nil")
	}

	parsedIdentity, err := ref.ParseExternalToIdentity()
	if err != nil {
		t.Errorf("ParseExternalToIdentity() unexpected error: %v", err)
	}

	expectedIdentity := &DataprocMetastoreServiceIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Service:  "my-service",
	}

	actualIdentity, ok := parsedIdentity.(*DataprocMetastoreServiceIdentity)
	if !ok {
		t.Fatalf("ParseExternalToIdentity() returned wrong type %T, want *DataprocMetastoreServiceIdentity", parsedIdentity)
	}

	if actualIdentity.Project != expectedIdentity.Project ||
		actualIdentity.Location != expectedIdentity.Location ||
		actualIdentity.Service != expectedIdentity.Service {
		t.Errorf("ParseExternalToIdentity() = %+v, want %+v", actualIdentity, expectedIdentity)
	}
}
