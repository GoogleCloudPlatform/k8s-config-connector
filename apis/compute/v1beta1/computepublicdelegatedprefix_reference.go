/*
Copyright 2026 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"

	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/reference"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &ComputePublicDelegatedPrefixRef{}

type ComputePublicDelegatedPrefixRef struct {
	// The value of an externally managed ComputePublicDelegatedPrefix resource.
	// Should be in the format "projects/{{projectId}}/regions/{{region}}/publicDelegatedPrefixes/{{publicDelegatedPrefixId}}" or "projects/{{projectId}}/global/publicDelegatedPrefixes/{{publicDelegatedPrefixId}}"
	External string `json:"external,omitempty"`

	// The name of a ComputePublicDelegatedPrefix resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputePublicDelegatedPrefix resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ComputePublicDelegatedPrefixRef) GetGVK() schema.GroupVersionKind {
	return ComputePublicDelegatedPrefixGVK
}

func (r *ComputePublicDelegatedPrefixRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputePublicDelegatedPrefixRef) GetExternal() string {
	return r.External
}

func (r *ComputePublicDelegatedPrefixRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ComputePublicDelegatedPrefixRef) ValidateExternal(ref string) error {
	id := &PublicDelegatedPrefixIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *ComputePublicDelegatedPrefixRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		id := &PublicDelegatedPrefixIdentity{}
		if err := id.FromExternal(r.External); err != nil {
			return err
		}
		r.External = id.String()
		return nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = defaultNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ComputePublicDelegatedPrefixGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return fmt.Errorf("reading referenced %s %s: %w", ComputePublicDelegatedPrefixGVK, key, err)
	}

	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef != "" {
		r.External = actualExternalRef
		return nil
	}

	// Get external from status.selfLink. This ensures backward compatibility for TF/DCL-based resources that lack status.externalRef.
	selfLink, _, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil {
		return fmt.Errorf("reading status.selfLink: %w", err)
	}
	if selfLink == "" {
		return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
	}

	// Parse selfLink to identity to ensure it is in consistent format
	// selfLink format: https://www.googleapis.com/compute/v1/projects/...
	// We want to extract Identity: projects/...
	// common.FixStaleComputeExternalFormat might work if it handles generic compute URL?
	// But PublicDelegatedPrefix identity is standard.
	// We can use FromExternal if we strip prefix?
	// Or use common generic parser?

	// Check if common.FixStaleComputeExternalFormat handles it.
	// It basically strips https://.../v1/
	external := common.FixStaleComputeExternalFormat(selfLink)
	id := &PublicDelegatedPrefixIdentity{}
	if err := id.FromExternal(external); err != nil {
		return fmt.Errorf("parsing selfLink %q: %w", selfLink, err)
	}

	r.External = id.String()
	return nil
}

func ParseComputePublicDelegatedPrefixExternal(external string) (*PublicDelegatedPrefixIdentity, error) {
	id := &PublicDelegatedPrefixIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}
