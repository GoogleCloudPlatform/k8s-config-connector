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

package metadata_test

import (
	"testing"

	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	testservicemetadataloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemetadataloader"

	dclunstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func Test_GVK_STV_Conversion(t *testing.T) {
	tests := []struct {
		name string
		gvk  schema.GroupVersionKind
		stv  dclunstruct.ServiceTypeVersion
	}{
		{
			name: "test1_foo",
			gvk: schema.GroupVersionKind{
				Group:   "test1.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test1Foo",
			},
			stv: dclunstruct.ServiceTypeVersion{
				Service: "test1",
				Version: "beta",
				Type:    "Foo",
			},
		},
		{
			name: "test1_bar",
			gvk: schema.GroupVersionKind{
				Group:   "test1.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test1Bar",
			},
			stv: dclunstruct.ServiceTypeVersion{
				Service: "test1",
				Version: "beta",
				Type:    "Bar",
			},
		},
		{
			name: "acronym naming",
			gvk: schema.GroupVersionKind{
				Group:   "test1.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test1IDPConfig",
			},
			stv: dclunstruct.ServiceTypeVersion{
				Service: "test1",
				Version: "beta",
				Type:    "IdpConfig",
			},
		},
		{
			name: "service name is renamed in KCC",
			gvk: schema.GroupVersionKind{
				Group:   "test3.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test3Qux",
			},
			stv: dclunstruct.ServiceTypeVersion{
				Service: "DCLTest3",
				Version: "beta",
				Type:    "Qux",
			},
		},
		{
			name: "alpha resource within Test7 service",
			gvk: schema.GroupVersionKind{
				Group:   "test7.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test7AlphaResource",
			},
			stv: dclunstruct.ServiceTypeVersion{
				Service: "test7",
				Version: "alpha",
				Type:    "AlphaResource",
			},
		},
		{
			name: "beta resource within Test7 service",
			gvk: schema.GroupVersionKind{
				Group:   "test7.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test7BetaResource",
			},
			stv: dclunstruct.ServiceTypeVersion{
				Service: "test7",
				Version: "beta",
				Type:    "BetaResource",
			},
		},
		{
			name: "ga resource within Test7 service",
			gvk: schema.GroupVersionKind{
				Group:   "test7.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test7GaResource",
			},
			stv: dclunstruct.ServiceTypeVersion{
				Service: "test7",
				Version: "ga",
				Type:    "GaResource",
			},
		},
	}
	loader := testservicemetadataloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			gvk, err := dclmetadata.ToGroupVersionKind(tc.stv, loader)
			if err != nil {
				t.Fatalf("unexpected error converting to GroupVersionKind for ServiceTypeVersion %v: %v", tc.stv, err)
			}
			if gvk != tc.gvk {
				t.Fatalf("got converted GroupVersionKind %v, but want %v", gvk, tc.gvk)
			}
			stv, err := dclmetadata.ToServiceTypeVersion(tc.gvk, loader)
			if err != nil {
				t.Fatalf("unexpected error converting to ServiceTypeVersion for GroupVersionKind %v: %v", tc.gvk, err)
			}
			if stv != tc.stv {
				t.Fatalf("got converted ServiceTypeVersion %v, but want %v", stv, tc.stv)
			}
		})
	}
}

func TestIsDCLBasedResourceKind(t *testing.T) {
	loader := testservicemetadataloader.NewForUnitTest()
	tests := []struct {
		name   string
		gvk    schema.GroupVersionKind
		result bool
	}{
		{
			name: "supported test resource Test1Foo",
			gvk: schema.GroupVersionKind{
				Group:   "test1.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test1Foo",
			},
			result: true,
		},
		{
			name: "supported test resource Test2Baz",
			gvk: schema.GroupVersionKind{
				Group:   "test2.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test2Baz",
			},
			result: true,
		},
		{
			name: "unsupported test resource in test1 service",
			gvk: schema.GroupVersionKind{
				Group:   "test1.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test1NewResource",
			},
			result: false,
		},
		{
			name: "unsupported service",
			gvk: schema.GroupVersionKind{
				Group:   "unsupportedservice.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "UnsupportedServiceFoo",
			},
			result: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			res := dclmetadata.IsDCLBasedResourceKind(tc.gvk, loader)
			if res != tc.result {
				t.Fatalf("expect to have %v, but got %v", tc.result, res)
			}
		})
	}
}
