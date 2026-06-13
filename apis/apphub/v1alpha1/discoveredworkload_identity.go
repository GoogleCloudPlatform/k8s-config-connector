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
	_ identity.IdentityV2 = &DiscoveredWorkloadIdentity{}
	_ identity.Resource   = &AppHubDiscoveredWorkload{}
)

var DiscoveredWorkloadIdentityFormat = gcpurls.Template[DiscoveredWorkloadIdentity]("apphub.googleapis.com", "projects/{project}/locations/{location}/discoveredWorkloads/{discoveredWorkload}")

// +k8s:deepcopy-gen=false
type DiscoveredWorkloadIdentity struct {
	Project            string
	Location           string
	DiscoveredWorkload string
}

func (i *DiscoveredWorkloadIdentity) String() string {
	return DiscoveredWorkloadIdentityFormat.ToString(*i)
}

func (i *DiscoveredWorkloadIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveredWorkloadIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AppHubDiscoveredWorkload external=%q was not known (use %s): %w", ref, DiscoveredWorkloadIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AppHubDiscoveredWorkload external=%q was not known (use %s)", ref, DiscoveredWorkloadIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveredWorkloadIdentity) Host() string {
	return DiscoveredWorkloadIdentityFormat.Host()
}

func getIdentityFromDiscoveredWorkloadSpec(ctx context.Context, reader client.Reader, obj client.Object) (*DiscoveredWorkloadIdentity, error) {
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

	identity := &DiscoveredWorkloadIdentity{
		Project:            projectID,
		Location:           location,
		DiscoveredWorkload: resourceID,
	}
	return identity, nil
}

func (obj *AppHubDiscoveredWorkload) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveredWorkloadSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveredWorkloadIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change AppHubDiscoveredWorkload identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *AppHubDiscoveredWorkload) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
