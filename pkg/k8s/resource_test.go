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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestResource_IsResourceIDConfigured(t *testing.T) {
	tests := []struct {
		name           string
		resource       *k8s.Resource
		expectedResult bool
		hasError       bool
	}{
		{
			name: "resource ID unspecified",
			resource: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-resource",
				},
			},
			expectedResult: false,
		},
		{
			name: "resource ID empty",
			resource: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-resource",
				},
				Spec: map[string]interface{}{
					"resourceID": "",
				},
			},
			hasError: true,
		},
		{
			name: "resource ID non-empty",
			resource: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-resource",
				},
				Spec: map[string]interface{}{
					"resourceID": "test-id",
				},
			},
			expectedResult: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result, err := tc.resource.IsResourceIDConfigured()
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil but want an error")
				}
			} else if err != nil {
				t.Fatalf("error setting metadata name as resource ID: "+
					"%v", err)
			}
			if got, want := result, tc.expectedResult; got != want {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}

func TestResource_IsSpecOrStatusUpdateRequired(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		resource       *k8s.Resource
		resourceStatus map[string]interface{}
		original       *k8s.Resource
		originalStatus map[string]interface{}
		expectedResult bool
	}{
		{
			name: "spec needs to update",
			resource: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Generation: 1,
				},
				Spec: map[string]interface{}{
					"foo": "someValue",
				},
			},
			resourceStatus: map[string]interface{}{
				"observedGeneration": float64(1),
			},
			original: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Generation: 1,
				},
				Spec: map[string]interface{}{},
			},
			originalStatus: map[string]interface{}{
				"observedGeneration": float64(1),
			},
			expectedResult: true,
		},
		{
			name: "status needs to update",
			resource: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Generation: 1,
				},
				Spec: map[string]interface{}{
					"foo": "someValue",
				},
			},
			resourceStatus: map[string]interface{}{
				"bar": "someValue",
			},
			original: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Generation: 1,
				},
				Spec: map[string]interface{}{
					"foo": "someValue",
				},
			},
			originalStatus: map[string]interface{}{},
			expectedResult: true,
		},
		{
			name: "observed generation needs to update",
			resource: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Generation: 2,
				},
				Spec: map[string]interface{}{
					"foo": "someValue",
				},
			},
			resourceStatus: map[string]interface{}{
				"bar":                "someValue",
				"observedGeneration": float64(2),
			},
			original: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Generation: 2,
				},
				Spec: map[string]interface{}{
					"foo": "someValue",
				},
			},
			originalStatus: map[string]interface{}{
				"bar":                "someValue",
				"observedGeneration": float64(1),
			},
			expectedResult: true,
		},
		{
			name: "no update is needed",
			resource: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Generation: 2,
				},
				Spec: map[string]interface{}{
					"foo": "someValue",
				},
			},
			resourceStatus: map[string]interface{}{
				"bar":                "someValue",
				"observedGeneration": float64(2),
			},
			original: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Generation: 2,
				},
				Spec: map[string]interface{}{
					"foo": "someValue",
				},
			},
			originalStatus: map[string]interface{}{
				"bar":                "someValue",
				"observedGeneration": float64(2),
			},
			expectedResult: false,
		},
		{
			name: "status is nil",
			resource: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Generation: 1,
				},
				Spec: map[string]interface{}{
					"foo": "someValue",
				},
			},
			original: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Generation: 1,
				},
				Spec: map[string]interface{}{
					"foo": "someValue",
				},
			},
			expectedResult: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.resource.SetStatus(tc.resourceStatus)
			tc.original.SetStatus(tc.originalStatus)
			actual := k8s.IsSpecOrStatusUpdateRequired(tc.resource, tc.original)
			if actual != tc.expectedResult {
				t.Fatalf("got %v, want %v", actual, tc.expectedResult)
			}
		})
	}
}
