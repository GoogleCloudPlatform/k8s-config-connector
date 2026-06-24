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
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeRouteIdentity{}
	_ identity.Resource   = &ComputeRoute{}
)

var ComputeRouteIdentityFormat = gcpurls.Template[ComputeRouteIdentity]("compute.googleapis.com", "projects/{project}/global/routes/{route}")

// ComputeRouteIdentity is the identity of a GCP ComputeRoute resource.
// +k8s:deepcopy-gen=false
type ComputeRouteIdentity struct {
	Project string
	Route   string
}

func (i *ComputeRouteIdentity) String() string {
	return ComputeRouteIdentityFormat.ToString(*i)
}

func (i *ComputeRouteIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeRouteIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeRoute external=%q was not known (use %s): %w", ref, ComputeRouteIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeRoute external=%q was not known (use %s)", ref, ComputeRouteIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeRouteIdentity) Host() string {
	return ComputeRouteIdentityFormat.Host()
}

func (i *ComputeRouteIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s", i.Project)
}

func ParseComputeRouteExternal(external string) (*ComputeRouteIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeRoute external value")
	}
	id := &ComputeRouteIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeRouteSpec(ctx context.Context, reader client.Reader, obj *ComputeRoute) (*ComputeRouteIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeRouteIdentity{
		Project: projectID,
		Route:   resourceID,
	}
	return identity, nil
}

func (obj *ComputeRoute) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeRouteSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &ComputeRouteIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeRoute identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
