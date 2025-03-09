// Copyright 2022 Google LLC
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

package krmtotf

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	tfversion "github.com/hashicorp/terraform-provider-google-beta/version"
)

// Inject the KCC identifier into the user agent for HTTP requests to GCP APIs issued from terraform provider.
// This is achieved by setting the following global variable provided by terraform provider.
// This function should only be called once in the program.
func SetUserAgentForTerraformProvider() {
	tfversion.ProviderVersion = gcp.KCCUserAgent()
}
