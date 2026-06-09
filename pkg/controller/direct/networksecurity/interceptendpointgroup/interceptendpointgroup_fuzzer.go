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
// proto.message: google.cloud.networksecurity.v1.InterceptEndpointGroup
// api.group: networksecurity.cnrm.cloud.google.com

package interceptendpointgroup

import (
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(interceptEndpointGroupFuzzer())
}

func interceptEndpointGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.InterceptEndpointGroup{},
		NetworkSecurityInterceptEndpointGroupSpec_FromProto, NetworkSecurityInterceptEndpointGroupSpec_ToProto,
		NetworkSecurityInterceptEndpointGroupObservedState_FromProto, NetworkSecurityInterceptEndpointGroupObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".intercept_deployment_group")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".connected_deployment_group")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".reconciling")
	f.StatusFields.Insert(".associations")

	f.UnimplementedFields.Insert(".name")

	return f
}
