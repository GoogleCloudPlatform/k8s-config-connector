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

package v1beta1

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
	_ identity.IdentityV2 = &SpannerInstanceIdentity{}
	_ identity.Resource   = &SpannerInstance{}
)

var SpannerInstanceIdentityFormat = gcpurls.Template[SpannerInstanceIdentity]("spanner.googleapis.com", "projects/{project}/instances/{instance}")

// +k8s:deepcopy-gen=false
type SpannerInstanceIdentity struct {
	Project  string
	Instance string
}

func (i *SpannerInstanceIdentity) String() string {
	return SpannerInstanceIdentityFormat.ToString(*i)
}

func (i *SpannerInstanceIdentity) FromExternal(ref string) error {
	parsed, match, err := SpannerInstanceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of SpannerInstance external=%q was not known (use %s): %w", ref, SpannerInstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of SpannerInstance external=%q was not known (use %s)", ref, SpannerInstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *SpannerInstanceIdentity) Host() string {
	return SpannerInstanceIdentityFormat.Host()
}

func getIdentityFromSpannerInstanceSpec(ctx context.Context, reader client.Reader, obj client.Object) (*SpannerInstanceIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &SpannerInstanceIdentity{
		Project:  projectID,
		Instance: resourceID,
	}
	return identity, nil
}

func (obj *SpannerInstance) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromSpannerInstanceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &SpannerInstanceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change SpannerInstance identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// NewSpannerInstanceIdentity is a helper to create a SpannerInstanceIdentity from a SpannerInstance object.
// Deprecated: use GetIdentity instead.
func NewSpannerInstanceIdentity(ctx context.Context, reader client.Reader, obj *SpannerInstance, u client.Object) (*SpannerInstanceIdentity, error) {
	return getIdentityFromSpannerInstanceSpec(ctx, reader, u)
}

// SpannerInstanceConfigPrefix is used by SpannerInstanceConfig reference
func (i *SpannerInstanceIdentity) SpannerInstanceConfigPrefix() string {
	return fmt.Sprintf("projects/%s/instanceConfigs/", i.Project)
}

// Parent returns the parent of the SpannerInstance.
// Deprecated: use Project field instead.
func (i *SpannerInstanceIdentity) Parent() string {
	return "projects/" + i.Project
}

// ID returns the ID of the SpannerInstance.
// Deprecated: use Instance field instead.
func (i *SpannerInstanceIdentity) ID() string {
	return i.Instance
}
