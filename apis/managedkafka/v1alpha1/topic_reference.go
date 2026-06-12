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

var _ refs.Ref = &TopicRef{}

func init() {
	refs.Register(&TopicRef{})
}

// TopicRef is a reference to a ManagedKafkaTopic.
type TopicRef struct {
	// A reference to an externally managed ManagedKafkaTopic resource.
	// Should be in the format "projects/{project}/locations/{location}/clusters/{cluster}/topics/{topic}".
	External string `json:"external,omitempty"`

	// The name of a ManagedKafkaTopic resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ManagedKafkaTopic resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *TopicRef) GetGVK() schema.GroupVersionKind {
	return ManagedKafkaTopicGVK
}

func (r *TopicRef) GetNamespacedName() client.ObjectKey {
	return client.ObjectKey{Name: r.Name, Namespace: r.Namespace}
}

func (r *TopicRef) GetExternal() string {
	return r.External
}

func (r *TopicRef) SetExternal(external string) {
	r.External = external
}

func (r *TopicRef) ValidateExternal(external string) error {
	return (&TopicIdentity{}).FromExternal(external)
}

func (r *TopicRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &TopicIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *TopicRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, func(u *unstructured.Unstructured) string {
		id, err := getIdentityFromManagedKafkaTopicSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return id.String()
	})
}
