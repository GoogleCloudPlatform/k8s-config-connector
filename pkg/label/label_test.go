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

package label_test

import (
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
)

func TestNewGcpFromK8sLabelsBasicMap(t *testing.T) {
	labels := map[string]string{
		"key1":        "val1",
		"key2":        "val2",
		"test.io/foo": "bar",
	}
	result := label.NewGcpFromK8sLabels(labels)
	expectedResult := map[string]string{
		"key1":            "val1",
		"key2":            "val2",
		"managed-by-cnrm": "true",
	}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("results mismatch: got '%v', want '%v'", result, expectedResult)
	}
}

func TestNilMap(t *testing.T) {
	result := label.NewGcpFromK8sLabels(nil)
	expectedResult := map[string]string{
		"managed-by-cnrm": "true",
	}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("results mismatch: got '%v', want '%v'", result, expectedResult)
	}
}

func TestRemoveLabelsWithKRMPrefix(t *testing.T) {
	labels := map[string]string{
		"test.io/foo": "bar",
		"key1":        "val1",
	}
	result := label.RemoveLabelsWithKRMPrefix(labels)
	expectedResult := map[string]string{
		"key1": "val1",
	}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("results mismatch: got '%v', want '%v'", result, expectedResult)
	}
}

func TestNewGCPLabelsFromK8sLabelsBasicMap(t *testing.T) {
	labels := map[string]string{
		"key1":        "val1",
		"key2":        "val2",
		"test.io/foo": "bar",
	}
	result := label.NewGCPLabelsFromK8SLabels(labels, label.GetDefaultLabels())
	expectedResult := map[string]interface{}{
		"key1":            "val1",
		"key2":            "val2",
		"managed-by-cnrm": "true",
	}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("results mismatch: got '%v', want '%v'", result, expectedResult)
	}
	result = label.NewGCPLabelsFromK8SLabels(labels)
	expectedResult = map[string]interface{}{
		"key1": "val1",
		"key2": "val2",
	}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("results mismatch: got '%v', want '%v'", result, expectedResult)
	}
	result = label.NewGCPLabelsFromK8SLabels(labels, nil)
	expectedResult = map[string]interface{}{
		"key1": "val1",
		"key2": "val2",
	}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("results mismatch: got '%v', want '%v'", result, expectedResult)
	}
	result = label.NewGCPLabelsFromK8SLabels(labels, map[string]string{"foo": "bar"})
	expectedResult = map[string]interface{}{
		"key1": "val1",
		"key2": "val2",
		"foo":  "bar",
	}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("results mismatch: got '%v', want '%v'", result, expectedResult)
	}
}
