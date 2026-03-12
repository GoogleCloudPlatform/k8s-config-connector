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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &BigQueryAnalyticsHubDataExchangeRef{}

// BigQueryAnalyticsHubDataExchangeRef defines the resource reference to BigQueryAnalyticsHubDataExchange, which "External" field
// holds the GCP identifier for the KRM object.
type BigQueryAnalyticsHubDataExchangeRef struct {
	// A reference to an externally managed BigQueryAnalyticsHubDataExchange resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/dataExchanges/{{dataExchangeID}}".
	External string `json:"external,omitempty"`

	// The name of a BigQueryAnalyticsHubDataExchange resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigQueryAnalyticsHubDataExchange resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&BigQueryAnalyticsHubDataExchangeRef{})
}

func (r *BigQueryAnalyticsHubDataExchangeRef) GetGVK() schema.GroupVersionKind {
	return BigQueryAnalyticsHubDataExchangeGVK
}

func (r *BigQueryAnalyticsHubDataExchangeRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *BigQueryAnalyticsHubDataExchangeRef) GetExternal() string {
	return r.External
}

func (r *BigQueryAnalyticsHubDataExchangeRef) SetExternal(ref string) {
	r.External = ref
}

func (r *BigQueryAnalyticsHubDataExchangeRef) ValidateExternal(ref string) error {
	id := &BigQueryAnalyticsHubDataExchangeIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *BigQueryAnalyticsHubDataExchangeRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &BigQueryAnalyticsHubDataExchangeIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *BigQueryAnalyticsHubDataExchangeRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromBigQueryAnalyticsHubDataExchangeSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

func (r *BigQueryAnalyticsHubDataExchangeRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}
