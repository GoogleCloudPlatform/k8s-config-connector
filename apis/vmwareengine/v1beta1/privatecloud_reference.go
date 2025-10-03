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

var VMwareEnginePrivateCloudGVK = schema.GroupVersionKind{
	Group:   "vmwareengine.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "VMwareEnginePrivateCloud",
}

var _ refsv1beta1.ExternalNormalizer = &PrivateCloudRef{}

// PrivateCloudRef defines the resource reference to VMwareEnginePrivateCloud, which "External" field
// holds the GCP identifier for the KRM object.
type PrivateCloudRef struct {
	// A reference to an externally managed VMwareEnginePrivateCloud resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/privateClouds/{{privatecloudID}}".
	External string `json:"external,omitempty"`

	// The name of a VMwareEnginePrivateCloud resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VMwareEnginePrivateCloud resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on VMwareEnginePrivateCloud.
// If the "External" is given in the other resource's spec.VMwareEnginePrivateCloudRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual VMwareEnginePrivateCloud object from the cluster.
func (r *PrivateCloudRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", VMwareEnginePrivateCloudGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := ParsePrivateCloudExternal(r.External); err != nil {
			return "", err
		}
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(VMwareEnginePrivateCloudGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", VMwareEnginePrivateCloudGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = actualExternalRef
	return r.External, nil
}

// ParsePrivateCloudExternal parses the external reference of a PrivateCloud.
func ParsePrivateCloudExternal(external string) (parent string, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "privateClouds" {
		return "", "", fmt.Errorf("format of privatecloud external=%q was not known (use projects/{{projectID}}/locations/{{location}}/privateClouds/{{privatecloudID}})", external)
	}
	parent = strings.Join(tokens[:len(tokens)-2], "/")
	resourceID = tokens[5]
	return parent, resourceID, nil
}
