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
	_ identity.IdentityV2 = &VertexAIFeatureMonitorIdentity{}
	_ identity.Resource   = &VertexAIFeatureMonitor{}
)

var VertexAIFeatureMonitorIdentityFormat = gcpurls.Template[VertexAIFeatureMonitorIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/featureGroups/{featureGroup}/featureMonitors/{featureMonitor}")

// VertexAIFeatureMonitorIdentity is the identity of a GCP VertexAIFeatureMonitor resource.
// +k8s:deepcopy-gen=false
type VertexAIFeatureMonitorIdentity struct {
	Project        string
	Location       string
	FeatureGroup   string
	FeatureMonitor string
}

func (i *VertexAIFeatureMonitorIdentity) String() string {
	return VertexAIFeatureMonitorIdentityFormat.ToString(*i)
}

func (i *VertexAIFeatureMonitorIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIFeatureMonitorIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIFeatureMonitor external=%q was not known (use %s): %w", ref, VertexAIFeatureMonitorIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIFeatureMonitor external=%q was not known (use %s)", ref, VertexAIFeatureMonitorIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAIFeatureMonitorIdentity) Host() string {
	return VertexAIFeatureMonitorIdentityFormat.Host()
}

func getIdentityFromVertexAIFeatureMonitorSpec(ctx context.Context, reader client.Reader, obj client.Object) (*VertexAIFeatureMonitorIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	vertexaiObj := obj.(*VertexAIFeatureMonitor)
	location := common.ValueOf(vertexaiObj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	featureGroupRef := vertexaiObj.Spec.FeatureGroupRef
	if featureGroupRef == nil {
		return nil, fmt.Errorf("spec.featureGroupRef must be specified")
	}
	if err := featureGroupRef.Normalize(ctx, reader, vertexaiObj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.featureGroupRef: %w", err)
	}
	featureGroupIdentityRaw, err := featureGroupRef.ParseExternalToIdentity()
	if err != nil {
		return nil, fmt.Errorf("parsing featureGroupRef: %w", err)
	}
	featureGroupIdentity, ok := featureGroupIdentityRaw.(*VertexAIFeatureGroupIdentity)
	if !ok {
		return nil, fmt.Errorf("expected *VertexAIFeatureGroupIdentity from featureGroupRef")
	}
	parentFeatureGroup := featureGroupIdentity.FeatureGroup
	identity := &VertexAIFeatureMonitorIdentity{
		Project:        projectID,
		Location:       location,
		FeatureGroup:   parentFeatureGroup,
		FeatureMonitor: resourceID,
	}
	return identity, nil
}

func (obj *VertexAIFeatureMonitor) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAIFeatureMonitorSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAIFeatureMonitorIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAIFeatureMonitor identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
