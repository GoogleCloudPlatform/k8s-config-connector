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
// proto.message: google.cloud.eventarc.v1.Pipeline
// api.group: eventarc.cnrm.cloud.google.com

package eventarc

import (
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(eventarcPipelineFuzzer())
}

func eventarcPipelineFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Pipeline{},
		EventarcPipelineSpec_FromProto, EventarcPipelineSpec_ToProto,
		EventarcPipelineObservedState_FromProto, EventarcPipelineObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".destinations")
	f.SpecField(".destinations.network_config")
	f.SpecField(".destinations.network_config.network_attachment")
	f.SpecField(".destinations.http_endpoint")
	f.SpecField(".destinations.http_endpoint.uri")
	f.SpecField(".destinations.http_endpoint.message_binding_template")
	f.SpecField(".destinations.workflow")
	f.SpecField(".destinations.message_bus")
	f.SpecField(".destinations.topic")
	f.SpecField(".destinations.authentication_config")
	f.SpecField(".destinations.authentication_config.google_oidc")
	f.SpecField(".destinations.authentication_config.google_oidc.service_account")
	f.SpecField(".destinations.authentication_config.google_oidc.audience")
	f.SpecField(".destinations.authentication_config.oauth_token")
	f.SpecField(".destinations.authentication_config.oauth_token.service_account")
	f.SpecField(".destinations.authentication_config.oauth_token.scope")
	f.SpecField(".destinations.output_payload_format")
	f.SpecField(".destinations.output_payload_format.protobuf")
	f.SpecField(".destinations.output_payload_format.protobuf.schema_definition")
	f.SpecField(".destinations.output_payload_format.avro")
	f.SpecField(".destinations.output_payload_format.avro.schema_definition")
	f.SpecField(".destinations.output_payload_format.json")
	f.SpecField(".mediations")
	f.SpecField(".mediations.transformation")
	f.SpecField(".mediations.transformation.transformation_template")
	f.SpecField(".crypto_key_name")
	f.SpecField(".input_payload_format")
	f.SpecField(".input_payload_format.protobuf")
	f.SpecField(".input_payload_format.protobuf.schema_definition")
	f.SpecField(".input_payload_format.avro")
	f.SpecField(".input_payload_format.avro.schema_definition")
	f.SpecField(".input_payload_format.json")
	f.SpecField(".logging_config")
	f.SpecField(".logging_config.log_severity")
	f.SpecField(".retry_policy")
	f.SpecField(".retry_policy.max_attempts")
	f.SpecField(".retry_policy.min_retry_delay")
	f.SpecField(".retry_policy.max_retry_delay")

	f.StatusField(".uid")
	f.StatusField(".etag")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_LabelsAnnotations(".annotations")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")

	return f
}
