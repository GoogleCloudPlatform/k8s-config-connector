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
	_ identity.IdentityV2 = &VertexAIBatchPredictionJobIdentity{}
	_ identity.Resource   = &VertexAIBatchPredictionJob{}
)

var VertexAIBatchPredictionJobIdentityFormat = gcpurls.Template[VertexAIBatchPredictionJobIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/batchPredictionJobs/{batchPredictionJob}")

// VertexAIBatchPredictionJobIdentity is the identity of a GCP VertexAIBatchPredictionJob resource.
// +k8s:deepcopy-gen=false
type VertexAIBatchPredictionJobIdentity struct {
	Project            string
	Location           string
	BatchPredictionJob string
}

func (i *VertexAIBatchPredictionJobIdentity) String() string {
	return VertexAIBatchPredictionJobIdentityFormat.ToString(*i)
}

func (i *VertexAIBatchPredictionJobIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIBatchPredictionJobIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIBatchPredictionJob external=%q was not known (use %s): %w", ref, VertexAIBatchPredictionJobIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIBatchPredictionJob external=%q was not known (use %s)", ref, VertexAIBatchPredictionJobIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAIBatchPredictionJobIdentity) Host() string {
	return VertexAIBatchPredictionJobIdentityFormat.Host()
}

func getIdentityFromVertexAIBatchPredictionJobSpec(ctx context.Context, reader client.Reader, obj client.Object) (*VertexAIBatchPredictionJobIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	vertexaiObj := obj.(*VertexAIBatchPredictionJob)
	location := common.ValueOf(vertexaiObj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	identity := &VertexAIBatchPredictionJobIdentity{
		Project:            projectID,
		Location:           location,
		BatchPredictionJob: resourceID,
	}
	return identity, nil
}

func (obj *VertexAIBatchPredictionJob) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAIBatchPredictionJobSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAIBatchPredictionJobIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAIBatchPredictionJob identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
