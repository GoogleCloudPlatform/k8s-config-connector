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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ConnectorsConnectionRef{}

// ConnectorsConnectionRef defines the resource reference to ConnectorsConnection, which "External" field
// holds the GCP identifier for the KRM object.
type ConnectorsConnectionRef struct {
	// A reference to an externally managed ConnectorsConnection resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/connections/{{connection}}".
	External string `json:"external,omitempty"`

	// The name of a ConnectorsConnection resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ConnectorsConnection resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ConnectorsConnectionRef{})
}

func (r *ConnectorsConnectionRef) GetGVK() schema.GroupVersionKind {
	return ConnectorsConnectionGVK
}

func (r *ConnectorsConnectionRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ConnectorsConnectionRef) GetExternal() string {
	return r.External
}

func (r *ConnectorsConnectionRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ConnectorsConnectionRef) ValidateExternal(ref string) error {
	id := &ConnectorsConnectionIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ConnectorsConnectionRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ConnectorsConnectionIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ConnectorsConnectionRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromConnectorsConnectionSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

var _ refs.Ref = &ConnectorsConnectorVersionRef{}

// ConnectorsConnectorVersionRef defines the resource reference to ConnectorsConnectorVersion.
type ConnectorsConnectorVersionRef struct {
	// A reference to an externally managed ConnectorsConnectorVersion resource.
	// Should be in the format "projects/{{projectID}}/locations/global/providers/{{providerID}}/connectors/{{connectorID}}/versions/{{versionID}}".
	External string `json:"external,omitempty"`

	// The name of a ConnectorsConnectorVersion resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ConnectorsConnectorVersion resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ConnectorsConnectorVersionRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "connectors.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "ConnectorsConnectorVersion",
	}
}

func (r *ConnectorsConnectorVersionRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ConnectorsConnectorVersionRef) GetExternal() string {
	return r.External
}

func (r *ConnectorsConnectorVersionRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ConnectorsConnectorVersionRef) ValidateExternal(ref string) error {
	return nil
}

func (r *ConnectorsConnectorVersionRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("must specify external on ConnectorsConnectorVersion reference (namespaced references not supported as the resource is not managed by Config Connector)")
	}
	return nil
}
