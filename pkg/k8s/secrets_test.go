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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetSecretVersionsFromAnnotations(t *testing.T) {
	tests := []struct {
		name                   string
		resource               *k8s.Resource
		expectedSecretVersions map[string]string
	}{
		{
			name: "secret versions exist in the annotation",
			resource: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"cnrm.cloud.google.com/observed-secret-versions": "{\"secret-1\":\"2\"}",
					},
				},
			},
			expectedSecretVersions: map[string]string{
				"secret-1": "2",
			},
		},
		{
			name: "no secret versions in the annotation",
			resource: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{},
				},
			},
			expectedSecretVersions: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			secretVersions, err := k8s.GetSecretVersionsFromAnnotations(tc.resource)
			if err != nil {
				t.Fatalf("error getting observed secret versions from annotation: %v", err)
			}
			if got, want := secretVersions, tc.expectedSecretVersions; !reflect.DeepEqual(got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}

func TestUpdateObservedSecretVersionsAnnotation(t *testing.T) {
	tests := []struct {
		name                  string
		resource              *k8s.Resource
		updatedSecretVersions map[string]string
		hasSensitiveFields    bool
	}{
		{
			name: "update secret versions exist in the annotation successfully",
			resource: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"cnrm.cloud.google.com/observed-secret-versions": "{\"secret-1\":\"2\"}",
					},
				},
			},
			updatedSecretVersions: map[string]string{
				"secret-1": "5",
			},
			hasSensitiveFields: true,
		},
		{
			name: "secret versions removed when hasSensitiveField is false",
			resource: &k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"cnrm.cloud.google.com/observed-secret-versions": "{\"secret-1\":\"2\"}",
					},
				},
			},
			updatedSecretVersions: map[string]string{
				"secret-1": "5",
			},
			hasSensitiveFields: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := k8s.UpdateOrRemoveObservedSecretVersionsAnnotation(tc.resource, tc.updatedSecretVersions, tc.hasSensitiveFields)
			if err != nil {
				t.Fatalf("error updating observed secret versions from annotation: %v", err)
			}

			retrievedSecretVersions, err := k8s.GetSecretVersionsFromAnnotations(tc.resource)
			if err != nil {
				t.Fatalf("error getting observed secret versions from annotation for verification: %v", err)
			}
			if !tc.hasSensitiveFields {
				if retrievedSecretVersions != nil {
					t.Fatalf("secret versions in the annotation should be nil, but got %v", retrievedSecretVersions)
				}
				return
			}
			if got, want := retrievedSecretVersions, tc.updatedSecretVersions; !reflect.DeepEqual(got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}
