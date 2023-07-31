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

package krmtotf_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"

	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestSetBlueprintAttribution(t *testing.T) {
	tests := []struct {
		name            string
		withAttribution bool
		input           *terraform.InstanceState
	}{
		{
			name:            "existing input, with attribution",
			withAttribution: true,
			input:           &terraform.InstanceState{ID: "test"},
		},
		{
			name:            "nil input, with attribution",
			withAttribution: true,
			input:           nil,
		},
		{
			name:            "existing input, no attribution",
			withAttribution: false,
			input:           &terraform.InstanceState{ID: "test"},
		},
		{
			name:            "nil input, no attribution",
			withAttribution: false,
			input:           nil,
		},
	}
	expectedAttribution := "blueprints/test"
	provider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resource := &krmtotf.Resource{}
			if tc.withAttribution {
				k8s.SetAnnotation(k8s.BlueprintAttributionAnnotation, "test", resource)
			}
			output := krmtotf.SetBlueprintAttribution(tc.input, resource, provider)
			blueprint, found := getModuleName(output, provider)
			if tc.withAttribution {
				if !found {
					t.Fatalf("blueprint attribution not set")
				}
				if blueprint != expectedAttribution {
					t.Fatalf("actual: %v, expected: %v", blueprint, expectedAttribution)
				}
			} else {
				if found {
					t.Fatalf("module name set to %v despite no attribution expected", blueprint)
				}
			}
		})
	}
}

func getModuleName(state *terraform.InstanceState, provider *tfschema.Provider) (string, bool) {
	if state == nil || state.ProviderMeta.IsNull() {
		return "", false
	}
	meta := krmtotf.CtyValToMap(state.ProviderMeta, tfschema.InternalMap(provider.ProviderMetaSchema).CoreConfigSchema().ImpliedType())
	val, ok := meta["module_name"]
	if !ok {
		return "", false
	}
	return val.(string), true
}
