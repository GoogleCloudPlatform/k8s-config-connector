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
	_ identity.IdentityV2 = &NetworkSecurityUrlListIdentity{}
	_ identity.Resource   = &NetworkSecurityUrlList{}
)

var NetworkSecurityUrlListIdentityFormat = gcpurls.Template[NetworkSecurityUrlListIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/urlLists/{url_list}")

// +k8s:deepcopy-gen=false
type NetworkSecurityUrlListIdentity struct {
	Project  string
	Location string
	UrlList  string
}

func (i *NetworkSecurityUrlListIdentity) String() string {
	return NetworkSecurityUrlListIdentityFormat.ToString(*i)
}

func (i *NetworkSecurityUrlListIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkSecurityUrlListIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkSecurityUrlList external=%q was not known (use %s): %w", ref, NetworkSecurityUrlListIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkSecurityUrlList external=%q was not known (use %s)", ref, NetworkSecurityUrlListIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkSecurityUrlListIdentity) Host() string {
	return NetworkSecurityUrlListIdentityFormat.Host()
}

func getIdentityFromNetworkSecurityUrlListSpec(ctx context.Context, reader client.Reader, obj client.Object) (*NetworkSecurityUrlListIdentity, error) {
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

	identity := &NetworkSecurityUrlListIdentity{
		Project:  projectID,
		Location: location,
		UrlList:  resourceID,
	}
	return identity, nil
}

func (obj *NetworkSecurityUrlList) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecurityUrlListSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &NetworkSecurityUrlListIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkSecurityUrlList identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
