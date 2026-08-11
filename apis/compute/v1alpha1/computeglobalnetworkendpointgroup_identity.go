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
// See the License for the_identity.go specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeGlobalNetworkEndpointGroupIdentity{}
)

var ComputeGlobalNetworkEndpointGroupIdentityFormat = gcpurls.Template[ComputeGlobalNetworkEndpointGroupIdentity](
	"compute.googleapis.com",
	"projects/{project}/global/networkEndpointGroups/{globalNetworkEndpointGroup}",
)

// ComputeGlobalNetworkEndpointGroupIdentity is the identity of a GCP ComputeGlobalNetworkEndpointGroup resource.
// +k8s:deepcopy-gen=false
type ComputeGlobalNetworkEndpointGroupIdentity struct {
	Project                    string
	GlobalNetworkEndpointGroup string
}

func (i *ComputeGlobalNetworkEndpointGroupIdentity) String() string {
	return ComputeGlobalNetworkEndpointGroupIdentityFormat.ToString(*i)
}

func (i *ComputeGlobalNetworkEndpointGroupIdentity) FromExternal(ref string) error {
	parsed, match, err := ComputeGlobalNetworkEndpointGroupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ComputeGlobalNetworkEndpointGroup external=%q was not known (use %s): %w", ref, ComputeGlobalNetworkEndpointGroupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeGlobalNetworkEndpointGroup external=%q was not known (use %s)", ref, ComputeGlobalNetworkEndpointGroupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeGlobalNetworkEndpointGroupIdentity) Host() string {
	return ComputeGlobalNetworkEndpointGroupIdentityFormat.Host()
}

func getIdentityFromComputeGlobalNetworkEndpointGroupSpec(ctx context.Context, reader client.Reader, obj client.Object) (*ComputeGlobalNetworkEndpointGroupIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &ComputeGlobalNetworkEndpointGroupIdentity{
		Project:                    projectID,
		GlobalNetworkEndpointGroup: resourceID,
	}
	return identity, nil
}
