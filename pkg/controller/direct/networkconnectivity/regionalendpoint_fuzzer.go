// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package networkconnectivity

import (
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzNetworkConnectivityRegionalEndpoint())
}

func fuzzNetworkConnectivityRegionalEndpoint() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.RegionalEndpoint{},
		NetworkConnectivityRegionalEndpointSpec_FromProto, NetworkConnectivityRegionalEndpointSpec_ToProto,
		NetworkConnectivityRegionalEndpointObservedState_FromProto, NetworkConnectivityRegionalEndpointObservedState_ToProto,
	)
	f.UnimplementedFields.Insert(".name")

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".access_type")
	f.SpecFields.Insert(".address")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".target_google_api")
	f.SpecFields.Insert(".network")
	f.SpecFields.Insert(".subnetwork")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".psc_forwarding_rule")
	f.StatusFields.Insert(".ip_address")
	return f
}
