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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &CloudSupportSupportEventSubscriptionRef{}

// CloudSupportSupportEventSubscriptionRef is a reference to a GCP CloudSupportSupportEventSubscription.
type CloudSupportSupportEventSubscriptionRef struct {
	// A reference to an externally managed CloudSupportSupportEventSubscription resource.
	// Should be in the format "organizations/{{organization}}/supportEventSubscriptions/{{supportEventSubscription}}".
	External string `json:"external,omitempty"`

	// The name of a CloudSupportSupportEventSubscription resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CloudSupportSupportEventSubscription resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&CloudSupportSupportEventSubscriptionRef{}, &CloudSupportSupportEventSubscription{})
}

func (r *CloudSupportSupportEventSubscriptionRef) GetGVK() schema.GroupVersionKind {
	return CloudSupportSupportEventSubscriptionGVK
}

func (r *CloudSupportSupportEventSubscriptionRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CloudSupportSupportEventSubscriptionRef) GetExternal() string {
	return r.External
}

func (r *CloudSupportSupportEventSubscriptionRef) SetExternal(external string) {
	r.External = external
}

func (r *CloudSupportSupportEventSubscriptionRef) ValidateExternal(ref string) error {
	id := &CloudSupportSupportEventSubscriptionIdentity{}
	return id.FromExternal(ref)
}

func (r *CloudSupportSupportEventSubscriptionRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CloudSupportSupportEventSubscriptionIdentity{}
	err := id.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CloudSupportSupportEventSubscriptionRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
