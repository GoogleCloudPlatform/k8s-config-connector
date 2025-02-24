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
	"fmt"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &ApigeeOrganizationRef{}

// ApigeeOrganizationRef is a reference to a ApigeeOrganization resource.
type ApigeeOrganizationRef struct {
	// A reference to an externally managed ApigeeOrganization resource.
	// Should be in the format "organizations/{{organizationID}}".
	External string `json:"external,omitempty"`

	// The name of a ApigeeOrganization resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ApigeeOrganization resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ApigeeOrganizationRef) GetGVK() schema.GroupVersionKind {
	return ApigeeOrganizationGVK
}

func (r *ApigeeOrganizationRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ApigeeOrganizationRef) GetExternal() string {
	return r.External
}

func (r *ApigeeOrganizationRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ApigeeOrganizationRef) ValidateExternal(ref string) error {
	id := &ApigeeOrganizationIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *ApigeeOrganizationRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	// TODO: Use general-purpose refsv1beta1.Normalize function once direct controller is implemented.
	// For now, we can build the external reference by reading status.ProjectId.
	if r.GetExternal() == "" {
		key := r.GetNamespacedName()
		if key.Namespace == "" {
			key.Namespace = defaultNamespace
		}
		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(r.GetGVK())
		if err := reader.Get(ctx, key, u); err != nil {
			if apierrors.IsNotFound(err) {
				return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
			}
			return fmt.Errorf("reading referenced %s %s: %w", r.GetGVK(), key, err)
		}

		projectID, _, err := unstructured.NestedString(u.Object, "status", "projectId")
		if err != nil {
			return fmt.Errorf("reading status.externalRef: %w", err)
		}
		if projectID == "" {
			return k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
		}
		r.SetExternal(ApigeeOrganizationIDToken + "/" + projectID)
	}

	id := &ApigeeOrganizationIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}
