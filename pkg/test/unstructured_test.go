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

package test

import "testing"

const exampleResource = `apiVersion: resourcemanager.cnrm.cloud.google.com/v1beta1
kind: Project
metadata:
  name: test-project
spec:
  billingAccountRef:
    external: 1234567890
  name: test-project
`

const exampleResourceWithLabels = `apiVersion: resourcemanager.cnrm.cloud.google.com/v1beta1
kind: Project
metadata:
  labels:
    foo: "bar"
  name: test-project
spec:
  billingAccountRef:
    external: 1234567890
  name: test-project
`

// ToUnstruct should add the cnrm-test label,
func TestToUnstructAddsTestLabel(t *testing.T) {
	for _, example := range []string{exampleResourceWithLabels, exampleResource} {
		b := []byte(example)
		got := ToUnstruct(t, b)
		labels := got.GetLabels()
		if labels["cnrm-test"] != "true" {
			t.Errorf("ToUnstruct(%v) label cnrm-test not set:  got '%v', want '%v'", example, labels["cnrm-test"], "true")
		}
	}
}

// ToUnstructWithNamespace should add the cnrm-test label,
// and appropriately set the namespace.
func TestToUnstructWithNamespaceAddsNamespace(t *testing.T) {
	for _, example := range []string{exampleResourceWithLabels, exampleResource} {
		b := []byte(example)
		namespace := "foo"
		got := ToUnstructWithNamespace(t, b, namespace)
		if got.GetNamespace() != namespace {
			t.Errorf("ToUnstructWithNamespace(%v) namespace not set:  got '%v', want '%v'", example, got.GetNamespace(), namespace)
		}
		labels := got.GetLabels()
		if labels["cnrm-test"] != "true" {
			t.Errorf("ToUnstructWithNamespace(%v) label cnrm-test not set:  got '%v', want '%v'", example, labels["cnrm-test"], "true")
		}
	}
}
