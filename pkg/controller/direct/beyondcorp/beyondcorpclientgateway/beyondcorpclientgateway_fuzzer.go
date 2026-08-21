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

package beyondcorpclientgateway

import (
	pb "cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(beyondCorpClientGatewayFuzzer())
}

func beyondCorpClientGatewayFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ClientGateway{},
		BeyondCorpClientGatewaySpec_FromProto, BeyondCorpClientGatewaySpec_ToProto,
		BeyondCorpClientGatewayObservedState_FromProto, BeyondCorpClientGatewayObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".state")
	f.StatusField(".id")
	f.StatusField(".client_connector_service")

	return f
}
