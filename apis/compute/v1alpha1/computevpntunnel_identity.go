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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &ComputeVPNTunnelIdentity{}
)

var ComputeVPNTunnelIdentityFormat = gcpurls.Template[ComputeVPNTunnelIdentity](
	"compute.googleapis.com",
	"projects/{project}/regions/{region}/vpnTunnels/{vpntunnel}",
)

// ComputeVPNTunnelIdentity is the identity of a GCP ComputeVPNTunnel resource.
// +k8s:deepcopy-gen=false
type ComputeVPNTunnelIdentity struct {
	Project   string
	Region    string
	VPNTunnel string
}

func (i *ComputeVPNTunnelIdentity) String() string {
	return ComputeVPNTunnelIdentityFormat.ToString(*i)
}

func (i *ComputeVPNTunnelIdentity) FromExternal(ref string) error {
	parsed, match, err := ComputeVPNTunnelIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ComputeVPNTunnel external=%q was not known (use %s): %w", ref, ComputeVPNTunnelIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeVPNTunnel external=%q was not known (use %s)", ref, ComputeVPNTunnelIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeVPNTunnelIdentity) Host() string {
	return ComputeVPNTunnelIdentityFormat.Host()
}
