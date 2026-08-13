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

// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1alpha1

package compute

import (
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
)

// NetworkEndpoint_ToProto converts a KRM ComputeGlobalNetworkEndpointSpec to a computepb.NetworkEndpoint.
func NetworkEndpoint_ToProto(spec *krm.ComputeGlobalNetworkEndpointSpec) *computepb.NetworkEndpoint {
	if spec == nil {
		return nil
	}
	port := spec.Port
	ep := &computepb.NetworkEndpoint{
		Port: &port,
	}
	if spec.Fqdn != nil && *spec.Fqdn != "" {
		fqdn := *spec.Fqdn
		ep.Fqdn = &fqdn
	}
	if spec.IPAddress != nil && *spec.IPAddress != "" {
		ip := *spec.IPAddress
		ep.IpAddress = &ip
	}
	return ep
}

// NetworkEndpoint_FromProto converts a computepb.NetworkEndpoint to KRM spec fields.
// It returns the port, fqdn and ipAddress.
func NetworkEndpoint_FromProto(ep *computepb.NetworkEndpoint) (port int32, fqdn string, ipAddress string) {
	if ep == nil {
		return 0, "", ""
	}
	port = ep.GetPort()
	fqdn = ep.GetFqdn()
	ipAddress = ep.GetIpAddress()
	return port, fqdn, ipAddress
}
