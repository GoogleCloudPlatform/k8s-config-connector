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

import (
	"testing"

	"github.com/ghodss/yaml"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// addTestLabels adds labels to indicate the resource
// was created for testing purposes. These labels enable
// tools like cleanup utilities.
func addTestLabels(u *unstructured.Unstructured) {
	labels := u.GetLabels()
	if labels == nil {
		labels = map[string]string{}
	}
	AddTestLabelsToMap(labels)
	u.SetLabels(labels)
}

// AddTestLabelsToMap adds test labels to a
// map directly.
func AddTestLabelsToMap(m map[string]string) {
	m["cnrm-test"] = "true"
}

func ToUnstruct(t *testing.T, bytes []byte) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	err := yaml.Unmarshal(bytes, u)
	if err != nil {
		t.Logf("bytes in string:\n%+v\n", string(bytes))
		t.Fatalf("error unmarshalling bytes to unstruct: %v", err)
	}
	addTestLabels(u)
	return u
}

func ToUnstructWithNamespace(t *testing.T, b []byte, namespace string) *unstructured.Unstructured {
	u := ToUnstruct(t, b)
	if u.GetNamespace() == "" {
		u.SetNamespace(namespace)
	}
	return u
}
