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
	_ identity.IdentityV2 = &CloudNumberRegistryRealmIdentity{}
	_ identity.Resource   = &CloudNumberRegistryRealm{}
)

var CloudNumberRegistryRealmIdentityFormat = gcpurls.Template[CloudNumberRegistryRealmIdentity]("cloudnumberregistry.googleapis.com", "projects/{project}/locations/{location}/realms/{realm}")

// CloudNumberRegistryRealmIdentity is the identity of a GCP CloudNumberRegistryRealm resource.
// +k8s:deepcopy-gen=false
type CloudNumberRegistryRealmIdentity struct {
	Project  string
	Location string
	Realm    string
}

func (i *CloudNumberRegistryRealmIdentity) String() string {
	return CloudNumberRegistryRealmIdentityFormat.ToString(*i)
}

func (i *CloudNumberRegistryRealmIdentity) FromExternal(ref string) error {
	parsed, match, err := CloudNumberRegistryRealmIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CloudNumberRegistryRealm external=%q was not known (use %s): %w", ref, CloudNumberRegistryRealmIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CloudNumberRegistryRealm external=%q was not known (use %s)", ref, CloudNumberRegistryRealmIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CloudNumberRegistryRealmIdentity) Host() string {
	return CloudNumberRegistryRealmIdentityFormat.Host()
}

func (i *CloudNumberRegistryRealmIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromCloudNumberRegistryRealmSpec(ctx context.Context, reader client.Reader, obj *CloudNumberRegistryRealm) (*CloudNumberRegistryRealmIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := common.ValueOf(obj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &CloudNumberRegistryRealmIdentity{
		Project:  projectID,
		Location: location,
		Realm:    resourceID,
	}
	return identity, nil
}

func (obj *CloudNumberRegistryRealm) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCloudNumberRegistryRealmSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &CloudNumberRegistryRealmIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change CloudNumberRegistryRealm identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
