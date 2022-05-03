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

package testk8s

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"
)

func GetNamespace(t *testing.T, kubeClient client.Client, namespaceName string) *corev1.Namespace {
	t.Helper()
	namespace := corev1.Namespace{}
	if err := kubeClient.Get(context.TODO(), types.NamespacedName{Name: namespaceName}, &namespace); err != nil {
		t.Fatalf("error getting namespace: %v", err)
	}
	return &namespace
}

func RemoveDeletionDefenderFinalizer(t *testing.T, obj metav1.Object, gvk schema.GroupVersionKind, c client.Client) {
	t.Helper()
	nn := k8s.GetNamespacedName(obj)
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	if err := c.Get(context.Background(), nn, u); err != nil {
		if errors.IsNotFound(err) {
			return
		}
		t.Fatalf("error getting %v %v: %v", gvk, nn, err)
	}
	if !k8s.HasFinalizer(u, k8s.DeletionDefenderFinalizerName) {
		return
	}
	k8s.RemoveFinalizer(u, k8s.DeletionDefenderFinalizerName)
	if err := c.Update(context.Background(), u); err != nil {
		t.Fatalf("error updating %v %v: %v", gvk, nn, err)
	}
}

func RemoveDeletionDefenderFinalizerForUnstructured(t *testing.T, u *unstructured.Unstructured, c client.Client) {
	t.Helper()
	RemoveDeletionDefenderFinalizer(t, u, u.GroupVersionKind(), c)
}

func MapToFieldPathSet(t *testing.T, m map[string]interface{}) *fieldpath.Set {
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatal("error marshaling field path set JSON:", err)
	}
	res := fieldpath.NewSet()
	if err := res.FromJSON(bytes.NewReader(b)); err != nil {
		t.Fatal("error constructing expected set from JSON:", err)
	}
	return res
}
