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
	_ identity.IdentityV2 = &DiscoveryEngineDataConnectorIdentity{}
	_ identity.Resource   = &DiscoveryEngineDataConnector{}
)

var DiscoveryEngineDataConnectorIdentityFormat = gcpurls.Template[DiscoveryEngineDataConnectorIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/collections/{collection}/dataConnector")

// DiscoveryEngineDataConnectorIdentity is the identity of a GCP DiscoveryEngineDataConnector resource.
// +k8s:deepcopy-gen=false
type DiscoveryEngineDataConnectorIdentity struct {
	Project    string
	Location   string
	Collection string
}

func (i *DiscoveryEngineDataConnectorIdentity) String() string {
	return DiscoveryEngineDataConnectorIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineDataConnectorIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineDataConnectorIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineDataConnector external=%q was not known (use %s): %w", ref, DiscoveryEngineDataConnectorIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineDataConnector external=%q was not known (use %s)", ref, DiscoveryEngineDataConnectorIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineDataConnectorIdentity) Host() string {
	return DiscoveryEngineDataConnectorIdentityFormat.Host()
}

func (i *DiscoveryEngineDataConnectorIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/collections/%s", i.Project, i.Location, i.Collection)
}

func getIdentityFromDiscoveryEngineDataConnectorSpec(ctx context.Context, reader client.Reader, obj client.Object) (*DiscoveryEngineDataConnectorIdentity, error) {
	connector, ok := obj.(*DiscoveryEngineDataConnector)
	if !ok {
		return nil, fmt.Errorf("object is not a DiscoveryEngineDataConnector")
	}

	if connector.Spec.Location == nil || *connector.Spec.Location == "" {
		return nil, fmt.Errorf("cannot resolve location: must be specified")
	}
	location := *connector.Spec.Location

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	if connector.Spec.Collection == nil || *connector.Spec.Collection == "" {
		return nil, fmt.Errorf("cannot resolve collection: must be specified")
	}
	collection := *connector.Spec.Collection

	identity := &DiscoveryEngineDataConnectorIdentity{
		Project:    projectID,
		Location:   location,
		Collection: collection,
	}
	return identity, nil
}

func (obj *DiscoveryEngineDataConnector) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineDataConnectorSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if obj.Status.ExternalRef != nil {
		statusIdentity := &DiscoveryEngineDataConnectorIdentity{}
		if err := statusIdentity.FromExternal(*obj.Status.ExternalRef); err != nil {
			return nil, err
		}

		if statusIdentity.Location != specIdentity.Location || statusIdentity.Collection != specIdentity.Collection {
			return nil, fmt.Errorf("cannot change DiscoveryEngineDataConnector parent identity (old=%q, new parent=%s/%s/%s)", statusIdentity.String(), specIdentity.Project, specIdentity.Location, specIdentity.Collection)
		}

		return statusIdentity, nil
	}

	return specIdentity, nil
}
