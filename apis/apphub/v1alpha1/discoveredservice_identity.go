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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &AppHubDiscoveredServiceIdentity{}
	_ identity.Resource   = &AppHubDiscoveredService{}
)

var AppHubDiscoveredServiceIdentityFormat = gcpurls.Template[AppHubDiscoveredServiceIdentity]("apphub.googleapis.com", "projects/{project}/locations/{location}/discoveredServices/{discoveredService}")

// +k8s:deepcopy-gen=false
type AppHubDiscoveredServiceIdentity struct {
	Project           string
	Location          string
	DiscoveredService string
}

func (i *AppHubDiscoveredServiceIdentity) String() string {
	return AppHubDiscoveredServiceIdentityFormat.ToString(*i)
}

func (i *AppHubDiscoveredServiceIdentity) FromExternal(ref string) error {
	normalizedRef := ref
	if strings.Contains(normalizedRef, "/discoveredservices/") {
		normalizedRef = strings.ReplaceAll(normalizedRef, "/discoveredservices/", "/discoveredServices/")
	}
	parsed, match, err := AppHubDiscoveredServiceIdentityFormat.Parse(normalizedRef)
	if err != nil {
		return fmt.Errorf("format of AppHubDiscoveredService external=%q was not known (use %s): %w", ref, AppHubDiscoveredServiceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AppHubDiscoveredService external=%q was not known (use %s)", ref, AppHubDiscoveredServiceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *AppHubDiscoveredServiceIdentity) Host() string {
	return AppHubDiscoveredServiceIdentityFormat.Host()
}

func ParseAppHubDiscoveredServiceIdentity(external string) (*AppHubDiscoveredServiceIdentity, error) {
	id := &AppHubDiscoveredServiceIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func (i *AppHubDiscoveredServiceIdentity) ID() string {
	return i.DiscoveredService
}

func getIdentityFromAppHubDiscoveredServiceSpec(ctx context.Context, reader client.Reader, obj client.Object) (*AppHubDiscoveredServiceIdentity, error) {
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

	identity := &AppHubDiscoveredServiceIdentity{
		Project:           projectID,
		Location:          location,
		DiscoveredService: resourceID,
	}
	return identity, nil
}

func (obj *AppHubDiscoveredService) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromAppHubDiscoveredServiceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &AppHubDiscoveredServiceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change AppHubDiscoveredService identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *AppHubDiscoveredService) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
