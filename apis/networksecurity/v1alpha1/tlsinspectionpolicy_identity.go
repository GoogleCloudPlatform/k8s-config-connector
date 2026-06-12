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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NetworkSecurityTLSInspectionPolicyIdentity{}
	_ identity.Resource   = &NetworkSecurityTLSInspectionPolicy{}
)

var networkSecurityTLSInspectionPolicyURL = gcpurls.Template[NetworkSecurityTLSInspectionPolicyIdentity](
	"networksecurity.googleapis.com",
	"projects/{project}/locations/{location}/tlsInspectionPolicies/{tlsinspectionpolicy}",
)

// NetworkSecurityTLSInspectionPolicyIdentity defines the resource reference to NetworkSecurityTLSInspectionPolicy, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type NetworkSecurityTLSInspectionPolicyIdentity struct {
	Project             string `json:"project"`
	Location            string `json:"location"`
	TlsInspectionPolicy string `json:"tls_inspection_policy"`
}

func (i *NetworkSecurityTLSInspectionPolicyIdentity) FromExternal(ref string) error {
	out, match, err := networkSecurityTLSInspectionPolicyURL.Parse(ref)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of NetworkSecurityTLSInspectionPolicy external=%q was not known (use %s)", ref, networkSecurityTLSInspectionPolicyURL.CanonicalForm())
	}
	*i = *out
	return nil
}

func (i *NetworkSecurityTLSInspectionPolicyIdentity) String() string {
	return networkSecurityTLSInspectionPolicyURL.ToString(*i)
}

func (i *NetworkSecurityTLSInspectionPolicyIdentity) Host() string {
	return networkSecurityTLSInspectionPolicyURL.Host()
}

// NewNetworkSecurityTLSInspectionPolicyIdentity builds a NetworkSecurityTLSInspectionPolicyIdentity from the Config Connector NetworkSecurityTLSInspectionPolicy object.
func NewNetworkSecurityTLSInspectionPolicyIdentity(ctx context.Context, reader client.Reader, obj *NetworkSecurityTLSInspectionPolicy) (*NetworkSecurityTLSInspectionPolicyIdentity, error) {
	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := common.ValueOf(obj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualIdentity := &NetworkSecurityTLSInspectionPolicyIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actualIdentity.Project != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.Project, projectID)
		}
		if actualIdentity.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualIdentity.Location, location)
		}
		if actualIdentity.TlsInspectionPolicy != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.TlsInspectionPolicy)
		}
	}
	return &NetworkSecurityTLSInspectionPolicyIdentity{
		Project:             projectID,
		Location:            location,
		TlsInspectionPolicy: resourceID,
	}, nil
}

func (obj *NetworkSecurityTLSInspectionPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return NewNetworkSecurityTLSInspectionPolicyIdentity(ctx, reader, obj)
}

// ExternalIdentifier is the GCP identifier for the resource.
func (obj *NetworkSecurityTLSInspectionPolicy) ExternalIdentifier(ctx context.Context, reader client.Reader) (string, error) {
	identity, err := NewNetworkSecurityTLSInspectionPolicyIdentity(ctx, reader, obj)
	if err != nil {
		return "", err
	}
	return identity.String(), nil
}
