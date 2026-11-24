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
	"strings"
	"testing"
)

func TestFirestoreIndexIdentityUsesResourceIDForAcquisition(t *testing.T) {
	resourceID := "CICAgOjXh4EK"
	obj := &FirestoreIndex{
		Spec: FirestoreIndexSpec{
			Collection: "Customer",
			Database:   ptrTo("my-database"),
			ResourceID: &resourceID,
		},
	}
	obj.SetName("customers-by-email")
	obj.SetAnnotations(map[string]string{"cnrm.cloud.google.com/project-id": "my-project"})

	id, err := obj.GetIdentity(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetIdentity() returned error: %v", err)
	}

	got := id.(*FirestoreIndexIdentity).String()
	want := "projects/my-project/databases/my-database/collectionGroups/Customer/indexes/CICAgOjXh4EK"
	if got != want {
		t.Fatalf("GetIdentity() = %q, want %q", got, want)
	}
}

func TestFirestoreIndexIdentityRejectsResourceIDChange(t *testing.T) {
	resourceID := "new-index-id"
	obj := &FirestoreIndex{
		Spec: FirestoreIndexSpec{
			Collection: "Customer",
			Database:   ptrTo("my-database"),
			ResourceID: &resourceID,
		},
		Status: FirestoreIndexStatus{
			Name: ptrTo("projects/my-project/databases/my-database/collectionGroups/Customer/indexes/existing-index-id"),
		},
	}
	obj.SetAnnotations(map[string]string{"cnrm.cloud.google.com/project-id": "my-project"})

	_, err := obj.GetIdentity(context.Background(), nil)
	if err == nil {
		t.Fatal("GetIdentity() succeeded, want resourceID change error")
	}
	if !strings.Contains(err.Error(), "cannot change FirestoreIndex resourceID") {
		t.Fatalf("GetIdentity() error = %q, want resourceID change error", err)
	}
}

func ptrTo[T any](v T) *T {
	return &v
}
