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
package v1alpha1

// +kcc:proto=google.cloud.securesourcemanager.v1.Instance.HostConfig
type Instance_HostConfig struct {
	// Output only. HTML hostname.
	Html *string `json:"html,omitempty"`

	// Output only. API hostname. This is the hostname to use for **Host: Data
	//  Plane** endpoints.
	Api *string `json:"api,omitempty"`

	// Output only. Git HTTP hostname.
	GitHttp *string `json:"gitHttp,omitempty"`

	// Output only. Git SSH hostname.
	GitSsh *string `json:"gitSsh,omitempty"`
}
