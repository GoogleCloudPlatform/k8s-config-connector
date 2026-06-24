// Copyright 2025 Google LLC
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
	_ identity.IdentityV2 = &BigLakeCatalogIdentity{}
	_ identity.Resource   = &BigLakeCatalog{}
)

var BigLakeCatalogIdentityFormat = gcpurls.Template[BigLakeCatalogIdentity]("biglake.googleapis.com", "projects/{project}/locations/{location}/catalogs/{catalog}")

// BigLakeCatalogIdentity is the identity of a BigLakeCatalog.
// +k8s:deepcopy-gen=false
type BigLakeCatalogIdentity struct {
	Project  string
	Location string
	Catalog  string
}

func (i *BigLakeCatalogIdentity) String() string {
	return BigLakeCatalogIdentityFormat.ToString(*i)
}

func (i *BigLakeCatalogIdentity) FromExternal(ref string) error {
	parsed, match, err := BigLakeCatalogIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigLakeCatalog external=%q was not known (use %s): %w", ref, BigLakeCatalogIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigLakeCatalog external=%q was not known (use %s)", ref, BigLakeCatalogIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BigLakeCatalogIdentity) Host() string {
	return BigLakeCatalogIdentityFormat.Host()
}

func getIdentityFromBigLakeCatalogSpec(ctx context.Context, reader client.Reader, obj *BigLakeCatalog) (*BigLakeCatalogIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &BigLakeCatalogIdentity{
		Project:  projectID,
		Location: location,
		Catalog:  resourceID,
	}
	return identity, nil
}

func (obj *BigLakeCatalog) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigLakeCatalogSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Validate against the ID stored in status.externalRef, if any
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &BigLakeCatalogIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, specIdentity.String())
		}
	}
	return specIdentity, nil
}
