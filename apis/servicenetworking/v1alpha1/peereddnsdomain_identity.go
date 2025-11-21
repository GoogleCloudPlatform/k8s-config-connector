// Copyright 2025 Google LLC
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

// +tool:krm-identity
// proto.message: mockgcp.cloud.servicenetworking.v1.PeeredDnsDomain
// proto.service: mockgcp.cloud.servicenetworking.v1.ServicesProjectsGlobalNetworksPeeredDnsDomainsServer
// crd.kind: ServiceNetworkingPeeredDnsDomain
// crd.version: v1alpha1

import (
	"context"
	"fmt"
	"strings"

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &PeeredDNSDomainIdentity{}

// +k8s:deepcopy-gen=false
type PeeredDNSDomainIdentity struct {
	Network *computev1beta1.NetworkIdentity
	Name    string
}

func (i *PeeredDNSDomainIdentity) String() string {
	return fmt.Sprintf("services/servicenetworking.googleapis.com/projects/%s/global/networks/%s/peeredDnsDomains/%s", i.Network.Parent().ProjectID, i.Network.ID(), i.Name)
}

func (i *PeeredDNSDomainIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/")
	if len(tokens) == 9 &&
		tokens[0] == "services" &&
		tokens[1] == "servicenetworking.googleapis.com" &&
		tokens[2] == "projects" &&
		tokens[4] == "global" &&
		tokens[5] == "networks" &&
		tokens[7] == "peeredDnsDomains" {

		network := &computev1beta1.NetworkIdentity{}
		if _, err := computev1beta1.ParseComputeNetworkExternal("projects/" + tokens[3] + "/global/networks/" + tokens[6]); err != nil {
			return fmt.Errorf("format of PeeredDNSDomain ref=%q was not known (use %q)", ref, "services/servicenetworking.googleapis.com/projects/<project>/global/networks/<networkID>/peeredDnsDomains/<name>")
		}

		name := tokens[8]

		i.Network = network
		i.Name = name

		return nil
	}

	return fmt.Errorf("format of PeeredDNSDomain ref=%q was not known (use %q)", ref, "services/servicenetworking.googleapis.com/projects/<project>/global/networks/<networkID>/peeredDnsDomains/<name>")
}

var _ identity.Resource = &ServiceNetworkingPeeredDNSDomain{}

func (obj *ServiceNetworkingPeeredDNSDomain) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get parent ID
	networkID, err := obj.GetParentIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Get resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	id := &PeeredDNSDomainIdentity{
		Network: networkID.(*computev1beta1.NetworkIdentity),
		Name:    resourceID,
	}

	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		previousID := &PeeredDNSDomainIdentity{}
		if err := previousID.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if id.String() != previousID.String() {
			return nil, fmt.Errorf("cannot update ServiceNetworkingPeeredDNSDomain identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
		}
	}

	return id, nil
}

func (obj *ServiceNetworkingPeeredDNSDomain) GetParentIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Normalize parent reference
	networkRef := *obj.Spec.NetworkRef
	if err := networkRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	// Get parent identity
	networkID := &computev1beta1.NetworkIdentity{}
	if err := networkID.FromExternal(networkRef.External); err != nil {
		return nil, err
	}
	return networkID, nil
}
