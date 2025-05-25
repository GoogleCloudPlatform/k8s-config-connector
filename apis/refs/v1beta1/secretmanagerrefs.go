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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
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
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "secrets" {
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
	resource, err := k8s.NewResource(secret)
	if err != nil {
		return nil, fmt.Errorf("error converting unstructured to resource: %w", err)
	}
	if !k8s.IsResourceReady(resource) {
		return nil, k8s.NewReferenceNotReadyError(secret.GroupVersionKind(), key)
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

func ResolveSecretIDForObject(ctx context.Context, reader client.Reader, obj *unstructured.Unstructured) (*SecretManagerSecret, error) {
	ref := &SecretManagerSecretRef{}
	refObj, found, err := unstructured.NestedMap(obj.Object, "spec", "secretRef")
	if err != nil {
		return nil, fmt.Errorf("getting spec.secretRef: %w", err)
	}
	if !found {
		return nil, fmt.Errorf("spec.secretRef not found in object")
	}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(refObj, ref); err != nil {
		return nil, fmt.Errorf("converting spec.secretRef: %w", err)
	}

	return ResolveSecretManagerSecretRef(ctx, reader, obj, ref)
}

type SecretManagerSecretVersionRef struct {
	//  If provided must be in the format `projects/*/secrets/*/versions/*`.
	External string `json:"external,omitempty"`
	// The `name` field of a `SecretManagerSecretVersion` resource.
	Name string `json:"name,omitempty"`
	// The `metadata.namespace` field of a `SecretManagerSecretVersion` resource.
	Namespace string `json:"namespace,omitempty"`
}

// ResolveSecretManagerSecretVersionRef will resolve a SecretManagerSecretVersionRef to a SecretManagerSecretVersion.
func ResolveSecretManagerSecretVersionRef(ctx context.Context, reader client.Reader, src client.Object, ref *SecretManagerSecretVersionRef) (*SecretManagerSecretVersionRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on SecretManagerSecretVersionRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on SecretManagerSecretVersionRef")
	}

	// External should be in the "projects/*/secrets/*/versions/*" format
	if ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "secrets" && tokens[4] == "versions" {
			ref = &SecretManagerSecretVersionRef{
				External: fmt.Sprintf("projects/%s/secrets/%s/versions/%s", tokens[1], tokens[3], tokens[5]),
			}
			return ref, nil
		}
		return nil, fmt.Errorf("format of secretManagerSecretVersionRef external=%q was not known (use projects/<projectId>/secrets/<secretID>/versions/<versionID>)", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	// Fetch object from k8s cluster to construct the external form
	secretversion := &unstructured.Unstructured{}
	secretversion.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "secretmanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "SecretManagerSecretVersion",
	})
	if err := reader.Get(ctx, key, secretversion); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced SecretManagerSecretVersion %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced SecretManagerSecretVersion %v: %w", key, err)
	}

	secretversionResourceID, err := GetResourceID(secretversion)
	if err != nil {
		return nil, err
	}

	secretversionProjectID, err := ResolveProjectID(ctx, reader, secretversion)
	if err != nil {
		return nil, err
	}

	secretID, err := ResolveSecretIDForObject(ctx, reader, secretversion)
	if err != nil {
		return nil, err
	}

	ref = &SecretManagerSecretVersionRef{
		External: fmt.Sprintf("projects/%s/secrets/%s/versions/%s", secretversionProjectID, secretID.ResourceID, secretversionResourceID),
	}

	return ref, nil
}
