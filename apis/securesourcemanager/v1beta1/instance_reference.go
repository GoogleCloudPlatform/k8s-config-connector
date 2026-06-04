// Copyright 2024 Google LLC
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
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &SecureSourceManagerInstanceRef{}

// SecureSourceManagerInstanceRef is a reference to a SecureSourceManagerInstance.
type SecureSourceManagerInstanceRef struct {
	// A reference to an externally managed SecureSourceManagerInstance resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/instances/{{instanceID}}".
	External string `json:"external,omitempty"`

	// The name of a SecureSourceManagerInstance resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SecureSourceManagerInstance resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&SecureSourceManagerInstanceRef{})
}

func (r *SecureSourceManagerInstanceRef) GetGVK() schema.GroupVersionKind {
	return SecureSourceManagerInstanceGVK
}

func (r *SecureSourceManagerInstanceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *SecureSourceManagerInstanceRef) GetExternal() string {
	return r.External
}

func (r *SecureSourceManagerInstanceRef) SetExternal(ref string) {
	r.External = ref
}

func (r *SecureSourceManagerInstanceRef) ValidateExternal(ref string) error {
	id := &SecureSourceManagerInstanceIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *SecureSourceManagerInstanceRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &SecureSourceManagerInstanceIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *SecureSourceManagerInstanceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromSecureSourceManagerInstanceSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// ConvertToProjectNumber converts the external reference to use a project number.
func (r *SecureSourceManagerInstanceRef) ConvertToProjectNumber(ctx context.Context, projectMapper *projects.ProjectMapper) error {
	if r == nil {
		return nil
	}

	id := &SecureSourceManagerInstanceIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return fmt.Errorf("parsing reference %q: %w", r.External, err)
	}
	projectNumber, err := projectMapper.LookupProjectNumber(ctx, id.Project)
	if err != nil {
		return fmt.Errorf("error looking up project number for project %q: %w", id.Project, err)
	}
	id.Project = strconv.FormatInt(projectNumber, 10)
	r.External = id.String()
	return nil
}
