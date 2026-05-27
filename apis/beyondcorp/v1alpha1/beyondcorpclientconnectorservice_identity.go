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

var _ identity.IdentityV2 = &BeyondCorpClientConnectorServiceIdentity{}

var beyondCorpClientConnectorServiceURL = gcpurls.Template[BeyondCorpClientConnectorServiceIdentity](
	"beyondcorp.googleapis.com",
	"projects/{project}/locations/{location}/clientConnectorServices/{clientConnectorService}",
)

// BeyondCorpClientConnectorServiceIdentity defines the resource reference to BeyondCorpClientConnectorService, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type BeyondCorpClientConnectorServiceIdentity struct {
	Project                string
	Location               string
	ClientConnectorService string
}

func (i *BeyondCorpClientConnectorServiceIdentity) FromExternal(ref string) error {
	out, match, err := beyondCorpClientConnectorServiceURL.Parse(ref)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of BeyondCorpClientConnectorService external=%q was not known (use %s)", ref, beyondCorpClientConnectorServiceURL.CanonicalForm())
	}
	*i = *out
	return nil
}

func (i *BeyondCorpClientConnectorServiceIdentity) String() string {
	return beyondCorpClientConnectorServiceURL.ToString(*i)
}

func (i *BeyondCorpClientConnectorServiceIdentity) Host() string {
	return beyondCorpClientConnectorServiceURL.Host()
}

// NewBuilds a BeyondCorpClientConnectorServiceIdentity from the Config Connector BeyondCorpClientConnectorService object.
func NewBeyondCorpClientConnectorServiceIdentity(ctx context.Context, reader client.Reader, obj *BeyondCorpClientConnectorService) (*BeyondCorpClientConnectorServiceIdentity, error) {
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
		actualIdentity := &BeyondCorpClientConnectorServiceIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actualIdentity.Project != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.Project, projectID)
		}
		if actualIdentity.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualIdentity.Location, location)
		}
		if actualIdentity.ClientConnectorService != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.ClientConnectorService)
		}
	}
	return &BeyondCorpClientConnectorServiceIdentity{
		Project:                projectID,
		Location:               location,
		ClientConnectorService: resourceID,
	}, nil
}
