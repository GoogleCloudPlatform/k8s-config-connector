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

// TODO: Move this to autoscalingpolicy_types.go once it is generated/implemented
var DataprocAutoscalingPolicyGVK = schema.GroupVersionKind{
	Group:   "dataproc.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "DataprocAutoscalingPolicy",
}

var _ refs.Ref = &DataprocAutoscalingPolicyRef{}

// DataprocAutoscalingPolicyRef is a reference to a DataprocAutoscalingPolicy resource.
type DataprocAutoscalingPolicyRef struct {
	// A reference to an externally managed DataprocAutoscalingPolicy resource.
	// Should be in the format "projects/{{projectID}}/regions/{{region}}/autoscalingPolicies/{{autoscalingPolicy}}".
	External string `json:"external,omitempty"`

	// The name of a DataprocAutoscalingPolicy resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DataprocAutoscalingPolicy resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&DataprocAutoscalingPolicyRef{})
}

func (r *DataprocAutoscalingPolicyRef) GetGVK() schema.GroupVersionKind {
	return DataprocAutoscalingPolicyGVK
}

func (r *DataprocAutoscalingPolicyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DataprocAutoscalingPolicyRef) GetExternal() string {
	return r.External
}

func (r *DataprocAutoscalingPolicyRef) SetExternal(ref string) {
	r.External = ref
}

func (r *DataprocAutoscalingPolicyRef) ValidateExternal(ref string) error {
	id := &DataprocAutoscalingPolicyIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *DataprocAutoscalingPolicyRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DataprocAutoscalingPolicyIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DataprocAutoscalingPolicyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromDataprocAutoscalingPolicySpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
