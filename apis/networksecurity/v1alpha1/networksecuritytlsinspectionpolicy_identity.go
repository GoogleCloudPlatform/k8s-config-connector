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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NetworkSecurityTLSInspectionPolicyIdentity{}
	_ identity.Resource   = &NetworkSecurityTLSInspectionPolicy{}
)

var NetworkSecurityTLSInspectionPolicyIdentityFormat = gcpurls.Template[NetworkSecurityTLSInspectionPolicyIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/tlsInspectionPolicies/{tlsInspectionPolicy}")

// NetworkSecurityTLSInspectionPolicyIdentity is the identity of a GCP NetworkSecurityTLSInspectionPolicy resource.
// +k8s:deepcopy-gen=false
type NetworkSecurityTLSInspectionPolicyIdentity struct {
	Project             string
	Location            string
	TlsInspectionPolicy string
}

func (i *NetworkSecurityTLSInspectionPolicyIdentity) String() string {
	return NetworkSecurityTLSInspectionPolicyIdentityFormat.ToString(*i)
}

func (i *NetworkSecurityTLSInspectionPolicyIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkSecurityTLSInspectionPolicyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkSecurityTLSInspectionPolicy external=%q was not known (use %s): %w", ref, NetworkSecurityTLSInspectionPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkSecurityTLSInspectionPolicy external=%q was not known (use %s)", ref, NetworkSecurityTLSInspectionPolicyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkSecurityTLSInspectionPolicyIdentity) Host() string {
	return NetworkSecurityTLSInspectionPolicyIdentityFormat.Host()
}

func (i *NetworkSecurityTLSInspectionPolicyIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromNetworkSecurityTLSInspectionPolicySpec(ctx context.Context, reader client.Reader, obj *NetworkSecurityTLSInspectionPolicy) (*NetworkSecurityTLSInspectionPolicyIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &NetworkSecurityTLSInspectionPolicyIdentity{
		Project:             projectID,
		Location:            location,
		TlsInspectionPolicy: resourceID,
	}
	return identity, nil
}

func (obj *NetworkSecurityTLSInspectionPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecurityTLSInspectionPolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
