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
	_ identity.IdentityV2 = &VertexAICustomJobIdentity{}
	_ identity.Resource   = &VertexAICustomJob{}
)

var VertexAICustomJobIdentityFormat = gcpurls.Template[VertexAICustomJobIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/customJobs/{customJob}")

// +k8s:deepcopy-gen=false
type VertexAICustomJobIdentity struct {
	Project   string
	Location  string
	CustomJob string
}

func (i *VertexAICustomJobIdentity) String() string {
	return VertexAICustomJobIdentityFormat.ToString(*i)
}

func (i *VertexAICustomJobIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAICustomJobIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAICustomJob external=%q was not known (use %s): %w", ref, VertexAICustomJobIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAICustomJob external=%q was not known (use %s)", ref, VertexAICustomJobIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAICustomJobIdentity) Host() string {
	return VertexAICustomJobIdentityFormat.Host()
}

func getIdentityFromVertexAICustomJobSpec(ctx context.Context, reader client.Reader, obj *VertexAICustomJob) (*VertexAICustomJobIdentity, error) {
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
		return nil, fmt.Errorf("cannot resolve projectID: %w", err)
	}

	identity := &VertexAICustomJobIdentity{
		Project:   projectID,
		Location:  location,
		CustomJob: resourceID,
	}
	return identity, nil
}

func NewVertexAICustomJobIdentity(ctx context.Context, reader client.Reader, obj *VertexAICustomJob) (*VertexAICustomJobIdentity, error) {
	return getIdentityFromVertexAICustomJobSpec(ctx, reader, obj)
}

func (obj *VertexAICustomJob) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAICustomJobSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		if !strings.Contains(externalRef, "/") {
			if externalRef != specIdentity.CustomJob {
				return nil, fmt.Errorf("cannot change VertexAICustomJob identity (old=%q, new=%q)", externalRef, specIdentity.CustomJob)
			}
		} else {
			statusIdentity := &VertexAICustomJobIdentity{}
			if err := statusIdentity.FromExternal(externalRef); err != nil {
				return nil, err
			}

			if statusIdentity.String() != specIdentity.String() {
				return nil, fmt.Errorf("cannot change VertexAICustomJob identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
			}
		}
	}

	return specIdentity, nil
}

func (obj *VertexAICustomJob) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
