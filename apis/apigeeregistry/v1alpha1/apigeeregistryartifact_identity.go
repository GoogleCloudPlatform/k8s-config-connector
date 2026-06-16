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
	_ identity.IdentityV2 = &ApigeeRegistryArtifactIdentity{}
	_ identity.Resource   = &ApigeeRegistryArtifact{}
)

var ApigeeRegistryArtifactIdentityFormat = gcpurls.Template[ApigeeRegistryArtifactIdentity]("apigeeregistry.googleapis.com", "projects/{project}/locations/{location}/artifacts/{artifact}")

// +k8s:deepcopy-gen=false

// ApigeeRegistryArtifactIdentity is the identity of a Google Cloud ApigeeRegistryArtifact resource.
type ApigeeRegistryArtifactIdentity struct {
	Project  string
	Location string
	Artifact string
}

func (i *ApigeeRegistryArtifactIdentity) String() string {
	return ApigeeRegistryArtifactIdentityFormat.ToString(*i)
}

func (i *ApigeeRegistryArtifactIdentity) FromExternal(ref string) error {
	parsed, match, err := ApigeeRegistryArtifactIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ApigeeRegistryArtifact external=%q was not known (use %s): %w", ref, ApigeeRegistryArtifactIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ApigeeRegistryArtifact external=%q was not known (use %s)", ref, ApigeeRegistryArtifactIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ApigeeRegistryArtifactIdentity) Host() string {
	return ApigeeRegistryArtifactIdentityFormat.Host()
}

func getIdentityFromApigeeRegistryArtifactSpec(ctx context.Context, reader client.Reader, obj *ApigeeRegistryArtifact) (*ApigeeRegistryArtifactIdentity, error) {
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

	identity := &ApigeeRegistryArtifactIdentity{
		Project:  projectID,
		Location: location,
		Artifact: resourceID,
	}
	return identity, nil
}

func (obj *ApigeeRegistryArtifact) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromApigeeRegistryArtifactSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ApigeeRegistryArtifactIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ApigeeRegistryArtifact identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
