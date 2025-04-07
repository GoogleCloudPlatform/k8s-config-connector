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

package v1alpha1

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

var _ refsv1beta1.ExternalNormalizer = &ConsumerGroupRef{}

// ConsumerGroupRef defines the resource reference to ManagedKafkaConsumerGroup, which "External" field
// holds the GCP identifier for the KRM object.
type ConsumerGroupRef struct {
	// A reference to an externally managed ManagedKafkaConsumerGroup resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/consumergroups/{{consumergroupID}}".
	External string `json:"external,omitempty"`

	// The name of a ManagedKafkaConsumerGroup resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ManagedKafkaConsumerGroup resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on ManagedKafkaConsumerGroup.
// If the "External" is given in the other resource's spec.ManagedKafkaConsumerGroupRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ManagedKafkaConsumerGroup object from the cluster.
func (r *ConsumerGroupRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ManagedKafkaConsumerGroupGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := ParseConsumerGroupExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(ManagedKafkaConsumerGroupGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", ManagedKafkaConsumerGroupGVK, key, err)
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
