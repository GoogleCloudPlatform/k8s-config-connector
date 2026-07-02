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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.IdentityV2 = &NetworkSecurityTLSInspectionPolicyIdentity{}
var _ identity.Resource = &NetworkSecurityTLSInspectionPolicy{}

// +k8s:deepcopy-gen=false
type NetworkSecurityTLSInspectionPolicyIdentity struct {
	Project             string
	Location            string
	TlsInspectionPolicy string
}

var NetworkSecurityTLSInspectionPolicyIdentityFormat = gcpurls.Template[NetworkSecurityTLSInspectionPolicyIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/tlsInspectionPolicies/{tlsinspectionpolicy}")

func (i *NetworkSecurityTLSInspectionPolicyIdentity) String() string {
	return NetworkSecurityTLSInspectionPolicyIdentityFormat.ToString(*i)
}

func (i *NetworkSecurityTLSInspectionPolicyIdentity) FromExternal(external string) error {
	parsed, match, err := NetworkSecurityTLSInspectionPolicyIdentityFormat.Parse(external)
	if err != nil {
		return fmt.Errorf("format of NetworkSecurityTLSInspectionPolicy external=%q was not known (use %s): %w", external, NetworkSecurityTLSInspectionPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkSecurityTLSInspectionPolicy external=%q was not known (use %s)", external, NetworkSecurityTLSInspectionPolicyIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *NetworkSecurityTLSInspectionPolicyIdentity) Host() string {
	return NetworkSecurityTLSInspectionPolicyIdentityFormat.Host()
}

func getIdentityFromNetworkSecurityTLSInspectionPolicySpec(ctx context.Context, reader client.Reader, obj client.Object) (*NetworkSecurityTLSInspectionPolicyIdentity, error) {
	project, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, err
	}

	var resourceID string
	if u, ok := obj.(*unstructured.Unstructured); ok {
		resourceID, _, _ = unstructured.NestedString(u.Object, "spec", "resourceID")
	} else if typed, ok := obj.(*NetworkSecurityTLSInspectionPolicy); ok {
		if typed.Spec.ResourceID != nil {
			resourceID = *typed.Spec.ResourceID
		}
	} else {
		return nil, fmt.Errorf("unexpected object type: %T", obj)
	}

	if resourceID == "" {
		resourceID = obj.GetName()
	}

	return &NetworkSecurityTLSInspectionPolicyIdentity{
		Project:             project,
		Location:            location,
		TlsInspectionPolicy: resourceID,
	}, nil
}

func (r *NetworkSecurityTLSInspectionPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	id, err := getIdentityFromNetworkSecurityTLSInspectionPolicySpec(ctx, reader, r)
	if err != nil {
		return nil, err
	}

	if r.Status.ExternalRef != nil {
		externalID := &NetworkSecurityTLSInspectionPolicyIdentity{}
		if err := externalID.FromExternal(*r.Status.ExternalRef); err != nil {
			return nil, err
		}
		if id.String() != externalID.String() {
			return nil, fmt.Errorf("identity mismatch: spec yields %q, but status.externalRef is %q", id.String(), externalID.String())
		}
	}

	return id, nil
}
