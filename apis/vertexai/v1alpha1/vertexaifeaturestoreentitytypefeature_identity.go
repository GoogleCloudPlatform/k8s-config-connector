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
	_ identity.IdentityV2 = &VertexAIFeaturestoreEntityTypeFeatureIdentity{}
	_ identity.Resource   = &VertexAIFeaturestoreEntityTypeFeature{}
)

var VertexAIFeaturestoreEntityTypeFeatureIdentityFormat = gcpurls.Template[VertexAIFeaturestoreEntityTypeFeatureIdentity](
	"aiplatform.googleapis.com",
	"projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entityType}/features/{feature}",
)

var entityTypeParentFormat = gcpurls.Template[entityTypeParent](
	"aiplatform.googleapis.com",
	"projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entityType}",
)

type entityTypeParent struct {
	Project      string
	Location     string
	Featurestore string
	EntityType   string
}

// VertexAIFeaturestoreEntityTypeFeatureIdentity is the identity of a GCP VertexAIFeaturestoreEntityTypeFeature resource.
// +k8s:deepcopy-gen=false
type VertexAIFeaturestoreEntityTypeFeatureIdentity struct {
	Project      string
	Location     string
	Featurestore string
	EntityType   string
	Feature      string
}

func (i *VertexAIFeaturestoreEntityTypeFeatureIdentity) String() string {
	return VertexAIFeaturestoreEntityTypeFeatureIdentityFormat.ToString(*i)
}

func (i *VertexAIFeaturestoreEntityTypeFeatureIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIFeaturestoreEntityTypeFeatureIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("parsing external %q: %w", ref, err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIFeaturestoreEntityTypeFeature external=%q was not known (use %s)", ref, VertexAIFeaturestoreEntityTypeFeatureIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *VertexAIFeaturestoreEntityTypeFeatureIdentity) Host() string {
	return VertexAIFeaturestoreEntityTypeFeatureIdentityFormat.Host()
}

func (i *VertexAIFeaturestoreEntityTypeFeatureIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s", i.Project, i.Location, i.Featurestore, i.EntityType)
}

func getIdentityFromVertexAIFeaturestoreEntityTypeFeatureSpec(ctx context.Context, reader client.Reader, obj *VertexAIFeaturestoreEntityTypeFeature) (*VertexAIFeaturestoreEntityTypeFeatureIdentity, error) {
	if obj.Spec.Entitytype == "" {
		return nil, fmt.Errorf("spec.entitytype is required")
	}

	parent, match, err := entityTypeParentFormat.Parse(obj.Spec.Entitytype)
	if err != nil {
		return nil, fmt.Errorf("parsing spec.entitytype %q: %w", obj.Spec.Entitytype, err)
	}
	if !match {
		return nil, fmt.Errorf("format of VertexAIFeaturestoreEntityType parent entitytype=%q was not known (use %s)", obj.Spec.Entitytype, entityTypeParentFormat.CanonicalForm())
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	return &VertexAIFeaturestoreEntityTypeFeatureIdentity{
		Project:      parent.Project,
		Location:     parent.Location,
		Featurestore: parent.Featurestore,
		EntityType:   parent.EntityType,
		Feature:      resourceID,
	}, nil
}

func (obj *VertexAIFeaturestoreEntityTypeFeature) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromVertexAIFeaturestoreEntityTypeFeatureSpec(ctx, reader, obj)
}
