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
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	ctyjson "github.com/hashicorp/go-cty/cty/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// terraformIDReplacer converts a Kubernetes object name into
// an idiomatic Terraform ID.
//
// Kubernetes object names themselves only allow alphanumerics, "-" and "_"
// (see https://github.com/kubernetes/community/blob/master/contributors/design-proposals/architecture/identifiers.md)
//
// Idiomatic terraform prefers only alphanumerics and underscores.
var terraformIDReplacer = strings.NewReplacer("-", "_", ".", "_")

// InstanceStateToMap converts state into a map[string]interface{}, using the schema as defined in r
// to coerce values to the appropriate type.
func InstanceStateToMap(r *schema.Resource, state *terraform.InstanceState) map[string]interface{} {
	ctyType := ctyTypeForResource(r)
	stateVal, err := state.AttrsAsObjectValue(ctyType)
	if err != nil {
		panic(fmt.Errorf("error parsing instance state as cty.Value: %w", err))
	}
	return CtyValToMap(stateVal, ctyType)
}

func MapToInstanceState(r *schema.Resource, m map[string]interface{}) *terraform.InstanceState {
	state := terraform.NewInstanceStateShimmedFromValue(MapToCtyVal(m, r.CoreConfigSchema().ImpliedType()), r.SchemaVersion)
	// Patch the schema version as a string; this seems to be a bug in the underlying Terraform code.
	state.Meta["schema_version"] = fmt.Sprintf("%v", r.SchemaVersion)
	return state
}

func ResourceConfigToMap(config *terraform.ResourceConfig) map[string]interface{} {
	return config.Raw
}

func MapToResourceConfig(r *schema.Resource, m map[string]interface{}) *terraform.ResourceConfig {
	schema := r.CoreConfigSchema()
	return terraform.NewResourceConfigShimmed(MapToCtyVal(m, schema.ImpliedType()), schema)
}

func CtyValToMap(val cty.Value, t cty.Type) map[string]interface{} {
	b, err := ctyjson.Marshal(val, t)
	if err != nil {
		panic(fmt.Errorf("error marshaling cty.Value as JSON: %w", err))
	}
	var ret map[string]interface{}
	if err := json.Unmarshal(b, &ret); err != nil {
		panic(fmt.Errorf("error unmarshalling JSON as map[string]interface{}: %w", err))
	}
	return ret
}

func MapToCtyVal(m map[string]interface{}, t cty.Type) cty.Value {
	b, err := json.Marshal(&m)
	if err != nil {
		panic(fmt.Errorf("error marshaling map as JSON: %w", err))
	}
	ret, err := ctyjson.Unmarshal(b, t)
	if err != nil {
		panic(fmt.Errorf("error unmarshalling JSON as cty.Value: %w", err))
	}
	return ret
}

func MapToCtyValWithSchema(m map[string]interface{}, s map[string]*schema.Schema) cty.Value {
	b, err := json.Marshal(&m)
	if err != nil {
		panic(fmt.Errorf("error marshaling map as JSON: %w", err))
	}
	ret, err := ctyjson.Unmarshal(b, schema.InternalMap(s).CoreConfigSchema().ImpliedType())
	if err != nil {
		panic(fmt.Errorf("error unmarshalling JSON as cty.Value: %w", err))
	}
	return ret
}

func ctyTypeForResource(r *schema.Resource) cty.Type {
	return r.CoreConfigSchema().ImpliedType()
}

// KRMNameToTerraformID converts a Kubernetes object name into
// an idiomatic Terraform ID.
func KRMNameToTerraformID(name string) string {
	return terraformIDReplacer.Replace(name)
}
