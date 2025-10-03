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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tfversion "github.com/hashicorp/terraform-provider-google-beta/version"
)

// Inject the KCC identifier into the user agent for HTTP requests to GCP APIs issued from terraform provider.
// This is achieved by setting the following global variable provided by terraform provider.
// This function should only be called once in the program.
//
// Note that SetBlueprintAttribution will be used to add the blueprint attribution part into the user agent per resource
// if the resource has the 'cnrm.cloud.google.com/blueprint' annotation.
func SetUserAgentForTerraformProvider() {
	tfversion.ProviderVersion = gcp.KCCUserAgent()
}

// SetBlueprintAttribution sets the module name to the blueprint name on the given instance state if the resource has the 'cnrm.cloud.google.com/blueprint' annotation.
// As a result, the blueprint name will be added into the user agent for requests to the particular GCP resource.
func SetBlueprintAttribution(s *terraform.InstanceState, r *Resource, p *tfschema.Provider) *terraform.InstanceState {
	bp, found := k8s.GetAnnotation(k8s.BlueprintAttributionAnnotation, r)
	if !found {
		return s
	}
	ret := s
	if s == nil {
		ret = &terraform.InstanceState{}
	}
	meta := map[string]interface{}{
		"module_name": fmt.Sprintf("blueprints/%v", bp),
	}
	ret.ProviderMeta = MapToCtyValWithSchema(meta, p.ProviderMetaSchema)
	return ret
}
