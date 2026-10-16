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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeRegionNetworkEndpointGroupIdentity{}
)

var ComputeRegionNetworkEndpointGroupIdentityFormat = gcpurls.Template[ComputeRegionNetworkEndpointGroupIdentity](
	"compute.googleapis.com",
	"projects/{Project}/regions/{Region}/networkEndpointGroups/{ComputeRegionNetworkEndpointGroup}",
)

// ComputeRegionNetworkEndpointGroupIdentity is the identity of a GCP ComputeRegionNetworkEndpointGroup resource.
// +k8s:deepcopy-gen=false
type ComputeRegionNetworkEndpointGroupIdentity struct {
	Project                           string
	Region                            string
	ComputeRegionNetworkEndpointGroup string
}

func (i *ComputeRegionNetworkEndpointGroupIdentity) String() string {
	return ComputeRegionNetworkEndpointGroupIdentityFormat.ToString(*i)
}

func (i *ComputeRegionNetworkEndpointGroupIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeRegionNetworkEndpointGroupIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeRegionNetworkEndpointGroup external=%q was not known (use %s): %w", ref, ComputeRegionNetworkEndpointGroupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeRegionNetworkEndpointGroup external=%q was not known (use %s)", ref, ComputeRegionNetworkEndpointGroupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeRegionNetworkEndpointGroupIdentity) Host() string {
	return ComputeRegionNetworkEndpointGroupIdentityFormat.Host()
}

func getIdentityFromComputeRegionNetworkEndpointGroupSpec(ctx context.Context, reader client.Reader, obj client.Object) (*ComputeRegionNetworkEndpointGroupIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return nil, fmt.Errorf("expected *unstructured.Unstructured, got %T", obj)
	}

	region, _, err := unstructured.NestedString(u.Object, "spec", "region")
	if err != nil {
		return nil, fmt.Errorf("cannot resolve region: %w", err)
	}
	if region == "" {
		region, _, _ = unstructured.NestedString(u.Object, "spec", "location")
	}
	if region == "" {
		return nil, fmt.Errorf("region is required but not found in spec")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &ComputeRegionNetworkEndpointGroupIdentity{
		Project:                           projectID,
		Region:                            region,
		ComputeRegionNetworkEndpointGroup: resourceID,
	}
	return identity, nil
}
