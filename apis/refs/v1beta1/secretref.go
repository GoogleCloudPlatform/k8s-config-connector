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
	"strings"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SecretManagerSecretRef struct {
	//  If provided must be in the format `projects/*/secrets/*`.
	External string `json:"external,omitempty"`
	// The `name` field of a `SecretManagerSecret` resource.
	Name string `json:"name,omitempty"`
	// The `metadata.namespace` field of a `SecretManagerSecret` resource.
	Namespace string `json:"namespace,omitempty"`
}

type SecretManagerSecret struct {
	Ref        *SecretManagerSecretRef
	ResourceID string
}

// ResolveSecretManagerSecretRef will resolve a partial SecretManagerSecretRef to a SecretManagerSecret.
func ResolveSecretManagerSecretRef(ctx context.Context, reader client.Reader, src client.Object, ref *SecretManagerSecretRef) (*SecretManagerSecret, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on SecretManagerSecretRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on SecretManagerSecretRef")
	}

	// External should be in the "projects/*/secrets/*" format
	if ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 4 && tokens[0] == "project" && tokens[2] == "secrets" {
			ref = &SecretManagerSecretRef{
				External: fmt.Sprintf("projects/%s/secrets/%s", tokens[1], tokens[3]),
			}
			return &SecretManagerSecret{Ref: ref, ResourceID: tokens[3]}, nil
		}
		return nil, fmt.Errorf("format of secretManagerSecretRef external=%q was not known (use projects/<projectId>/secrets/<secretID>)", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	// Fetch object from k8s cluster to construct the external form
	secret := &unstructured.Unstructured{}
	secret.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "secretmanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "SecretManagerSecret",
	})
	if err := reader.Get(ctx, key, secret); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced SecretManagerSecret %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced SecretManagerSecret %v: %w", key, err)
	}

	secretResourceID, err := GetResourceID(secret)
	if err != nil {
		return nil, err
	}

	secretProjectID, err := ResolveProjectID(ctx, reader, secret)
	if err != nil {
		return nil, err
	}

	ref = &SecretManagerSecretRef{
		External: fmt.Sprintf("projects/%s/secrets/%s", secretProjectID, secretResourceID),
	}

	return &SecretManagerSecret{Ref: ref, ResourceID: secretResourceID}, nil
}

func ResolveSecretIDForObject(ctx context.Context, reader client.Reader, obj *unstructured.Unstructured) (string, error) {
	secretRefExternal, _, err := unstructured.NestedString(obj.Object, "spec", "secretRef", "external")
	if err != nil {
		return "", fmt.Errorf("error fetching secretRef.external %w", err)
	}
	if secretRefExternal != "" {
		secretRef := SecretManagerSecretRef{
			External: secretRefExternal,
		}

		secret, err := ResolveSecretManagerSecretRef(ctx, reader, obj, &secretRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse secretRef.external %q in %v %v/%v: %w", secretRefExternal, obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return secret.ResourceID, nil
	}

	secretRefName, _, err := unstructured.NestedString(obj.Object, "spec", "secretRef", "name")
	if err != nil {
		return "", fmt.Errorf("error fetching secretRef.name %w", err)
	}
	if secretRefName != "" {
		secretRefNamespace, _, err := unstructured.NestedString(obj.Object, "spec", "secretRef", "namespace")
		if err != nil {
			return "", fmt.Errorf("error fetching secretRef.namespace %w", err)
		}

		secretRef := SecretManagerSecretRef{
			Name:      secretRefName,
			Namespace: secretRefNamespace,
		}
		if secretRef.Namespace == "" {
			secretRef.Namespace = obj.GetNamespace()
		}

		secret, err := ResolveSecretManagerSecretRef(ctx, reader, obj, &secretRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse secretRef in %v %v/%v: %w", obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return secret.ResourceID, nil
	}

	return "", fmt.Errorf("cannot find secret for %v %v/%v", obj.GetKind(), obj.GetNamespace(), obj.GetName())
}
