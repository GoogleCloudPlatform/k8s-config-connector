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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// SecurityCenterMuteConfigRef represents a reference to a SecurityCenterMuteConfig resource.
type SecurityCenterMuteConfigRef struct {
	// A reference to an externally managed SecurityCenterMuteConfig resource. Should be in the format "organizations/{{organizationID}}/muteConfigs/{{muteConfig}}".
	// +optional
	External string `json:"external,omitempty"`

	// The name of a SecurityCenterMuteConfig resource.
	// +optional
	Name string `json:"name,omitempty"`

	// The namespace of a SecurityCenterMuteConfig resource.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

var _ refs.Ref = &SecurityCenterMuteConfigRef{}

func init() {
	refs.Register(&SecurityCenterMuteConfigRef{}, &SecurityCenterMuteConfig{})
}

func (r *SecurityCenterMuteConfigRef) GetGVK() schema.GroupVersionKind {
	return SecurityCenterMuteConfigGVK
}

func (r *SecurityCenterMuteConfigRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *SecurityCenterMuteConfigRef) GetExternal() string {
	return r.External
}

func (r *SecurityCenterMuteConfigRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

func (r *SecurityCenterMuteConfigRef) ValidateExternal(external string) error {
	id := &SecurityCenterMuteConfigIdentity{}
	if err := id.FromExternal(external); err != nil {
		return err
	}
	return nil
}

func (r *SecurityCenterMuteConfigRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &SecurityCenterMuteConfigIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *SecurityCenterMuteConfigRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		typed, err := common.ToStructuredType[*SecurityCenterMuteConfig](u)
		if err != nil {
			return ""
		}
		id, err := getIdentityFromSecurityCenterMuteConfigSpec(ctx, reader, typed)
		if err != nil {
			return ""
		}
		return id.String()
	})
}
