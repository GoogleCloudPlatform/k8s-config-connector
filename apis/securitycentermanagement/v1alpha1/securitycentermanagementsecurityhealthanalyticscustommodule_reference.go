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

// SecurityCenterManagementSecurityHealthAnalyticsCustomModuleRef represents a reference to a SecurityCenterManagementSecurityHealthAnalyticsCustomModule resource.
type SecurityCenterManagementSecurityHealthAnalyticsCustomModuleRef struct {
	// A reference to an externally managed SecurityCenterManagementSecurityHealthAnalyticsCustomModule resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/securityHealthAnalyticsCustomModules/{{securityHealthAnalyticsCustomModule}}".
	// +optional
	External string `json:"external,omitempty"`

	// The name of a SecurityCenterManagementSecurityHealthAnalyticsCustomModule resource.
	// +optional
	Name string `json:"name,omitempty"`

	// The namespace of a SecurityCenterManagementSecurityHealthAnalyticsCustomModule resource.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

var _ refs.Ref = &SecurityCenterManagementSecurityHealthAnalyticsCustomModuleRef{}

func init() {
	refs.Register(&SecurityCenterManagementSecurityHealthAnalyticsCustomModuleRef{}, &SecurityCenterManagementSecurityHealthAnalyticsCustomModule{})
}

func (r *SecurityCenterManagementSecurityHealthAnalyticsCustomModuleRef) GetGVK() schema.GroupVersionKind {
	return SecurityCenterManagementSecurityHealthAnalyticsCustomModuleGVK
}

func (r *SecurityCenterManagementSecurityHealthAnalyticsCustomModuleRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *SecurityCenterManagementSecurityHealthAnalyticsCustomModuleRef) GetExternal() string {
	return r.External
}

func (r *SecurityCenterManagementSecurityHealthAnalyticsCustomModuleRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

func (r *SecurityCenterManagementSecurityHealthAnalyticsCustomModuleRef) ValidateExternal(external string) error {
	id := &SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity{}
	if err := id.FromExternal(external); err != nil {
		return err
	}
	return nil
}

func (r *SecurityCenterManagementSecurityHealthAnalyticsCustomModuleRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *SecurityCenterManagementSecurityHealthAnalyticsCustomModuleRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		typed, err := common.ToStructuredType[*SecurityCenterManagementSecurityHealthAnalyticsCustomModule](u)
		if err != nil {
			return ""
		}
		id, err := getIdentityFromSecurityCenterManagementSecurityHealthAnalyticsCustomModuleSpec(ctx, reader, typed)
		if err != nil {
			return ""
		}
		return id.String()
	})
}
