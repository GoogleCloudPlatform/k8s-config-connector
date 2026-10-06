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
	_ identity.IdentityV2 = &NetAppKMSConfigIdentity{}
	_ identity.Resource   = &NetAppKMSConfig{}
)

var NetAppKMSConfigIdentityFormat = gcpurls.Template[NetAppKMSConfigIdentity]("netapp.googleapis.com", "projects/{project}/locations/{location}/kmsConfigs/{kmsConfig}")

// NetAppKMSConfigIdentity is the identity of a GCP NetAppKMSConfig resource.
// +k8s:deepcopy-gen=false
type NetAppKMSConfigIdentity struct {
	Project   string
	Location  string
	KMSConfig string
}

func (i *NetAppKMSConfigIdentity) String() string {
	return NetAppKMSConfigIdentityFormat.ToString(*i)
}

func (i *NetAppKMSConfigIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func (i *NetAppKMSConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := NetAppKMSConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetAppKMSConfig external=%q was not known (use %s): %w", ref, NetAppKMSConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetAppKMSConfig external=%q was not known (use %s)", ref, NetAppKMSConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetAppKMSConfigIdentity) Host() string {
	return NetAppKMSConfigIdentityFormat.Host()
}

func getIdentityFromNetAppKMSConfigSpec(ctx context.Context, reader client.Reader, obj *NetAppKMSConfig) (*NetAppKMSConfigIdentity, error) {
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

	identity := &NetAppKMSConfigIdentity{
		Project:   projectID,
		Location:  location,
		KMSConfig: resourceID,
	}
	return identity, nil
}

func (obj *NetAppKMSConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetAppKMSConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &NetAppKMSConfigIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetAppKMSConfig identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
