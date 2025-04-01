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
	fuzztesting.RegisterKRMFuzzer(fuzzNetworkConnectivityInteralRange())
}

func fuzzNetworkConnectivityInteralRange() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.InternalRange{},
		NetworkConnectivityInternalRangeSpec_FromProto, NetworkConnectivityInternalRangeSpec_ToProto,
		NetworkConnectivityInternalRangeObservedState_FromProto, NetworkConnectivityInternalRangeObservedState_ToProto,
	)
	f.UnimplementedFields.Insert(".name")

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".ip_cidr_range")
	f.SpecFields.Insert(".migration")
	f.SpecFields.Insert(".migration.source")
	f.SpecFields.Insert(".migration.target")
	f.SpecFields.Insert(".network")
	f.SpecFields.Insert(".overlaps")
	f.SpecFields.Insert(".peering")
	f.SpecFields.Insert(".prefix_length")
	f.SpecFields.Insert(".target_cidr_range")
	f.SpecFields.Insert(".usage")

	f.StatusFields.Insert(".users")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	return f
}
