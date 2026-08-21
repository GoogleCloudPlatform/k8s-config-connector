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

package k8s_test

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
)

func TestEqualIgnoreTransitionTime(t *testing.T) {
	condition := v1alpha1.Condition{}
	cType := reflect.TypeOf(&condition).Elem()
	if cType.NumField() != 5 {
		t.Fatalf("number of fields in type '%v/%v' has increased, this test needs to be updated",
			cType.PkgPath(), cType.Name())
	}
	testCases := []struct {
		Name           string
		ConditionOne   v1alpha1.Condition
		ConditionTwo   v1alpha1.Condition
		ExpectedResult bool
	}{
		{
			Name:           "Equal structs",
			ConditionOne:   v1alpha1.Condition{LastTransitionTime: "2018-11-08", Message: "Message", Reason: "Reason", Status: "True", Type: "Ready"},
			ConditionTwo:   v1alpha1.Condition{LastTransitionTime: "2018-11-08", Message: "Message", Reason: "Reason", Status: "True", Type: "Ready"},
			ExpectedResult: true,
		},
		{
			Name:           "Different times, all other values equal",
			ConditionOne:   v1alpha1.Condition{LastTransitionTime: "2018-11-08", Message: "Message", Reason: "Reason", Status: "True", Type: "Ready"},
			ConditionTwo:   v1alpha1.Condition{LastTransitionTime: "2018-11-09", Message: "Message", Reason: "Reason", Status: "True", Type: "Ready"},
			ExpectedResult: true,
		},
		{
			Name:           "Different message",
			ConditionOne:   v1alpha1.Condition{LastTransitionTime: "2018-11-08", Message: "Message", Reason: "Reason", Status: "True", Type: "Ready"},
			ConditionTwo:   v1alpha1.Condition{LastTransitionTime: "2018-11-08", Message: "Message2", Reason: "Reason", Status: "True", Type: "Ready"},
			ExpectedResult: false,
		},
		{
			Name:           "Different reason",
			ConditionOne:   v1alpha1.Condition{LastTransitionTime: "2018-11-08", Message: "Message", Reason: "Reason", Status: "True", Type: "Ready"},
			ConditionTwo:   v1alpha1.Condition{LastTransitionTime: "2018-11-08", Message: "Message", Reason: "Reason2", Status: "True", Type: "Ready"},
			ExpectedResult: false,
		},
		{
			Name:           "Different status",
			ConditionOne:   v1alpha1.Condition{LastTransitionTime: "2018-11-08", Message: "Message", Reason: "Reason", Status: "True", Type: "Ready"},
			ConditionTwo:   v1alpha1.Condition{LastTransitionTime: "2018-11-08", Message: "Message", Reason: "Reason", Status: "False", Type: "Ready"},
			ExpectedResult: false,
		},
		{
			Name:           "Different type",
			ConditionOne:   v1alpha1.Condition{LastTransitionTime: "2018-11-08", Message: "Message", Reason: "Reason", Status: "True", Type: "Ready"},
			ConditionTwo:   v1alpha1.Condition{LastTransitionTime: "2018-11-08", Message: "Message", Reason: "Reason", Status: "True", Type: "Ready2"},
			ExpectedResult: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := k8s.ConditionsEqualIgnoreTransitionTime(tc.ConditionOne, tc.ConditionTwo)
			if result != tc.ExpectedResult {
				functionName := runtime.FuncForPC(reflect.ValueOf(k8s.ConditionsEqualIgnoreTransitionTime).Pointer()).Name()
				t.Errorf("unexpected result for '%v': got '%v', want '%v'", functionName, result, tc.ExpectedResult)
			}
		})
	}
}

func TestConditionSlicesEqual(t *testing.T) {
	c1 := v1alpha1.Condition{LastTransitionTime: "2018-11-08", Message: "Message", Reason: "Reason", Status: "True"}
	c1DifferentTransitionTime := v1alpha1.Condition{LastTransitionTime: "2018-11-09", Message: "Message", Reason: "Reason", Status: "True"}
	c2 := v1alpha1.Condition{LastTransitionTime: "2018-11-08", Message: "Different Message", Reason: "Reason", Status: "True"}
	testCases := []struct {
		Name           string
		ConditionsOne  []v1alpha1.Condition
		ConditionsTwo  []v1alpha1.Condition
		ExpectedResult bool
	}{
		{
			Name:           "Nil slices",
			ConditionsOne:  nil,
			ConditionsTwo:  nil,
			ExpectedResult: true,
		},
		{
			Name:           "Empty slices",
			ConditionsOne:  []v1alpha1.Condition{},
			ConditionsTwo:  []v1alpha1.Condition{},
			ExpectedResult: true,
		},
		{
			Name:           "Equal slices, size one",
			ConditionsOne:  []v1alpha1.Condition{c1},
			ConditionsTwo:  []v1alpha1.Condition{c1},
			ExpectedResult: true,
		},
		{
			Name:           "Equal slices, with different transition times, size one",
			ConditionsOne:  []v1alpha1.Condition{c1},
			ConditionsTwo:  []v1alpha1.Condition{c1DifferentTransitionTime},
			ExpectedResult: true,
		},
		{
			Name:           "Different slices, size one",
			ConditionsOne:  []v1alpha1.Condition{c1},
			ConditionsTwo:  []v1alpha1.Condition{c2},
			ExpectedResult: false,
		},
		{
			Name:           "Equal slices, size two",
			ConditionsOne:  []v1alpha1.Condition{c1, c2},
			ConditionsTwo:  []v1alpha1.Condition{c1, c2},
			ExpectedResult: true,
		},
		{
			Name:           "Equal slices, with different transition times, size two",
			ConditionsOne:  []v1alpha1.Condition{c1, c2},
			ConditionsTwo:  []v1alpha1.Condition{c1DifferentTransitionTime, c2},
			ExpectedResult: true,
		},
		{
			Name:           "Different slices, size two",
			ConditionsOne:  []v1alpha1.Condition{c1, c1},
			ConditionsTwo:  []v1alpha1.Condition{c1, c2},
			ExpectedResult: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := k8s.ConditionSlicesEqual(tc.ConditionsOne, tc.ConditionsTwo)
			if result != tc.ExpectedResult {
				functionName := runtime.FuncForPC(reflect.ValueOf(k8s.ConditionSlicesEqual).Pointer()).Name()
				t.Errorf("unexpected result for '%v': got '%v', want '%v'", functionName, result, tc.ExpectedResult)
			}
		})
	}
}
