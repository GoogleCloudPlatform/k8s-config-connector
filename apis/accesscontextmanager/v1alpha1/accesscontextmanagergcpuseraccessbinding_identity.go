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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func getIdentityFromAccessContextManagerGCPUserAccessBindingSpec(ctx context.Context, reader client.Reader, obj *AccessContextManagerGCPUserAccessBinding) (*AccessContextManagerGCPUserAccessBindingIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	var organizationID string
	if obj.Spec.OrganizationRef != nil {
		external := obj.Spec.OrganizationRef.External
		if external != "" && !strings.Contains(external, "/") {
			external = "organizations/" + external
		}
		ref := &refs.OrganizationRef{
			External: external,
		}
		org, err := refs.ResolveOrganization(ctx, reader, obj, ref)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve organization: %w", err)
		}
		organizationID = org.OrganizationID
	} else {
		org, err := refs.ResolveOrganizationFromAnnotation(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve organization: %w", err)
		}
		organizationID = strings.TrimPrefix(org.OrganizationID, "organizations/")
	}

	return &AccessContextManagerGCPUserAccessBindingIdentity{
		Organization:         organizationID,
		GcpUserAccessBinding: resourceID,
	}, nil
}
