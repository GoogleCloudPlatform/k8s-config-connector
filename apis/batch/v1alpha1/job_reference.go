// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"context"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &BatchJobRef{}

type BatchJobRef struct {
	// External represents a fully-qualified GCP resource name.
	//   Format: projects/{{project}}/locations/{{location}}/jobs/{{value}}
	// +optional
	External string `json:"external,omitempty"`
	// Name represents the name of the GCP resource.
	// +optional
	Name string `json:"name,omitempty"`
	// Namespace represents the namespace of the GCP resource.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

func (r *BatchJobRef) SetExternal(ref string) {
	r.External = ref
}

func (r *BatchJobRef) ValidateExternal(ref string) error {
	id := &JobIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *BatchJobRef) GetGVK() schema.GroupVersionKind {
	return BatchJobGVK
}

func (r *BatchJobRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *BatchJobRef) IsExternal() bool {
	return r.External != ""
}

func (r *BatchJobRef) GetExternal() string {
	return r.External
}

func (r *BatchJobRef) String() string {
	if r.External != "" {
		return r.External
	}
	return r.Name
}

func (r *BatchJobRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}
