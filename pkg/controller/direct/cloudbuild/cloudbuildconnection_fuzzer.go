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
// proto.message: google.devtools.cloudbuild.v2.Connection
// api.group: cloudbuild.cnrm.cloud.google.com

package cloudbuild

import (
	cloudbuildpb "cloud.google.com/go/cloudbuild/apiv2/cloudbuildpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(cloudbuildConnectionFuzzer())
}

func cloudbuildConnectionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&cloudbuildpb.Connection{},
		CloudBuildConnectionSpec_v1alpha1_FromProto, CloudBuildConnectionSpec_v1alpha1_ToProto,
		CloudBuildConnectionObservedState_v1alpha1_FromProto, CloudBuildConnectionObservedState_v1alpha1_ToProto,
	)

	f.Unimplemented_Identity(".name")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".installation_state")
	f.SpecField(".disabled")
	f.StatusField(".reconciling")
	f.SpecField(".annotations")
	f.StatusField(".etag")

	f.SpecField(".github_config")
	f.SpecField(".github_config.authorizer_credential")
	f.SpecField(".github_config.authorizer_credential.oauth_token_secret_version")
	f.StatusField(".github_config.authorizer_credential.username")
	f.SpecField(".github_config.app_installation_id")

	f.SpecField(".github_enterprise_config")
	f.SpecField(".github_enterprise_config.host_uri")
	f.SpecField(".github_enterprise_config.api_key")
	f.SpecField(".github_enterprise_config.app_id")
	f.SpecField(".github_enterprise_config.app_slug")
	f.SpecField(".github_enterprise_config.private_key_secret_version")
	f.SpecField(".github_enterprise_config.webhook_secret_secret_version")
	f.SpecField(".github_enterprise_config.app_installation_id")
	f.SpecField(".github_enterprise_config.service_directory_config")
	f.SpecField(".github_enterprise_config.service_directory_config.service")
	f.SpecField(".github_enterprise_config.ssl_ca")
	f.StatusField(".github_enterprise_config.server_version")

	f.SpecField(".gitlab_config")
	f.SpecField(".gitlab_config.host_uri")
	f.SpecField(".gitlab_config.webhook_secret_secret_version")
	f.SpecField(".gitlab_config.read_authorizer_credential")
	f.SpecField(".gitlab_config.read_authorizer_credential.user_token_secret_version")
	f.StatusField(".gitlab_config.read_authorizer_credential.username")
	f.SpecField(".gitlab_config.authorizer_credential")
	f.SpecField(".gitlab_config.authorizer_credential.user_token_secret_version")
	f.StatusField(".gitlab_config.authorizer_credential.username")
	f.SpecField(".gitlab_config.service_directory_config")
	f.SpecField(".gitlab_config.service_directory_config.service")
	f.SpecField(".gitlab_config.ssl_ca")
	f.StatusField(".gitlab_config.server_version")

	f.SpecField(".bitbucket_data_center_config")
	f.SpecField(".bitbucket_data_center_config.host_uri")
	f.SpecField(".bitbucket_data_center_config.webhook_secret_secret_version")
	f.SpecField(".bitbucket_data_center_config.read_authorizer_credential")
	f.SpecField(".bitbucket_data_center_config.read_authorizer_credential.user_token_secret_version")
	f.StatusField(".bitbucket_data_center_config.read_authorizer_credential.username")
	f.SpecField(".bitbucket_data_center_config.authorizer_credential")
	f.SpecField(".bitbucket_data_center_config.authorizer_credential.user_token_secret_version")
	f.StatusField(".bitbucket_data_center_config.authorizer_credential.username")
	f.SpecField(".bitbucket_data_center_config.service_directory_config")
	f.SpecField(".bitbucket_data_center_config.service_directory_config.service")
	f.SpecField(".bitbucket_data_center_config.ssl_ca")
	f.StatusField(".bitbucket_data_center_config.server_version")

	f.SpecField(".bitbucket_cloud_config")
	f.SpecField(".bitbucket_cloud_config.workspace")
	f.SpecField(".bitbucket_cloud_config.webhook_secret_secret_version")
	f.SpecField(".bitbucket_cloud_config.read_authorizer_credential")
	f.SpecField(".bitbucket_cloud_config.read_authorizer_credential.user_token_secret_version")
	f.StatusField(".bitbucket_cloud_config.read_authorizer_credential.username")
	f.SpecField(".bitbucket_cloud_config.authorizer_credential")
	f.SpecField(".bitbucket_cloud_config.authorizer_credential.user_token_secret_version")
	f.StatusField(".bitbucket_cloud_config.authorizer_credential.username")

	return f
}
