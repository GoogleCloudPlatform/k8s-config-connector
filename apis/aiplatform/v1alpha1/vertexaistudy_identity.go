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
	_ identity.IdentityV2 = &VertexAIStudyIdentity{}
	_ identity.Resource   = &VertexAIStudy{}
)

var VertexAIStudyIdentityFormat = gcpurls.Template[VertexAIStudyIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/studies/{study}")

// +k8s:deepcopy-gen=false
type VertexAIStudyIdentity struct {
	Project  string
	Location string
	Study    string
}

func (i *VertexAIStudyIdentity) String() string {
	return VertexAIStudyIdentityFormat.ToString(*i)
}

func (i *VertexAIStudyIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func (i *VertexAIStudyIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIStudyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIStudy external=%q was not known (use %s): %w", ref, VertexAIStudyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIStudy external=%q was not known (use %s)", ref, VertexAIStudyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAIStudyIdentity) Host() string {
	return VertexAIStudyIdentityFormat.Host()
}

func getIdentityFromVertexAIStudySpec(ctx context.Context, reader client.Reader, obj client.Object) (*VertexAIStudyIdentity, error) {
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

	identity := &VertexAIStudyIdentity{
		Project:  projectID,
		Location: location,
		Study:    resourceID,
	}
	return identity, nil
}

func (obj *VertexAIStudy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAIStudySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAIStudyIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		return statusIdentity, nil
	}

	return specIdentity, nil
}
