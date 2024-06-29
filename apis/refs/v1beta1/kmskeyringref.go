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

// +reference-type
// apiVersion: kms.cnrm.cloud.google.com/v1beta1
// kind: KMSKeyRing
// external: projects/{project}/locations/{location}/keyRings/{name}

type KMSKeyRingRef struct {
	/* The keyRing selflink of form "projects/{project}/locations/{location}/keyRings/{name}", when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `KMSKeyRing` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `KMSKeyRing` resource. */
	Namespace string `json:"namespace,omitempty"`
}

func (ref *KMSKeyRingRef) NormalizeReference(ctx context.Context, reader client.Reader, src client.Object, project Project) error {
	if ref == nil {
		return nil
	}

	if ref.Name == "" && ref.External == "" {
		return fmt.Errorf("must specify either name or external on reference")
	}
	if ref.Name != "" && ref.External != "" {
		return fmt.Errorf("cannot specify both name and external on reference")
	}

	if ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "keyRings" {
			*ref = KMSKeyRingRef{
				External: fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", tokens[1], tokens[3], tokens[5]),
			}
			return nil
		}
		return fmt.Errorf("format of kmsKeyRingRef external=%q was not known (use projects/{project}/locations/{location}/keyRings/{name})", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	referencedObj := &unstructured.Unstructured{}
	referencedObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "kms.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "KMSKeyRing",
	})
	if err := reader.Get(ctx, key, referencedObj); err != nil {
		if apierrors.IsNotFound(err) {
			return fmt.Errorf("referenced KMSKeyRing %v not found", key)
		}
		return fmt.Errorf("error reading referenced KMSKeyRing %v: %w", key, err)
	}

	referencedID, err := GetResourceID(referencedObj)
	if err != nil {
		return err
	}

	referencedLocation, err := GetLocation(referencedObj)
	if err != nil {
		return err
	}

	referencedProjectID, err := ResolveProjectID(ctx, reader, referencedObj)
	if err != nil {
		return err
	}

	*ref = KMSKeyRingRef{
		External: fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", referencedProjectID, referencedLocation, referencedID),
	}
	return nil
}
