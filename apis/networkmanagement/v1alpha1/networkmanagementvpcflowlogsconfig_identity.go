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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NetworkManagementVpcFlowLogsConfigIdentity{}
	_ identity.Resource   = &NetworkManagementVpcFlowLogsConfig{}
)

var NetworkManagementVpcFlowLogsConfigIdentityFormat = gcpurls.Template[NetworkManagementVpcFlowLogsConfigIdentity]("networkmanagement.googleapis.com", "projects/{project}/locations/{location}/vpcFlowLogsConfigs/{vpcflowlogsconfig}")

// NetworkManagementVpcFlowLogsConfigIdentity is the identity of a GCP NetworkManagementVpcFlowLogsConfig resource.
// +k8s:deepcopy-gen=false
type NetworkManagementVpcFlowLogsConfigIdentity struct {
	Project           string
	Location          string
	VpcFlowLogsConfig string
}

func (i *NetworkManagementVpcFlowLogsConfigIdentity) String() string {
	return NetworkManagementVpcFlowLogsConfigIdentityFormat.ToString(*i)
}

func (i *NetworkManagementVpcFlowLogsConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkManagementVpcFlowLogsConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkManagementVpcFlowLogsConfig external=%q was not known (use %s): %w", ref, NetworkManagementVpcFlowLogsConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkManagementVpcFlowLogsConfig external=%q was not known (use %s)", ref, NetworkManagementVpcFlowLogsConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkManagementVpcFlowLogsConfigIdentity) Host() string {
	return NetworkManagementVpcFlowLogsConfigIdentityFormat.Host()
}

func (i *NetworkManagementVpcFlowLogsConfigIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromNetworkManagementVpcFlowLogsConfigSpec(ctx context.Context, reader client.Reader, obj *NetworkManagementVpcFlowLogsConfig) (*NetworkManagementVpcFlowLogsConfigIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &NetworkManagementVpcFlowLogsConfigIdentity{
		Project:           projectID,
		Location:          location,
		VpcFlowLogsConfig: resourceID,
	}
	return identity, nil
}

func (obj *NetworkManagementVpcFlowLogsConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkManagementVpcFlowLogsConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Since NetworkManagementVpcFlowLogsConfig Status does not contain Name or ExternalRef yet, we return the specIdentity directly without cross-check.
	return specIdentity, nil
}
