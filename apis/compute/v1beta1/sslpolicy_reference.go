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

var _ refsv1beta1.Ref = &ComputeSSLPolicyRef{}

var ComputeSSLPolicyGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputeSSLPolicy",
}

// A reference to a ComputeSSLPolicy resource.
type ComputeSSLPolicyRef struct {
	// Allowed value: The `selfLink` field of a `ComputeSSLPolicy` resource.
	External string `json:"external,omitempty"`

	// The name of a ComputeSSLPolicy resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeSSLPolicy resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ComputeSSLPolicyRef) GetGVK() schema.GroupVersionKind {
	return ComputeSSLPolicyGVK
}

func (r *ComputeSSLPolicyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeSSLPolicyRef) GetExternal() string {
	return r.External
}

func (r *ComputeSSLPolicyRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeSSLPolicyRef) ValidateExternal(ref string) error {
	return nil
}

func (r *ComputeSSLPolicyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.GetExternal() != "" {
		return r.ValidateExternal(r.GetExternal())
	}
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

	// Get external from status.externalRef. This is the most trustworthy place.
	externalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return fmt.Errorf("reading status.externalRef: %w", err)
	}
	if externalRef == "" {
		if externalRef, err = sslPolicyLegacyExternalRef(ctx, reader, u); err != nil {
			return err
		}
	}
	if externalRef == "" {
		return k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.SetExternal(externalRef)
	return nil
}

func sslPolicyLegacyExternalRef(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (string, error) {
	selfLink, found, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil {
		return "", fmt.Errorf("reading status.selfLink: %w", err)
	}
	if !found || selfLink == "" {
		return "", nil
	}
	return selfLink, nil
}
