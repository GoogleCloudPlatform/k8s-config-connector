// Copyright 2024 Google LLC
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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func GetCloudIDSEndpointResourceOverrides() ResourceOverrides {
	return ResourceOverrides{
		Kind: "CloudIDSEndpoint",
		Overrides: []ResourceOverride{
			mapEndpointIpToEndpointIP(),
		},
	}
}

func mapEndpointIpToEndpointIP() ResourceOverride {
	return ResourceOverride{
		CRDDecorate: func(crd *apiextensions.CustomResourceDefinition) error {
			for _, version := range crd.Spec.Versions {
				schema := version.Schema.OpenAPIV3Schema
				status, ok := schema.Properties["status"]
				if !ok {
					return fmt.Errorf("status field not found for version %s", version.Name)
				}
				observedState, ok := status.Properties["observedState"]
				if !ok {
					return fmt.Errorf("observedState field not found for version %s", version.Name)
				}
				endpointIp, ok := observedState.Properties["endpointIp"]
				if !ok {
					return fmt.Errorf("endpointIp field not found for version %s", version.Name)
				}

				// Rename endpointIp to endpointIP
				observedState.Properties["endpointIP"] = endpointIp
				delete(observedState.Properties, "endpointIp")
			}

			return nil
		},
		PostActuationTransform: func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
			observedState, found := reconciled.Status["observedState"]
			if !found {
				// if there is no observedState there is nothing to do!
				return nil
			}
			observedStateM, ok := observedState.(map[string]interface{})
			if !ok {
				return fmt.Errorf("cannot parse observedState map")
			}
			endpointIp, found := observedStateM["endpointIp"]
			if !found {
				// field endpointIp not populated
				return nil
			}

			observedStateM["endpointIP"] = endpointIp
			delete(observedStateM, "endpointIp")

			reconciled.Status["observedState"] = observedStateM

			return nil
		},
	}
}
