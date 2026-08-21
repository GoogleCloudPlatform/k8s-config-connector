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

package resourceoverrides_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	fooKind = "Foo"
	barKind = "Bar"
	bazKind = "Baz"
)

func TestResourceOverridesHandler(t *testing.T) {
	t.Parallel()
	handler := resourceoverrides.NewResourceOverridesHandler()
	handler.Register(getFooOverrides())
	handler.Register(getBazOverrides())

	fooUnstruc := &unstructured.Unstructured{}
	fooUnstruc.SetKind(fooKind)
	barUnstruc := &unstructured.Unstructured{}
	barUnstruc.SetKind(barKind)
	bazUnstruc := &unstructured.Unstructured{}
	bazUnstruc.SetKind(bazKind)

	tests := []struct {
		name              string
		kind              string
		crd               *apiextensions.CustomResourceDefinition
		u                 *unstructured.Unstructured
		original          *k8s.Resource
		reconciled        *k8s.Resource
		hasOverrides      bool
		hasConfigValidate bool
	}{
		{
			name: "resource Foo with overrides",
			kind: fooKind,
			crd: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Names: apiextensions.CustomResourceDefinitionNames{
						Kind: fooKind,
					},
				},
			},
			u: fooUnstruc,
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
			},
			hasConfigValidate: true,
			hasOverrides:      true,
		},
		{
			name: "resource Bar with no override",
			kind: barKind,
			crd: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Names: apiextensions.CustomResourceDefinitionNames{
						Kind: barKind,
					},
				},
			},
			u: barUnstruc,
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: barKind,
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: barKind,
				},
			},
			hasOverrides:      false,
			hasConfigValidate: false,
		},
		{
			name: "resource baz has overrides with no ConfigValidate",
			kind: bazKind,
			crd: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Names: apiextensions.CustomResourceDefinitionNames{
						Kind: bazKind,
					},
				},
			},
			u: bazUnstruc,
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: bazKind,
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: bazKind,
				},
			},
			hasOverrides:      true,
			hasConfigValidate: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.hasOverrides != handler.HasOverrides(tc.kind) {
				t.Fatalf("expect HasOverrides() for kind %v to return %v", tc.kind, tc.hasOverrides)
			}
			if tc.hasConfigValidate != handler.HasConfigValidate(tc.kind) {
				t.Fatalf("expect HasConfigValidate() for kind %v to return %v", tc.kind, tc.hasConfigValidate)
			}
			if err := handler.CRDDecorate(tc.crd); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if err := handler.ConfigValidate(tc.u); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if err := handler.PreActuationTransform(tc.original); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if err := handler.PostActuationTransform(tc.original, tc.reconciled, nil, nil); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func getBazOverrides() resourceoverrides.ResourceOverrides {
	ro := resourceoverrides.ResourceOverrides{
		Kind: bazKind,
	}
	o1 := resourceoverrides.ResourceOverride{}
	o1.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		return nil
	}
	return ro
}

func getFooOverrides() resourceoverrides.ResourceOverrides {
	ro := resourceoverrides.ResourceOverrides{
		Kind: fooKind,
	}
	o1 := resourceoverrides.ResourceOverride{}
	o1.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		return nil
	}
	o1.ConfigValidate = func(r *unstructured.Unstructured) error {
		return nil
	}
	o1.PostActuationTransform = func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
		return nil
	}
	o1.PreActuationTransform = func(r *k8s.Resource) error {
		return nil
	}
	// ResourceOverride only modifies the CRD
	o2 := resourceoverrides.ResourceOverride{}
	o2.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		return nil
	}
	ro.Overrides = append(ro.Overrides, o1)
	ro.Overrides = append(ro.Overrides, o2)
	return ro
}
