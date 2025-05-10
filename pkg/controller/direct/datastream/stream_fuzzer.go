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
// proto.message: google.cloud.datastream.v1.Stream
// api.group: datastream.cnrm.cloud.google.com

package datastream

import (
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(datastreamStreamFuzzer())
}

func datastreamStreamFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Stream{},
		DatastreamStreamSpec_FromProto, DatastreamStreamSpec_ToProto,
		DatastreamStreamObservedState_FromProto, DatastreamStreamObservedState_ToProto,
	)

	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".source_config")
	f.SpecFields.Insert(".destination_config")
	f.SpecFields.Insert(".state")
	f.SpecFields.Insert(".backfill_all")
	f.SpecFields.Insert(".backfill_none")
	f.SpecFields.Insert(".customer_managed_encryption_key")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".errors")
	f.StatusFields.Insert(".last_recovery_time")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
