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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NetworkServicesEdgeCacheServiceIdentity{}
	_ identity.Resource   = &NetworkServicesEdgeCacheService{}
)

var NetworkServicesEdgeCacheServiceIdentityFormat = gcpurls.Template[NetworkServicesEdgeCacheServiceIdentity]("networkservices.googleapis.com", "projects/{project}/locations/global/edgeCacheServices/{edgeCacheService}")

// NetworkServicesEdgeCacheServiceIdentity is the identity of a GCP NetworkServicesEdgeCacheService resource.
// +k8s:deepcopy-gen=false
type NetworkServicesEdgeCacheServiceIdentity struct {
	Project          string
	EdgeCacheService string
}

func (i *NetworkServicesEdgeCacheServiceIdentity) String() string {
	return NetworkServicesEdgeCacheServiceIdentityFormat.ToString(*i)
}

func (i *NetworkServicesEdgeCacheServiceIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkServicesEdgeCacheServiceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkServicesEdgeCacheService external=%q was not known (use %s): %w", ref, NetworkServicesEdgeCacheServiceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkServicesEdgeCacheService external=%q was not known (use %s)", ref, NetworkServicesEdgeCacheServiceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkServicesEdgeCacheServiceIdentity) Host() string {
	return NetworkServicesEdgeCacheServiceIdentityFormat.Host()
}

func (i *NetworkServicesEdgeCacheServiceIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/global"
}

func NewNetworkServicesEdgeCacheServiceIdentity(ctx context.Context, reader client.Reader, obj *NetworkServicesEdgeCacheService) (*NetworkServicesEdgeCacheServiceIdentity, error) {
	resourceID, err := refsv1beta1.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &NetworkServicesEdgeCacheServiceIdentity{
		Project:          projectID,
		EdgeCacheService: resourceID,
	}
	return identity, nil
}

func (obj *NetworkServicesEdgeCacheService) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := NewNetworkServicesEdgeCacheServiceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	return specIdentity, nil
}
