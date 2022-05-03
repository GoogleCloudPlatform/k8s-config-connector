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
	v1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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
	status := kccResource.Object["status"].(map[string]interface{})
	conditionsList := status["conditions"].([]interface{})
	if len(conditionsList) < 1 {
		t.Error("error getting instance conditions")
	}
	readyCondition := conditionsList[0].(map[string]interface{})
	// Temp fix: IAMPolicy drops reason, Bigtable drops message + reason
	message, _ := readyCondition["message"].(string)
	reason, _ := readyCondition["reason"].(string)
	return []condition.Condition{
		{
			LastTransitionTime: readyCondition["lastTransitionTime"].(string),
			Message:            message,
			Reason:             reason,
			Status:             v1.ConditionStatus(readyCondition["status"].(string)),
			Type:               readyCondition["type"].(string),
		},
	}
}
