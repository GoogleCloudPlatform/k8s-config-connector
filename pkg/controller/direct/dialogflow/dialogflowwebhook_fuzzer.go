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
// api.group: dialogflow.cnrm.cloud.google.com

package dialogflow

import (
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dialogflowWebhookFuzzer())
}

func dialogflowWebhookFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer[*pb.Webhook, krm.DialogflowWebhookSpec, fuzztesting.NoStatus](&pb.Webhook{},
		DialogflowWebhookSpec_FromProto,
		DialogflowWebhookSpec_ToProto,
		nil,
		nil,
	)

	f.Unimplemented_Identity(".name")

	f.Unimplemented_NotYetTriaged(".generic_web_service.secret_version_for_username_password")
	f.Unimplemented_NotYetTriaged(".generic_web_service.secret_versions_for_request_headers")
	f.Unimplemented_NotYetTriaged(".generic_web_service.service_account_auth_config")
	f.Unimplemented_NotYetTriaged(".generic_web_service.oauth_config.secret_version_for_client_secret")

	f.Unimplemented_NotYetTriaged(".service_directory.generic_web_service.secret_version_for_username_password")
	f.Unimplemented_NotYetTriaged(".service_directory.generic_web_service.secret_versions_for_request_headers")
	f.Unimplemented_NotYetTriaged(".service_directory.generic_web_service.service_account_auth_config")
	f.Unimplemented_NotYetTriaged(".service_directory.generic_web_service.oauth_config.secret_version_for_client_secret")

	f.SpecField(".display_name")
	f.SpecField(".generic_web_service")
	f.SpecField(".service_directory")
	f.SpecField(".timeout")
	f.SpecField(".disabled")

	return f
}
