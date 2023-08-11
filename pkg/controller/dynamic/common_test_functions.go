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

package dynamic

import (
	"io/ioutil"
	"log"
	"testing"

	condition "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	"github.com/ghodss/yaml"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func UnmarshalFileToCRD(t *testing.T, fileName string) *apiextensions.CustomResourceDefinition {
	t.Helper()
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatalf("error reading file '%v': %v", fileName, err)
	}
	o := &apiextensions.CustomResourceDefinition{}
	err = yaml.Unmarshal(bytes, o)
	if err != nil {
		t.Fatalf("error unmarshalling bytes to CRD: %v", err)
	}
	return o
}

func UnmarshalToCRD(fileName string) *apiextensions.CustomResourceDefinition {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error reading file '%v': %v", fileName, err)
	}
	o := &apiextensions.CustomResourceDefinition{}
	err = yaml.Unmarshal(bytes, o)
	if err != nil {
		log.Fatalf("error unmarshalling bytes to CRD: %v", err)
	}
	return o
}

func GetConditions(t *testing.T, kccResource *unstructured.Unstructured) []condition.Condition {
	// Simple types with the fields we care about, so that we can use the libraries
	type withConditions struct {
		Conditions []condition.Condition `json:"conditions"`
	}
	type withStatusConditions struct {
		Status withConditions `json:"status"`
	}
	var obj withStatusConditions

	// Convert into the above simplifed types
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(kccResource.UnstructuredContent(), &obj); err != nil {
		t.Errorf("error converting to object with status.conditions: %v", err)
	}

	return obj.Status.Conditions
}
