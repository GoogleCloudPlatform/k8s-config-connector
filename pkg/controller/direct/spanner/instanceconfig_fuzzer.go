// Copyright 2024 Google LLC
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
// proto.message: google.spanner.admin.instance.v1.InstanceConfig
// api.group: spanner.cnrm.cloud.google.com

package spanner

import (
	pb "cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(spannerInstanceConfigFuzzer())
}

func spannerInstanceConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.InstanceConfig{},
		SpannerInstanceConfigSpec_FromProto, SpannerInstanceConfigSpec_ToProto,
		SpannerInstanceConfigObservedState_FromProto, SpannerInstanceConfigObservedState_ToProto,
	)

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".replicas")
	f.SpecFields.Insert(".base_config")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".etag")
	f.SpecFields.Insert(".leader_options")

	f.StatusFields.Insert(".config_type")
	f.StatusFields.Insert(".optional_replicas")
	f.StatusFields.Insert(".reconciling")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".free_instance_availability")
	f.StatusFields.Insert(".quorum_type")
	f.StatusFields.Insert(".storage_limit_per_processing_unit")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
