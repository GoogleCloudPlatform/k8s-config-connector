// Copyright 2026 Google LLC
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
	fuzztesting.RegisterKRMFuzzer(fuzzNetworkConnectivityMulticloudDataTransferConfig())
}

func fuzzNetworkConnectivityMulticloudDataTransferConfig() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.MulticloudDataTransferConfig{},
		NetworkConnectivityMulticloudDataTransferConfigSpec_FromProto, NetworkConnectivityMulticloudDataTransferConfigSpec_ToProto,
		NetworkConnectivityMulticloudDataTransferConfigObservedState_FromProto, NetworkConnectivityMulticloudDataTransferConfigObservedState_ToProto,
	)
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Etag()

	f.SpecField(".description")
	f.SpecField(".labels")

	f.StatusField(".create_time")
	f.StatusField(".destinations_active_count")
	f.StatusField(".destinations_count")
	f.StatusField(".services")
	f.StatusField(".uid")
	f.StatusField(".update_time")

	return f
}
