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

package resourceoverrides

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func GetSecretManagerSecretResourceOverrides() ResourceOverrides {
	return ResourceOverrides{
		Kind: "SecretManagerSecret",
		Overrides: []ResourceOverride{{
			PostActuationTransform: func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
				// observedFields already copies createTime into observedState.
				// The legacy status mapping also emits a top-level copy, which is
				// not part of the SecretManagerSecret status schema.
				delete(reconciled.Status, "createTime")
				return nil
			},
		}},
	}
}
