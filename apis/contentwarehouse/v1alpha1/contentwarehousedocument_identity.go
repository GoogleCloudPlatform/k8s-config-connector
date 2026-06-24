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
	_ identity.IdentityV2 = &ContentWarehouseDocumentIdentity{}
	_ identity.Resource   = &ContentWarehouseDocument{}
)

var ContentWarehouseDocumentIdentityFormat = gcpurls.Template[ContentWarehouseDocumentIdentity]("contentwarehouse.googleapis.com", "projects/{project}/locations/{location}/documents/{document}")

// +k8s:deepcopy-gen=false
type ContentWarehouseDocumentIdentity struct {
	Project  string
	Location string
	Document string
}

func (i *ContentWarehouseDocumentIdentity) String() string {
	return ContentWarehouseDocumentIdentityFormat.ToString(*i)
}

func (i *ContentWarehouseDocumentIdentity) FromExternal(ref string) error {
	parsed, match, err := ContentWarehouseDocumentIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ContentWarehouseDocument external=%q was not known (use %s): %w", ref, ContentWarehouseDocumentIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ContentWarehouseDocument external=%q was not known (use %s)", ref, ContentWarehouseDocumentIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ContentWarehouseDocumentIdentity) Host() string {
	return ContentWarehouseDocumentIdentityFormat.Host()
}

func getIdentityFromContentWarehouseDocumentSpec(ctx context.Context, reader client.Reader, obj client.Object) (*ContentWarehouseDocumentIdentity, error) {
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

	identity := &ContentWarehouseDocumentIdentity{
		Project:  projectID,
		Location: location,
		Document: resourceID,
	}
	return identity, nil
}

func (obj *ContentWarehouseDocument) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromContentWarehouseDocumentSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &ContentWarehouseDocumentIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ContentWarehouseDocument identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *ContentWarehouseDocument) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
