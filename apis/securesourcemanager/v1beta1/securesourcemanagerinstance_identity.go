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

package v1beta1

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
	_ identity.IdentityV2 = &SecureSourceManagerInstanceIdentity{}
	_ identity.Resource   = &SecureSourceManagerInstance{}
)

var SecureSourceManagerInstanceIdentityFormat = gcpurls.Template[SecureSourceManagerInstanceIdentity]("securesourcemanager.googleapis.com", "projects/{project}/locations/{location}/instances/{instance}")

// SecureSourceManagerInstanceIdentity is the identity of a GCP SecureSourceManagerInstance resource.
// +k8s:deepcopy-gen=false
type SecureSourceManagerInstanceIdentity struct {
	Project  string
	Location string
	Instance string
}

func (i *SecureSourceManagerInstanceIdentity) String() string {
	return SecureSourceManagerInstanceIdentityFormat.ToString(*i)
}

func (i *SecureSourceManagerInstanceIdentity) FromExternal(ref string) error {
	parsed, match, err := SecureSourceManagerInstanceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of SecureSourceManagerInstance external=%q was not known (use %s): %w", ref, SecureSourceManagerInstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of SecureSourceManagerInstance external=%q was not known (use %s)", ref, SecureSourceManagerInstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *SecureSourceManagerInstanceIdentity) Host() string {
	return SecureSourceManagerInstanceIdentityFormat.Host()
}

func (i *SecureSourceManagerInstanceIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromSecureSourceManagerInstanceSpec(ctx context.Context, reader client.Reader, obj *SecureSourceManagerInstance) (*SecureSourceManagerInstanceIdentity, error) {
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

	identity := &SecureSourceManagerInstanceIdentity{
		Project:  projectID,
		Location: location,
		Instance: resourceID,
	}
	return identity, nil
}

func (obj *SecureSourceManagerInstance) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromSecureSourceManagerInstanceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &SecureSourceManagerInstanceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change SecureSourceManagerInstance identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
