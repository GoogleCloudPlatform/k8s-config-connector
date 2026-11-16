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
// proto.message: google.cloud.dialogflow.cx.v3.Webhook
// api.group: dialogflowcx.cnrm.cloud.google.com

package dialogflowcx

import (
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflowcx/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(dialogflowCXWebhookFuzzer())
}

func dialogflowCXWebhookFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer[*pb.Webhook, krm.DialogflowCXWebhookSpec](&pb.Webhook{},
		DialogflowCXWebhookSpec_FromProto,
		DialogflowCXWebhookSpec_ToProto,
	)

	// Identity field
	f.Unimplemented_Identity(".name")

	/*
		Spec field comparison:
		KRM Spec field          Proto Field
		------------------      ------------------
		disabled                .disabled
		displayName             .display_name
		genericWebService       .generic_web_service
		serviceDirectory        .service_directory
		timeout                 .timeout
	*/
	f.SpecField(".disabled")
	f.SpecField(".display_name")
	f.SpecField(".generic_web_service.allowed_ca_certs")
	f.SpecField(".generic_web_service.request_headers")
	f.SpecField(".generic_web_service.uri")
	f.SpecField(".service_directory.generic_web_service.allowed_ca_certs")
	f.SpecField(".service_directory.generic_web_service.request_headers")
	f.SpecField(".service_directory.generic_web_service.uri")
	f.SpecField(".service_directory.service")
	f.SpecField(".timeout")

	// Fields in pb.Webhook.GenericWebService not mapped in KRM GenericWebService
	f.Unimplemented_NotYetTriaged(".generic_web_service.username")
	f.Unimplemented_NotYetTriaged(".generic_web_service.password")
	f.Unimplemented_NotYetTriaged(".generic_web_service.secret_version_for_username_password")
	f.Unimplemented_NotYetTriaged(".generic_web_service.secret_versions_for_request_headers")
	f.Unimplemented_NotYetTriaged(".generic_web_service.oauth_config")
	f.Unimplemented_NotYetTriaged(".generic_web_service.service_account_auth_config")
	f.Unimplemented_NotYetTriaged(".generic_web_service.service_agent_auth")
	f.Unimplemented_NotYetTriaged(".generic_web_service.webhook_type")
	f.Unimplemented_NotYetTriaged(".generic_web_service.http_method")
	f.Unimplemented_NotYetTriaged(".generic_web_service.request_body")
	f.Unimplemented_NotYetTriaged(".generic_web_service.parameter_mapping")

	f.Unimplemented_NotYetTriaged(".service_directory.generic_web_service.username")
	f.Unimplemented_NotYetTriaged(".service_directory.generic_web_service.password")
	f.Unimplemented_NotYetTriaged(".service_directory.generic_web_service.secret_version_for_username_password")
	f.Unimplemented_NotYetTriaged(".service_directory.generic_web_service.secret_versions_for_request_headers")
	f.Unimplemented_NotYetTriaged(".service_directory.generic_web_service.oauth_config")
	f.Unimplemented_NotYetTriaged(".service_directory.generic_web_service.service_account_auth_config")
	f.Unimplemented_NotYetTriaged(".service_directory.generic_web_service.service_agent_auth")
	f.Unimplemented_NotYetTriaged(".service_directory.generic_web_service.webhook_type")
	f.Unimplemented_NotYetTriaged(".service_directory.generic_web_service.http_method")
	f.Unimplemented_NotYetTriaged(".service_directory.generic_web_service.request_body")
	f.Unimplemented_NotYetTriaged(".service_directory.generic_web_service.parameter_mapping")

	return f
}
