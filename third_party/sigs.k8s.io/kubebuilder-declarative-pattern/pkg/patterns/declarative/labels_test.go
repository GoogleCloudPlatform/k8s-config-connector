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

package declarative

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
	addonv1alpha1 "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

func Test_AddLabels(t *testing.T) {
	inputManifest := `---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: test-app
  name: frontend
spec:
  replicas: 1
  selector:
    matchLabels:
      app: test-app
  strategy: {}
  template:
    metadata:
      labels:
        app: test-app
    spec:
      containers:
      - image: busybox
        name: busybox`
	var testcast = []struct {
		name          string
		inputManifest string
		input         map[string]string
		expected      map[string]string
	}{
		{
			name:          "success pattern",
			inputManifest: inputManifest,
			input:         map[string]string{"foo": "foo-val"},
			expected:      map[string]string{"app": "test-app", "foo": "foo-val"},
		},
		{
			name:          "success pattern which have two labels in input",
			inputManifest: inputManifest,
			input:         map[string]string{"foo": "foo-val", "bar": "bar-val"},
			expected:      map[string]string{"app": "test-app", "foo": "foo-val", "bar": "bar-val"},
		},
	}
	for _, test := range testcast {
		t.Run(test.name, func(t *testing.T) {

			dummyDeclarative := &TestResource{
				TypeMeta: metav1.TypeMeta{
					Kind:       "TestResource",
					APIVersion: "addons.example.org/v1alpha1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-instance",
				},
			}

			ctx := context.Background()

			objects, err := manifest.ParseObjects(ctx, test.inputManifest)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			fn := AddLabels(test.input)
			_ = fn(ctx, dummyDeclarative, objects)

			for _, o := range objects.Items {
				if diff := cmp.Diff(test.expected, o.UnstructuredObject().GetLabels()); diff != "" {
					t.Fatalf("result mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func Test_SourceLabel(t *testing.T) {
	var testcast = []struct {
		name        string
		prepareFunc func() *runtime.Scheme
		expected    map[string]string
	}{
		{
			name: "success pattern",
			prepareFunc: func() *runtime.Scheme {
				GroupVersion := schema.GroupVersion{Group: "addons.example.org", Version: "v1alpha1"}
				SchemeBuilder := &scheme.Builder{GroupVersion: GroupVersion}
				SchemeBuilder.Register(&TestResource{})
				addToScheme := SchemeBuilder.AddToScheme

				scheme := runtime.NewScheme()
				utilruntime.Must(clientgoscheme.AddToScheme(scheme))
				utilruntime.Must(addToScheme(scheme))
				return scheme
			},
			expected: map[string]string{"addons.example.org/testresource": "test-instance"},
		},
		{
			name: "failure pattern which is not registerd to scheme",
			prepareFunc: func() *runtime.Scheme {
				GroupVersion := schema.GroupVersion{Group: "addons.example.org", Version: "v1alpha1"}
				SchemeBuilder := &scheme.Builder{GroupVersion: GroupVersion}

				addToScheme := SchemeBuilder.AddToScheme

				scheme := runtime.NewScheme()
				utilruntime.Must(clientgoscheme.AddToScheme(scheme))
				utilruntime.Must(addToScheme(scheme))
				return scheme
			},
			expected: map[string]string{},
		},
	}
	for _, test := range testcast {
		scheme := test.prepareFunc()

		t.Run(test.name, func(t *testing.T) {

			dummyDeclarative := &TestResource{
				TypeMeta: metav1.TypeMeta{
					Kind:       "TestResource",
					APIVersion: "addons.example.org/v1alpha1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-instance",
				},
			}

			ctx := context.Background()
			fn := SourceLabel(scheme)
			actual := fn(ctx, dummyDeclarative)

			if diff := cmp.Diff(test.expected, actual); diff != "" {
				t.Fatalf("result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// Below define struct for testing.

type TestRepository struct {
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
