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
// proto.message: google.cloud.connectors.v1.Connection
// api.group: connectors.cnrm.cloud.google.com

package connectors

import (
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/connectors/pb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(connectorsConnectionFuzzer())
}

func connectorsConnectionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Connection{},
		ConnectorsConnectionSpec_FromProto, ConnectorsConnectionSpec_ToProto,
		ConnectorsConnectionObservedState_FromProto, ConnectorsConnectionObservedState_ToProto,
	)

	f.SpecField(".labels")
	f.SpecField(".description")
	f.SpecField(".connector_version")
	f.SpecField(".config_variables")
	f.SpecField(".auth_config")
	f.SpecField(".lock_config")
	f.SpecField(".destination_configs")
	f.SpecField(".service_account")
	f.SpecField(".suspended")
	f.SpecField(".node_config")
	f.SpecField(".ssl_config")

	f.StatusField(".name")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".status")
	f.StatusField(".image_location")
	f.StatusField(".service_directory")
	f.StatusField(".envoy_image_location")

	return f
}
