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
	_ identity.IdentityV2 = &NetAppActiveDirectoryIdentity{}
	_ identity.Resource   = &NetAppActiveDirectory{}
)

var NetAppActiveDirectoryIdentityFormat = gcpurls.Template[NetAppActiveDirectoryIdentity]("netapp.googleapis.com", "projects/{project}/locations/{location}/activeDirectories/{activeDirectory}")

// NetAppActiveDirectoryIdentity is the identity of a GCP NetAppActiveDirectory resource.
// +k8s:deepcopy-gen=false
type NetAppActiveDirectoryIdentity struct {
	Project         string
	Location        string
	ActiveDirectory string
}

func (i *NetAppActiveDirectoryIdentity) String() string {
	return NetAppActiveDirectoryIdentityFormat.ToString(*i)
}

func (i *NetAppActiveDirectoryIdentity) FromExternal(ref string) error {
	parsed, match, err := NetAppActiveDirectoryIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetAppActiveDirectory external=%q was not known (use %s): %w", ref, NetAppActiveDirectoryIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetAppActiveDirectory external=%q was not known (use %s)", ref, NetAppActiveDirectoryIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetAppActiveDirectoryIdentity) Host() string {
	return NetAppActiveDirectoryIdentityFormat.Host()
}

func (i *NetAppActiveDirectoryIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromNetAppActiveDirectorySpec(ctx context.Context, reader client.Reader, obj *NetAppActiveDirectory) (*NetAppActiveDirectoryIdentity, error) {
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

	id := &NetAppActiveDirectoryIdentity{
		Project:         projectID,
		Location:        location,
		ActiveDirectory: resourceID,
	}
	return id, nil
}

func (obj *NetAppActiveDirectory) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetAppActiveDirectorySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &NetAppActiveDirectoryIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetAppActiveDirectory identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
