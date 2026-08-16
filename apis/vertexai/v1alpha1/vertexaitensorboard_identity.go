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
	_ identity.IdentityV2 = &VertexAITensorboardIdentity{}
	_ identity.Resource   = &VertexAITensorboard{}
)

var VertexAITensorboardIdentityFormat = gcpurls.Template[VertexAITensorboardIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/tensorboards/{tensorboard}")

// VertexAITensorboardIdentity is the identity of a GCP VertexAITensorboard resource.
// +k8s:deepcopy-gen=false
type VertexAITensorboardIdentity struct {
	Project     string
	Location    string
	Tensorboard string
}

func (i *VertexAITensorboardIdentity) String() string {
	return VertexAITensorboardIdentityFormat.ToString(*i)
}

func (i *VertexAITensorboardIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAITensorboardIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAITensorboard external=%q was not known (use %s): %w", ref, VertexAITensorboardIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAITensorboard external=%q was not known (use %s)", ref, VertexAITensorboardIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAITensorboardIdentity) Host() string {
	return VertexAITensorboardIdentityFormat.Host()
}

func (i *VertexAITensorboardIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromVertexAITensorboardSpec(ctx context.Context, reader client.Reader, obj *VertexAITensorboard) (*VertexAITensorboardIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// If resourceID is a full GCP URI, parse the actual tensorboard name from it.
	if strings.HasPrefix(resourceID, "projects/") {
		parsed, match, err := VertexAITensorboardIdentityFormat.Parse(resourceID)
		if err == nil && match {
			resourceID = parsed.Tensorboard
		}
	}

	location := obj.Spec.Region
	if location == "" {
		return nil, fmt.Errorf("cannot resolve region")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &VertexAITensorboardIdentity{
		Project:     projectID,
		Location:    location,
		Tensorboard: resourceID,
	}
	return identity, nil
}

func (obj *VertexAITensorboard) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAITensorboardSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAITensorboardIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		// If resourceID is not set in spec, then specIdentity.Tensorboard was defaulted to obj.GetName()
		// and the actual resource was created with a server-generated ID.
		// In this case, we update specIdentity.Tensorboard to match the server-generated ID.
		if obj.Spec.ResourceID == nil || *obj.Spec.ResourceID == "" {
			specIdentity.Tensorboard = statusIdentity.Tensorboard
		}

		// Validate identity fields. We avoid direct string comparison of the full Identity because Project
		// can be project ID string in specIdentity, but project number in statusIdentity.
		if statusIdentity.Location != specIdentity.Location || statusIdentity.Tensorboard != specIdentity.Tensorboard {
			return nil, fmt.Errorf("cannot change VertexAITensorboard identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
