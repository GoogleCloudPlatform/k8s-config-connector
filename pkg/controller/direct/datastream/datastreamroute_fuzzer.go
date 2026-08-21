// Copyright 2024 Google LLC
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

// +tool:fuzz-gen
// proto.message: google.cloud.datastream.v1.Route
// api.group: datastream.cnrm.cloud.google.com

package datastream

import (
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(datastreamRouteFuzzer())
}

func datastreamRouteFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Route{},
		DatastreamRouteSpec_FromProto, DatastreamRouteSpec_ToProto,
		DatastreamRouteObservedState_FromProto, DatastreamRouteObservedState_ToProto,
	)

	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".destination_address")
	f.SpecFields.Insert(".destination_port")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	f.UnimplementedFields.Insert(".name")

	return f
}
