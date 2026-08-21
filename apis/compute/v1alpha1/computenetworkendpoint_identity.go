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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeNetworkEndpointIdentity{}
	_ identity.Resource   = &ComputeNetworkEndpoint{}
)

var ComputeNetworkEndpointWithInstanceIdentityFormat = gcpurls.Template[ComputeNetworkEndpointIdentity](
	"compute.googleapis.com",
	"projects/{Project}/zones/{Zone}/networkEndpointGroups/{ComputeNetworkEndpointGroup}/{Instance}/{IpAddress}/{Port}",
)

var ComputeNetworkEndpointWithoutInstanceIdentityFormat = gcpurls.Template[ComputeNetworkEndpointIdentity](
	"compute.googleapis.com",
	"projects/{Project}/zones/{Zone}/networkEndpointGroups/{ComputeNetworkEndpointGroup}//{IpAddress}/{Port}",
)

// ComputeNetworkEndpointIdentity is the identity of a GCP ComputeNetworkEndpoint resource.
// +k8s:deepcopy-gen=false
type ComputeNetworkEndpointIdentity struct {
	Project                     string
	Zone                        string
	ComputeNetworkEndpointGroup string
	Instance                    string
	IpAddress                   string
	Port                        string
}

func (i *ComputeNetworkEndpointIdentity) String() string {
	if i.Instance != "" {
		return ComputeNetworkEndpointWithInstanceIdentityFormat.ToString(*i)
	}
	return ComputeNetworkEndpointWithoutInstanceIdentityFormat.ToString(*i)
}

func (i *ComputeNetworkEndpointIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)

	// First, try with instance
	parsed, match, err := ComputeNetworkEndpointWithInstanceIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeNetworkEndpoint external=%q was not known (use %s or %s): %w", ref, ComputeNetworkEndpointWithInstanceIdentityFormat.CanonicalForm(), ComputeNetworkEndpointWithoutInstanceIdentityFormat.CanonicalForm(), err)
	}
	if match {
		*i = *parsed
		return nil
	}

	// Next, try without instance
	parsed, match, err = ComputeNetworkEndpointWithoutInstanceIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeNetworkEndpoint external=%q was not known (use %s or %s): %w", ref, ComputeNetworkEndpointWithInstanceIdentityFormat.CanonicalForm(), ComputeNetworkEndpointWithoutInstanceIdentityFormat.CanonicalForm(), err)
	}
	if match {
		*i = *parsed
		return nil
	}

	return fmt.Errorf("format of ComputeNetworkEndpoint external=%q was not known (use %s or %s)", ref, ComputeNetworkEndpointWithInstanceIdentityFormat.CanonicalForm(), ComputeNetworkEndpointWithoutInstanceIdentityFormat.CanonicalForm())
}

func (i *ComputeNetworkEndpointIdentity) Host() string {
	return ComputeNetworkEndpointWithInstanceIdentityFormat.Host()
}

func (i *ComputeNetworkEndpointIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/zones/%s/networkEndpointGroups/%s", i.Project, i.Zone, i.ComputeNetworkEndpointGroup)
}

func getIdentityFromComputeNetworkEndpointSpec(ctx context.Context, reader client.Reader, obj *ComputeNetworkEndpoint) (*ComputeNetworkEndpointIdentity, error) {
	if obj.Spec.NetworkEndpointGroupRef == nil {
		return nil, fmt.Errorf("spec.networkEndpointGroupRef is required")
	}
	negRef := obj.Spec.NetworkEndpointGroupRef.DeepCopy()
	if err := negRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("cannot normalize networkEndpointGroupRef: %w", err)
	}
	negIdentity, err := negRef.ParseExternalToIdentity()
	if err != nil {
		return nil, fmt.Errorf("cannot parse networkEndpointGroupRef external to identity: %w", err)
	}
	negID := negIdentity.(*computev1beta1.ComputeNetworkEndpointGroupIdentity)

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	var instance string
	if obj.Spec.InstanceRef != nil {
		instanceRef := obj.Spec.InstanceRef.DeepCopy()
		if err := instanceRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, fmt.Errorf("cannot normalize instanceRef: %w", err)
		}
		instIdentity, err := instanceRef.ParseExternalToIdentity()
		if err != nil {
			return nil, fmt.Errorf("cannot parse instanceRef external to identity: %w", err)
		}
		instance = instIdentity.(*computev1beta1.ComputeInstanceIdentity).Instance
	}

	return &ComputeNetworkEndpointIdentity{
		Project:                     negID.Project,
		Zone:                        negID.Zone,
		ComputeNetworkEndpointGroup: negID.ComputeNetworkEndpointGroup,
		Instance:                    instance,
		IpAddress:                   obj.Spec.IpAddress,
		Port:                        resourceID,
	}, nil
}

func (obj *ComputeNetworkEndpoint) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeNetworkEndpointSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	return specIdentity, nil
}
