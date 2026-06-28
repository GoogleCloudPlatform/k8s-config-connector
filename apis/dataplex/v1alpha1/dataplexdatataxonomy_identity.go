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
	_ identity.IdentityV2 = &DataTaxonomyIdentity{}
	_ identity.Resource   = &DataplexDataTaxonomy{}
)

var DataTaxonomyIdentityFormat = gcpurls.Template[DataTaxonomyIdentity]("dataplex.googleapis.com", "projects/{project}/locations/{location}/dataTaxonomies/{dataTaxonomy}")

// +k8s:deepcopy-gen=false
type DataTaxonomyIdentity struct {
	Project      string
	Location     string
	DataTaxonomy string
}

func (i *DataTaxonomyIdentity) String() string {
	return DataTaxonomyIdentityFormat.ToString(*i)
}

func (i *DataTaxonomyIdentity) FromExternal(ref string) error {
	parsed, match, err := DataTaxonomyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DataplexDataTaxonomy external=%q was not known (use %s): %w", ref, DataTaxonomyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DataplexDataTaxonomy external=%q was not known (use %s)", ref, DataTaxonomyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DataTaxonomyIdentity) Host() string {
	return DataTaxonomyIdentityFormat.Host()
}

func getIdentityFromDataplexDataTaxonomySpec(ctx context.Context, reader client.Reader, obj client.Object) (*DataTaxonomyIdentity, error) {
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

	identity := &DataTaxonomyIdentity{
		Project:      projectID,
		Location:     location,
		DataTaxonomy: resourceID,
	}
	return identity, nil
}

func (obj *DataplexDataTaxonomy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDataplexDataTaxonomySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DataTaxonomyIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DataplexDataTaxonomy identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
