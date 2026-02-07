// Copyright 2025 Google LLC
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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &FirestoreDatabaseRef{}

// FirestoreDatabaseRef is a reference to a FirestoreDatabase resource.
type FirestoreDatabaseRef struct {
	// A reference to an externally managed Firestore database resource.
	// Should be in the format "projects/{{projectID}}/databases/{{databaseID}}".
	External string `json:"external,omitempty"`

	// The name of a FirestoreDatabase resource.
	Name string `json:"name,omitempty"`

	// The namespace of a FirestoreDatabase resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *FirestoreDatabaseRef) GetGVK() schema.GroupVersionKind {
	return FirestoreDatabaseGVK
}

func (r *FirestoreDatabaseRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *FirestoreDatabaseRef) GetExternal() string {
	return r.External
}

func (r *FirestoreDatabaseRef) SetExternal(ref string) {
	r.External = ref
}

func (r *FirestoreDatabaseRef) ValidateExternal(ref string) error {
	id := &FirestoreDatabaseIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *FirestoreDatabaseRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}

func (r *FirestoreDatabaseRef) GetExternalFromCustomFields() []string {
	return nil
}
