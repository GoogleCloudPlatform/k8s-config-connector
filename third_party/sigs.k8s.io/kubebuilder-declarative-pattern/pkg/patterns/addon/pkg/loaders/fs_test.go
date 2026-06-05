/*
Copyright 2022 The Kubernetes Authors.

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

package loaders

import (
	"context"
	"reflect"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	addonv1alpha1 "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
)

func Test_NewManifestLoader(t *testing.T) {
	var testcast = []struct {
		name     string
		channel  string
		expected interface{}
	}{
		{
			name:     "file system pattern",
			channel:  "channels",
			expected: &FSRepository{},
		},
		{
			name:     "http pattern",
			channel:  "http://example.com",
			expected: &HTTPRepository{},
		},
		{
			name:     "http pattern",
			channel:  "https://example.com",
			expected: &HTTPRepository{},
		},
		{
			name:     "git pattern",
			channel:  "git//example.com",
			expected: &GitRepository{},
		},
		{
			name:     "git pattern",
			channel:  "example.com/dummy.git",
			expected: &GitRepository{},
		},
	}
	for _, test := range testcast {
		t.Run(test.name, func(t *testing.T) {
			actual, _ := NewManifestLoader(test.channel)
			actualType := reflect.TypeOf(actual.repo)
			expectedType := reflect.TypeOf(test.expected)
			if expectedType != actualType {
				t.Fatalf("expected %+v but got %+v", expectedType, actualType)
			}
		})
	}
}

func Test_ResolveManifest(t *testing.T) {
	var testcast = []struct {
		name       string
		testObject *TestResource
		expected   map[string]string
	}{
		{
			name: "success pattern",
			testObject: &TestResource{
				Spec: TestResourceSpec{
					CommonSpec: addonv1alpha1.CommonSpec{
						Version: "1.0.0",
						Channel: "stable",
					},
				},
			},
			expected: map[string]string{
				"/fake/path": "fake-manifest-for-testresource",
			},
		},
		{
			name: "Resource have no version",
			testObject: &TestResource{
				Spec: TestResourceSpec{
					CommonSpec: addonv1alpha1.CommonSpec{
						Channel: "stable",
					},
				},
			},
			expected: map[string]string{
				"/fake/path": "fake-manifest-for-testresource",
			},
		},
		{
			name: "Resource have no version and channnel",
			testObject: &TestResource{
				Spec: TestResourceSpec{
					CommonSpec: addonv1alpha1.CommonSpec{},
				},
			},
			expected: map[string]string{
				"/fake/path": "fake-manifest-for-testresource",
			},
		},
	}
	for _, test := range testcast {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			l := newTestLoader()
			obj := test.testObject
			actual, _ := l.ResolveManifest(ctx, obj)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("expected %+v but got %+v", test.expected, actual)
			}
		})
	}
}

// Below define struct for testing.

func newTestLoader() *ManifestLoader {
	return &ManifestLoader{repo: &TestRepository{}}
}

var _ Repository = &TestRepository{}

type TestRepository struct {
}

func (t *TestRepository) LoadChannel(ctx context.Context, name string) (*Channel, error) {
	r := &Channel{
		Manifests: []Version{
			{
				Package: "testresource",
				Version: "1.0.0",
			},
		},
	}
	return r, nil
}

func (t *TestRepository) LoadManifest(ctx context.Context, packageName string, id string) (map[string]string, error) {
	var r = make(map[string]string)
	r["/fake/path"] = "fake-manifest-for-testresource"
	return r, nil
}

type TestResourceSpec struct {
	addonv1alpha1.CommonSpec `json:",inline"`
	addonv1alpha1.PatchSpec  `json:",inline"`
}

type TestResourceStatus struct {
	addonv1alpha1.CommonStatus `json:",inline"`
}

type TestResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestResourceSpec   `json:"spec,omitempty"`
	Status TestResourceStatus `json:"status,omitempty"`
}

var _ addonv1alpha1.CommonObject = &TestResource{}
var _ addonv1alpha1.Patchable = &TestResource{}

func (o *TestResource) ComponentName() string {
	return "testresource"
}

func (o *TestResource) CommonSpec() addonv1alpha1.CommonSpec {
	return o.Spec.CommonSpec
}

func (o *TestResource) PatchSpec() addonv1alpha1.PatchSpec {
	return o.Spec.PatchSpec
}

func (o *TestResource) GetCommonStatus() addonv1alpha1.CommonStatus {
	return o.Status.CommonStatus
}

func (o *TestResource) SetCommonStatus(s addonv1alpha1.CommonStatus) {
	o.Status.CommonStatus = s
}

type TestResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestResource `json:"items"`
}

func (in *TestResource) DeepCopyObject() runtime.Object {
	// just stub
	return nil
}
