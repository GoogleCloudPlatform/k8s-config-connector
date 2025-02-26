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

// +kcc:proto=google.cloud.notebooks.v1.ContainerImage
type ContainerImage struct {
	// Required. The path to the container image repository. For example:
	//  `gcr.io/{project_id}/{image_name}`
	// +kcc:proto:field=google.cloud.notebooks.v1.ContainerImage.repository
	Repository *string `json:"repository,omitempty"`

	// The tag of the container image. If not specified, this defaults
	//  to the latest tag.
	// +kcc:proto:field=google.cloud.notebooks.v1.ContainerImage.tag
	Tag *string `json:"tag,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.Environment
type Environment struct {

	// Display name of this environment for the UI.
	// +kcc:proto:field=google.cloud.notebooks.v1.Environment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A brief description of this environment.
	// +kcc:proto:field=google.cloud.notebooks.v1.Environment.description
	Description *string `json:"description,omitempty"`

	// Use a Compute Engine VM image to start the notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Environment.vm_image
	VmImage *VmImage `json:"vmImage,omitempty"`

	// Use a container image to start the notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Environment.container_image
	ContainerImage *ContainerImage `json:"containerImage,omitempty"`

	// Path to a Bash script that automatically runs after a notebook instance
	//  fully boots up. The path must be a URL or
	//  Cloud Storage path. Example: `"gs://path-to-file/file-name"`
	// +kcc:proto:field=google.cloud.notebooks.v1.Environment.post_startup_script
	PostStartupScript *string `json:"postStartupScript,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.VmImage
type VmImage struct {
	// Required. The name of the Google Cloud project that this VM image belongs to.
	//  Format: `{project_id}`
	// +kcc:proto:field=google.cloud.notebooks.v1.VmImage.project
	Project *string `json:"project,omitempty"`

	// Use VM image name to find the image.
	// +kcc:proto:field=google.cloud.notebooks.v1.VmImage.image_name
	ImageName *string `json:"imageName,omitempty"`

	// Use this VM image family to find the image; the newest image in this
	//  family will be used.
	// +kcc:proto:field=google.cloud.notebooks.v1.VmImage.image_family
	ImageFamily *string `json:"imageFamily,omitempty"`
}

// +kcc:proto=google.cloud.notebooks.v1.Environment
type EnvironmentObservedState struct {
	// Output only. Name of this environment.
	//  Format:
	//  `projects/{project_id}/locations/{location}/environments/{environment_id}`
	// +kcc:proto:field=google.cloud.notebooks.v1.Environment.name
	Name *string `json:"name,omitempty"`

	// Output only. The time at which this environment was created.
	// +kcc:proto:field=google.cloud.notebooks.v1.Environment.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
