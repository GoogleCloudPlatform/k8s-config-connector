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

	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ApigeeNATAddressIdentity{}
	_ identity.Resource   = &ApigeeNATAddress{}
)

var ApigeeNATAddressIdentityFormat = gcpurls.Template[ApigeeNATAddressIdentity]("apigee.googleapis.com", "organizations/{organization}/instances/{instance}/natAddresses/{natAddress}")

// ApigeeNATAddressIdentity is the identity of a GCP ApigeeNATAddress resource.
// +k8s:deepcopy-gen=false
type ApigeeNATAddressIdentity struct {
	Organization string
	Instance     string
	NatAddress   string
}

func (i *ApigeeNATAddressIdentity) String() string {
	return ApigeeNATAddressIdentityFormat.ToString(*i)
}

func (i *ApigeeNATAddressIdentity) FromExternal(ref string) error {
	trimmed := identity.StripReferencePrefixes(ref, "apigee.googleapis.com")
	parsed, match, err := ApigeeNATAddressIdentityFormat.Parse(trimmed)
	if err != nil {
		return fmt.Errorf("format of ApigeeNATAddress external=%q was not known (use %s): %w", ref, ApigeeNATAddressIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ApigeeNATAddress external=%q was not known (use %s)", ref, ApigeeNATAddressIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ApigeeNATAddressIdentity) Host() string {
	return ApigeeNATAddressIdentityFormat.Host()
}

func (i *ApigeeNATAddressIdentity) ID() string {
	return i.NatAddress
}

func (i *ApigeeNATAddressIdentity) ParentString() string {
	return fmt.Sprintf("organizations/%s/instances/%s", i.Organization, i.Instance)
}

func getIdentityFromApigeeNATAddressSpec(ctx context.Context, reader client.Reader, obj *ApigeeNATAddress) (*ApigeeNATAddressIdentity, error) {
	if obj.Spec.InstanceId == "" {
		return nil, fmt.Errorf("spec.instanceId is required")
	}

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	var orgID, instanceID string
	trimmed := identity.StripReferencePrefixes(obj.Spec.InstanceId, "apigee.googleapis.com")
	if strings.Contains(trimmed, "/") {
		tokens := strings.Split(trimmed, "/")
		if len(tokens) == 3 && tokens[1] == "instances" {
			trimmed = "organizations/" + trimmed
		}
		instanceIdentity := &apigeev1beta1.ApigeeInstanceIdentity{}
		if err := instanceIdentity.FromExternal(trimmed); err != nil {
			return nil, fmt.Errorf("invalid instanceId %q: %w", obj.Spec.InstanceId, err)
		}
		orgID = instanceIdentity.ParentID.ResourceID
		instanceID = instanceIdentity.ResourceID
	} else {
		if reader == nil {
			return nil, fmt.Errorf("cannot resolve ApigeeInstance %q without reader", obj.Spec.InstanceId)
		}
		instanceRef := &apigeev1beta1.ApigeeInstanceRef{
			Name:      obj.Spec.InstanceId,
			Namespace: obj.GetNamespace(),
		}
		if err := instanceRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, fmt.Errorf("resolving instanceId %q: %w", obj.Spec.InstanceId, err)
		}
		instanceIdentity := &apigeev1beta1.ApigeeInstanceIdentity{}
		if err := instanceIdentity.FromExternal(identity.StripReferencePrefixes(instanceRef.External, "apigee.googleapis.com")); err != nil {
			return nil, fmt.Errorf("parsing resolved instance reference %q: %w", instanceRef.External, err)
		}
		orgID = instanceIdentity.ParentID.ResourceID
		instanceID = instanceIdentity.ResourceID
	}

	return &ApigeeNATAddressIdentity{
		Organization: orgID,
		Instance:     instanceID,
		NatAddress:   resourceID,
	}, nil
}

func (obj *ApigeeNATAddress) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromApigeeNATAddressSpec(ctx, reader, obj)
}
