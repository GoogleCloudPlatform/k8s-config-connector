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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// DataCatalogPolicyTagIdentityURL is the format for the externalRef of a DataCatalogPolicyTag.
	DataCatalogPolicyTagIdentityURL = "projects/{project}/locations/{location}/taxonomies/{taxonomy}/policyTags/{policyTag}"
)

var (
	_ identity.IdentityV2 = &DataCatalogPolicyTagIdentity{}
	_ identity.Resource   = &DataCatalogPolicyTag{}
)

var DataCatalogPolicyTagIdentityFormat = gcpurls.Template[DataCatalogPolicyTagIdentity](
	"datacatalog.googleapis.com",
	DataCatalogPolicyTagIdentityURL,
)

// DataCatalogPolicyTagIdentity represents the identity of a DataCatalogPolicyTag.
// +k8s:deepcopy-gen=false
type DataCatalogPolicyTagIdentity struct {
	Project   string
	Location  string
	Taxonomy  string
	PolicyTag string
}

func (i *DataCatalogPolicyTagIdentity) String() string {
	return DataCatalogPolicyTagIdentityFormat.ToString(*i)
}

func (i *DataCatalogPolicyTagIdentity) FromExternal(ref string) error {
	parsed, match, err := DataCatalogPolicyTagIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DataCatalogPolicyTag external=%q was not known (use %s): %w", ref, DataCatalogPolicyTagIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DataCatalogPolicyTag external=%q was not known (use %s)", ref, DataCatalogPolicyTagIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DataCatalogPolicyTagIdentity) Host() string {
	return DataCatalogPolicyTagIdentityFormat.Host()
}

func getIdentityFromDataCatalogPolicyTagSpec(ctx context.Context, reader client.Reader, obj client.Object) (*DataCatalogPolicyTagIdentity, error) {
	// Get resource ID
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	// Get taxonomy identity
	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return nil, fmt.Errorf("object was not unstructured")
	}

	taxonomyRefStr, _, err := unstructured.NestedString(u.Object, "spec", "taxonomyRef", "external")
	if err != nil {
		return nil, fmt.Errorf("reading spec.taxonomyRef.external: %w", err)
	}

	if taxonomyRefStr == "" {
		// Try to resolve it if name/namespace are provided
		taxonomyRefName, _, _ := unstructured.NestedString(u.Object, "spec", "taxonomyRef", "name")
		taxonomyRefNamespace, _, _ := unstructured.NestedString(u.Object, "spec", "taxonomyRef", "namespace")
		if taxonomyRefName != "" {
			if taxonomyRefNamespace == "" {
				taxonomyRefNamespace = obj.GetNamespace()
			}
			taxonomyRef := &TaxonomyRef{
				Name:      taxonomyRefName,
				Namespace: taxonomyRefNamespace,
			}
			taxonomyRefStr, err = taxonomyRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
			if err != nil {
				return nil, err
			}
		}
	}

	if taxonomyRefStr == "" {
		return nil, fmt.Errorf("cannot resolve taxonomyRef")
	}

	taxonomyParent, taxonomyID, err := ParseTaxonomyExternal(taxonomyRefStr)
	if err != nil {
		return nil, err
	}

	specIdentity := &DataCatalogPolicyTagIdentity{
		Project:   taxonomyParent.ProjectID,
		Location:  taxonomyParent.Region,
		Taxonomy:  taxonomyID,
		PolicyTag: resourceID,
	}
	return specIdentity, nil
}

func (obj *DataCatalogPolicyTag) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDataCatalogPolicyTagSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DataCatalogPolicyTagIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DataCatalogPolicyTag identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
