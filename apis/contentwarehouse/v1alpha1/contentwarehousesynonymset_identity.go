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
	_ identity.IdentityV2 = &ContentWarehouseSynonymSetIdentity{}
	_ identity.Resource   = &ContentWarehouseSynonymSet{}
)

var ContentWarehouseSynonymSetIdentityFormat = gcpurls.Template[ContentWarehouseSynonymSetIdentity]("contentwarehouse.googleapis.com", "projects/{project}/locations/{location}/synonymSets/{context}")

// ContentWarehouseSynonymSetIdentity is the identity of a GCP ContentWarehouseSynonymSet resource.
// +k8s:deepcopy-gen=false
type ContentWarehouseSynonymSetIdentity struct {
	Project  string
	Location string
	Context  string
}

func (i *ContentWarehouseSynonymSetIdentity) String() string {
	return ContentWarehouseSynonymSetIdentityFormat.ToString(*i)
}

func (i *ContentWarehouseSynonymSetIdentity) FromExternal(ref string) error {
	parsed, match, err := ContentWarehouseSynonymSetIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ContentWarehouseSynonymSet external=%q was not known (use %s): %w", ref, ContentWarehouseSynonymSetIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ContentWarehouseSynonymSet external=%q was not known (use %s)", ref, ContentWarehouseSynonymSetIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ContentWarehouseSynonymSetIdentity) Host() string {
	return ContentWarehouseSynonymSetIdentityFormat.Host()
}

func getIdentityFromContentWarehouseSynonymSetSpec(ctx context.Context, reader client.Reader, obj *ContentWarehouseSynonymSet) (*ContentWarehouseSynonymSetIdentity, error) {
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

	identity := &ContentWarehouseSynonymSetIdentity{
		Project:  projectID,
		Location: location,
		Context:  resourceID,
	}
	return identity, nil
}

func (obj *ContentWarehouseSynonymSet) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromContentWarehouseSynonymSetSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ContentWarehouseSynonymSetIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ContentWarehouseSynonymSet identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
