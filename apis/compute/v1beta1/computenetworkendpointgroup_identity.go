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
	_ identity.IdentityV2 = &ComputeNetworkEndpointGroupIdentity{}
)

var ComputeNetworkEndpointGroupIdentityFormat = gcpurls.Template[ComputeNetworkEndpointGroupIdentity](
	"compute.googleapis.com",
	"projects/{Project}/zones/{Zone}/networkEndpointGroups/{ComputeNetworkEndpointGroup}",
)

// ComputeNetworkEndpointGroupIdentity is the identity of a GCP ComputeNetworkEndpointGroup resource.
// +k8s:deepcopy-gen=false
type ComputeNetworkEndpointGroupIdentity struct {
	Project                     string
	Zone                        string
	ComputeNetworkEndpointGroup string
}

func (i *ComputeNetworkEndpointGroupIdentity) String() string {
	return ComputeNetworkEndpointGroupIdentityFormat.ToString(*i)
}

func (i *ComputeNetworkEndpointGroupIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeNetworkEndpointGroupIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeNetworkEndpointGroup external=%q was not known (use %s): %w", ref, ComputeNetworkEndpointGroupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeNetworkEndpointGroup external=%q was not known (use %s)", ref, ComputeNetworkEndpointGroupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeNetworkEndpointGroupIdentity) Host() string {
	return ComputeNetworkEndpointGroupIdentityFormat.Host()
}

func getIdentityFromComputeNetworkEndpointGroupSpec(ctx context.Context, reader client.Reader, obj client.Object) (*ComputeNetworkEndpointGroupIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return nil, fmt.Errorf("expected *unstructured.Unstructured, got %T", obj)
	}

	location, _, err := unstructured.NestedString(u.Object, "spec", "location")
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}
	if location == "" {
		return nil, fmt.Errorf("location is required but not found in spec")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &ComputeNetworkEndpointGroupIdentity{
		Project:                     projectID,
		Zone:                        location,
		ComputeNetworkEndpointGroup: resourceID,
	}
	return identity, nil
}
