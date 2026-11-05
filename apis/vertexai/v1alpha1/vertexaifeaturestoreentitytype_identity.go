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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &VertexAIFeaturestoreEntityTypeIdentity{}
	_ identity.Resource   = &VertexAIFeaturestoreEntityType{}
)

var VertexAIFeaturestoreEntityTypeIdentityFormat = gcpurls.Template[VertexAIFeaturestoreEntityTypeIdentity](
	"aiplatform.googleapis.com",
	"projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entityType}",
)

var featurestoreParentFormat = gcpurls.Template[entityTypeParentFeaturestore](
	"aiplatform.googleapis.com",
	"projects/{project}/locations/{location}/featurestores/{featurestore}",
)

type entityTypeParentFeaturestore struct {
	Project      string
	Location     string
	Featurestore string
}

// VertexAIFeaturestoreEntityTypeIdentity is the identity of a GCP VertexAIFeaturestoreEntityType resource.
// +k8s:deepcopy-gen=false
type VertexAIFeaturestoreEntityTypeIdentity struct {
	Project      string
	Location     string
	Featurestore string
	EntityType   string
}

func (i *VertexAIFeaturestoreEntityTypeIdentity) String() string {
	return VertexAIFeaturestoreEntityTypeIdentityFormat.ToString(*i)
}

func (i *VertexAIFeaturestoreEntityTypeIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIFeaturestoreEntityTypeIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("parsing external %q: %w", ref, err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIFeaturestoreEntityType external=%q was not known (use %s)", ref, VertexAIFeaturestoreEntityTypeIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *VertexAIFeaturestoreEntityTypeIdentity) Host() string {
	return VertexAIFeaturestoreEntityTypeIdentityFormat.Host()
}

func (i *VertexAIFeaturestoreEntityTypeIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", i.Project, i.Location, i.Featurestore)
}

func getIdentityFromVertexAIFeaturestoreEntityTypeSpec(ctx context.Context, reader client.Reader, obj *VertexAIFeaturestoreEntityType) (*VertexAIFeaturestoreEntityTypeIdentity, error) {
	if obj.Spec.Featurestore == "" {
		return nil, fmt.Errorf("spec.featurestore is required")
	}

	parent, match, err := featurestoreParentFormat.Parse(obj.Spec.Featurestore)
	if err != nil {
		return nil, fmt.Errorf("parsing spec.featurestore %q: %w", obj.Spec.Featurestore, err)
	}
	if !match {
		return nil, fmt.Errorf("format of VertexAIFeaturestoreEntityType parent featurestore=%q was not known (use %s)", obj.Spec.Featurestore, featurestoreParentFormat.CanonicalForm())
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	return &VertexAIFeaturestoreEntityTypeIdentity{
		Project:      parent.Project,
		Location:     parent.Location,
		Featurestore: parent.Featurestore,
		EntityType:   resourceID,
	}, nil
}

func (obj *VertexAIFeaturestoreEntityType) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromVertexAIFeaturestoreEntityTypeSpec(ctx, reader, obj)
}
