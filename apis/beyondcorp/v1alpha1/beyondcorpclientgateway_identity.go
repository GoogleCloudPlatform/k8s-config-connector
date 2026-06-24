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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BeyondCorpClientGatewayIdentity{}
	_ identity.Resource   = &BeyondCorpClientGateway{}
)

var BeyondCorpClientGatewayIdentityFormat = gcpurls.Template[BeyondCorpClientGatewayIdentity]("beyondcorp.googleapis.com", "projects/{project}/locations/{location}/clientGateways/{clientGateway}")

// +k8s:deepcopy-gen=false
type BeyondCorpClientGatewayIdentity struct {
	Project       string
	Location      string
	ClientGateway string
}

func (i *BeyondCorpClientGatewayIdentity) String() string {
	return BeyondCorpClientGatewayIdentityFormat.ToString(*i)
}

func (i *BeyondCorpClientGatewayIdentity) FromExternal(ref string) error {
	parsed, match, err := BeyondCorpClientGatewayIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BeyondCorpClientGateway external=%q was not known (use %s): %w", ref, BeyondCorpClientGatewayIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BeyondCorpClientGateway external=%q was not known (use %s)", ref, BeyondCorpClientGatewayIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BeyondCorpClientGatewayIdentity) Host() string {
	return BeyondCorpClientGatewayIdentityFormat.Host()
}

func ParseClientGatewayIdentity(external string) (*BeyondCorpClientGatewayIdentity, error) {
	id := &BeyondCorpClientGatewayIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func (i *BeyondCorpClientGatewayIdentity) ID() string {
	return i.ClientGateway
}

func getIdentityFromBeyondCorpClientGatewaySpec(ctx context.Context, reader client.Reader, obj client.Object) (*BeyondCorpClientGatewayIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &BeyondCorpClientGatewayIdentity{
		Project:       projectID,
		Location:      location,
		ClientGateway: resourceID,
	}
	return identity, nil
}

func (obj *BeyondCorpClientGateway) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBeyondCorpClientGatewaySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &BeyondCorpClientGatewayIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BeyondCorpClientGateway identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
