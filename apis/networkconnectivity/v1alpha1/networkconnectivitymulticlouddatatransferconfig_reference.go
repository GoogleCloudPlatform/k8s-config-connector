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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &NetworkConnectivityMulticloudDataTransferConfigRef{}

// NetworkConnectivityMulticloudDataTransferConfigRef is a reference to a GCP NetworkConnectivityMulticloudDataTransferConfig.
type NetworkConnectivityMulticloudDataTransferConfigRef struct {
	// A reference to an externally managed NetworkConnectivityMulticloudDataTransferConfig resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/multicloudDataTransferConfigs/{{multicloudDataTransferConfig}}".
	External string `json:"external,omitempty"`

	// The name of a NetworkConnectivityMulticloudDataTransferConfig resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkConnectivityMulticloudDataTransferConfig resource.
	Namespace string `json:"namespace,omitempty"`
}

// GetGVK returns the GroupVersionKind for NetworkConnectivityMulticloudDataTransferConfig
func (r *NetworkConnectivityMulticloudDataTransferConfigRef) GetGVK() schema.GroupVersionKind {
	return NetworkConnectivityMulticloudDataTransferConfigGVK
}

// GetNamespacedName returns the NamespacedName for NetworkConnectivityMulticloudDataTransferConfig
func (r *NetworkConnectivityMulticloudDataTransferConfigRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

// GetExternal returns the External string for NetworkConnectivityMulticloudDataTransferConfig
func (r *NetworkConnectivityMulticloudDataTransferConfigRef) GetExternal() string {
	return r.External
}

// SetExternal sets the External string for NetworkConnectivityMulticloudDataTransferConfig
func (r *NetworkConnectivityMulticloudDataTransferConfigRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

// ValidateExternal checks if the External string is valid
func (r *NetworkConnectivityMulticloudDataTransferConfigRef) ValidateExternal(external string) error {
	identity := &NetworkConnectivityMulticloudDataTransferConfigIdentity{}
	return identity.FromExternal(external)
}

// ParseExternalToIdentity parses the External string to an identity
func (r *NetworkConnectivityMulticloudDataTransferConfigRef) ParseExternalToIdentity() (any, error) {
	identity := &NetworkConnectivityMulticloudDataTransferConfigIdentity{}
	err := identity.FromExternal(r.External)
	return identity, err
}

// Normalize normalizes the reference
func (r *NetworkConnectivityMulticloudDataTransferConfigRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*NetworkConnectivityMulticloudDataTransferConfig](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromNetworkConnectivityMulticloudDataTransferConfigSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	})
}

func init() {
	refs.Register(&NetworkConnectivityMulticloudDataTransferConfigRef{}, &NetworkConnectivityMulticloudDataTransferConfig{})
}
