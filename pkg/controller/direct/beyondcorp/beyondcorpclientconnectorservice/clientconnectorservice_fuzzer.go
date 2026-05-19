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
// proto.message: google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService
// api.group: beyondcorp.cnrm.cloud.google.com

package beyondcorpclientconnectorservice

import (
	pb "cloud.google.com/go/beyondcorp/clientconnectorservices/apiv1/clientconnectorservicespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(beyondCorpClientConnectorServiceFuzzer())
}

func beyondCorpClientConnectorServiceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ClientConnectorService{},
		BeyondCorpClientConnectorServiceSpec_FromProto, BeyondCorpClientConnectorServiceSpec_ToProto,
		BeyondCorpClientConnectorServiceObservedState_FromProto, BeyondCorpClientConnectorServiceObservedState_ToProto,
	)

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".ingress")
	f.SpecFields.Insert(".egress")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".state")

	f.IdentityField(".name")

	return f
}
