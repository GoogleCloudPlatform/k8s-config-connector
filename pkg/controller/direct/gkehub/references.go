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

package gkehub

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Membership struct {
	id string
}

type Feature struct {
	id string
}

// resolveMembershipRef returns a membership that has membershipId as "projects/*/locations/*/memberships/{membershipId}".
func resolveMembershipRef(ctx context.Context, reader client.Reader, obj *krm.GKEHubFeatureMembership, projectID string) (*Membership, error) {
	ref := &obj.Spec.MembershipRef
	if err := ref.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}

	return &Membership{id: ref.External}, nil
}

// resolveFeatureRef returns a feature that has featureID as "projects/*/locations/*/features/{featureId}".
func resolveFeatureRef(ctx context.Context, reader client.Reader, obj *krm.GKEHubFeatureMembership, projectID string) (*Feature, error) {
	name := obj.Spec.FeatureRef.Name
	namespace := obj.Spec.FeatureRef.Namespace
	external := obj.Spec.FeatureRef.External

	if external != "" {
		if name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on feature reference")
		}

		tokens := strings.Split(external, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "features" {
			return &Feature{id: external}, nil
		}
		return nil, fmt.Errorf("format of feature external=%q was not known (use projects/*/locations/*/features/{featureId})", external)
	}

	if name == "" {
		return nil, fmt.Errorf("must specify either name or external on feature reference")
	}

	key := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	feature := &unstructured.Unstructured{}
	feature.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "gkehub.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "GKEHubFeature",
	})
	if err := reader.Get(ctx, key, feature); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced feature %v: %w", key, err)
	}

	featureName, _, err := unstructured.NestedString(feature.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from feature %v: %w", key, err)
	}
	if featureName == "" {
		featureName = feature.GetName()
	}
	featureLocation := obj.Spec.Location
	if featureLocation == "" {
		featureLocation = "global"
	}
	return &Feature{
		id: fmt.Sprintf("projects/%s/locations/%s/features/%s", projectID, featureLocation, featureName),
	}, nil
}
