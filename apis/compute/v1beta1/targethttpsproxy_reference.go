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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &ComputeTargetHTTPSProxyRef{}

type ComputeTargetHTTPSProxyRef struct {
	// Allowed value: string of the format `projects/{{project}}/global/targetHttpsProxies/{{value}}` or `projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{value}}`, where {{value}} is the `name` field of a `ComputeTargetHTTPSProxy` resource.
	External string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

func (r *ComputeTargetHTTPSProxyRef) GetGVK() schema.GroupVersionKind {
	return ComputeTargetHTTPSProxyGVK
}

func (r *ComputeTargetHTTPSProxyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeTargetHTTPSProxyRef) GetExternal() string {
	return r.External
}

func (r *ComputeTargetHTTPSProxyRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ComputeTargetHTTPSProxyRef) ValidateExternal(ref string) error {
	id := &ComputeTargetHTTPSProxyIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeTargetHTTPSProxyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}
