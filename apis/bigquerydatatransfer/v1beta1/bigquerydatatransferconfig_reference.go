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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ v1beta1.Ref = &BigQueryDataTransferConfigRef{}

func init() {
	v1beta1.Register(&BigQueryDataTransferConfigRef{})
}

type BigQueryDataTransferConfigRef struct {
	/* A reference to an externally managed BigQueryDataTransferConfig resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/transferConfigs/{{transferConfigID}}" */
	External string `json:"external,omitempty"`

	/* The name of a BigQueryDataTransferConfig resource. */
	Name string `json:"name,omitempty"`

	/* The namespace of a BigQueryDataTransferConfig resource. */
	Namespace string `json:"namespace,omitempty"`
}

func (r *BigQueryDataTransferConfigRef) GetGVK() schema.GroupVersionKind {
	return BigQueryDataTransferConfigGVK
}

func (r *BigQueryDataTransferConfigRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *BigQueryDataTransferConfigRef) GetExternal() string {
	return r.External
}

func (r *BigQueryDataTransferConfigRef) SetExternal(external string) {
	r.External = external
}

func (r *BigQueryDataTransferConfigRef) ValidateExternal(external string) error {
	return (&BigQueryDataTransferConfigIdentity{}).FromExternal(external)
}

func (r *BigQueryDataTransferConfigRef) ParseExternalToIdentity() (identity.Identity, error) {
	identity := &BigQueryDataTransferConfigIdentity{}
	if err := identity.FromExternal(r.External); err != nil {
		return nil, err
	}
	return identity, nil
}

func (r *BigQueryDataTransferConfigRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return v1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		if id, err := getIdentityFromBigQueryDataTransferConfigSpec(ctx, reader, u); err == nil {
			return id.String()
		}
		return ""
	})
}
