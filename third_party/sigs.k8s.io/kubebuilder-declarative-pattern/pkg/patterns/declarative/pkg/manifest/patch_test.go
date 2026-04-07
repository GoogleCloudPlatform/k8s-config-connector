/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package manifest_test

import (
	"context"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

func Test_Patch(t *testing.T) {
	var testcases = []struct {
		name     string
		base     string
		patches  string
		expected string
	}{
		{
			name: "single applied patch",
			base: `---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: foo-operator
  namespace: kube-system`,
			patches: `---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: foo-operator
  namespace: kube-system
  labels:
     foo: bar`,
			expected: `---
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    foo: bar
  name: foo-operator
  namespace: kube-system`,
		},
		{
			name: "mismatched name, no patch",
			base: `---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: foo-operator
  namespace: kube-system`,
			patches: `---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: baz-operator
  namespace: kube-system
  labels:
     foo: bar`,
			expected: `---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: foo-operator
  namespace: kube-system`,
		},
		{
			name: "multiple patches",
			base: `---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: foo-operator
  namespace: kube-system
---
apiVersion: app.k8s.io/v1beta1
kind: Application
metadata:
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  assemblyPhase: Pending
---
apiVersion: apps/v1beta2
kind: Deployment
metadata:
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  replicas: 1`,
			patches: `---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: foo-operator
  namespace: kube-system
  labels:
     foo: bar
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: foo-operator
  namespace: kube-system
  labels:
     baz: foo
---
apiVersion: app.k8s.io/v1beta1
kind: Application
metadata:
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  descriptor:
    keywords:
    - addon
    - dashboard`,
			expected: `---
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    baz: foo
    foo: bar
  name: foo-operator
  namespace: kube-system
---
apiVersion: app.k8s.io/v1beta1
kind: Application
metadata:
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  assemblyPhase: Pending
  descriptor:
    keywords:
    - addon
    - dashboard
---
apiVersion: apps/v1beta2
kind: Deployment
metadata:
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  replicas: 1`,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			base, err := manifest.ParseObjects(ctx, tc.base)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			expected, err := manifest.ParseObjects(ctx, tc.expected)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			patchesManifest, err := manifest.ParseObjects(ctx, tc.patches)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}

			score := func(o *manifest.Object) int { return 0 }
			expected.Sort(score)
			base.Sort(score)

			var patches []*unstructured.Unstructured
			for _, p := range patchesManifest.Items {
				patches = append(patches, p.UnstructuredObject())
			}

			err = base.Patch(ctx, patches)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			actual := base

			actualLen := len(actual.Items)
			expectedLen := len(expected.Items)
			if expectedLen != actualLen {
				t.Fatalf("invalud number of objects, expected %d, found %d", expectedLen, actualLen)
			}

			for i, _ := range actual.Items {
				actualBytes, err := actual.Items[i].JSON()
				if err != nil {
					t.Fatalf("unexpected err: %v", err)
				}
				expectedBytes, err := expected.Items[i].JSON()
				if err != nil {
					t.Fatalf("unexpected err: %v", err)
				}
				actualStr := string(actualBytes)
				expectedStr := string(expectedBytes)
				if expectedStr != actualStr {
					t.Fatalf("unexpected result, expected ========\n%v\n\nactual ========\n%v\n", expectedStr, actualStr)
				}
			}
		})
	}

}
