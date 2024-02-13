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

package lifecyclehandler

import (
	"testing"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	mgr manager.Manager
)

func TestIsOrphaned(t *testing.T) {
	tests := []struct {
		name             string
		resource         *k8s.Resource
		parentObjectName string
		parentConfigs    []corekccv1alpha1.TypeConfig
		isOrphaned       bool
	}{
		{
			name: "the parent k8s resource is gone",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"barRef": map[string]interface{}{
						"name": "bar-parent",
					},
				},
			},
			parentConfigs: []corekccv1alpha1.TypeConfig{
				{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
					Parent: true,
				},
			},
			isOrphaned: true,
		},
		{
			name: "the parent k8s resource is around",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"barRef": map[string]interface{}{
						"name": "bar-parent",
					},
				},
			},
			parentConfigs: []corekccv1alpha1.TypeConfig{
				{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
					Parent: true,
				},
			},
			parentObjectName: "bar-parent",
			isOrphaned:       false,
		},
		{
			name: "external parent resource",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"barRef": map[string]interface{}{
						"external": "bar-parent",
					},
				},
			},
			parentConfigs: []corekccv1alpha1.TypeConfig{
				{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
					Parent: true,
				},
			},
			isOrphaned: false,
		},
		{
			name: "some resource can have multiple parent resource types, but only one parent field is specified",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"barRef": map[string]interface{}{
						"name": "bar-parent",
					},
				},
			},
			parentConfigs: []corekccv1alpha1.TypeConfig{
				{
					Key: "fooRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Foo",
					},
					Parent: true,
				},
				{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
					Parent: true,
				},
			},
			parentObjectName: "bar-parent",
			isOrphaned:       false,
		},
		{
			name: "resources have no parent",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{},
			},
			parentConfigs: []corekccv1alpha1.TypeConfig{},
			isOrphaned:    false,
		},
	}
	c := mgr.GetClient()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testId := testvariable.NewUniqueId()
			tc.resource.SetNamespace(testId)
			if err := testcontroller.EnsureNamespaceExists(c, testId); err != nil {
				t.Fatal(err)
			}
			if tc.parentObjectName != "" {
				references := []*unstructured.Unstructured{
					test.NewBarUnstructured(tc.parentObjectName, testId, corev1.ConditionTrue),
				}
				test.EnsureObjectsExist(t, references, c)
			}
			res, _, err := IsOrphaned(tc.resource, tc.parentConfigs, c)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if res != tc.isOrphaned {
				t.Fatalf("got resource isOrphaned as %v, want %v", res, tc.isOrphaned)
			}
		})
	}
}

func TestMain(m *testing.M) {
	testmain.TestMainForUnitTests(m, &mgr)
}
