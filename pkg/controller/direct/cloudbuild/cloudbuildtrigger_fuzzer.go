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
// proto.message: google.devtools.cloudbuild.v1.BuildTrigger
// api.group: cloudbuild.cnrm.cloud.google.com

package cloudbuild

import (
	pb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(cloudbuildTriggerFuzzer())
}

func cloudbuildTriggerFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.BuildTrigger{},
		CloudBuildTriggerSpec_FromProto, CloudBuildTriggerSpec_ToProto,
		CloudBuildTriggerObservedState_FromProto, CloudBuildTriggerObservedState_ToProto,
	)

	f.Unimplemented_Identity(".resource_name")
	f.StatusField(".id")
	f.SpecField(".description")
	f.Unimplemented_NotYetTriaged(".name")
	f.SpecField(".tags")
	f.SpecField(".trigger_template")
	f.SpecField(".github")
	f.SpecField(".pubsub_config")
	f.SpecField(".webhook_config")
	f.Unimplemented_NotYetTriaged(".autodetect")
	f.SpecField(".build")
	f.SpecField(".filename")
	f.SpecField(".git_file_source")
	f.StatusField(".create_time")
	f.SpecField(".disabled")
	f.SpecField(".substitutions")
	f.SpecField(".ignored_files")
	f.SpecField(".included_files")
	f.SpecField(".filter")
	f.SpecField(".source_to_build")
	f.SpecField(".service_account")
	f.SpecField(".repository_event_config")

	f.Unimplemented_NotYetTriaged(".build.name")
	f.Unimplemented_NotYetTriaged(".build.id")
	f.Unimplemented_NotYetTriaged(".build.project_id")
	f.Unimplemented_NotYetTriaged(".build.status")
	f.Unimplemented_NotYetTriaged(".build.status_detail")
	f.Unimplemented_NotYetTriaged(".build.results")
	f.Unimplemented_NotYetTriaged(".build.create_time")
	f.Unimplemented_NotYetTriaged(".build.start_time")
	f.Unimplemented_NotYetTriaged(".build.finish_time")
	f.Unimplemented_NotYetTriaged(".build.source_provenance")
	f.Unimplemented_NotYetTriaged(".build.build_trigger_id")
	f.Unimplemented_NotYetTriaged(".build.log_url")
	f.Unimplemented_NotYetTriaged(".build.timing")
	f.Unimplemented_NotYetTriaged(".build.approval")
	f.Unimplemented_NotYetTriaged(".build.warnings")
	f.Unimplemented_NotYetTriaged(".build.failure_info")
	f.Unimplemented_NotYetTriaged(".build.git_config")
	f.Unimplemented_NotYetTriaged(".build.dependencies")
	f.SpecField(".build.source")
	f.Unimplemented_NotYetTriaged(".build.source.storage_source.source_fetcher")
	f.Unimplemented_NotYetTriaged(".build.source.storage_source_manifest")
	f.Unimplemented_NotYetTriaged(".build.source.git_source")
	f.Unimplemented_NotYetTriaged(".build.source.connected_repository")
	f.SpecField(".build.available_secrets")
	f.Unimplemented_NotYetTriaged(".build.available_secrets.inline")
	f.SpecField(".build.secrets")
	f.SpecField(".build.options")
	f.Unimplemented_NotYetTriaged(".build.options.pool")
	f.Unimplemented_NotYetTriaged(".build.options.default_logs_bucket_behavior")
	f.Unimplemented_NotYetTriaged(".build.options.automap_substitutions")
	f.Unimplemented_NotYetTriaged(".build.options.enable_structured_logging")
	f.Unimplemented_NotYetTriaged(".build.steps[].timing")
	f.Unimplemented_NotYetTriaged(".build.steps[].pull_timing")
	f.Unimplemented_NotYetTriaged(".build.steps[].status")
	f.Unimplemented_NotYetTriaged(".build.steps[].exit_code")
	f.Unimplemented_NotYetTriaged(".build.steps[].automap_substitutions")
	f.Unimplemented_NotYetTriaged(".build.service_account")
	f.SpecField(".build.artifacts")
	f.Unimplemented_NotYetTriaged(".build.artifacts.objects.timing")
	f.Unimplemented_NotYetTriaged(".build.artifacts.go_modules")
	f.Unimplemented_NotYetTriaged(".build.artifacts.maven_artifacts")
	f.Unimplemented_NotYetTriaged(".build.artifacts.npm_packages")
	f.Unimplemented_NotYetTriaged(".build.artifacts.python_packages")

	f.Unimplemented_NotYetTriaged(".trigger_template.project_id")
	f.Unimplemented_NotYetTriaged(".trigger_template.substitutions")

	f.Unimplemented_NotYetTriaged(".webhook_config.secret")

	f.Unimplemented_NotYetTriaged(".repository_event_config.repository_type")

	f.Unimplemented_NotYetTriaged(".source_to_build.github_enterprise_config")
	f.Unimplemented_NotYetTriaged(".source_to_build.repository")
	f.Unimplemented_NotYetTriaged(".source_to_build.bitbucket_server_config")

	f.Unimplemented_NotYetTriaged(".git_file_source.bitbucket_server_config")
	f.Unimplemented_NotYetTriaged(".git_file_source.github_enterprise_config")

	f.Unimplemented_NotYetTriaged(".github.installation_id")
	f.Unimplemented_NotYetTriaged(".github.pull_request.invert_regex")

	return f
}
