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

// +tool:fuzz-gen
// proto.message: google.cloud.netapp.v1.HostGroup
// api.group: netapp.cnrm.cloud.google.com

package netapp

import (
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(netAppHostGroupFuzzer())
}

func netAppHostGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.HostGroup{},
		NetAppHostGroupSpec_FromProto, NetAppHostGroupSpec_ToProto,
		NetAppHostGroupObservedState_FromProto, NetAppHostGroupObservedState_ToProto,
	)

	f.SpecField(".type")
	f.SpecField(".hosts")
	f.SpecField(".os_type")
	f.SpecField(".description")
	f.SpecField(".labels")

	f.StatusField(".state")
	f.StatusField(".create_time")

	f.IdentityField(".name")

	return f
}
