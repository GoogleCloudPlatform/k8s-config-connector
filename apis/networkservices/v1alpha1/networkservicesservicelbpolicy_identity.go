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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NetworkServicesServiceLBPolicyIdentity{}
	_ identity.Resource   = &NetworkServicesServiceLBPolicy{}
)

var NetworkServicesServiceLBPolicyIdentityFormat = gcpurls.Template[NetworkServicesServiceLBPolicyIdentity]("networkservices.googleapis.com", "projects/{project}/locations/{location}/serviceLbPolicies/{serviceLbPolicy}")

// NetworkServicesServiceLBPolicyIdentity is the identity of a GCP NetworkServicesServiceLBPolicy resource.
// +k8s:deepcopy-gen=false
type NetworkServicesServiceLBPolicyIdentity struct {
	Project         string
	Location        string
	ServiceLbPolicy string
}

func (i *NetworkServicesServiceLBPolicyIdentity) String() string {
	return NetworkServicesServiceLBPolicyIdentityFormat.ToString(*i)
}

func (i *NetworkServicesServiceLBPolicyIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkServicesServiceLBPolicyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkServicesServiceLBPolicy external=%q was not known (use %s): %w", ref, NetworkServicesServiceLBPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkServicesServiceLBPolicy external=%q was not known (use %s)", ref, NetworkServicesServiceLBPolicyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkServicesServiceLBPolicyIdentity) Host() string {
	return NetworkServicesServiceLBPolicyIdentityFormat.Host()
}

func (i *NetworkServicesServiceLBPolicyIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func NewNetworkServicesServiceLBPolicyIdentity(ctx context.Context, reader client.Reader, obj *NetworkServicesServiceLBPolicy) (*NetworkServicesServiceLBPolicyIdentity, error) {
	resourceID, err := refsv1beta1.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location, err := refsv1beta1.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &NetworkServicesServiceLBPolicyIdentity{
		Project:         projectID,
		Location:        location,
		ServiceLbPolicy: resourceID,
	}
	return identity, nil
}

func (obj *NetworkServicesServiceLBPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := NewNetworkServicesServiceLBPolicyIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if obj.Status.ExternalRef != nil && *obj.Status.ExternalRef != "" {
		statusIdentity := &NetworkServicesServiceLBPolicyIdentity{}
		if err := statusIdentity.FromExternal(*obj.Status.ExternalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkServicesServiceLBPolicy identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
