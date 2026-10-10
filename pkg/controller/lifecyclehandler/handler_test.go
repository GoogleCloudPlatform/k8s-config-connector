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

package lifecyclehandler

import (
	"context"
	"errors"
	"strings"
	"testing"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
)

func TestIsOrphaned(t *testing.T) {
	tests := []struct {
		name             string
		resource         *k8s.Resource
		parentObjectName string
		parentConfigs    []corekccv1alpha1.TypeConfig
		isOrphaned       bool
	}{
		{
			name: "the parent k8s resource is gone",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"barRef": map[string]interface{}{
						"name": "bar-parent",
					},
				},
			},
			parentConfigs: []corekccv1alpha1.TypeConfig{
				{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
					Parent: true,
				},
			},
			isOrphaned: true,
		},
		{
			name: "the parent k8s resource is around",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"barRef": map[string]interface{}{
						"name": "bar-parent",
					},
				},
			},
			parentConfigs: []corekccv1alpha1.TypeConfig{
				{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
					Parent: true,
				},
			},
			parentObjectName: "bar-parent",
			isOrphaned:       false,
		},
		{
			name: "external parent resource",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"barRef": map[string]interface{}{
						"external": "bar-parent",
					},
				},
			},
			parentConfigs: []corekccv1alpha1.TypeConfig{
				{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
					Parent: true,
				},
			},
			isOrphaned: false,
		},
		{
			name: "some resource can have multiple parent resource types, but only one parent field is specified",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"barRef": map[string]interface{}{
						"name": "bar-parent",
					},
				},
			},
			parentConfigs: []corekccv1alpha1.TypeConfig{
				{
					Key: "fooRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Foo",
					},
					Parent: true,
				},
				{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
					Parent: true,
				},
			},
			parentObjectName: "bar-parent",
			isOrphaned:       false,
		},
		{
			name: "resources have no parent",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{},
			},
			parentConfigs: []corekccv1alpha1.TypeConfig{},
			isOrphaned:    false,
		},
	}
	ctx := context.TODO()
	h := test.NewKubeHarness(ctx, t)
	c := h.GetClient()

	h.CreateDummyCRD(schema.GroupVersionKind{Group: "test1.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "Test1Foo"})
	h.CreateDummyCRD(schema.GroupVersionKind{Group: "test1.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "Test1Bar"})

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testID := testvariable.NewUniqueID()
			tc.resource.SetNamespace(testID)
			h.EnsureNamespaceExists(testID)
			if tc.parentObjectName != "" {
				references := []*unstructured.Unstructured{
					test.NewBarUnstructured(tc.parentObjectName, testID, corev1.ConditionTrue),
				}
				test.EnsureObjectsExist(t, references, c)
			}
			res, _, err := IsOrphaned(tc.resource, tc.parentConfigs, c)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if res != tc.isOrphaned {
				t.Fatalf("got resource isOrphaned as %v, want %v", res, tc.isOrphaned)
			}
		})
	}
}

func Test_reasonForUnresolvableDeps(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		want    string
		wantErr bool
	}{
		{
			name:    "reference not ready",
			err:     &k8s.ReferenceNotReadyError{},
			want:    k8s.DependencyNotReady,
			wantErr: false,
		},
		{
			name:    "NewReferenceNotReadyErrorForResource gives DependencyNotReady",
			err:     k8s.NewReferenceNotReadyErrorForResource(&k8s.Resource{}),
			want:    k8s.DependencyNotReady,
			wantErr: false,
		},
		{
			name:    "transitive dependency not ready",
			err:     &k8s.TransitiveDependencyNotReadyError{},
			want:    k8s.DependencyNotReady,
			wantErr: false,
		},
		{
			name:    "reference not found",
			err:     &k8s.ReferenceNotFoundError{},
			want:    k8s.DependencyNotFound,
			wantErr: false,
		},
		{
			name:    "secret not found",
			err:     &k8s.SecretNotFoundError{},
			want:    k8s.DependencyNotFound,
			wantErr: false,
		},
		{
			name:    "transitive dependency not found",
			err:     &k8s.TransitiveDependencyNotFoundError{},
			want:    k8s.DependencyNotFound,
			wantErr: false,
		},
		{
			name:    "NewTransitiveDependencyNotFoundError gives DependencyNotFound",
			err:     k8s.NewTransitiveDependencyNotFoundError(schema.GroupVersionKind{}, types.NamespacedName{}),
			want:    k8s.DependencyNotFound,
			wantErr: false,
		},
		{
			name:    "key in secret not found",
			err:     &k8s.KeyInSecretNotFoundError{},
			want:    k8s.DependencyInvalid,
			wantErr: false,
		},
		{
			name:    "unrecognized error",
			err:     errors.New("some error"),
			want:    "",
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := reasonForUnresolvableDeps(test.err)
			if test.wantErr {
				if err == nil {
					t.Errorf("reasonForUnresolvableDeps() error = nil, want err")
				}
				return
			}
			if err != nil {
				t.Errorf("reasonForUnresolvableDeps() error = %v, want nil", err)
				return
			}
			if got != test.want {
				t.Errorf("reasonForUnresolvableDeps() got = %v, want %v", got, test.want)
			}
		})
	}
}

func TestHandlePaused(t *testing.T) {
	ctx := context.TODO()
	h := test.NewKubeHarness(ctx, t)
	c := h.GetClient()

	h.CreateDummyCRD(schema.GroupVersionKind{Group: "test1.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "Test1Bar"})

	testID := testvariable.NewUniqueID()
	h.EnsureNamespaceExists(testID)

	bar := test.NewBarUnstructured("test-bar", testID, "")
	test.EnsureObjectsExist(t, []*unstructured.Unstructured{bar}, c)

	recorder := record.NewFakeRecorder(100)
	handler := NewLifecycleHandler(c, recorder)

	resource, err := k8s.NewResource(bar)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := handler.HandlePaused(ctx, resource); err != nil {
		t.Fatalf("HandlePaused failed: %v", err)
	}

	// Verify an event was recorded on state transition
	select {
	case event := <-recorder.Events:
		if !strings.Contains(event, k8s.Paused) {
			t.Errorf("unexpected event: %s, want event containing %s", event, k8s.Paused)
		}
	default:
		t.Errorf("expected an event to be recorded, but got none")
	}

	// Fetch object from apiserver and verify condition
	updatedBar := &unstructured.Unstructured{}
	updatedBar.SetGroupVersionKind(bar.GroupVersionKind())
	if err := c.Get(ctx, types.NamespacedName{Namespace: testID, Name: "test-bar"}, updatedBar); err != nil {
		t.Fatalf("failed to get updated bar: %v", err)
	}

	updatedResource, err := k8s.NewResource(updatedBar)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cond, found := k8s.GetReadyCondition(updatedResource)
	if !found {
		t.Fatalf("ready condition not found")
	}
	if cond.Status != corev1.ConditionFalse {
		t.Errorf("ready condition status = %v, want %v", cond.Status, corev1.ConditionFalse)
	}
	if cond.Reason != k8s.Paused {
		t.Errorf("ready condition reason = %v, want %v", cond.Reason, k8s.Paused)
	}
	if cond.Message != k8s.PausedMessage {
		t.Errorf("ready condition message = %v, want %v", cond.Message, k8s.PausedMessage)
	}

	// Subsequent HandlePaused call on already-paused resource should be idempotent and not emit duplicate events
	if err := handler.HandlePaused(ctx, updatedResource); err != nil {
		t.Fatalf("HandlePaused second call failed: %v", err)
	}

	select {
	case event := <-recorder.Events:
		t.Errorf("expected no event on idempotent HandlePaused call, but got: %s", event)
	default:
		// expected
	}
}
