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
	_ identity.IdentityV2 = &VertexAIRagDataSchemaIdentity{}
	_ identity.Resource   = &VertexAIRagDataSchema{}
)

var VertexAIRagDataSchemaIdentityFormat = gcpurls.Template[VertexAIRagDataSchemaIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/ragCorpora/{ragCorpus}/ragDataSchemas/{ragDataSchema}")

// VertexAIRagDataSchemaIdentity is the identity of a GCP VertexAIRagDataSchema resource.
// +k8s:deepcopy-gen=false
type VertexAIRagDataSchemaIdentity struct {
	Project       string
	Location      string
	RagCorpus     string
	RagDataSchema string
}

func (i *VertexAIRagDataSchemaIdentity) String() string {
	return VertexAIRagDataSchemaIdentityFormat.ToString(*i)
}

func (i *VertexAIRagDataSchemaIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIRagDataSchemaIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIRagDataSchema external=%q was not known (use %s): %w", ref, VertexAIRagDataSchemaIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIRagDataSchema external=%q was not known (use %s)", ref, VertexAIRagDataSchemaIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAIRagDataSchemaIdentity) Host() string {
	return VertexAIRagDataSchemaIdentityFormat.Host()
}

func getIdentityFromVertexAIRagDataSchemaSpec(ctx context.Context, reader client.Reader, obj client.Object) (*VertexAIRagDataSchemaIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	vertexaiObj := obj.(*VertexAIRagDataSchema)
	location := common.ValueOf(vertexaiObj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	ragCorpusRef := vertexaiObj.Spec.RagCorpusRef
	if ragCorpusRef == nil {
		return nil, fmt.Errorf("spec.ragCorpusRef must be specified")
	}
	if err := ragCorpusRef.Normalize(ctx, reader, vertexaiObj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.ragCorpusRef: %w", err)
	}
	ragCorpusIdentityRaw, err := ragCorpusRef.ParseExternalToIdentity()
	if err != nil {
		return nil, fmt.Errorf("parsing ragCorpusRef: %w", err)
	}
	ragCorpusIdentity, ok := ragCorpusIdentityRaw.(*VertexAIRagCorpusIdentity)
	if !ok {
		return nil, fmt.Errorf("expected *VertexAIRagCorpusIdentity from ragCorpusRef")
	}
	parentRagCorpus := ragCorpusIdentity.RagCorpus
	identity := &VertexAIRagDataSchemaIdentity{
		Project:       projectID,
		Location:      location,
		RagCorpus:     parentRagCorpus,
		RagDataSchema: resourceID,
	}
	return identity, nil
}

func (obj *VertexAIRagDataSchema) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAIRagDataSchemaSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAIRagDataSchemaIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAIRagDataSchema identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
