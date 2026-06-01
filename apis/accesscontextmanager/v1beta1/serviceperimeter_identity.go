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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ServicePerimeterIdentity{}
	_ identity.Resource   = &AccessContextManagerServicePerimeter{}
)

var ServicePerimeterIdentityFormat = gcpurls.Template[ServicePerimeterIdentity]("accesscontextmanager.googleapis.com", "accessPolicies/{accessPolicy}/servicePerimeters/{servicePerimeter}")

// +k8s:deepcopy-gen=false
type ServicePerimeterIdentity struct {
	AccessPolicy     string
	ServicePerimeter string
}

func (i *ServicePerimeterIdentity) String() string {
	return ServicePerimeterIdentityFormat.ToString(*i)
}

func (i *ServicePerimeterIdentity) FromExternal(ref string) error {
	parsed, match, err := ServicePerimeterIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AccessContextManagerServicePerimeter external=%q was not known (use %s): %w", ref, ServicePerimeterIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AccessContextManagerServicePerimeter external=%q was not known (use %s)", ref, ServicePerimeterIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ServicePerimeterIdentity) Host() string {
	return ServicePerimeterIdentityFormat.Host()
}

func getIdentityFromServicePerimeterSpec(ctx context.Context, reader client.Reader, obj *AccessContextManagerServicePerimeter) (*ServicePerimeterIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	if obj.Spec.AccessPolicyRef == nil {
		return nil, fmt.Errorf("spec.accessPolicyRef is required")
	}

	accessPolicyExternal, err := obj.Spec.AccessPolicyRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, fmt.Errorf("cannot resolve accessPolicyRef: %w", err)
	}

	accessPolicyID, err := ParseAccessPolicyExternal(accessPolicyExternal)
	if err != nil {
		return nil, fmt.Errorf("cannot parse accessPolicyRef external: %w", err)
	}

	identity := &ServicePerimeterIdentity{
		AccessPolicy:     accessPolicyID,
		ServicePerimeter: resourceID,
	}
	return identity, nil
}

func (obj *AccessContextManagerServicePerimeter) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromServicePerimeterSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ServicePerimeterIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change AccessContextManagerServicePerimeter identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
