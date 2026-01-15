// Copyright 2025 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var ExternalAddressIdentityFormat = gcpurls.Template[ExternalAddressIdentity]("vmwareengine.googleapis.com", "projects/{project}/locations/{location}/privateClouds/{privateCloud}/externalAddresses/{externalAddress}")

// ExternalAddressIdentity defines the resource reference to VMwareEngineExternalAddress, which "External" field
// holds the GCP identifier for the KRM object.
type ExternalAddressIdentity struct {
	Project         string
	Location        string
	PrivateCloud    string
	ExternalAddress string
}

func (i *ExternalAddressIdentity) String() string {
	return ExternalAddressIdentityFormat.ToString(*i)
}

func (i *ExternalAddressIdentity) ID() string {
	return i.ExternalAddress
}

func (i *ExternalAddressIdentity) Parent() *ExternalAddressParent {
	return &ExternalAddressParent{
		Project:      i.Project,
		Location:     i.Location,
		PrivateCloud: i.PrivateCloud,
	}
}

func (i *ExternalAddressIdentity) FromExternal(ref string) error {
	parsed, match, err := ExternalAddressIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ExternalAddress external=%q was not known (use %s): %w", ref, ExternalAddressIdentityFormat, err)
	}
	if !match {
		return fmt.Errorf("format of ExternalAddress external=%q was not known (use %s)", ref, ExternalAddressIdentityFormat)
	}

	*i = *parsed
	return nil
}

type ExternalAddressParent struct {
	Project      string
	Location     string
	PrivateCloud string
}

func (p *ExternalAddressParent) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", p.Project, p.Location, p.PrivateCloud)
}

// New builds a ExternalAddressIdentity from the Config Connector ExternalAddress object.
func NewExternalAddressIdentity(ctx context.Context, reader client.Reader, obj *VMwareEngineExternalAddress) (*ExternalAddressIdentity, error) {
	// Get Parent
	privateCloud, err := obj.Spec.PrivateCloudRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
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
		id := &ExternalAddressIdentity{}
		if err := id.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if id.Parent().String() != privateCloud {
			return nil, fmt.Errorf("spec.privateCloudRef changed, expect %s, got %s", id.Parent().String(), privateCloud)
		}
		if id.ExternalAddress != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, id.ExternalAddress)
		}
		return id, nil
	}

	// Construct from spec
	fullPath := fmt.Sprintf("%s/externalAddresses/%s", privateCloud, resourceID)
	id := &ExternalAddressIdentity{}
	if err := id.FromExternal(fullPath); err != nil {
		return nil, err
	}
	return id, nil
}

func ParseExternalAddressExternal(external string) (parent *ExternalAddressParent, resourceID string, err error) {
	id := &ExternalAddressIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, "", err
	}
	return id.Parent(), id.ExternalAddress, nil
}
