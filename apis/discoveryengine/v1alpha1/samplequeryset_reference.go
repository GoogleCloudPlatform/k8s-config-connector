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

package v1alpha1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &DiscoveryEngineSampleQuerySetRef{}

// DiscoveryEngineSampleQuerySetRef defines the resource reference to DiscoveryEngineSampleQuerySet, which "External" field
// holds the GCP identifier for the KRM object.
type DiscoveryEngineSampleQuerySetRef struct {
	// A reference to an externally managed DiscoveryEngineSampleQuerySet resource.
	// Should be of the format `projects/{{projectID}}/locations/{{location}}/sampleQuerySets/{{sample_query_set}}`.
	External string `json:"external,omitempty"`

	// The `name` of a `DiscoveryEngineSampleQuerySet` resource in Kubernetes.
	Name string `json:"name,omitempty"`

	// The `namespace` of a `DiscoveryEngineSampleQuerySet` resource in Kubernetes.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provisions the "External" value for other resource that depends on `DiscoveryEngineSampleQuerySet`.
// If the "External" is given in the other resource's spec.DiscoveryEngineSampleQuerySetRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual `DiscoveryEngineSampleQuerySet` object from the cluster.
func (r *DiscoveryEngineSampleQuerySetRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", DiscoveryEngineSampleQuerySetGVK.Kind)
	}
	// From given External
	if r.External != "" {
		id := &DiscoveryEngineSampleQuerySetIdentity{}
		if err := id.FromExternal(r.External); err != nil {
			return "", err
		}
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := client.ObjectKey{Name: r.Name, Namespace: r.Namespace}
	u := &DiscoveryEngineSampleQuerySet{}
	if err := reader.Get(ctx, key, u); err != nil {
		if client.IgnoreNotFound(err) == nil {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("error reading resource %q: %w", key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef := common.ValueOf(u.Status.ExternalRef)
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	return actualExternalRef, nil
}
