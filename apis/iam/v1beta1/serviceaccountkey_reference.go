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

package v1beta1

import (
	"context"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &ServiceAccountKeyRef{}

var IAMServiceAccountGVK = schema.GroupVersionKind{
	Group:   "iam.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "IAMServiceAccount",
}

// ServiceAccountKeyRef defines the resource reference to IAMServiceAccountKey, which "External" field
// holds the GCP identifier for the KRM object.
type ServiceAccountKeyRef struct {
	// A reference to an externally managed IAMServiceAccountKey resource.
	// Should be in the format "projects/{{projectID}}/serviceAccounts/{{serviceAccountID}}/keys/{{keyID}}".
	External string `json:"external,omitempty"`

	// The name of an IAMServiceAccountKey resource.
	Name string `json:"name,omitempty"`

	// The namespace of an IAMServiceAccountKey resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ServiceAccountKeyRef) GetGVK() schema.GroupVersionKind {
	return IAMServiceAccountGVK
}

func (r *ServiceAccountKeyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ServiceAccountKeyRef) GetExternal() string {
	return r.External
}

func (r *ServiceAccountKeyRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ServiceAccountKeyRef) ValidateExternal(ref string) error {
	id := &ServiceAccountKeyIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

// NormalizedExternal provision the "External" value for other resource that depends on IAMServiceAccountKey.
// If the "External" is given in the other resource's spec.ServiceAccountKeyRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual IAMServiceAccountKey object from the cluster.
func (r *ServiceAccountKeyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		resourceID, err := refsv1beta1.GetResourceID(u)
		if err != nil || resourceID == "" {
			// ServiceAccountKye is currently non-adoptable.
			// As such we expect the key ref to not work, for now.
			return ""
		}

		externelRef, present, err := unstructured.NestedString(u.Object, "spec", "serviceAccountRef", "external")
		if err != nil {
			return ""
		}
		if present {
			parentIdentity := &ServiceAccountIdentity{}
			if err := parentIdentity.FromExternal(externelRef); err != nil {
				return ""
			}
			identity := ServiceAccountKeyIdentity{
				ServiceAccountIdentity: *parentIdentity,
				Id:                     resourceID,
			}
			return identity.String()
		}

		parentName, present, err := unstructured.NestedString(u.Object, "spec", "serviceAccountRef", "name")
		if err != nil || !present {
			return ""
		}
		parentNamespance, present, err := unstructured.NestedString(u.Object, "spec", "serviceAccountRef", "namespace")
		if err != nil || !present {
			return ""
		}
		key := types.NamespacedName{
			Name:      parentName,
			Namespace: parentNamespance,
		}
		p := &unstructured.Unstructured{}
		p.SetGroupVersionKind(IAMServiceAccountGVK)
		if err := reader.Get(ctx, key, p); err != nil {
			if apierrors.IsNotFound(err) {
				return ""
			}
			return ""
		}
		project, err := refsv1beta1.ResolveProjectID(ctx, reader, p)
		if err != nil {
			return ""
		}
		account, err := refsv1beta1.GetResourceID(p)
		if err != nil {
			return ""
		}
		parentIdentity := &ServiceAccountIdentity{
			Project: project,
			Account: account,
		}
		identity := ServiceAccountKeyIdentity{
			ServiceAccountIdentity: *parentIdentity,
			Id:                     resourceID,
		}
		return identity.String()
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
