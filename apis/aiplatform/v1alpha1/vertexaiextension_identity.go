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
	_ identity.IdentityV2 = &VertexAIExtensionIdentity{}
	_ identity.Resource   = &VertexAIExtension{}
)

var (
	VertexAIExtensionIdentityFormat         = gcpurls.Template[VertexAIExtensionIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/extensions/{extension}")
	VertexAIExtensionIdentityFormatRelative = gcpurls.Template[VertexAIExtensionIdentity]("", "projects/{project}/locations/{location}/extensions/{extension}")
)

// VertexAIExtensionIdentity is the identity of a GCP VertexAIExtension resource.
// +k8s:deepcopy-gen=false
type VertexAIExtensionIdentity struct {
	Project   string
	Location  string
	Extension string
}

func (i *VertexAIExtensionIdentity) String() string {
	return VertexAIExtensionIdentityFormat.ToString(*i)
}

func (i *VertexAIExtensionIdentity) FromExternal(ref string) error {
	// Normalize the reference by stripping optional scheme, regional hosts, and API versions.
	s := ref
	s = strings.TrimPrefix(s, "https:")
	s = strings.TrimPrefix(s, "http:")
	s = strings.TrimPrefix(s, "//")

	if idx := strings.Index(s, "aiplatform.googleapis.com/"); idx != -1 {
		s = s[idx+len("aiplatform.googleapis.com/"):]
	}

	s = strings.TrimPrefix(s, "v1beta1/")
	s = strings.TrimPrefix(s, "v1/")
	s = strings.TrimPrefix(s, "v1alpha1/")

	parsed, match, err := VertexAIExtensionIdentityFormatRelative.Parse(s)
	if err != nil {
		return fmt.Errorf("format of VertexAIExtension external=%q was not known (use %s): %w", ref, VertexAIExtensionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIExtension external=%q was not known (use %s)", ref, VertexAIExtensionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAIExtensionIdentity) Host() string {
	return VertexAIExtensionIdentityFormat.Host()
}

func (i *VertexAIExtensionIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromVertexAIExtensionSpec(ctx context.Context, reader client.Reader, obj *VertexAIExtension) (*VertexAIExtensionIdentity, error) {
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

	identity := &VertexAIExtensionIdentity{
		Project:   projectID,
		Location:  location,
		Extension: resourceID,
	}
	return identity, nil
}

func (obj *VertexAIExtension) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAIExtensionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAIExtensionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.Project != specIdentity.Project {
			return nil, fmt.Errorf("cannot change VertexAIExtension project (old=%q, new=%q)", statusIdentity.Project, specIdentity.Project)
		}
		if statusIdentity.Location != specIdentity.Location {
			return nil, fmt.Errorf("cannot change VertexAIExtension location (old=%q, new=%q)", statusIdentity.Location, specIdentity.Location)
		}
		return statusIdentity, nil
	}

	return specIdentity, nil
}
