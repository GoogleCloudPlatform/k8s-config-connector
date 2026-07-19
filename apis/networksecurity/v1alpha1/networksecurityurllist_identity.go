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
	_ identity.IdentityV2 = &NetworkSecurityURLListIdentity{}
	_ identity.Resource   = &NetworkSecurityURLList{}
)

var NetworkSecurityURLListIdentityFormat = gcpurls.Template[NetworkSecurityURLListIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/urlLists/{urllist}")

// NetworkSecurityURLListIdentity is the identity of a GCP NetworkSecurityURLList resource.
// +k8s:deepcopy-gen=false
type NetworkSecurityURLListIdentity struct {
	Project  string
	Location string
	UrlList  string
}

func (i *NetworkSecurityURLListIdentity) String() string {
	return NetworkSecurityURLListIdentityFormat.ToString(*i)
}

func (i *NetworkSecurityURLListIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkSecurityURLListIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkSecurityURLList external=%q was not known (use %s): %w", ref, NetworkSecurityURLListIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkSecurityURLList external=%q was not known (use %s)", ref, NetworkSecurityURLListIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkSecurityURLListIdentity) Host() string {
	return NetworkSecurityURLListIdentityFormat.Host()
}

func (i *NetworkSecurityURLListIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromNetworkSecurityURLListSpec(ctx context.Context, reader client.Reader, obj *NetworkSecurityURLList) (*NetworkSecurityURLListIdentity, error) {
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

	identity := &NetworkSecurityURLListIdentity{
		Project:  projectID,
		Location: location,
		UrlList:  resourceID,
	}
	return identity, nil
}

func (obj *NetworkSecurityURLList) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecurityURLListSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}

// ExternalIdentifier returns the GCP external identifier (the GCP URL).
func (obj *NetworkSecurityURLList) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
