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
	_ identity.IdentityV2 = &DiscoveryEngineSampleQuerySetIdentity{}
	_ identity.Resource   = &DiscoveryEngineSampleQuerySet{}
)

var DiscoveryEngineSampleQuerySetIdentityFormat = gcpurls.Template[DiscoveryEngineSampleQuerySetIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/sampleQuerySets/{sampleQuerySet}")

// +k8s:deepcopy-gen=false
type DiscoveryEngineSampleQuerySetIdentity struct {
	Project        string
	Location       string
	SampleQuerySet string
}

// Deprecated: use DiscoveryEngineSampleQuerySetIdentity instead
type SampleQuerySetIdentity = DiscoveryEngineSampleQuerySetIdentity

func (i *DiscoveryEngineSampleQuerySetIdentity) String() string {
	return DiscoveryEngineSampleQuerySetIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineSampleQuerySetIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineSampleQuerySetIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineSampleQuerySet external=%q was not known (use %s): %w", ref, DiscoveryEngineSampleQuerySetIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineSampleQuerySet external=%q was not known (use %s)", ref, DiscoveryEngineSampleQuerySetIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineSampleQuerySetIdentity) Host() string {
	return DiscoveryEngineSampleQuerySetIdentityFormat.Host()
}

func (i *DiscoveryEngineSampleQuerySetIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func (i *DiscoveryEngineSampleQuerySetIdentity) ID() string {
	return i.SampleQuerySet
}

func (i *DiscoveryEngineSampleQuerySetIdentity) Parent() *SampleQuerySetParent {
	return &SampleQuerySetParent{
		ProjectID: i.Project,
		Location:  i.Location,
	}
}

type SampleQuerySetParent struct {
	ProjectID string
	Location  string
}

func (p *SampleQuerySetParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func getIdentityFromDiscoveryEngineSampleQuerySetSpec(ctx context.Context, reader client.Reader, obj *DiscoveryEngineSampleQuerySet) (*DiscoveryEngineSampleQuerySetIdentity, error) {
	projectRef, err := refs.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := obj.Spec.Location

	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	return &DiscoveryEngineSampleQuerySetIdentity{
		Project:        projectID,
		Location:       location,
		SampleQuerySet: resourceID,
	}, nil
}

// NewSampleQuerySetIdentity builds a SampleQuerySetIdentity from the Config Connector SampleQuerySet object.
func NewSampleQuerySetIdentity(ctx context.Context, reader client.Reader, obj *DiscoveryEngineSampleQuerySet) (*SampleQuerySetIdentity, error) {
	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	return identity.(*DiscoveryEngineSampleQuerySetIdentity), nil
}

func ParseSampleQuerySetExternal(external string) (parent *SampleQuerySetParent, resourceID string, err error) {
	id := &DiscoveryEngineSampleQuerySetIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, "", err
	}
	parent = &SampleQuerySetParent{
		ProjectID: id.Project,
		Location:  id.Location,
	}
	return parent, id.SampleQuerySet, nil
}

func NewSampleQuerySetIdentityFromExternal(external string) (*SampleQuerySetIdentity, error) {
	id := &DiscoveryEngineSampleQuerySetIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func (obj *DiscoveryEngineSampleQuerySet) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineSampleQuerySetSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveryEngineSampleQuerySetIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineSampleQuerySet identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (c *DiscoveryEngineSampleQuerySet) ExternalIdentifier() *string {
	return c.Status.ExternalRef
}
