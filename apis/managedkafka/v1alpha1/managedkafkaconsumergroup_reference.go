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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ConsumerGroupRef{}

func init() {
	refs.Register(&ConsumerGroupRef{})
}

// ConsumerGroupRef defines the resource reference to ManagedKafkaConsumerGroup, which "External" field
// holds the GCP identifier for the KRM object.
type ConsumerGroupRef struct {
	// A reference to an externally managed ManagedKafkaConsumerGroup resource.
	// Should be in the format "projects/{project}/locations/{location}/clusters/{cluster}/consumerGroups/{consumerGroup}".
	External string `json:"external,omitempty"`

	// The name of a ManagedKafkaConsumerGroup resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ManagedKafkaConsumerGroup resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ConsumerGroupRef) GetGVK() schema.GroupVersionKind {
	return ManagedKafkaConsumerGroupGVK
}

func (r *ConsumerGroupRef) GetNamespacedName() client.ObjectKey {
	return client.ObjectKey{Name: r.Name, Namespace: r.Namespace}
}

func (r *ConsumerGroupRef) GetExternal() string {
	return r.External
}

func (r *ConsumerGroupRef) SetExternal(external string) {
	r.External = external
}

func (r *ConsumerGroupRef) ValidateExternal(external string) error {
	return (&ManagedKafkaConsumerGroupIdentity{}).FromExternal(external)
}

func (r *ConsumerGroupRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ManagedKafkaConsumerGroupIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ConsumerGroupRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, func(u *unstructured.Unstructured) string {
		id, err := GetManagedKafkaConsumerGroupSpecIdentity(ctx, reader, u)
		if err != nil {
			return ""
		}
		return id.String()
	})
}
