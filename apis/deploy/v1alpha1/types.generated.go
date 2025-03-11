// Copyright 2025 Google LLC
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

package v1alpha1


// +kcc:proto=google.cloud.deploy.v1.CustomTargetSkaffoldActions
type CustomTargetSkaffoldActions struct {
	// Optional. The Skaffold custom action responsible for render operations. If
	//  not provided then Cloud Deploy will perform the render operations via
	//  `skaffold render`.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetSkaffoldActions.render_action
	RenderAction *string `json:"renderAction,omitempty"`

	// Required. The Skaffold custom action responsible for deploy operations.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetSkaffoldActions.deploy_action
	DeployAction *string `json:"deployAction,omitempty"`

	// Optional. List of Skaffold modules Cloud Deploy will include in the
	//  Skaffold Config as required before performing diagnose.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetSkaffoldActions.include_skaffold_modules
	IncludeSkaffoldModules []SkaffoldModules `json:"includeSkaffoldModules,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CustomTargetType
type CustomTargetType struct {
	// Optional. Name of the `CustomTargetType`. Format is
	//  `projects/{project}/locations/{location}/customTargetTypes/{customTargetType}`.
	//  The `customTargetType` component must match
	//  `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.name
	Name *string `json:"name,omitempty"`

	// Optional. Description of the `CustomTargetType`. Max length is 255
	//  characters.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.description
	Description *string `json:"description,omitempty"`

	// Optional. User annotations. These attributes can only be set and used by
	//  the user, and not by Cloud Deploy. See
	//  https://google.aip.dev/128#annotations for more details such as format and
	//  size limitations.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Labels are attributes that can be set and used by both the
	//  user and by Cloud Deploy. Labels must meet the following constraints:
	//
	//  * Keys and values can contain only lowercase letters, numeric characters,
	//  underscores, and dashes.
	//  * All characters must use UTF-8 encoding, and international characters are
	//  allowed.
	//  * Keys must start with a lowercase letter or international character.
	//  * Each resource is limited to a maximum of 64 labels.
	//
	//  Both keys and values are additionally constrained to be <= 128 bytes.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. This checksum is computed by the server based on the value of
	//  other fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.etag
	Etag *string `json:"etag,omitempty"`

	// Configures render and deploy for the `CustomTargetType` using Skaffold
	//  custom actions.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.custom_actions
	CustomActions *CustomTargetSkaffoldActions `json:"customActions,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.SkaffoldModules
type SkaffoldModules struct {
	// Optional. The Skaffold Config modules to use from the specified source.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.configs
	Configs []string `json:"configs,omitempty"`

	// Remote git repository containing the Skaffold Config modules.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.git
	Git *SkaffoldModules_SkaffoldGitSource `json:"git,omitempty"`

	// Cloud Storage bucket containing the Skaffold Config modules.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.google_cloud_storage
	GoogleCloudStorage *SkaffoldModules_SkaffoldGCSSource `json:"googleCloudStorage,omitempty"`

	// Cloud Build V2 repository containing the Skaffold Config modules.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.google_cloud_build_repo
	GoogleCloudBuildRepo *SkaffoldModules_SkaffoldGcbRepoSource `json:"googleCloudBuildRepo,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCBRepoSource
type SkaffoldModules_SkaffoldGcbRepoSource struct {
	// Required. Name of the Cloud Build V2 Repository.
	//  Format is
	//  projects/{project}/locations/{location}/connections/{connection}/repositories/{repository}.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCBRepoSource.repository
	Repository *string `json:"repository,omitempty"`

	// Optional. Relative path from the repository root to the Skaffold Config
	//  file.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCBRepoSource.path
	Path *string `json:"path,omitempty"`

	// Optional. Branch or tag to use when cloning the repository.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCBRepoSource.ref
	Ref *string `json:"ref,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCSSource
type SkaffoldModules_SkaffoldGCSSource struct {
	// Required. Cloud Storage source paths to copy recursively. For example,
	//  providing "gs://my-bucket/dir/configs/*" will result in Skaffold copying
	//  all files within the "dir/configs" directory in the bucket "my-bucket".
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCSSource.source
	Source *string `json:"source,omitempty"`

	// Optional. Relative path from the source to the Skaffold file.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCSSource.path
	Path *string `json:"path,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGitSource
type SkaffoldModules_SkaffoldGitSource struct {
	// Required. Git repository the package should be cloned from.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGitSource.repo
	Repo *string `json:"repo,omitempty"`

	// Optional. Relative path from the repository root to the Skaffold file.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGitSource.path
	Path *string `json:"path,omitempty"`

	// Optional. Git branch or tag to use when cloning the repository.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGitSource.ref
	Ref *string `json:"ref,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CustomTargetType
type CustomTargetTypeObservedState struct {
	// Output only. Resource id of the `CustomTargetType`.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.custom_target_type_id
	CustomTargetTypeID *string `json:"customTargetTypeID,omitempty"`

	// Output only. Unique identifier of the `CustomTargetType`.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Time at which the `CustomTargetType` was created.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Most recent time at which the `CustomTargetType` was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
