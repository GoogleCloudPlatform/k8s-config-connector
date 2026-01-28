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

	"k8s.io/apimachinery/pkg/runtime/schema"

	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/reference"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &ComputeSubnetworkRef{}

type ComputeSubnetworkRef struct {
	// The value of an externally managed ComputeSubnetwork resource.
	// Should be in the format "https://www.googleapis.com/compute/{{version}}/projects/{{projectId}}/regions/{{region}}/subnetworks/{{subnetworkId}}" or "projects/{{projectId}}/regions/{{region}}/subnetworks/{{subnetworkId}}"
	External string `json:"external,omitempty"`

	// The `name` field of a `ComputeSubnetwork` resource.
	Name string `json:"name,omitempty"`

	// The `namespace` field of a `ComputeSubnetwork` resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ComputeSubnetworkRef) GetGVK() schema.GroupVersionKind {
	return ComputeSubnetworkGVK
}

func (r *ComputeSubnetworkRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeSubnetworkRef) GetExternal() string {
	return r.External
}

func (r *ComputeSubnetworkRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ComputeSubnetworkRef) ValidateExternal(ref string) error {
	id := &SubnetworkIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *ComputeSubnetworkRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		_, err := ParseComputeSubnetworkExternal(r.External)
		if err != nil {
			return err
		}
		external := common.FixStaleComputeExternalFormat(r.External)
		r.External = external
		return nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = defaultNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ComputeSubnetworkGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return fmt.Errorf("reading referenced %s %s: %w", ComputeSubnetworkGVK, key, err)
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

	external := common.FixStaleComputeExternalFormat(selfLink)
	r.External = external
	return nil
}
