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
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var CloudSecurityComplianceCloudControlGVK = schema.GroupVersionKind{
	Group:   "cloudsecuritycompliance.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "CloudSecurityComplianceCloudControl",
}

var _ refs.Ref = &CloudSecurityComplianceCloudControlRef{}

// CloudSecurityComplianceCloudControlRef is a reference to a CloudSecurityComplianceCloudControl.
type CloudSecurityComplianceCloudControlRef struct {
	// A reference to an externally managed CloudSecurityComplianceCloudControl resource.
	// Should be in the format "organizations/{{organizationID}}/locations/{{location}}/cloudControls/{{cloudControlID}}".
	External string `json:"external,omitempty"`

	// The name of a CloudSecurityComplianceCloudControl resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CloudSecurityComplianceCloudControl resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&CloudSecurityComplianceCloudControlRef{})
}

func (r *CloudSecurityComplianceCloudControlRef) GetGVK() schema.GroupVersionKind {
	return CloudSecurityComplianceCloudControlGVK
}

func (r *CloudSecurityComplianceCloudControlRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CloudSecurityComplianceCloudControlRef) GetExternal() string {
	return r.External
}

func (r *CloudSecurityComplianceCloudControlRef) SetExternal(ref string) {
	r.External = ref
}

func (r *CloudSecurityComplianceCloudControlRef) ValidateExternal(ref string) error {
	id := &CloudSecurityComplianceCloudControlIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CloudSecurityComplianceCloudControlRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CloudSecurityComplianceCloudControlIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CloudSecurityComplianceCloudControlRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromCloudSecurityComplianceCloudControlSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
