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
	_ identity.IdentityV2 = &VertexAIMemoryIdentity{}
	_ identity.Resource   = &VertexAIMemory{}
)

var VertexAIMemoryIdentityFormat = gcpurls.Template[VertexAIMemoryIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/reasoningEngines/{reasoningEngine}/memories/{memory}")

// VertexAIMemoryIdentity is the identity of a GCP VertexAIMemory resource.
// +k8s:deepcopy-gen=false
type VertexAIMemoryIdentity struct {
	Project         string
	Location        string
	ReasoningEngine string
	Memory          string
}

func (i *VertexAIMemoryIdentity) String() string {
	return VertexAIMemoryIdentityFormat.ToString(*i)
}

func (i *VertexAIMemoryIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIMemoryIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIMemory external=%q was not known (use %s): %w", ref, VertexAIMemoryIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIMemory external=%q was not known (use %s)", ref, VertexAIMemoryIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAIMemoryIdentity) Host() string {
	return VertexAIMemoryIdentityFormat.Host()
}

func getIdentityFromVertexAIMemorySpec(ctx context.Context, reader client.Reader, obj client.Object) (*VertexAIMemoryIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	vertexaiObj := obj.(*VertexAIMemory)
	location := common.ValueOf(vertexaiObj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	reasoningEngineRef := vertexaiObj.Spec.ReasoningEngineRef
	if reasoningEngineRef == nil {
		return nil, fmt.Errorf("spec.reasoningEngineRef must be specified")
	}
	if err := reasoningEngineRef.Normalize(ctx, reader, vertexaiObj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.reasoningEngineRef: %w", err)
	}
	reasoningEngineIdentityRaw, err := reasoningEngineRef.ParseExternalToIdentity()
	if err != nil {
		return nil, fmt.Errorf("parsing reasoningEngineRef: %w", err)
	}
	reasoningEngineIdentity, ok := reasoningEngineIdentityRaw.(*VertexAIReasoningEngineIdentity)
	if !ok {
		return nil, fmt.Errorf("expected *VertexAIReasoningEngineIdentity from reasoningEngineRef")
	}
	parentReasoningEngine := reasoningEngineIdentity.ReasoningEngine
	identity := &VertexAIMemoryIdentity{
		Project:         projectID,
		Location:        location,
		ReasoningEngine: parentReasoningEngine,
		Memory:          resourceID,
	}
	return identity, nil
}

func (obj *VertexAIMemory) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAIMemorySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAIMemoryIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAIMemory identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
