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
	_ identity.IdentityV2 = &ComputeURLMapIdentity{}
	_ identity.Resource   = &ComputeURLMap{}
)

var ComputeURLMapGlobalIdentityFormat = gcpurls.Template[ComputeURLMapIdentity]("compute.googleapis.com", "projects/{project}/global/urlMaps/{urlMap}")
var ComputeURLMapRegionalIdentityFormat = gcpurls.Template[ComputeURLMapIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/urlMaps/{urlMap}")

// +k8s:deepcopy-gen=false
type ComputeURLMapIdentity struct {
	Project string
	Region  string
	UrlMap  string
}

func (i *ComputeURLMapIdentity) String() string {
	if i.Region == "" {
		return ComputeURLMapGlobalIdentityFormat.ToString(*i)
	}
	return ComputeURLMapRegionalIdentityFormat.ToString(*i)
}

func (i *ComputeURLMapIdentity) FromExternal(ref string) error {
	if parsed, match, err := ComputeURLMapRegionalIdentityFormat.Parse(ref); match {
		if err != nil {
			return err
		}
		*i = *parsed
		return nil
	}
	if parsed, match, err := ComputeURLMapGlobalIdentityFormat.Parse(ref); match {
		if err != nil {
			return err
		}
		*i = *parsed
		return nil
	}

	return fmt.Errorf("format of ComputeURLMap external=%q was not known (use %s or %s)", ref, ComputeURLMapGlobalIdentityFormat.CanonicalForm(), ComputeURLMapRegionalIdentityFormat.CanonicalForm())
}

func (i *ComputeURLMapIdentity) Host() string {
	return ComputeURLMapGlobalIdentityFormat.Host()
}

func getIdentityFromComputeURLMapSpec(ctx context.Context, reader client.Reader, obj client.Object) (*ComputeURLMapIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	identity := &ComputeURLMapIdentity{
		Project: projectID,
		UrlMap:  resourceID,
	}
	if location != "global" {
		identity.Region = location
	}
	return identity, nil
}

func (obj *ComputeURLMap) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeURLMapSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ComputeURLMapIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeURLMap identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
