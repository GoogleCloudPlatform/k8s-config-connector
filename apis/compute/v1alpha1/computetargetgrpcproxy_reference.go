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
// See the License for the_reference.go specific language governing permissions and
// limitations under the License.

package v1alpha1

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

var _ refs.Ref = &ComputeTargetGRPCProxyRef{}

// ComputeTargetGRPCProxyRef is a reference to a ComputeTargetGRPCProxy.
type ComputeTargetGRPCProxyRef struct {
	// A reference to an externally managed ComputeTargetGRPCProxy resource.
	// Should be in the format "projects/{{projectID}}/global/targetGrpcProxies/{{computetargetgrpcproxyID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeTargetGRPCProxy resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeTargetGRPCProxy resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeTargetGRPCProxyRef{})
}

func (r *ComputeTargetGRPCProxyRef) GetGVK() schema.GroupVersionKind {
	return ComputeTargetGRPCProxyGVK
}

func (r *ComputeTargetGRPCProxyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeTargetGRPCProxyRef) GetExternal() string {
	return r.External
}

func (r *ComputeTargetGRPCProxyRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeTargetGRPCProxyRef) ValidateExternal(ref string) error {
	id := &ComputeTargetGRPCProxyIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeTargetGRPCProxyRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeTargetGRPCProxyIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeTargetGRPCProxyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*ComputeTargetGRPCProxy](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromComputeTargetGRPCProxySpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
