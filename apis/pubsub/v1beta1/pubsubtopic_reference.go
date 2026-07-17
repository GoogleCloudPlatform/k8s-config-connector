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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var PubSubTopicGVK = schema.GroupVersionKind{
	Group:   "pubsub.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "PubSubTopic",
}

var _ refs.Ref = &PubSubTopicRef{}

// PubSubTopicRef is a reference to a PubSubTopic.
type PubSubTopicRef struct {
	// A reference to an externally managed PubSubTopic resource. Should be in the format "projects/{{projectID}}/topics/{{topicID}}".
	External string `json:"external,omitempty"`

	// The name of a PubSubTopic resource.
	Name string `json:"name,omitempty"`

	// The namespace of a PubSubTopic resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&PubSubTopicRef{}, &PubSubTopic{})
}

func (r *PubSubTopicRef) GetGVK() schema.GroupVersionKind {
	return PubSubTopicGVK
}

func (r *PubSubTopicRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *PubSubTopicRef) GetExternal() string {
	return r.External
}

func (r *PubSubTopicRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *PubSubTopicRef) ValidateExternal(ref string) error {
	id := &PubSubTopicIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *PubSubTopicRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &PubSubTopicIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *PubSubTopicRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*PubSubTopic](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromPubSubTopicSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// NormalizedExternal provision the "External" value.
// Kept for backward compatibility with older callers.
func (r *PubSubTopicRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}
