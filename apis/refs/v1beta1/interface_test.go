// Copyright 2024 Google LLC
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

package v1beta1

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var (
	testGVK = schema.GroupVersionKind{
		Group:   "test.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "TestResource",
	}
)

type mockRef struct {
	Name      string
	Namespace string
	External  string
}

func (m *mockRef) GetGVK() schema.GroupVersionKind {
	return testGVK
}

func (m *mockRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      m.Name,
		Namespace: "default",
	}
}

func (m *mockRef) GetExternal() string {
	return m.External
}

func (m *mockRef) SetExternal(ref string) {
	m.External = ref
}

func (m *mockRef) ValidateExternal(ref string) error {
	tokens := strings.Split(ref, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "resources" {
		return fmt.Errorf("format of %s %s is not known", m.GetGVK().Kind, ref)
	}
	return nil
}

func (m *mockRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return Normalize(ctx, reader, m, defaultNamespace)
}

func (m *mockRef) NormalizeOnTemplate(ctx context.Context, reader client.Reader, defaultNamespace string, tpl string) error {
	return NormalizeOnTemplate(ctx, reader, m, defaultNamespace, tpl)
}

func TestNormalize(t *testing.T) {
	testObj := &unstructured.Unstructured{}
	testObj.SetGroupVersionKind(testGVK)
	testObj.SetName("test")
	testObj.SetNamespace("default")
	unstructured.SetNestedField(testObj.Object, "projects/projectId/locations/us/resources/test", "status", "externalRef")

	testCases := []struct {
		name             string
		ref              *mockRef
		objs             []client.Object
		defaultNamespace string
		wantExternal     string
		wantErr          string
	}{
		{
			name: "external is set",
			ref: &mockRef{
				External: "projects/projectId/locations/us/resources/test",
			},
			wantExternal: "projects/projectId/locations/us/resources/test",
		},
		{
			name: "external is set but has invalid format",
			ref: &mockRef{
				External: "resources/test",
			},
			wantErr: "format of TestResource resources/test is not known",
		},
		{
			name: "successful normalization",
			ref: &mockRef{
				Name: "test",
			},
			objs:             []client.Object{testObj},
			defaultNamespace: "default",
			wantExternal:     "projects/projectId/locations/us/resources/test",
		},
		{
			name: "referenced object not found",
			ref: &mockRef{
				Name: "non-existent",
			},
			objs:             []client.Object{},
			defaultNamespace: "default",
			wantErr:          "reference TestResource default/non-existent is not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			fakeClient := fake.NewClientBuilder().WithObjects(tc.objs...).Build()

			err := tc.ref.Normalize(ctx, fakeClient, tc.defaultNamespace)

			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("got nil error, want %q", tc.wantErr)
				}
				if !strings.Contains(err.Error(), tc.wantErr) {
					t.Errorf("got error %q, want it to contain %q", err.Error(), tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			if tc.ref.External != tc.wantExternal {
				t.Errorf("got external %q, want %q", tc.ref.External, tc.wantExternal)
			}
		})
	}
}

func TestNormalizeOnTemplate(t *testing.T) {
	testObj := &unstructured.Unstructured{}
	testObj.SetGroupVersionKind(testGVK)
	testObj.SetName("test")
	testObj.SetNamespace("default")
	unstructured.SetNestedField(testObj.Object, "projects/projectId/locations/us/resources/test", "status", "externalRef")
	unstructured.SetNestedField(testObj.Object, "https://link/to/resource", "status", "selfLink")
	unstructured.SetNestedField(testObj.Object, "https://link", "status", "selfLinkWithoutId")
	unstructured.SetNestedField(testObj.Object, "10.0.0.1", "status", "observedState", "address")

	testObjNoExternalRef := &unstructured.Unstructured{}
	testObjNoExternalRef.SetGroupVersionKind(testGVK)
	testObjNoExternalRef.SetName("test-no-external-ref")
	testObjNoExternalRef.SetNamespace("default")

	testCases := []struct {
		name             string
		ref              *mockRef
		tpl              string
		objs             []client.Object
		defaultNamespace string
		wantExternal     string
		wantErr          string
	}{
		{
			name: "empty template",
			ref: &mockRef{
				Name: "test",
			},
			tpl:          "",
			objs:         []client.Object{testObj},
			wantExternal: "projects/projectId/locations/us/resources/test",
		},
		{
			name: "referenced object not found",
			ref: &mockRef{
				Name: "non-existent",
			},
			tpl:     "{{selfLink}}",
			objs:    []client.Object{},
			wantErr: "reference TestResource default/non-existent is not found"},
		{
			name: "normalize fails",
			ref: &mockRef{
				Name: "test-no-external-ref",
			},
			tpl:     "{{selfLink}}",
			objs:    []client.Object{testObjNoExternalRef},
			wantErr: "reference TestResource default/test-no-external-ref is not ready",
		},
		{
			name: "template field not found",
			ref: &mockRef{
				Name: "test",
			},
			tpl:     "{{nonExistentField}}",
			objs:    []client.Object{testObj},
			wantErr: "template field \"nonExistentField\" not found",
		},
		{
			name: "template with resourceId",
			ref: &mockRef{
				Name: "test",
			},
			tpl:          "myResources/{{resourceId}}",
			objs:         []client.Object{testObj},
			wantExternal: "myResources/test",
		},
		{
			name: "template with status field",
			ref: &mockRef{
				Name: "test",
			},
			tpl:          "{{selfLink}}",
			objs:         []client.Object{testObj},
			wantExternal: "https://link/to/resource",
		},
		{
			name: "template with status.observedState field",
			ref: &mockRef{
				Name: "test",
			},
			tpl:          "{{address}}",
			objs:         []client.Object{testObj},
			wantExternal: "10.0.0.1",
		},
		{
			name: "complex template",
			ref: &mockRef{
				Name: "test",
			},
			tpl:          "{{selfLinkWithoutId}}/myResources/{{resourceId}}",
			objs:         []client.Object{testObj},
			wantExternal: "https://link/myResources/test",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			fakeClient := fake.NewClientBuilder().WithObjects(tc.objs...).Build()

			err := tc.ref.NormalizeOnTemplate(ctx, fakeClient, tc.defaultNamespace, tc.tpl)

			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("got nil error, want %q", tc.wantErr)
				}
				if !strings.Contains(err.Error(), tc.wantErr) {
					t.Errorf("got error %q, want it to contain %q", err.Error(), tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			if tc.ref.External != tc.wantExternal {
				t.Errorf("got external %q, want %q", tc.ref.External, tc.wantExternal)
			}
		})
	}
}
