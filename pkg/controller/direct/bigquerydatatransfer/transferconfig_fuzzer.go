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
// proto.message: google.cloud.bigquery.datatransfer.v1.TransferConfig
// api.group: bigquerydatatransfer.cnrm.cloud.google.com

package bigquerydatatransfer

import (
	pb "cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(bigQueryDataTransferConfigFuzzer())
}

func bigQueryDataTransferConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TransferConfig{},
		BigQueryDataTransferConfigSpec_FromProto, BigQueryDataTransferConfigSpec_ToProto,
		BigQueryDataTransferConfigObservedState_FromProto, BigQueryDataTransferConfigObservedState_ToProto,
	)

	f.SpecFields.Insert(".destination_dataset_id")
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".data_source_id")
	f.SpecFields.Insert(".params")
	f.SpecFields.Insert(".schedule")
	f.SpecFields.Insert(".schedule_options")
	f.SpecFields.Insert(".data_refresh_window_days")
	f.SpecFields.Insert(".disabled")
	f.SpecFields.Insert(".notification_pubsub_topic")
	f.SpecFields.Insert(".email_preferences")
	f.SpecFields.Insert(".encryption_configuration")
	f.SpecFields.Insert(".schedule_options_v2")

	f.StatusFields.Insert(".name")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".next_run_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".dataset_region")
	f.StatusFields.Insert(".owner_info")
	f.StatusFields.Insert(".user_id")
	f.StatusFields.Insert(".error")

	f.UnimplementedFields.Insert(".error.details")

	return f
}
