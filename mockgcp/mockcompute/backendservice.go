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

package mockcompute

import (
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

func (s *MockService) populateBackendServiceDefaults(obj *pb.BackendService) {
	if obj.AffinityCookieTtlSec == nil {
		obj.AffinityCookieTtlSec = PtrTo(int32(0))
	}

	if obj.ConnectionDraining == nil {
		obj.ConnectionDraining = &pb.ConnectionDraining{}
	}

	if obj.ConnectionDraining.DrainingTimeoutSec == nil {
		obj.ConnectionDraining.DrainingTimeoutSec = PtrTo(int32(0))
	}

	if obj.Description == nil {
		obj.Description = PtrTo("")
	}

	if obj.LoadBalancingScheme == nil {
		obj.LoadBalancingScheme = PtrTo("EXTERNAL")
	}

	switch obj.GetProtocol() {
	case "HTTP", "HTTP2":
		if obj.Port == nil {
			obj.Port = PtrTo[int32](80)
		}
		if obj.PortName == nil {
			obj.PortName = PtrTo("http")
		}
	}

	if obj.SessionAffinity == nil {
		obj.SessionAffinity = PtrTo("NONE")
	}

	if obj.TimeoutSec == nil {
		obj.TimeoutSec = PtrTo(int32(30))
	}
}
