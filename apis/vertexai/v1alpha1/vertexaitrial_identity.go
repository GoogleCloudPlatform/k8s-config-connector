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
	_ identity.IdentityV2 = &VertexAITrialIdentity{}
	_ identity.Resource   = &VertexAITrial{}
)

var VertexAITrialIdentityFormat = gcpurls.Template[VertexAITrialIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/studies/{study}/trials/{trial}")

// VertexAITrialIdentity is the identity of a GCP VertexAITrial resource.
// +k8s:deepcopy-gen=false
type VertexAITrialIdentity struct {
	Project  string
	Location string
	Study    string
	Trial    string
}

func (i *VertexAITrialIdentity) String() string {
	return VertexAITrialIdentityFormat.ToString(*i)
}

func (i *VertexAITrialIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAITrialIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAITrial external=%q was not known (use %s): %w", ref, VertexAITrialIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAITrial external=%q was not known (use %s)", ref, VertexAITrialIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAITrialIdentity) Host() string {
	return VertexAITrialIdentityFormat.Host()
}

func getIdentityFromVertexAITrialSpec(ctx context.Context, reader client.Reader, obj client.Object) (*VertexAITrialIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	vertexaiObj := obj.(*VertexAITrial)
	location := common.ValueOf(vertexaiObj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	studyRef := vertexaiObj.Spec.StudyRef
	if studyRef == nil {
		return nil, fmt.Errorf("spec.studyRef must be specified")
	}
	if err := studyRef.Normalize(ctx, reader, vertexaiObj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.studyRef: %w", err)
	}
	studyIdentityRaw, err := studyRef.ParseExternalToIdentity()
	if err != nil {
		return nil, fmt.Errorf("parsing studyRef: %w", err)
	}
	studyIdentity, ok := studyIdentityRaw.(*VertexAIStudyIdentity)
	if !ok {
		return nil, fmt.Errorf("expected *VertexAIStudyIdentity from studyRef")
	}
	parentStudy := studyIdentity.Study
	identity := &VertexAITrialIdentity{
		Project:  projectID,
		Location: location,
		Study:    parentStudy,
		Trial:    resourceID,
	}
	return identity, nil
}

func (obj *VertexAITrial) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAITrialSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAITrialIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAITrial identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
