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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeRouterNATIdentity{}
	_ identity.Resource   = &ComputeRouterNAT{}
)

var ComputeRouterNATIdentityFormat = gcpurls.Template[ComputeRouterNATIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/routers/{router}/{computerouternat}")

// ComputeRouterNATIdentity is the identity of a GCP ComputeRouterNAT resource.
// +k8s:deepcopy-gen=false
type ComputeRouterNATIdentity struct {
	Project          string
	Region           string
	Router           string
	ComputeRouterNAT string
}

func (i *ComputeRouterNATIdentity) String() string {
	return ComputeRouterNATIdentityFormat.ToString(*i)
}

func (i *ComputeRouterNATIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeRouterNATIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeRouterNAT external=%q was not known (use %s): %w", ref, ComputeRouterNATIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeRouterNAT external=%q was not known (use %s)", ref, ComputeRouterNATIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeRouterNATIdentity) Host() string {
	return ComputeRouterNATIdentityFormat.Host()
}

func (i *ComputeRouterNATIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/regions/%s/routers/%s", i.Project, i.Region, i.Router)
}

func ParseComputeRouterNATExternal(external string) (*ComputeRouterNATIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeRouterNAT external value")
	}
	id := &ComputeRouterNATIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeRouterNATSpec(ctx context.Context, reader client.Reader, obj *ComputeRouterNAT) (*ComputeRouterNATIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	routerRef := &obj.Spec.RouterRef
	if err := routerRef.Normalize(ctx, reader, obj.Namespace); err != nil {
		return nil, fmt.Errorf("cannot normalize routerRef: %w", err)
	}
	routerExternal := routerRef.External
	if routerExternal == "" {
		return nil, fmt.Errorf("cannot resolve routerRef")
	}
	routerIdentity, err := ParseComputeRouterExternal(routerExternal)
	if err != nil {
		return nil, fmt.Errorf("cannot parse resolved routerRef external=%q: %w", routerExternal, err)
	}

	region := obj.Spec.Region
	if region == "" {
		region = routerIdentity.Region
	}
	if region == "" {
		return nil, fmt.Errorf("cannot resolve region: spec.region is empty")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeRouterNATIdentity{
		Project:          projectID,
		Region:           region,
		Router:           routerIdentity.Router,
		ComputeRouterNAT: resourceID,
	}
	return identity, nil
}

func (obj *ComputeRouterNAT) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeRouterNATSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
