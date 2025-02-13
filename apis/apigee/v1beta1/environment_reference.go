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

var _ refsv1beta1.Ref = &ApigeeEnvironmentRef{}

// ApigeeEnvironmentRef is a reference to a ApigeeEnvironment resource.
type ApigeeEnvironmentRef struct {
	// A reference to an externally managed ApigeeEnvironment resource.
	// Should be in the format "organizations/{{organizationID}}/environments/{{environmentID}}".
	External string `json:"external,omitempty"`

	// The name of a ApigeeEnvironment resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ApigeeEnvironment resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ApigeeEnvironmentRef) GetGVK() schema.GroupVersionKind {
	return ApigeeEnvironmentGVK
}

func (r *ApigeeEnvironmentRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ApigeeEnvironmentRef) GetExternal() string {
	return r.External
}

func (r *ApigeeEnvironmentRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ApigeeEnvironmentRef) ValidateExternal(ref string) error {
	id := &ApigeeEnvironmentIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *ApigeeEnvironmentRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	// TODO: Use general-purpose refsv1beta1.Normalize function once direct controller is implemented.
	// For now, we can build the external reference by reading spec fields.
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

		orgName, _, err := unstructured.NestedString(u.Object, "spec", "apigeeOrganizationRef", "name")
		if err != nil {
			return fmt.Errorf("reading spec.apigeeOrganizationRef.name: %w", err)
		}
		orgNamespace, _, err := unstructured.NestedString(u.Object, "spec", "apigeeOrganizationRef", "namespace")
		if err != nil {
			return fmt.Errorf("reading spec.apigeeOrganizationRef.namespace: %w", err)
		}
		orgExternal, _, err := unstructured.NestedString(u.Object, "spec", "apigeeOrganizationRef", "external")
		if err != nil {
			return fmt.Errorf("reading spec.apigeeOrganizationRef.external: %w", err)
		}
		// Normalize OrganizationRef
		orgRef := ApigeeOrganizationRef{
			Name:      orgName,
			Namespace: orgNamespace,
			External:  orgExternal,
		}
		if err := orgRef.Normalize(ctx, reader, defaultNamespace); err != nil {
			if k8s.IsReferenceNotReadyError(err) {
				return err
			}
			return fmt.Errorf("failed to normalize org ref: %w", err)
		}
		// Build EnvironmentID
		resourceID, _, err := unstructured.NestedString(u.Object, "spec", "resourceID")
		if err != nil {
			return fmt.Errorf("reading spec.resourceID: %w", err)
		}
		envID := resourceID
		if envID == "" {
			metadataName, _, err := unstructured.NestedString(u.Object, "metadata", "name")
			if err != nil {
				return fmt.Errorf("reading metadata.name: %w", err)
			}
			envID = metadataName
		}
		if envID == "" {
			return k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
		}
		r.SetExternal(orgRef.External + "/" + EnvironmentIDToken + "/" + envID)
	}

	id := &ApigeeEnvironmentIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}
