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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
)

// ApplyRouterNATCreationDefaults applies default values to a RouterNat proto on creation
// to ensure REST state parity with GCP baselines and backwards compatibility with legacy controller.
func ApplyRouterNATCreationDefaults(obj *pb.RouterNat) {
	if obj == nil {
		return
	}
	if obj.IcmpIdleTimeoutSec == nil {
		obj.IcmpIdleTimeoutSec = proto.Int32(30)
	}
	if obj.TcpEstablishedIdleTimeoutSec == nil {
		obj.TcpEstablishedIdleTimeoutSec = proto.Int32(1200)
	}
	if obj.TcpTimeWaitTimeoutSec == nil {
		obj.TcpTimeWaitTimeoutSec = proto.Int32(120)
	}
	if obj.TcpTransitoryIdleTimeoutSec == nil {
		obj.TcpTransitoryIdleTimeoutSec = proto.Int32(30)
	}
	if obj.UdpIdleTimeoutSec == nil {
		obj.UdpIdleTimeoutSec = proto.Int32(30)
	}
	if obj.EnableEndpointIndependentMapping == nil {
		obj.EnableEndpointIndependentMapping = proto.Bool(true)
	}
	if obj.Type == nil || *obj.Type == "" {
		obj.Type = proto.String("PUBLIC")
	}
}
