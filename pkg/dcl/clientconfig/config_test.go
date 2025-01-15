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

package clientconfig_test

import (
	"regexp"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/clientconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestSetUserAgentWithBlueprintAttribution(t *testing.T) {
	kind := "Test1Foo"
	apiVersion := "test1.cnrm.cloud.google.com/v1alpha1"
	bp := "test-blueprint"
	dclConfig := dcl.NewConfig(dcl.WithUserAgent(gcp.KCCUserAgent()))
	tests := []struct {
		name           string
		obj            metav1.Object
		dclConfig      dcl.Config
		userAgentRegex string
	}{
		{
			name: "resource with no blueprint attribution annotation",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
					},
				},
			},
			userAgentRegex: "kcc/controller-manager DeclarativeClientLib([0-9A-Za-z.:/_-]+)",
		},
		{
			name: "resource with blueprint attribution annotation set",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/blueprint": bp,
						},
					},
				},
			},
			userAgentRegex: "kcc/controller-manager blueprints/test-blueprint DeclarativeClientLib([0-9A-Za-z.:/_-]+)",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			newConfig := clientconfig.SetUserAgentWithBlueprintAttribution(dclConfig, tc.obj)
			actualUserAgent := newConfig.UserAgent()
			expr, err := regexp.Compile(tc.userAgentRegex)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !expr.MatchString(actualUserAgent) {
				t.Fatalf("the actual user agent %v doesn't match %v regex expression", actualUserAgent, tc.userAgentRegex)
			}
		})
	}
}
