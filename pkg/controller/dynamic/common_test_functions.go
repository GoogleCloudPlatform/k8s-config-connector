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
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	condition "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	"github.com/ghodss/yaml" //nolint:depguard
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

// GetStatus holds the required fields for computing if an object should be considered ready (fully reconciled).
type ObjectStatus struct {
	Generation         int64
	ObservedGeneration *int64
	Conditions         []condition.Condition
}

// GetObjectStatus extracts the required fields for computing if an object should be considered ready (fully reconciled).
func GetObjectStatus(object runtime.Object) (ObjectStatus, error) {
	// Simple types with the fields we care about, so that we can use the libraries
	type withConditions struct {
		ObservedGeneration *int64                `json:"observedGeneration,omitempty"`
		Conditions         []condition.Condition `json:"conditions"`
	}
	type withStatusConditions struct {
		Status withConditions `json:"status"`
	}

	// Extract information by converting to withConditions

	u, ok := object.(*unstructured.Unstructured)
	if !ok {
		m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(object)
		if err != nil {
			return ObjectStatus{}, fmt.Errorf("error from runtime.DefaultUnstructuredConverter.ToUnstructured(%T): %w", object, err)
		}
		u = &unstructured.Unstructured{Object: m}
	}

	generation := u.GetGeneration()
	var statusConditionsObj withStatusConditions
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &statusConditionsObj); err != nil {
		return ObjectStatus{}, fmt.Errorf("error converting to object with status.conditions: %w", err)
	}

	return	 ObjectStatus{
		Generation:         generation,
		ObservedGeneration: statusConditionsObj.Status.ObservedGeneration,
		Conditions:         statusConditionsObj.Status.Conditions,
	}, nil
}	
