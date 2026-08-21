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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &CloudBuildConnectionRef{}

type CloudBuildConnectionRef struct {
	// A reference to an externally managed CloudBuildConnection resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/connections/{{connection}}"
	External string `json:"external,omitempty"`

	// The name of a CloudBuildConnection resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CloudBuildConnection resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&CloudBuildConnectionRef{})
}

func (r *CloudBuildConnectionRef) GetGVK() schema.GroupVersionKind {
	return CloudBuildConnectionGVK
}

func (r *CloudBuildConnectionRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CloudBuildConnectionRef) GetExternal() string {
	return r.External
}

func (r *CloudBuildConnectionRef) SetExternal(external string) {
	r.External = external
}

func (r *CloudBuildConnectionRef) ValidateExternal(ref string) error {
	id := &CloudBuildConnectionIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CloudBuildConnectionRef) ParseExternalToIdentity(ref string) (any, error) {
	id := &CloudBuildConnectionIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return nil, err
	}
	return id, nil
}
func (r *CloudBuildConnectionRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromCloudBuildConnectionSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
