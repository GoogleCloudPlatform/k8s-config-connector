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
	"fmt"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &ServiceAccountKeyRef{}

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

// NormalizedExternal provision the "External" value for other resource that depends on IAMServiceAccountKey.
// If the "External" is given in the other resource's spec.ServiceAccountKeyRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual IAMServiceAccountKey object from the cluster.
func (r *ServiceAccountKeyRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", IAMServiceAccountKeyGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, err := ParseServiceAccountKeyExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(IAMServiceAccountKeyGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", IAMServiceAccountKeyGVK, key, err)
	}
	// Get external from status.name. This is the full resource name.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "name")
	if err != nil {
		return "", fmt.Errorf("reading status.name from %s %s: %w", IAMServiceAccountKeyGVK, key, err)
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = actualExternalRef
	return r.External, nil
}
