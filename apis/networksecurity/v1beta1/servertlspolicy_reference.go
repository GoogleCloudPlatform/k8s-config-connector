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
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &NetworkSecurityServerTLSPolicyRef{}

var NetworkSecurityServerTLSPolicyGVK = schema.GroupVersionKind{
	Group:   "networksecurity.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "NetworkSecurityServerTLSPolicy",
}

// A reference to a NetworkSecurityServerTLSPolicy resource.
type NetworkSecurityServerTLSPolicyRef struct {
	// Allowed value: string of the format `projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{value}}`, where {{value}} is the `name` field of a `NetworkSecurityServerTLSPolicy` resource.
	External string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

func (r *NetworkSecurityServerTLSPolicyRef) GetGVK() schema.GroupVersionKind {
	return NetworkSecurityServerTLSPolicyGVK
}

func (r *NetworkSecurityServerTLSPolicyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NetworkSecurityServerTLSPolicyRef) GetExternal() string {
	return r.External
}

func (r *NetworkSecurityServerTLSPolicyRef) SetExternal(ref string) {
	r.External = ref
}

func (r *NetworkSecurityServerTLSPolicyRef) ValidateExternal(ref string) error {
	if !strings.HasPrefix(ref, "projects/") {
		return fmt.Errorf("external reference format %q is not known; expected projects/<project>/locations/<location>/serverTlsPolicies/<name>", ref)
	}
	return nil
}

func (r *NetworkSecurityServerTLSPolicyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
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
		if externalRef, err = serverTLSPolicyLegacyExternalRef(ctx, reader, u); err != nil {
			return err
		}
	}
	if externalRef == "" {
		return k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.SetExternal(externalRef)
	return nil
}

func serverTLSPolicyLegacyExternalRef(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (string, error) {
	resourceID, err := refsv1beta1.GetResourceID(u)
	if err != nil {
		return "", err
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return "", err
	}

	location, err := refsv1beta1.GetLocation(u)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("projects/%s/locations/%s/serverTlsPolicies/%s", projectID, location, resourceID), nil
}
